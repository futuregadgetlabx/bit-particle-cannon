package cron

import (
	"bytes"
	"github.com/futuregadgetlabx/bit-particle-cannon/config"
	"github.com/futuregadgetlabx/bit-particle-cannon/registry"
	"github.com/futuregadgetlabx/bit-particle-cannon/util/lark"
	"github.com/futuregadgetlabx/bit-particle-cannon/util/leetcode"
	btmpl "github.com/futuregadgetlabx/bit-particle-cannon/util/template"
	"github.com/futuregadgetlabx/bit-particle-cannon/util/tianapi"
	"github.com/sirupsen/logrus"
	"html/template"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var difficultyMap = map[string]string{
	"Easy":   "🌞",
	"Medium": "⛅",
	"Hard":   "⛈️",
}

func Notification() {
	b := &btmpl.BasicTemplate{}

	for larkID, creds := range registry.Users {
		lcClient := leetcode.NewClient(creds)
		calendar, err := lcClient.GetUserCalendar()
		if err != nil {
			logrus.WithError(err).Error("leetcode client request error")
			return
		}
		buildCalendar(b, calendar, larkID)
	}
	lcClient := leetcode.NewClient("")
	question, err := lcClient.GetDailyQuestion()
	buildQuestion(b, question)
	tmpl, err := template.New("template").Parse(btmpl.Basic)
	var filledTmpl bytes.Buffer
	err = tmpl.Execute(&filledTmpl, b)
	if err != nil {
		logrus.WithError(err).Error("template execute error")
		return
	}
	err = lark.SendMsg(config.App.Lark.ChatID, "interactive", "chat_id", filledTmpl.String())
	if err != nil {
		logrus.WithError(err).Error("send message error")
	}
}
func buildCalendar(b *btmpl.BasicTemplate, calendar *leetcode.UserCalendarResp, larkID string) {
	uc := btmpl.UserCalendar{
		LarkID:          larkID,
		TotalActiveDays: calendar.Data.UserCalendar.TotalActiveDays,
		RecentStreak:    calendar.Data.UserCalendar.RecentStreak,
		LastSubmit:      parseLastSubmit(calendar.Data.UserCalendar.SubmissionCalendar),
	}
	b.UserCalendars = append(b.UserCalendars, uc)
}
func buildQuestion(b *btmpl.BasicTemplate, question *leetcode.DailyQuestionResp) {
	sentence := tianapi.GetSentence()
	b.Sentence = sentence[0]
	b.Source = sentence[1]
	q := question.Data.TodayRecord[0].Question
	b.Question.Title = q.QuestionID + ":" + q.TitleCn
	b.Question.Url = "https://leetcode.cn/problems/" + q.TitleSlug
	b.Question.AcRate = strconv.FormatFloat(q.AcRate*100, 'f', 2, 64) + "%"
	b.Question.Difficulty = q.Difficulty + " " + difficultyMap[q.Difficulty]
	var tagNames []string
	for _, tag := range q.TopicTags {
		tagNames = append(tagNames, tag.NameTranslated)
	}
	b.Question.Tags = strings.Join(tagNames, "、")
}

func parseLastSubmit(sc string) string {
	pattern := `[0-9]+`

	// 编译正则表达式
	reg := regexp.MustCompile(pattern)

	// 匹配字符串中的所有数字
	matches := reg.FindAllString(sc, -1)

	// 获取最后两个数字
	lsd := getLastSubmitDate(matches)
	ts, _ := strconv.Atoi(lsd[0])

	return time.Unix(int64(ts), 0).Format("2006/01/02")
}

func getLastSubmitDate(numbers []string) []string {
	length := len(numbers)
	if length < 2 {
		return numbers
	}
	return numbers[length-2 : length-1]
}
