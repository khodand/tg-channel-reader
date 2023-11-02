package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/khodand/tg-channel-reader/pkg/telegram"
)

var (
	malePatterns   = []string{"парн", "мужч", "парен", "мальч", "актер ", "актёр ", "мужск", "юнош"}
	femalePatterns = []string{"деву", "женщ", "актри", "дево", "жен ", "женс"}
	file           = "data.json"
)

type Data struct {
	Messages []telegram.Message `json:"messages"`
}

func main() {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)
	var telegramData Data
	if err := json.Unmarshal(byteValue, &telegramData); err != nil {
		panic(err)
	}
	_ = jsonFile.Close()

	dict := make(map[string]int64, 1000)

	for i := range telegramData.Messages {
		text := strings.ReplaceAll(telegramData.Messages[i].Text, "\n", " ")
		text = strings.ReplaceAll(text, ",", "")
		text = strings.ReplaceAll(text, ".", "")
		text = strings.ReplaceAll(text, "!", "")
		text = strings.ReplaceAll(text, "-", " ")
		text = strings.ReplaceAll(text, "•", " ")
		text = strings.ReplaceAll(text, "+", " ")
		text = strings.ReplaceAll(text, "”", " ")
		text = strings.ReplaceAll(text, "(", " ")
		text = strings.ReplaceAll(text, ")", " ")
		text = strings.ReplaceAll(text, "*", " ")
		text = strings.ReplaceAll(text, ":", " ")
		text = strings.ReplaceAll(text, "»", " ")
		text = strings.ReplaceAll(text, "«", " ")
		text = strings.ReplaceAll(text, ";", " ")
		text = strings.ReplaceAll(text, "\"", " ")
		text = strings.ReplaceAll(text, "'", " ")
		text = strings.ReplaceAll(text, "⃣", " ")
		text = strings.ReplaceAll(text, "	", " ")
		text = strings.ReplaceAll(text, "0", " ")
		text = strings.ReplaceAll(text, "1", " ")
		text = strings.ReplaceAll(text, "2", " ")
		text = strings.ReplaceAll(text, "3", " ")
		text = strings.ReplaceAll(text, "4", " ")
		text = strings.ReplaceAll(text, "5", " ")
		text = strings.ReplaceAll(text, "6", " ")
		text = strings.ReplaceAll(text, "7", " ")
		text = strings.ReplaceAll(text, "8", " ")
		text = strings.ReplaceAll(text, "9", " ")
		text = strings.TrimSpace(text)

		words := strings.Split(text, " ")
		for j := range words {
			switch words[j] {
			case "":
				continue
			case "в":
				continue
			case "и":
				continue
			case "на":
				continue
			case "с":
				continue
			case "не":
				continue
			case "для":
				continue
			}
			dict[words[j]]++
		}
	}

	type wS struct {
		text  string
		count int64
	}

	words := make([]wS, 0, len(dict))
	for k, v := range dict {
		words = append(words, wS{
			text:  k,
			count: v,
		})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})

	fmt.Println(len(words))
	//for i := 0; i < 20; i++ {
	//	fmt.Println(words[i].text, words[i].count)
	//}
	for i := range words {
		w := words[i]
		//if words[i].count < 10 {
		//	break
		//}
		lb := getLabel(w.text)
		if lb == "2" {
			//continue
			fmt.Println(lb, w.text, w.count)
		}
	}
}

func getLabel(text string) string {
	text = strings.ToLower(text)

	male, female := false, false
	for i := range malePatterns {
		if strings.Contains(text, malePatterns[i]) {
			male = true
			break
		}
	}

	for i := range femalePatterns {
		if strings.Contains(text, femalePatterns[i]) {
			female = true
			break
		}
	}

	switch {
	case male && female:
		return "0"
	case male:
		return "1"
	case female:
		return "2"
	default:
		return "3"
	}
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
