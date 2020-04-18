package persist

import (
	"fmt"
	"log"
	"os"
	"spider/model"
	"strings"
)

//存储
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out

			itemCount++
			switch itemJudged := item.(type) {
			case model.Novel:
				if itemJudged.Title != "" {
					saveNovelToTxt(itemJudged)
					log.Printf("Novel: %s Chapter: %s saved! ", itemJudged.Name, itemJudged.Title)
				} else {
					log.Printf("Item Saver :got item "+"#%d : %+v", itemCount, item)
				}
			}
		}
	}()
	return out
}

func saveNovelToTxt(m model.Novel) {
	filename := fmt.Sprintf("./books/%s_%s_%d.txt", m.Name, m.Author, m.BookId)

	file, err := os.OpenFile(filename, os.O_APPEND, 0644)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(filename)
	}
	if err != nil {
		log.Printf("can't create or open file %s,error: %s \n", filename, err.Error())
		return
	}
	defer file.Close()
	file.WriteString(" \n")
	file.WriteString("*********" + m.Title + "*********")
	content := m.Contents
	content = strings.ReplaceAll(content, "<p>", "")
	content = strings.ReplaceAll(content, "</p>", "\n")
	file.WriteString(content)
}
