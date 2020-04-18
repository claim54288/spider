package model

type Novel struct {
	BookId             int
	Author             string
	Name               string
	LetterNum          string
	RecommendedTotally int
	Clicked            int
	RecommendedWeekly  int
	Chapter            string //章
	TotalChapterNum    int
	Title              string //章标题
	Contents           string //章节内容
}
