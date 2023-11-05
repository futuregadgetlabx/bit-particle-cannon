package template

const Basic = `
{
  "elements": [
    {
      "tag": "markdown",
      "content": "***{{ .Sentence }} —— {{ .Source }}***"
    },
    {
      "tag": "hr"
    },
    {
      "tag": "markdown",
      "content": "<at id=all></at>\n🤖比特粒子炮提醒您，别忘记今天的算法练习。每日一题仅供参考，酌情挑战。"
    },
    {
      "tag": "div",
      "fields": [
        {
          "is_short": true,
          "text": {
            "tag": "lark_md",
            "content": "**📖 每日一题：**[{{ .Question.Title }}]({{ .Question.Url }})\n**⚠️ 题目难度：**{{ .Question.Difficulty }}\n**🚩 通过率：** {{ .Question.AcRate }}\n**📃 题目标签：**{{ .Question.Tags }}"
          }
        }
      ]
    },
    {
      "tag": "div",
      "text": {
        "content": "🎆 **记录** ",
        "tag": "lark_md"
      }
    },
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "<at id={{ .UserId }}></at> 🎈累计提交天数:{{ .UserCalendar.TotalActiveDays }} 🚀连续提交:{{ .UserCalendar.RecentStreak }} 📅上次提交: {{ .UserCalendar.LastSubmit }}",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "📄 如何添加LeetCode账号到机器人"
          },
          "type": "primary",
          "multi_url": {
            "url": "https://steinsgate.feishu.cn/docx/MpdzdpqxPoOkLYxykI0cPZbxn6c",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "tag": "note",
      "elements": [
        {
          "tag": "plain_text",
          "content": "来自未来ガジェット研究所 - 比特粒子炮"
        }
      ]
    }
  ],
  "header": {
    "template": "blue",
    "title": {
      "content": "📬力扣每日提醒",
      "tag": "plain_text"
    }
  }
}`

type BasicTemplate struct {
	UserId       string
	Sentence     string
	Source       string
	Question     Question
	UserCalendar UserCalendar
}

type Question struct {
	Title      string
	Url        string
	AcRate     string
	Difficulty string
	Tags       string
}

type UserCalendar struct {
	TotalActiveDays int
	RecentStreak    int
	LastSubmit      string
}

const LcNotifySuccess = `{
  "elements": [
    {
      "tag": "markdown",
      "content": "LeetCode用户{{ .Username }}添加成功"
    }
  ],
  "header": {
    "template": "green",
    "title": {
      "content": "🎉 添加成功",
      "tag": "plain_text"
    }
  }
}`

type LcNotify struct {
	Username string
}
