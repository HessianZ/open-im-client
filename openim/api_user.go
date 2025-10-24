package openim

import (
	"net/http"

	"github.com/openimsdk/protocol/msggateway"
	"github.com/openimsdk/protocol/user"
)

// POST /user/user_register
func (c *OpenIMClient) UserRegister(data *user.UserRegisterReq) error {
	resp := OpenIMResponse[*user.UserRegisterResp]{}
	return c.request(http.MethodPost, "/user/user_register", data, &resp)
}

// POST /user/account_check
func (c *OpenIMClient) AccountCheck(data *user.AccountCheckReq) (*user.AccountCheckResp, error) {
	resp := OpenIMResponse[*user.AccountCheckResp]{}
	err := c.request(http.MethodPost, "/user/account_check", data, &resp)
	return resp.Data, err
}

// 获取已注册用户列表
// POST /user/get_users
func (c *OpenIMClient) GetUsers(data *user.GetPaginationUsersReq) (*user.GetPaginationUsersResp, error) {
	resp := OpenIMResponse[*user.GetPaginationUsersResp]{}
	err := c.request(http.MethodPost, "/user/get_users", data, &resp)
	return resp.Data, err
}

// 获取已注册的用户 ID 列表。
// POST /user/get_all_users_uid
func (c *OpenIMClient) GetAllUsersUid() (*user.GetAllUserIDResp, error) {
	resp := OpenIMResponse[*user.GetAllUserIDResp]{}
	err := c.request(http.MethodPost, "/user/get_all_users_uid", nil, &resp)
	return resp.Data, err
}

// 获取指定用户详情列表
// POST /user/get_users_info
func (c *OpenIMClient) GetUsersInfo(data *user.GetDesignateUsersReq) (*user.GetDesignateUsersResp, error) {
	resp := OpenIMResponse[*user.GetDesignateUsersResp]{}
	err := c.request(http.MethodPost, "/user/get_users_info", data, &resp)
	return resp.Data, err
}

// 获取指定用户在线状态
// 获取指定用户各终端在线状态和 token。
// POST /user/get_users_online_status
func (c *OpenIMClient) GetUsersOnlineStatus(data *msggateway.GetUsersOnlineStatusReq) ([]*msggateway.GetUsersOnlineStatusResp_SuccessResult, error) {
	resp := OpenIMResponse[[]*msggateway.GetUsersOnlineStatusResp_SuccessResult]{}
	err := c.request(http.MethodPost, "/user/get_users_online_status", data, &resp)
	return resp.Data, err
}

// 获取指定用户在线token详情
// 获取指定用户处于线上的平台下的 token 数组列表。
// POST /user/get_users_online_token_detail
func (c *OpenIMClient) GetUsersOnlineTokenDetail(data *msggateway.GetUsersOnlineStatusReq) ([]*msggateway.SingleDetail, error) {
	resp := OpenIMResponse[[]*msggateway.SingleDetail]{}
	err := c.request(http.MethodPost, "/user/get_users_online_token_detail", data, &resp)
	return resp.Data, err
}

// 获取指定用户所订阅用户的在线状态
// POST /user/get_subscribe_users_status
func (c *OpenIMClient) GetSubscribeUsersStatus(data *user.GetSubscribeUsersStatusReq) (*user.GetSubscribeUsersStatusResp, error) {
	resp := OpenIMResponse[*user.GetSubscribeUsersStatusResp]{}
	err := c.request(http.MethodPost, "/user/get_subscribe_users_status", data, &resp)
	return resp.Data, err
}

// 修改指定用户全局免打扰状态
// POST /user/set_global_msg_recv_opt
func (c *OpenIMClient) SetGlobalMsgRecvOpt(data *user.SetGlobalRecvMessageOptReq) error {
	resp := OpenIMResponse[*user.SetGlobalRecvMessageOptResp]{}
	return c.request(http.MethodPost, "/user/set_global_msg_recv_opt", data, &resp)
}

// 修改用户的在线状态订阅关系
// 修改（订阅或者取消订阅）指定用户的在线状态订阅关系。
// POST /user/subscribe_users_status
func (c *OpenIMClient) SubscribeUsersStatus(data *user.SubscribeOrCancelUsersStatusReq) error {
	resp := OpenIMResponse[*user.SubscribeOrCancelUsersStatusResp]{}
	return c.request(http.MethodPost, "/user/subscribe_users_status", data, &resp)
}

// 修改用户信息
// 修改用户的头像、名称、Ex等信息。仅传需要修改的字段，也支持零值。
// POST /user/update_user_info_ex
func (c *OpenIMClient) UpdateUserInfoEx(data *user.UpdateUserInfoExReq) error {
	resp := OpenIMResponse[*user.UpdateUserInfoExResp]{}
	return c.request(http.MethodPost, "/user/update_user_info_ex", data, &resp)
}

// 获取系统号
// 管理员可以获取所创建的系统号列表，包括 ID、头像和名称。
// POST /user/search_notification_account
func (c *OpenIMClient) SearchNotificationAccount(data *user.SearchNotificationAccountReq) (*user.SearchNotificationAccountResp, error) {
	resp := OpenIMResponse[*user.SearchNotificationAccountResp]{}
	err := c.request(http.MethodPost, "/user/search_notification_account", data, &resp)
	return resp.Data, err
}

// 增加系统号
// 增加系统号，管理员可以以系统号身份发送通知消息。
// POST /user/add_notification_account
func (c *OpenIMClient) AddNotificationAccount(data *user.AddNotificationAccountReq) (*user.AddNotificationAccountResp, error) {
	resp := OpenIMResponse[*user.AddNotificationAccountResp]{}
	err := c.request(http.MethodPost, "/user/add_notification_account", data, &resp)
	return resp.Data, err
}

// 修改系统号信息
// 修改系统号信息，包括头像和名称。
// POST /user/update_notification_account
func (c *OpenIMClient) UpdateNotificationAccount(data *user.UpdateNotificationAccountInfoReq) error {
	resp := OpenIMResponse[*user.UpdateNotificationAccountInfoResp]{}
	return c.request(http.MethodPost, "/user/update_notification_account", data, &resp)
}
