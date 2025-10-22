package openim

import (
	"net/http"

	"github.com/openimsdk/protocol/conversation"
)

// 获取当前用户的会话列表
// POST /conversation/get_owner_conversation
func (c *OpenIMClient) GetOwnerConversation(data *conversation.GetOwnerConversationReq) (*conversation.GetOwnerConversationResp, error) {
	resp := OpenIMResponse[*conversation.GetOwnerConversationResp]{}
	err := c.request(http.MethodPost, "/conversation/get_owner_conversation", data, &resp)
	return resp.Data, err
}

// 获取排序后的会话列表
// POST /conversation/get_sorted_conversation_list
func (c *OpenIMClient) GetSortedConversationList(data *conversation.GetSortedConversationListReq) (*conversation.GetSortedConversationListResp, error) {
	resp := OpenIMResponse[*conversation.GetSortedConversationListResp]{}
	err := c.request(http.MethodPost, "/conversation/get_sorted_conversation_list", data, &resp)
	return resp.Data, err
}

// 多个用户对同一ID的会话设置字段
// POST /conversation/set_conversations
func (c *OpenIMClient) SetConversations(data *conversation.SetConversationsReq) (*conversation.SetConversationsResp, error) {
	resp := OpenIMResponse[*conversation.SetConversationsResp]{}
	err := c.request(http.MethodPost, "/conversation/set_conversations", data, &resp)
	return resp.Data, err
}
