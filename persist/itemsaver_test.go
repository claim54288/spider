package persist

//func TestSave(t *testing.T) {
//	expected := engine.Item{
//		Url:  "http://book.zongheng.com/book/849617.html",
//		Type: "zongheng",
//		Id:   "849617",
//		Payload: model.Novel{
//			BookId:             849617,
//			Author:             "关中老人",
//			Name:               "一脉承腔",
//			LetterNum:          "30.1万",
//			RecommendedTotally: 66,
//			Clicked:            0,
//			RecommendedWeekly:  0,
//			Chapter:            "",

//			TotalChapterNum:    0,
//			Title:              "",
//			Contents:           "",
//		},
//	}
//	err := saveToElastic(expected)
//	if err != nil {
//		panic(err)
//	}
//
//	client, err := elastic.NewClient(
//		elastic.SetURL("http://192.168.159.131:9200"),
//		elastic.SetSniff(false),
//	)
//	if err != nil {
//		panic(err)
//	}
//	resp, err := client.Get().Index("dating_novel").Type(expected.Type).Id(expected.Id).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	//t.Logf("%s", resp.Source)
//
//	var actual engine.Item
//	json.Unmarshal(*resp.Source, &actual)
//	actualNovel, _ := model.FromJsonObj(actual.Payload)
//	actual.Payload = actualNovel
//
//	if actual != expected {
//		t.Errorf("get %+v;expected %+v", actual, expected)
//	}
//}
