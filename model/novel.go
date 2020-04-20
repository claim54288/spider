package model

import "encoding/json"

type Novel struct {
	BookId             int
	Author             string
	Name               string
	LetterNum          int64
	RecommendedTotally int64
	Clicked            int64
	RecommendedWeekly  int64
	Chapter            string //章
	TotalChapterNum    int
	Title              string //章标题
	Contents           string //章节内容
}

func FromJsonObj(o interface{}) (Novel, error) {
	var novel Novel
	s, err := json.Marshal(o)
	if err != nil {
		return novel, err
	}
	err = json.Unmarshal(s, &novel)
	return novel, err
}
