package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var client = &http.Client{}

func main() {
	ids, err := parseIds("babichev.csv")
	fmt.Println(ids)
	if err != nil {
		fmt.Println(err)
		return
	}
	topics, err := readTopics(ids)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ids, len(ids))

	err = writeTopics("babichev.json", topics)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
}

type Post struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	UpdationDate int    `json:"updationDate"`
	CreationDate int    `json:"creationDate"`
}
type Topic struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	TaskURL string `json:"taskUrl"`
	Post    Post   `json:"post"`
}
type Data struct {
	Topic Topic `json:"topic"`
}
type Result struct {
	Data Data `json:"data"`
}
type TaskID struct {
	taskLink string
	id       int
}

func writeTopics(filename string, topics []Topic) error {
	bytes, err := json.Marshal(topics)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, bytes, 0644)
	return err
}

func readTopics(ids []TaskID) (topics []Topic, err error) {
	var topic *Topic
	for _, id := range ids {
		topic, err = readTopic(id.id)
		topic.TaskURL = id.taskLink
		if err != nil {
			return
		}
		topics = append(topics, *topic)
	}
	return
}

func parseIds(filename string) (ids []TaskID, err error) {
	var seq int
	var url string
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	re := regexp.MustCompile(`/\d+/`)
	reURL := regexp.MustCompile(`https://leetcode.com/problems/[^/]+`)

	for n, err := fmt.Fscan(f, &seq, &url); n != 0 && err == nil; n, err = fmt.Fscan(f, &seq, &url) {
		var taskID TaskID
		urlBytes := []byte(url)
		idRes := string(re.Find(urlBytes))
		taskID.id, err = strconv.Atoi(strings.Trim(idRes, "/"))
		taskID.taskLink = string(reURL.Find(urlBytes))
		ids = append(ids, taskID)
	}

	return
}

func readTopic(id int) (topic *Topic, err error) {
	var data Result
	cont, err := read(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(cont, &data)
	if err != nil {
		return
	}
	return &data.Data.Topic, nil
}

func read(id int) (content []byte, err error) {
	url := "https://leetcode.com/graphql"
	method := "POST"

	query := fmt.Sprintf(`{"operationName":"DiscussTopic","variables":{"topicId":%v},"query":"query DiscussTopic($topicId: Int!) {\n  topic(id: $topicId) {\n    id\n    title\n   post {\n      ...DiscussPost\n     }\n      }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  content\n  updationDate\n  creationDate\n   author {\n   username\n     }\n   }\n"}`, id)
	payload := strings.NewReader(query)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\"")
	req.Header.Add("accept", "*/*")
	req.Header.Add("X-NewRelic-ID", "UAQDVFVRGwEAXVlbBAg=")
	req.Header.Add("x-csrftoken", "7oEOdbDWkx27S6lB2kOkMLtI8F1xKYWAspX0y5mSqnuSL5YLdKoHEKYtjxkfH5Kw")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Cookie", "__cf_bm=5eceb9f0d8f9a4a277082586322ae39348641f5e-1617196161-1800-AW6CEdXadtEik4XNTUsLW2ArnhKaAUthhuY3YOVDdhNVpYVDdCT+SqW/aCaZv97PSFD71B3mo1GmHnCtDeqZ8KQ=; __cfduid=dc6299f72540ee04ac2d15004383f03411617177420; csrftoken=tlJEiXgBwEes5qpVtxTOH8nALwQmYj9p5LLYBSlzsxwIMdvr0a6W0oyqRarALXIw")
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return body, err
}
