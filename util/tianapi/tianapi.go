package tianapi

import (
	"encoding/json"
	"github.com/futuregadgetlabx/bit-particle-cannon/config"
	"io"
	"net/http"
)

type SentenceResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Saying string `json:"saying"`
		Transl string `json:"transl"`
		Source string `json:"source"`
	} `json:"result"`
}

func GetSentence() []string {
	resp, err := http.Get("https://apis.tianapi.com/lzmy/index?key=" + config.App.TianApi.Key)
	r, err := io.ReadAll(resp.Body)
	var sr SentenceResp
	err = json.Unmarshal(r, &sr)
	if err != nil {
		return []string{"书山有路勤为径，学海无涯苦作舟。", "增广贤文"}
	}
	return []string{sr.Result.Saying, sr.Result.Source}
}
