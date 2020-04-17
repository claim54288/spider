package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

func ParseProfile(contents *goquery.Document, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	index := 0
	contents.Find(".purple-btns div").Each(func(i int, s *goquery.Selection) {
		index++
		switch index {
		case 1:
			profile.Marriage = s.Text()
		case 2:
			age := s.Text()
			ageInt, _ := strconv.Atoi(age)
			profile.Age = ageInt
		case 3:
			profile.Xinzuo = s.Text()
		case 4:
			if strings.Contains(s.Text(), "kg") {
				profile.Weight, _ = strconv.Atoi(strings.TrimSuffix(s.Text(), "kg"))
				index++
			} else {
				profile.Height, _ = strconv.Atoi(strings.TrimSuffix(s.Text(), "cm"))
			}
		case 5:
			if strings.Contains(s.Text(), "工作地") {
				profile.WorkPlace = strings.TrimPrefix(s.Text(), "工作地:")
				index++
			} else {
				profile.Weight, _ = strconv.Atoi(strings.TrimSuffix(s.Text(), "kg"))
			}
		case 6:
			profile.WorkPlace = strings.TrimPrefix(s.Text(), "工作地:")
		case 7:
			profile.Income = strings.TrimPrefix(s.Text(), "月收入:")
		case 8:
			profile.Occupation = s.Text()
		case 9:
			profile.Education = s.Text()
		}
	})

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
