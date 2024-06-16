package view

import (
	"context"
	"encoding/json"

	"github.com/chenxuan520/feishuinfo/internal/config"
	"github.com/chenxuan520/feishuinfo/internal/logger"
	"github.com/chenxuan520/feishuinfo/internal/service"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type EventRoute struct {
	ChatService *service.ChatService
}

func (e *EventRoute) InitEvent() *dispatcher.EventDispatcher {
	//register event handle
	events := dispatcher.
		NewEventDispatcher(config.GlobalConfig.Feishu.Verification, config.GlobalConfig.Feishu.EncryptKey).
		OnP2MessageReceiveV1(e.MsgReceive)
	return events
}

func (e *EventRoute) MsgReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	if event.Event == nil || event.Event.Message == nil || event.Event.Sender == nil ||
		event.Event.Sender.SenderId.UserId == nil || event.Event.Message.MessageType == nil {
		logger.GetLogger().Error("Error:get wrong message")
		return nil
	}
	userId := *event.Event.Sender.SenderId.UserId

	text := larkim.MessagePostText{}
	err := json.Unmarshal([]byte(*event.Event.Message.Content), &text)
	if err != nil {
		return nil
	}
	switch *event.Event.Message.MessageType {
	case "text":
		e.ChatService.DealMsg(userId, text.Text)
	default:
	}
	return nil
}

func NewEventRoute() *EventRoute {
	return &EventRoute{
		ChatService: service.NewChatService(),
	}
}
