package handler

import (
	"github.com/gin-gonic/gin"
)

func HandleLark(c *gin.Context) {
	body := make(map[string]interface{})
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, body)

	//var e lark.Event
	//err := c.ShouldBindJSON(&e)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//userID := e.Event.Sender.SenderID.UserID
	//var msg lark.MsgContent
	//_ = json.Unmarshal([]byte(e.Event.Message.Content), &msg)
	//err = registry.Add(userID, msg.Text)
	//if err != nil {
	//	logrus.WithError(err).Error("add user error.")
	//	err = lark.SendMsg(userID, "text", "{\"text\":\"添加失败，凭证不合法\"}")
	//	if err != nil {
	//		logrus.WithError(err).Error("send message error")
	//	}
	//	c.JSON(http.StatusOK, err.Error())
	//	return
	//}
	//lcClient := leetcode.NewClient(msg.Text)
	//status, err := lcClient.GetUserStatus()
	//if err != nil {
	//	return
	//}
	//var lcNotify btmpl.LcNotify
	//lcNotify.Username = status.Data.UserStatus.Username
	//tmpl, err := template.New("template").Parse(btmpl.LcNotifySuccess)
	//var filledTmpl bytes.Buffer
	//err = tmpl.Execute(&filledTmpl, &lcNotify)
	//if err != nil {
	//	logrus.WithError(err).Error("template execute error")
	//	return
	//}
	//err = lark.SendMsg(userID, "interactive", filledTmpl.String())
	//if err != nil {
	//	logrus.WithError(err).Error("send message error")
	//}
	//
	//c.JSON(http.StatusOK, "ok")
}
