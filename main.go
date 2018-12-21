package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


type slackMessage struct {
	Channel string `json:"channel"`
	Text string		`json:"text"`
	Attachments []slackAttachment `json:"attachments"`
}

type slackAttachment struct {
	Text string `json:"text"`
	Color string `json:"color"`
}


func main(){

	b, err := ioutil.ReadFile("slack-token.txt")
	if err != nil{
		panic(err)
	}

	a := slackAttachment{
		Text: "test attachment",
		Color: "#800000",
	}
	list := []slackAttachment{}
	list = append(list, a)
	message := slackMessage{
		Channel: "gitlab",
		Text: "test message",
		Attachments: list,
	}

	attach, err := json.Marshal(message)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(attach))
	err1 := postForm("https://slack.com/api/chat.postMessage", attach, strings.TrimSpace(string(b)))
	if err1 != nil{
		panic(err1)
	}

}

func postForm(url string, values []byte, token string) error {
	reqBody := strings.NewReader(string(values))
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return err
	}

	slackToken := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", slackToken)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
		return err
	}
	fmt.Println(res.StatusCode)

	return err
}