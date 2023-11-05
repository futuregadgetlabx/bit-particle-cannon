package lark

import (
	"context"
	"github.com/futuregadgetlabx/bit-particle-cannon/config"
	"github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

var client *lark.Client

func newClient() *lark.Client {
	if client != nil {
		return client
	}
	// 启用SDK自动管理租户Token的获取与刷新，可调用lark.WithEnableTokenCache(true)进行设置
	client = lark.NewClient(config.App.Lark.Bot.AppId, config.App.Lark.Bot.AppSecret, lark.WithEnableTokenCache(true))
	return client
}

func SendMsg(receiver, msgType, content string) error {
	if client != nil {
		client = newClient()
	}
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(`user_id`).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiver).
			MsgType(msgType).
			Content(content).
			Build()).
		Build()
	return doSend(req)
}

func doSend(req *larkim.CreateMessageReq) error {
	resp, err := client.Im.Message.Create(context.Background(), req)
	if err != nil {
		logrus.WithError(err).Error("send message error.")
		return err
	}
	if resp.Code != 0 {
		logrus.WithError(err).Error("send message error.")
		return err
	}
	return nil
}
