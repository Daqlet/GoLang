package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var lastContestId int32
var SolutionCount = make(map[string]int)

func AddSubmission(submission Submission, chatId int64) {
	TYPE := "contest"
	if submission.ContestId > lastContestId {
		TYPE = "gym"
	}
	url := "https://codeforces.com/" + TYPE + "/" +
		strconv.Itoa(int(submission.ContestId)) +
		"/submission/" + strconv.Itoa(int(submission.Id))
	ParseAndSendCode(url, submission, chatId)
}

func ParseAndSendCode(url string, submission Submission, chatId int64) {
	Code := GetCode(url)
	path := "Codeforces/Contest" + strconv.Itoa(int(submission.ContestId))
	CreateFolder("Codeforces")
	CreateFolder(path)
	problemName := ParseProblemName(submission) + GetExtension(submission.ProgrammingLanguage)
	path += "/" + problemName
	err := os.WriteFile(path, []byte(Code), 0644)
	if err != nil {
		panic(err.Error())
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileBytes := tgbotapi.FileBytes{
		Name:  problemName,
		Bytes: file,
	}
	_, err = bot.Send(tgbotapi.NewDocumentUpload(int64(chatId), fileBytes))
	if err != nil {
		panic(err)
	}
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
		problem += GoodLetter(rune(submission.Problem.Name[i]))
	}
	if SolutionCount[problem] > 0 {
		problem += "_version_" + strconv.Itoa(SolutionCount[problem]+1)
	}
	SolutionCount[problem] += 1
	return problem
}

func GoodLetter(letter rune) string {
	if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
		return string(letter)
	}
	return "_"
}

func SetLastContestID() {
	Response, ok := http.Get("https://codeforces.com/problemset")
	if ok != nil {
		panic(ok.Error())
	}
	AsBytes, ok := ioutil.ReadAll(Response.Body)
	if ok != nil {
		panic(ok.Error())
	}
	AsString := string(AsBytes)
	keyword := "/problemset/problem/"
	Id := strings.Index(AsString, keyword)
	if Id == -1 {
		panic("Can not find ID of last contest")
	}
	lastContestId = 0
	for i := 0; i < 4; i += 1 {
		lastContestId = lastContestId*10 + rune(AsString[len(keyword)+Id+i]) - '0'
	}
}
