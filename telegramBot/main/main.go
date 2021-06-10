package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var lastContestId int = 10000
var botUrl string

func main() {
	bot, err := tgbotapi.NewBotAPI("1648899629:AAEpJWaRSxBBL0IoW5yJQ4_0uoHr7Qg1cyg")
	if err != nil {
		panic(err)
	}
	fmt.Print(bot)
	botToken := "1648899629:AAEpJWaRSxBBL0IoW5yJQ4_0uoHr7Qg1cyg"
	botApi := "https://api.telegram.org/bot"
	botUrl = botApi + botToken
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			panic(err)
		}
		fmt.Println(updates)
		for _, update := range updates {
			err = respond(botUrl, update)
			if err != nil {
				panic(err)
			}
			offset = update.Update_id + 1
		}
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	source, err := http.Get(botUrl + "/getUpdates?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer source.Body.Close()
	body, err_ := ioutil.ReadAll(source.Body)
	if err_ != nil {
		return nil, err_
	}
	result := Respond{}
	err_ = json.Unmarshal(body, &result)
	if err_ != nil {
		return nil, err_
	}
	return result.Update, err_
}

func respond(botURL string, update Update) error {
	handle := update.Message.Text
	err := findHandle(handle, update.Message.Chat.Id)
	if err != nil {
		message := BotMessage{}
		message.ChatID = update.Message.Chat.Id
		message.Text = "Incorrect handle"
		buffer, err := json.Marshal(message)
		if err != nil {
			return err
		}
		_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buffer))
		return err
	}
	return err
}

func findHandle(handle string, chatId int) error {
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
	for _, submission := range response.Submissions {
		if submission.Verdict == "OK" {
			AddSubmission(submission, chatId)
		}
	}
	return nil
}

func AddSubmission(submission Submission, chatId int) {
	TYPE := "contest"
	if submission.ContestId > lastContestId {
		TYPE = "gym"
	}
	url := "https://codeforces.com/" + TYPE + "/" +
		strconv.Itoa(int(submission.ContestId)) +
		"/submission/" + strconv.Itoa(int(submission.Id))
	ParseAndGetCode(url, submission, chatId)
}

func ParseAndGetCode(url string, submission Submission, chatId int) {
	Code := GetCode(url)
	path := "Codeforces/Contest" + strconv.Itoa(int(submission.ContestId))
	CreateFolder("Codeforces")
	CreateFolder(path)
	problemName := ParseProblemName(submission) + ".txt"
	path += "/" + problemName
	err := os.WriteFile(path, []byte(Code), 0644)
	if err != nil {
		panic(err.Error())
	}

	file, _ := os.Open(path)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", botUrl+"/SendDocument", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Do(r)

	/*

		message := BotMessage{}
		message.ChatID = chatId
		message.Text = string(file)
		buf, err := json.Marshal(message)
		if err != nil {
			panic(err)
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			panic(err)
		}
		/*
			message := BotMessage{}
			message.ChatID = ChatID
			message.Document =
			buffer, err := json.Marshal(message)
			if err != nil {
				return err
			}
			_, err = http.Post(botURL+"/sendDocument", "application/json", bytes.NewBuffer(buffer))
	*/
}

func GetCode(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	if response.Request.URL.String() != url {
		return "URL: " + url + " redirecting to the main page. " +
			"Couldn`t find a code\n" +
			"Probably, the submission it private!"
	}
	body, err := ioutil.ReadAll(response.Body)
	code := string(body)
	startPosition := strings.Index(code, "program-source-text")
	for code[startPosition] != '>' {
		startPosition += 1
	}
	startPosition += 1
	code = code[startPosition:]
	endPosition := strings.Index(code, "</pre>")
	code = code[:endPosition]
	code = ReplaceSymbols(code)
	return code
}

func ReplaceSymbols(code string) string {
	Old := []string{"&quot;", "&gt;", "&lt;", "&#39;", "&amp;"}
	New := []string{"\"", ">", "<", "'", "&"}
	for i := 0; i < 5; i++ {
		code = strings.ReplaceAll(code, Old[i], New[i])
	}
	return code
}

func CreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0700)
		if err != nil {
			panic(err.Error())
		}
	}
}

func ParseProblemName(submission Submission) string {
	problem := submission.Problem.Index + "."
	for i := 0; i < len(submission.Problem.Name); i++ {
		problem += goodLetter(rune(submission.Problem.Name[i]))
	}
	return problem
}

func goodLetter(letter rune) string {
	if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
		return string(letter)
	}
	return "_"
}
