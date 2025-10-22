package openim

import "net/http"

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
