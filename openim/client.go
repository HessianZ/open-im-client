package openim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type OpenIMClient struct {
	apiBaseUrl    string
	tokenProvider TokenProvider
	httpClient    *http.Client
	adminID       string
	adminSecret   string
	debug         bool
}

func NewOpenIMClient(apiBaseUrl, adminID, adminSecret string, debug bool) *OpenIMClient {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	return &OpenIMClient{
		apiBaseUrl:    apiBaseUrl,
		adminID:       adminID,
		adminSecret:   adminSecret,
		debug:         debug,
		httpClient:    httpClient,
		tokenProvider: &MemoryTokenStore{},
	}
}

// base http request
func (c *OpenIMClient) request(method, path string, reqBody any, respData any) (err error) {
	var reqReader io.Reader
	if method == "POST" && reqBody != nil {
		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		reqReader = bytes.NewReader(reqBytes)
	}
	req, err := http.NewRequest(method, c.apiBaseUrl+path, reqReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("operationID", uuid.NewString())
	if path != "/auth/get_admin_token" {
		err = c.loadToken(false)
		if err != nil {
			return
		}
		if token, err := c.tokenProvider.GetToken(); err != nil {
			return err
		} else {
			req.Header.Set("token", token)
		}
	}

	if c.debug {
		log.Printf("[debug] request: %s %s %+v", req.Method, req.URL.String(), req.Header)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed, HTTP error %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)

	if c.debug {
		log.Printf("[debug] respBody: %s %s %s", req.Method, req.URL.String(), respBody)
	}

	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}
	if respData == nil {
		respData = OpenIMResponse[struct{}]{}
	}
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Printf("Error unmarshalling response body: %v - %s", err, string(respBody))
		return
	}

	if respWithType, ok := respData.(OpenIMError); ok {
		if respWithType.GetErrCode() != 0 {
			log.Printf("OpenIMError: %+v", respWithType)
			if respWithType.GetErrCode() == ErrCode_TokenNotExists {
				// 获取token失败，需要重新获取
				c.tokenProvider.SetToken("", 0)
			}
			err = respWithType
		}
	}

	return
}

func (c *OpenIMClient) SetTokenProvider(provider TokenProvider) {
	c.tokenProvider = provider
}

func (c *OpenIMClient) loadToken(force bool) (err error) {
	token, err := c.tokenProvider.GetToken()

	if token != "" && !force {
		return
	}
	tokenData, err := c.GetAdminToken()
	if err != nil {
		return err
	}
	err = c.tokenProvider.SetToken(tokenData.Token, tokenData.ExpireTimeSeconds)
	if err != nil {
		return
	}
	return
}
