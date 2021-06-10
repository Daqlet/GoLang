package main

type Update struct {
	Update_id int     `json:"update_id"`
	Message   Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	Id int `json:"id"`
}

type Respond struct {
	Update []Update `json:"result"`
}

type BotMessage struct {
	ChatID   int      `json:"chat_id"`
	Text     string   `json:"text"`
	Document Document `json:"document"`
}

type Document struct {
	File_id   string `json:"file_id"`
	File_Name string `json:"file_name"`
}

type BotCommand struct {
	Command string `json:"command"`
}
