package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/engine"
	"spider/model"
)

func ParseStudentInfo(contents *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}
	student := model.Student{}
	item := engine.Item{}
	url := contents.Url.String()
	item.Url = url
	item.Id = getStudentId(regStudentId, url)
	item.Type = "student_info"
	contents.Find(".registration-card-table .tdcon").EachWithBreak(func(i int, s *goquery.Selection) bool {
		switch i {
		case 0:
			name := s.Text()
			if name == "" { //名字都没，信息有问题
				return false
			}
			student.Name = name
		case 1:
			student.Sex = s.Text()
		case 2:
			student.Birth = s.Text()
		case 3:
			student.Nation = s.Text()
		case 4:
			student.NativePlace = s.Text()
		case 5:
			student.IdCard = s.Text()
		case 6:
			student.MemberAppellation1 = s.Text()
		case 7:
			student.MemberName1 = s.Text()
		case 8:
			student.MemberJob1 = s.Text()
		case 9:
			student.MemberPhone1 = s.Text()
		case 10:
			student.MemberAppellation2 = s.Text()
		case 11:
			student.MemberName2 = s.Text()
		case 12:
			student.MemberJob2 = s.Text()
		case 13:
			student.MemberPhone2 = s.Text()
		case 14:
			student.MemberAppellation3 = s.Text()
		case 15:
			student.MemberName3 = s.Text()
		case 16:
			student.MemberJob3 = s.Text()
		case 17:
			student.MemberPhone3 = s.Text()
		case 18:
			student.Address = s.Text()
		case 19:
			student.Phone = s.Text()
		case 20:
			student.Police = s.Text()
		case 21:
			student.Graduate = s.Text()
		default:
			return true
		}
		return true
	})
	item.Payload = student
	result.Items = []engine.Item{item}
	return result
}
