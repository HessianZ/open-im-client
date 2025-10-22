package openim

import (
	"net/http"

	"github.com/openimsdk/protocol/relation"
)

// 导入好友
// 让指定用户（ownerUserID）和多个用户（friendUserIDs）建立好友关系。
// POST /friend/import_friend
func (c *OpenIMClient) ImportFriend(data *relation.ImportFriendReq) (*relation.ImportFriendResp, error) {
	resp := OpenIMResponse[*relation.ImportFriendResp]{}
	err := c.request(http.MethodPost, "/friend/import_friend", data, &resp)
	return resp.Data, err
}

// 修改好友信息
// 修改指定用户（ownerUserID）的某个好友（friendUserID）的备注、星标信息。仅传需要修改的字段，也支持零值。
// POST /friend/update_friends
func (c *OpenIMClient) UpdateFriends(data *relation.UpdateFriendsReq) (*relation.UpdateFriendsResp, error) {
	resp := OpenIMResponse[*relation.UpdateFriendsResp]{}
	err := c.request(http.MethodPost, "/friend/update_friends", data, &resp)
	return resp.Data, err
}

// 检查好友关系
// 检查两个用户是否存在好友关系。
// POST /friend/is_friend
func (c *OpenIMClient) IsFriend(data *relation.IsFriendReq) (*relation.IsFriendResp, error) {
	resp := OpenIMResponse[*relation.IsFriendResp]{}
	err := c.request(http.MethodPost, "/friend/is_friend", data, &resp)
	return resp.Data, err
}

// 删除好友
// 将 friendUserID 从 ownerUserID 的好友列表中单向删除。
// POST /friend/delete_friend
func (c *OpenIMClient) DeleteFriend(data *relation.DeleteFriendReq) (*relation.DeleteFriendResp, error) {
	resp := OpenIMResponse[*relation.DeleteFriendResp]{}
	err := c.request(http.MethodPost, "/friend/delete_friend", data, &resp)
	return resp.Data, err
}

// 获取好友列表
// 获取指定用户的好友列表。
// POST /friend/get_friend_list
func (c *OpenIMClient) GetFriendList(data *relation.GetPaginationFriendsReq) (*relation.GetPaginationFriendsResp, error) {
	resp := OpenIMResponse[*relation.GetPaginationFriendsResp]{}
	err := c.request(http.MethodPost, "/friend/get_friend_list", data, &resp)
	return resp.Data, err
}

// 获取发起的好友申请
// 作为主动添加者，获取该用户向其他用户发起的好友申请列表。
// POST /friend/get_self_friend_apply_list
func (c *OpenIMClient) GetSelfFriendApplyList(data *relation.GetPaginationFriendsApplyFromReq) (*relation.GetPaginationFriendsApplyFromResp, error) {
	resp := OpenIMResponse[*relation.GetPaginationFriendsApplyFromResp]{}
	err := c.request(http.MethodPost, "/friend/get_self_friend_apply_list", data, &resp)
	return resp.Data, err
}

// 获取收到的好友申请
// 作为被添加者，获取其他用户向该用户发出的好友申请列表。
// POST /friend/get_friend_apply_list
func (c *OpenIMClient) GetFriendApplyList(data *relation.GetPaginationFriendsApplyToReq) (*relation.GetPaginationFriendsApplyToResp, error) {
	resp := OpenIMResponse[*relation.GetPaginationFriendsApplyToResp]{}
	err := c.request(http.MethodPost, "/friend/get_friend_apply_list", data, &resp)
	return resp.Data, err
}

// 发起好友申请
// fromUserID 向 toUserID 发起好友申请。
// POST /friend/add_friend
func (c *OpenIMClient) AddFriend(data *relation.ApplyToAddFriendReq) (*relation.ApplyToAddFriendResp, error) {
	resp := OpenIMResponse[*relation.ApplyToAddFriendResp]{}
	err := c.request(http.MethodPost, "/friend/add_friend", data, &resp)
	return resp.Data, err
}

// 处理好友申请
// 同意或拒绝 fromUserID 向 toUserID 发起的好友申请。
// POST /friend/add_friend_response
func (c *OpenIMClient) AddFriendResponse(data *relation.RespondFriendApplyReq) (*relation.RespondFriendApplyResp, error) {
	resp := OpenIMResponse[*relation.RespondFriendApplyResp]{}
	err := c.request(http.MethodPost, "/friend/add_friend_response", data, &resp)
	return resp.Data, err
}

// 获取黑名单列表
// 获取指定用户的黑名单列表。
// POST /friend/get_black_list
func (c *OpenIMClient) GetBlackList(data *relation.GetPaginationBlacksReq) (*relation.GetPaginationBlacksResp, error) {
	resp := OpenIMResponse[*relation.GetPaginationBlacksResp]{}
	err := c.request(http.MethodPost, "/friend/get_black_list", data, &resp)
	return resp.Data, err
}

// 添加黑名单
// 将 blackUserID 添加到 ownerUserID 的黑名单中，blackUserID 不能再给 ownerUserID 发送消息。
// POST /friend/add_black
func (c *OpenIMClient) AddBlack(data *relation.AddBlackReq) (*relation.AddBlackResp, error) {
	resp := OpenIMResponse[*relation.AddBlackResp]{}
	err := c.request(http.MethodPost, "/friend/add_black", data, &resp)
	return resp.Data, err
}

// 移除黑名单
// 将 blackUserID 从 ownerUserID 的黑名单列表中移除，即恢复成正常关系。
// POST /friend/remove_black
func (c *OpenIMClient) RemoveBlack(data *relation.RemoveBlackReq) (*relation.RemoveBlackResp, error) {
	resp := OpenIMResponse[*relation.RemoveBlackResp]{}
	err := c.request(http.MethodPost, "/friend/remove_black", data, &resp)
	return resp.Data, err
}
