package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider/engine"
	"spider/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	var err error
	err = persist.SaveToElastic(s.Client, s.Index, item)
	log.Printf("Item %v saved.\n", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %+v:%v", item, err)
	}
	return err
}
