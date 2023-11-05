package leetcode

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const LcApi = "https://leetcode.cn/graphql/"
const LcApi2 = "https://leetcode.cn/graphql/noj-go/"

type LcClient struct {
	userSlug string
	creds    string
}

func NewClient(creds string) *LcClient {
	return &LcClient{
		creds: creds,
	}
}

func (l *LcClient) GetCalendarSubmitRecord() (*CalendarSubmitRecordResp, error) {
	body := &LcReqBody{
		Query:         "query calendarSubmitRecord {calendarSubmitRecord {dailySubmitStreakCount dailyQuestionStreakCount monthlyAcCount}}",
		Variables:     struct{}{},
		OperationName: "calendarSubmitRecord",
	}

	respBody, err := l.doRequest(LcApi, body)
	if err != nil {
		return nil, err
	}

	r, err := io.ReadAll(respBody)
	if err != nil {
		return nil, err
	}
	var csr CalendarSubmitRecordResp
	err = json.Unmarshal(r, &csr)
	if err != nil {
		return nil, err
	}

	return &csr, nil
}

func (l *LcClient) GetUserCalendar() (*UserCalendarResp, error) {
	userStatus, err := l.GetUserStatus()
	if err != nil {
		return nil, err
	}

	body := &LcReqBody{
		Query: "query userProfileCalendar($userSlug: String!, $year: Int) {userCalendar(userSlug: $userSlug, year: $year) {streak totalActiveDays submissionCalendar activeYears monthlyMedals {name obtainDate category config {icon iconGif iconGifBackground} progress id year month} recentStreak}}",
		//Query: "query calendarSubmitRecord {calendarSubmitRecord {dailySubmitStreakCount dailyQuestionStreakCount monthlyAcCount}}",
		Variables: struct {
			UserSlug string `json:"userSlug"`
		}{
			UserSlug: userStatus.Data.UserStatus.UserSlug,
		},
		OperationName: "userProfileCalendar",
	}
	respBody, err := l.doRequest(LcApi2, body)
	if err != nil {
		return nil, err
	}

	r, err := io.ReadAll(respBody)
	if err != nil {
		return nil, err
	}
	var uc UserCalendarResp
	err = json.Unmarshal(r, &uc)
	if err != nil {
		return nil, err
	}

	return &uc, nil
}

func (l *LcClient) GetUserStatus() (*UserStatusResp, error) {
	body := &LcReqBody{
		Query:         "query globalData {userStatus {isSignedIn checkedInToday useTranslation userSlug username realName}}",
		Variables:     struct{}{},
		OperationName: "globalData",
	}

	respBody, err := l.doRequest(LcApi2, body)
	if err != nil {
		return nil, err
	}

	r, err := io.ReadAll(respBody)
	if err != nil {
		return nil, err
	}
	var us UserStatusResp
	err = json.Unmarshal(r, &us)
	if err != nil {
		return nil, err
	}

	return &us, nil
}

func (l *LcClient) GetDailyQuestion() (*DailyQuestionResp, error) {
	body := &LcReqBody{
		Query:         "query questionOfToday {todayRecord {question {questionId frontendQuestionId: questionFrontendId difficulty titleCn: translatedTitle titleSlug acRate solutionNum topicTags {nameTranslated: translatedName}}}}",
		Variables:     struct{}{},
		OperationName: "questionOfToday",
	}

	respBody, err := l.doRequest(LcApi, body)
	if err != nil {
		return nil, err
	}

	r, err := io.ReadAll(respBody)
	if err != nil {
		return nil, err
	}
	var dq DailyQuestionResp
	err = json.Unmarshal(r, &dq)
	if err != nil {
		return nil, err
	}

	return &dq, nil
}
func (l *LcClient) doRequest(api string, body *LcReqBody) (io.ReadCloser, error) {
	b, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", l.creds)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unknown err")
	}

	return resp.Body, nil
}
