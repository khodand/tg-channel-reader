package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/khodand/tg-channel-reader/pkg/telegram"
)

var (
	malePatterns   = []string{"парн", "мужч", "парен", "мальч", "актер ", "актёр ", "мужск", "юнош"}
	femalePatterns = []string{"деву", "женщ", "актри", "дево", "жен ", "женс"}
	files          = []string{
		"castings.json",
		"castingspb.json",
		"marketplace.json",
		"primepeople.json",
		"result.json",
		"result1.json",
		"result2.json",
		"result3.json",
		"result4.json",
	}
)

func main() {
	var data Data
	for i := range files {
		jsonFile, err := os.Open("data/" + files[i])
		if err != nil {
			panic(err)
		}

		byteValue, _ := io.ReadAll(jsonFile)
		var telegramData telegram.Data
		if err := json.Unmarshal(byteValue, &telegramData); err != nil {
			panic(err)
		}
		_ = jsonFile.Close()

		messages := make([]telegram.Message, 0, len(telegramData.Messages))
		for j := range telegramData.Messages {
			messages = append(messages, *telegramData.Messages[j].ToMessage())
			messages[j].Text = strings.ToLower(messages[j].Text)
		}
		data.Messages = append(data.Messages, messages...)
	}

	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(file)
	w.Write([]string{"label", "full_text", "hashtag"})

	both, male, female, none := 0, 0, 0, 0
	for i := range data.Messages {
		text := data.Messages[i].Text
		label := getLabel(text)
		switch label {
		case "0":
			both++
		case "1":
			male++
		case "2":
			female++
		case "3":
			none++
		}
		w.Write([]string{label, text, strings.Join(data.Messages[i].Hashtags, " ")})
	}
	w.Flush()
	fmt.Printf("both: %d, male: %d, female: %d, none: %d \n", both, male, female, none)

	//fmt.Printf("len %d\n", len(data.Messages))
	//bytes, err := json.Marshal(data)
	//if err != nil {
	//	panic(err)
	//}
	//dataJson, err := os.Create("data.json")
	//if err != nil {
	//	return
	//}
	//_, _ = dataJson.Write(bytes)
	//if err != nil {
	//	return
	//}
	//dataJson.Close()
}

type Data struct {
	Messages []telegram.Message `json:"messages"`
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
