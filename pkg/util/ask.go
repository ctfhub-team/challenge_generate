package util

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/exp/maps"
)

var errorMessages = map[string]string{
	"interrupt": "程序已退出 (Program exited.)",
}

func HandleErr(e error) {
	if e != nil {
		for k, v := range errorMessages {
			if strings.Contains(e.Error(), k) {
				fmt.Println(v)
				os.Exit(0)
			}
		}
	}
}

func SelectOne(Message string, Options map[string]string) string {
	var answer string
	var keys = maps.Keys(Options)
	sort.Strings(keys)
	prompt := &survey.Select{
		Message: Message,
		Options: keys,
	}
	err := survey.AskOne(prompt, &answer)
	HandleErr(err)
	return Options[answer]
}

func SelectArray(Message string, Options []string) string {
	var answer string
	prompt := &survey.Select{
		Message: Message,
		Options: Options,
	}
	err := survey.AskOne(prompt, &answer)
	HandleErr(err)
	return answer
}

func InputLine(Message string) string {
	var answer string
	prompt := &survey.Input{
		Message: Message,
	}
	err := survey.AskOne(prompt, &answer)
	HandleErr(err)
	return answer
}
