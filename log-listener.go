package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/hpcloud/tail"
)

func logListenerStart(logPath string, keywords []string, handler func(string, string)) {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		log.Println(err)
		return
	}
	go listenWarningsLogLine(logPath, keywords, handler)
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
