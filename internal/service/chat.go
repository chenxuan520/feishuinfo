package service

import (
	"context"
	"fmt"
	"strings"
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

func (c *ChatService) DealMsg(userId, content string) {
	if userId == "" || content == "" {
		return
	}
	logger.GetLogger().Debug(fmt.Sprintf("DEBUG:userId:%s,content:%s", userId, content))
	err := c.RobotSendTextMsg(userId, "Hello World")
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("ERROR:DealMsg:send msg error:%s", err.Error()))
	}
}

func (c *ChatService) RobotSendTextMsg(receiveID, content string) error {
	if content == "" {
		logger.GetLogger().Debug("DEBUG:content is empty")
		return nil
	}
	// 这里需要将error中的"替换成\"，否则在发消息时会出现json反序列化错误
	content = strings.Replace(content, "\"", "\\\"", -1)

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
