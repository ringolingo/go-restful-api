package article

import "strings"

type Article struct {
	Id      int64    `json."Id"`
	Title   string `json:"Title"`
	Logline string `json:"Logline"`
	Content string `json:"Content"`
}

func FormatTitle(title string) string {
	return strings.ToUpper(title)
}

func FormatLogline(logline string) string {
	return strings.Title(logline)
}