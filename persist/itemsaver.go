package persist

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"spider/engine"
	"spider/model"
	"strings"
)

//存储
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.159.131:9200"),
		elastic.SetSniff(false), //维护集群状态
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out

			itemCount++
			switch itemJudged := item.Payload.(type) {
			case model.Novel:

				if itemJudged.Title != "" {
					//saveNovelToTxt(itemJudged)
					//log.Printf("Novel: %s Chapter: %s saved! ", itemJudged.Name, itemJudged.Title)
					//if err != nil {
					//	log.Printf("Item Saver:error saving item %v:%v", item, err)
					//}
				} else {
					err := SaveToElastic(client, index, item)
					if err != nil {
						log.Printf("Item Saver:error saving item %v:%v", item, err)
					}
					log.Printf("Item Saver :got item #%d : %+v", itemCount, item)
				}
			case model.Student:
				err := SaveToElastic(client, index, item)
				if err != nil {
					log.Printf("Item Saver:error saving item %v:%v", item, err)
				} else {
					log.Printf("Item Saver:item %v saved", item)
				}
			default:
				log.Printf("Item Saver :got item #%d : %+v", itemCount, item)
			}
		}
	}()
	return out, nil
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

func SaveToElastic(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	if item.Type == "zongheng" {
		indexService := client.Index().Index(index).Type(item.Type)
		if item.Id != "" {
			indexService.Id(item.Id)
		}
		_, err := indexService.BodyJson(item).Do(context.Background())
		return err
	}
	return nil
}
