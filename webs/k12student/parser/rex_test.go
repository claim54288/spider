package parser

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {
	re := "student_id=([a-zA-Z0-9]*)="
	testStr := "http://test.b.kai12.cn/archives/archives/old_registration_card?student_id=zGDY9R7KQ6w=&term_id=2016&semester=2&grade_id=101&student_name=%E4%B9%94%E9%95%BF%E4%BC%9F&class_id=13&class_name=%E4%B8%80%E5%B9%B4%E7%BA%A7(1)%E7%8F%AD&student_edit=1"
	recomp := regexp.MustCompile(re)
	s := recomp.FindStringSubmatch(testStr)
	fmt.Printf("%s\n", s[1])
}
