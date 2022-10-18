package main

import "strings"

const (
	LANG_ZH_WINDOW_TITLE = "COH 参谋长"
	LANG_EN_WINDOW_TITLE = "COH Chief of Staff"
)

var Lang map[string]string

func init() {
	Lang = map[string]string{
		"ZH_WINDOW_TITLE": LANG_ZH_WINDOW_TITLE,
		"EN_WINDOW_TITLE": LANG_EN_WINDOW_TITLE,
	}
}

func getString(lang, key string) string {
	if v, ok := Lang[strings.ToUpper(lang+"_"+key)]; ok {
		return v
	}
	return ""
}
