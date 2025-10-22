package openim

import "net/http"

type GetAdminTokenData struct {
	Token             string `json:"token"`
	ExpireTimeSeconds int64  `json:"expireTimeSeconds"`
}

// POST /auth/get_admin_token
func (c *OpenIMClient) GetAdminToken() (*GetAdminTokenData, error) {
	resp := OpenIMResponse[*GetAdminTokenData]{}
	data := map[string]string{
		"secret": c.adminSecret,
		"userID": c.adminID,
	}
	err := c.request(http.MethodPost, "/auth/get_admin_token", data, &resp)

	return resp.Data, err
}

type GetUserTokenData struct {
	Token             string `json:"token"`
	ExpireTimeSeconds int64  `json:"expireTimeSeconds"`
}

// POST /auth/get_user_token
func (c *OpenIMClient) GetUserToken(platformID PlatformID, userID string) (*GetUserTokenData, error) {
	resp := OpenIMResponse[*GetUserTokenData]{}
	data := map[string]any{
		"platformID": platformID, // User login terminal type, values range from 1-9
		"userID":     userID,
	}
	err := c.request(http.MethodPost, "/auth/get_user_token", data, &resp)
	return resp.Data, err
}

// 强制用户下线
// 强制用户从某个终端退出登录，客户端 SDK 会收到 onKickedOffline 回调事件。
// POST /auth/force_logout
func (c *OpenIMClient) ForceLogout(platformID PlatformID, userID string) error {
	data := map[string]any{
		"platformID": platformID, // User login terminal type, values range from 1-9
		"userID":     userID,
	}
	err := c.request(http.MethodPost, "/auth/force_logout", data, nil)
	return err
}
