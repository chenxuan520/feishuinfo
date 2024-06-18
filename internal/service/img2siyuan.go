package service

import ()

type Img2SiyuanModule struct {
	*AbstractModule
}

func (i *Img2SiyuanModule) Deal(userId, content string, messageType MessageType) {
	// TODO: 这部分感觉是一个伪需求,en,暂时搁置一下 //
}

func (i *Img2SiyuanModule) RegisterStr() string {
	return "img"
}

func (i *Img2SiyuanModule) HelpStr() string {
	return "img img_url,will get siyuan img link"
}

func (i *Img2SiyuanModule) RegisterType() []MessageType {
	return []MessageType{ImageMessage, TextMessage}
}

func NewImg2SiyuanModule() *Img2SiyuanModule {
	return &Img2SiyuanModule{
		AbstractModule: NewAbstractModule(),
	}
}
