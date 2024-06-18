package service

import (
	"context"
	"fmt"
	"time"

	"github.com/chenxuan520/feishuinfo/internal/logger"
	"github.com/chenxuan520/feishuinfo/internal/tools"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type ChatService struct {
}

var DefaultChatService *ChatService = nil

func NewChatService() *ChatService {
	if DefaultChatService == nil {
		DefaultChatService = &ChatService{}
	}
	return DefaultChatService
}

func (c *ChatService) RobotSendTextMsg(receiveID, content string) error {
	if content == "" {
		logger.GetLogger().Debug("DEBUG:content is empty")
		return nil
	}
	// 这里需要将content中的"替换成\"，否则在发消息时会出现json反序列化错误
	tools.ReplaceSpecialChar(content)

	content = larkim.NewTextMsgBuilder().Text(content).Build()
	uuid := time.Now().String()
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("user_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveID).
			MsgType("text").
			Content(content).
			Uuid(tools.MD5(uuid)).
			Build()).
		Build()
	resp, err := tools.GlobalLark.Im.Message.Create(context.Background(), req, larkcore.WithTenantKey(tools.GetAccessToken()))
	if err != nil {
		return err
	}
	if !resp.Success() {
		return fmt.Errorf("error:%d %s %s", resp.Code, resp.Msg, resp.RequestId())
	}
	return nil
}
