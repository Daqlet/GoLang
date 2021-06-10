package main

type Submission struct {
	Id                  int     `json:"id"`
	ContestId           int32   `json:"contestId"`
	Problem             Problem `json:"problem"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict"`
}

type Problem struct {
	ContestId int    `json:"constestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
}

type StatusResponse struct {
	Status      string       `json:"status"`
	Submissions []Submission `json:"result"`
	Comment     string       `json:"comment"`
}
