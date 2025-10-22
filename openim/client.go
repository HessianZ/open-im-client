package openim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type OpenIMError interface {
	GetErrCode() int
	GetErrMsg() string
	GetErrDetail() string
	Error() string
}
type OpenIMClient struct {
	apiBaseUrl  string
	token       string
	expireTime  time.Time
	tokenLock   sync.Mutex
	httpClient  *http.Client
	adminID     string
	adminSecret string
	debug       bool
}

type openIMErrorSt struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	ErrDlt  string `json:"errDlt"`
}

type OpenIMResponse[T any] struct {
	openIMErrorSt
	Data T `json:"data"`
}

func (r *OpenIMResponse[T]) GetErrCode() int {
	return r.openIMErrorSt.ErrCode
}
func (r *OpenIMResponse[T]) GetErrMsg() string {
	return r.openIMErrorSt.ErrMsg
}
func (r *OpenIMResponse[T]) GetErrDetail() string {
	return r.openIMErrorSt.ErrDlt
}

func (e *OpenIMResponse[T]) Error() string {
	return fmt.Sprintf("openim error[%d] %s - %s", e.ErrCode, e.ErrMsg, e.ErrDlt)
}

func NewOpenIMClient(apiBaseUrl, adminID, adminSecret string, debug bool) *OpenIMClient {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	return &OpenIMClient{
		apiBaseUrl:  apiBaseUrl,
		adminID:     adminID,
		adminSecret: adminSecret,
		debug:       debug,
		httpClient:  httpClient,
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
		req.Header.Set("token", c.token)
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
				c.token = ""
				c.expireTime = time.Time{}
			}
			err = respWithType
		}
	}

	return
}
func (c *OpenIMClient) loadToken(force bool) (err error) {
	c.tokenLock.Lock()
	defer c.tokenLock.Unlock()

	if c.token != "" && c.expireTime.After(time.Now()) && !force {
		return
	}
	tokenData, err := c.GetAdminToken()
	if err != nil {
		return err
	}
	c.token = tokenData.Token
	c.expireTime = time.Now().Add(time.Duration(tokenData.ExpireTimeSeconds) * time.Second)
	return
}

// POST /user/user_register
func (c *OpenIMClient) UserRegister(userID, nickname, faceURL string) error {
	data := map[string]any{
		"users": []map[string]any{
			{
				"userID":   userID,
				"nickname": nickname,
				"faceURL":  faceURL,
			},
		},
	}
	return c.request(http.MethodPost, "/user/user_register", data, nil)
}

// POST /user/update_user_info_ex
func (c *OpenIMClient) UpdateUserInfo(userID, nickname, faceURL, extra string) error {
	data := map[string]any{
		"userInfo": map[string]any{
			"userID":   userID,
			"nickname": nickname,
			"faceURL":  faceURL,
			"ex":       extra,
		},
	}
	return c.request(http.MethodPost, "/user/update_user_info_ex", data, nil)
}
