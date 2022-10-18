package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/hpcloud/tail"
)

func parseLogPath(logPath string) string {
	parsedlogPath := logPath
	sysVarList := []string{"USERPROFILE"}
	for _, v := range sysVarList {
		_v := "%" + v + "%"
		if strings.Contains(parsedlogPath, _v) {
			parsedlogPath = strings.ReplaceAll(parsedlogPath, _v, os.Getenv(v))
		}
	}
	return parsedlogPath
}

func logListenerStart(logPath string, keywords []string, handler func(string, string)) {
	parsedlogPath := parseLogPath(logPath)
	if _, err := os.Stat(parsedlogPath); os.IsNotExist(err) {
		log.Println(err)
		return
	}
	go listenWarningsLogLine(parsedlogPath, keywords, handler)
}

func listenWarningsLogLine(logPath string, keywords []string, handler func(string, string)) {
	t, _ := tail.TailFile(logPath, tail.Config{Follow: true, Poll: true})
	for line := range t.Lines {
		for _, keyword := range keywords {
			if strings.Contains(line.Text, keyword) {

				file, _ := os.Open(logPath)
				reader := bufio.NewReader(file)
				bts, _, _ := reader.ReadLine()
				file.Close()
				tmp := strings.Split(string(bts), " ")
				logCreatedAt := tmp[len(tmp)-2]

				handler(strings.TrimSpace(logCreatedAt), strings.TrimSpace(line.Text))
			}
		}
	}
}
