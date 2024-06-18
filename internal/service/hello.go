package service

type HelloModule struct {
	*AbstractModule
}

func (h *HelloModule) Deal(userId, content string, messageType MessageType) {
	h.RobotSendTextMsg(userId, "hello "+content)
}

func (h *HelloModule) RegisterStr() string {
	return "im"
}

func (h *HelloModule) HelpStr() string {
	return "say im xxx,will get hello xxx"
}

func (h *HelloModule) RegisterType() []MessageType {
	return []MessageType{TextMessage}
}

func NewHelloModule() *HelloModule {
	return &HelloModule{
		AbstractModule: NewAbstractModule(),
	}
}
