package leetcode

// LcReqBody LeetCode请求题，在Query中传递GraphQL
type LcReqBody struct {
	Query         string      `json:"query"`
	Variables     interface{} `json:"variables"`
	OperationName string      `json:"operationName"`
}

// CalendarSubmitRecordResp 提交日历简要数据
type CalendarSubmitRecordResp struct {
	Data struct {
		CalendarSubmitRecord struct {
			DailySubmitStreakCount   int `json:"dailySubmitStreakCount"`
			DailyQuestionStreakCount int `json:"dailyQuestionStreakCount"`
			MonthlyAcCount           int `json:"monthlyAcCount"`
		} `json:"calendarSubmitRecord"`
	} `json:"data"`
}

// UserCalendarResp 个人资料里日历数据
type UserCalendarResp struct {
	Data struct {
		UserCalendar struct {
			Streak             int    `json:"streak"`
			TotalActiveDays    int    `json:"totalActiveDays"`
			SubmissionCalendar string `json:"submissionCalendar"`
			ActiveYears        []int  `json:"activeYears"`
			MonthlyMedals      any    `json:"monthlyMedals"`
			RecentStreak       int    `json:"recentStreak"`
		} `json:"userCalendar"`
	} `json:"data"`
}

// UserStatusResp 用户信息
type UserStatusResp struct {
	Data struct {
		UserStatus struct {
			IsSignedIn     bool   `json:"isSignedIn"`
			Username       string `json:"username"`
			RealName       string `json:"realName"`
			UserSlug       string `json:"userSlug"`
			CheckedInToday bool   `json:"checkedInToday"`
		} `json:"userStatus"`
		JobsMyCompany any `json:"jobsMyCompany"`
	} `json:"data"`
}

// DailyQuestionResp 每日一问
type DailyQuestionResp struct {
	Data struct {
		TodayRecord []struct {
			Question struct {
				QuestionID         string  `json:"questionId"`
				FrontendQuestionID string  `json:"frontendQuestionId"`
				Difficulty         string  `json:"difficulty"`
				TitleCn            string  `json:"titleCn"`
				TitleSlug          string  `json:"titleSlug"`
				AcRate             float64 `json:"acRate"`
				SolutionNum        int     `json:"solutionNum"`
				TopicTags          []struct {
					NameTranslated string `json:"nameTranslated"`
				} `json:"topicTags"`
			} `json:"question"`
		} `json:"todayRecord"`
	} `json:"data"`
}
