package openim

import (
	"net/http"

	"github.com/HessianZ/open-im-client/openim/apistruct"
	"github.com/openimsdk/protocol/msg"
)

// POST /msg/send_msg
func (c *OpenIMClient) SendMsg(data apistruct.SendMsgReq) (*apistruct.SendMsgResp, error) {
	resp := OpenIMResponse[*apistruct.SendMsgResp]{}
	err := c.request(http.MethodPost, "/msg/send_msg", data, &resp)
	return resp.Data, err
}

// POST /msg/batch_send_msg
func (c *OpenIMClient) BatchSendMsg(data apistruct.BatchSendMsgReq) (*apistruct.BatchSendMsgResp, error) {
	resp := OpenIMResponse[*apistruct.BatchSendMsgResp]{}
	err := c.request(http.MethodPost, "/msg/batch_send_msg", data, &resp)
	return resp.Data, err
}

// @see https://github.com/openimsdk/open-im-server/tree/main/internal/api/msg.go:335
type SendBusinessNotificationReq struct {
	Key         string `json:"key"`                           // 根据业务分类，客户端可以通过改字段用不同方法处理data
	Data        string `json:"data"`                          // 业务数据
	SendUserID  string `json:"sendUserID" binding:"required"` // 系统通知号ID，或用户ID
	RecvUserID  string `json:"recvUserID"`                    // 接收者用户ID，与recvGroupID只能选其中一个
	RecvGroupID string `json:"recvGroupID"`                   // 接收群ID，与recvUserID只能选其中一个
	SendMsg     bool   `json:"sendMsg"`                       // 是否已消息形式发送，默认: false
	// 通知消息的可靠级别，1: 在线推送。2: 必达通知(断线重连或重新登录也会触发，用于必达的场景，该可靠性等级下，由于是顺序全量同步，建议不能发送过多，否则会影响客户端消息同步性能)，默认: 1
	ReliabilityLevel *int `json:"reliabilityLevel"`
}

// POST /msg/send_business_notification
func (c *OpenIMClient) SendBusinessNotification(data SendBusinessNotificationReq) (*msg.SendMsgResp, error) {
	resp := OpenIMResponse[*msg.SendMsgResp]{}
	err := c.request(http.MethodPost, "/msg/send_business_notification", data, &resp)
	return resp.Data, err
}
