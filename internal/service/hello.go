package service

type HelloModule struct {
	*AbstractModule
}

func (h *HelloModule) Deal(userId, content string) {
	h.RobotSendTextMsg(userId, "hello "+content)
}

func (h *HelloModule) RegisterStr() string {
	return "im"
}

func (h *HelloModule) HelpStr() string {
	return "say im xxx,will get hello xxx"
}

func NewHelloModule() *HelloModule {
	return &HelloModule{
		AbstractModule: NewAbstractModule(),
	}
}
