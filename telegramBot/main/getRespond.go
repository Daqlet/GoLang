package main

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
)

func GetRespond(handle string, update tgbotapi.Update) error {
	SetLastContestID()
	Client := http.Client{}
	Response, err := Client.Get("https://codeforces.com/api/user.status?handle=" + handle + "&from=1&count=20")
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		return err
	}
	response := StatusResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if response.Status == "FAILED" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.Comment)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := bot.Send(msg)
		if err != nil {
			panic(err)
		}
		return nil
	}
	for _, submission := range response.Submissions {
		if submission.Verdict == "OK" {
			AddSubmission(submission, update.Message.Chat.ID)
		}
	}
	SaveFilesInZip()
	return nil
}
