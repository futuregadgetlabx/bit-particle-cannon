package handler

import (
	"encoding/json"
	"fmt"
	"github.com/futuregadgetlabx/bit-particle-cannon/registry"
	"github.com/futuregadgetlabx/bit-particle-cannon/util/lark"
	"github.com/futuregadgetlabx/bit-particle-cannon/util/leetcode"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// HandleLark 监听飞书事件
// 所有返回都使用200，避免事件返回其他状态码导致飞书重复发送事件消息
func HandleLark(c *gin.Context) {
	var e lark.Event
	err := c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	userID := e.Event.Sender.SenderID.UserID
	var msg lark.MsgContent
	_ = json.Unmarshal([]byte(e.Event.Message.Content), &msg)
	err = registry.Add(userID, msg.Text)
	if err != nil {
		logrus.WithError(err).Error("add user error.")
		err = lark.SendMsg(userID, "text", "user_id", "{\"text\":\"添加失败，凭证不合法\"}")
		if err != nil {
			logrus.WithError(err).Error("send message error")
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, "ok")
		return
	}
	lcClient := leetcode.NewClient(msg.Text)
	status, err := lcClient.GetUserStatus()
	if err != nil {
		return
	}
	err = lark.SendMsg(userID, "text", "user_id",
		fmt.Sprintf("{\"text\":\"添加用户[%v]成功，凭证不合法\"}", status.Data.UserStatus.Username))
	if err != nil {
		logrus.WithError(err).Error("send message error")
		c.JSON(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
