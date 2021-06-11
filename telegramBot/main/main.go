package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var lastContestId int32
var bot, err_ = tgbotapi.NewBotAPI("MyApi")
var SolutionCount = make(map[string]int)

func main() {
	if err_ != nil {
		panic(err_)
	}
	bot.Debug = true
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60
	updates, err := bot.GetUpdatesChan(upd)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		err = respond(update.Message.Text, update)
		if err != nil {
			panic(err)
		}
	}
}

func respond(handle string, update tgbotapi.Update) error {
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
	return nil
}

func AddSubmission(submission Submission, chatId int64) {
	TYPE := "contest"
	if submission.ContestId > lastContestId {
		TYPE = "gym"
	}
	url := "https://codeforces.com/" + TYPE + "/" +
		strconv.Itoa(int(submission.ContestId)) +
		"/submission/" + strconv.Itoa(int(submission.Id))
	ParseAndGetCode(url, submission, chatId)
}

func ParseAndGetCode(url string, submission Submission, chatId int64) {
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
		problem += goodLetter(rune(submission.Problem.Name[i]))
	}
	if SolutionCount[problem] > 0 {
		problem += "_version_" + strconv.Itoa(SolutionCount[problem]+1)
	}
	SolutionCount[problem] += 1
	return problem
}

func goodLetter(letter rune) string {
	if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
		return string(letter)
	}
	return "_"
}

func GetExtension(Language string) string {
	switch Language {
	case "GNU C++17", "GNU C++14", "GNU C++11", "GNU C++", "MS C++ 2017", "MS C++", "GNU C++17 (64)":
		{
			return ".cpp"
		}
	case "GNU C11":
		{
			return ".c"
		}
	case "Node.js", "JavaScript":
		{
			return ".js"
		}
	case "Scala":
		{
			return ".scala"
		}
	case "Rust":
		{
			return ".rs"
		}
	case "Ruby 3":
		{
			return ".rb"
		}
	case "PyPy 3", "PyPy 2", "Python 3 + libs", "Python 3", "Python 2":
		{
			return ".py"
		}
	case "PHP":
		{
			return ".php"
		}
	case "Perl":
		{
			return ".pl"
		}
	case "PascalABC.NET", "FPC", "Delphi":
		{
			return ".pas"
		}
	case "Ocaml":
		{
			return ".ml"
		}
	case "Kotlin":
		{
			return ".kt"
		}
	case "Java 8", "Java 11":
		{
			return ".java"
		}
	case "Haskell":
		{
			return ".hs"
		}
	case "Go":
		{
			return ".go"
		}
	case "D":
		{
			return ".d"
		}
	case "Mono C#", ".NET Core C#":
		{
			return ".cs"
		}

	default:
		{
			return ".txt"
		}
	}
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
