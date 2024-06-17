package service

import (
	"fmt"

	"github.com/chenxuan520/feishuinfo/internal/logger"
	"github.com/chenxuan520/feishuinfo/internal/tools"
)

type Module interface {
	Deal(userId, content string)
	RegisterStr() string
	HelpStr() string
}

type AbstractModule struct {
	*ChatService
}

func (a *AbstractModule) Deal(userId, content string) {
	logger.GetLogger().Error("ERROR:AbstractModule:Deal:abstract module can not deal msg")
	a.RobotSendTextMsg(userId, "abstract module can not deal msg")
}

func (a *AbstractModule) RegisterStr() string {
	logger.GetLogger().Error("ERROR:AbstractModule:RegisterStr:abstract module can not register")
	return ""
}

func (a *AbstractModule) HelpStr() string {
	logger.GetLogger().Error("ERROR:AbstractModule:HelpStr:abstract module can not help")
	return ""
}

func NewAbstractModule() *AbstractModule {
	return &AbstractModule{
		ChatService: NewChatService(),
	}
}

type ModuleService struct {
	ChatService    *ChatService
	Modules        []Module
	DefaultContent string
}

var DefaultModuleService *ModuleService = nil

func (m *ModuleService) RegisterModule(module Module) {
	m.Modules = append(m.Modules, module)
}

func (m *ModuleService) DealMsg(userId, content string) {
	if userId == "" || content == "" {
		return
	}
	logger.GetLogger().Debug(fmt.Sprintf("DEBUG:userId:%s,content:%s", userId, content))
	for _, module := range m.Modules {
		// 将content第一个单词提取出来,匹配
		prefixWord, args := tools.ExtractFirstWord(content)
		if prefixWord == module.RegisterStr() {
			module.Deal(userId, args)
			return
		}
	}
	err := m.ChatService.RobotSendTextMsg(userId, m.DefaultContent)
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("ERROR:DealMsg:send default msg error:%s", err.Error()))
	}
}

func NewModuleService() *ModuleService {
	return DefaultModuleService
}

func InitModuleService() {
	if DefaultModuleService == nil {
		DefaultModuleService = &ModuleService{
			ChatService:    NewChatService(),
			Modules:        make([]Module, 0),
			DefaultContent: "Application For Feishu Quick Help\n",
		}

		// TODO:register model in there
		DefaultModuleService.RegisterModule(NewHelloModule())

		// add module help message
		for _, module := range DefaultModuleService.Modules {
			DefaultModuleService.DefaultContent += module.RegisterStr() + " " + module.HelpStr() + "\n"
		}
	}
}
