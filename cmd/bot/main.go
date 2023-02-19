package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"githab.com/telegrammbot/bot/internal/modules"
	// "githab.com/internal/models/internal/models"
)

// var some_var Message

// some_var.text = "ewdfsgsdfg"
// fmt.Println(some_var)

//////////////////////////////////////////////////////////////////////

// https://api.telegram.org/bot6161865465:AAHLcBcpSRWVDxIe1bK6R2l-feml36wHOAU/getMe
// {
//     "ok": true,
//     "result": {
//         "id": 6161865465,
//         "is_bot": true,
//         "first_name": "go_bot",
//         "username": "go_special_bot",
//         "can_join_groups": true,
//         "can_read_all_group_messages": false,
//         "supports_inline_queries": false
//     }
// }

//////////////////////////////////////////////////////////////////////
// ПРИМЕР НАШЕГО СООБЩЕНИЯ, ЗНАЯ СООБЩЕНИЯ МЫ СМОЖЕМ ЕГО РАСПАРСИТЬ

// resp.Header
// map[
// 	Access-Control-Allow-Methods:[GET, POST, OPTIONS]
// 	Access-Control-Allow-Origin:[*]
// 	Access-Control-Expose-Headers:[Content-Length,Content-Type,Date,Server,Connection]
// 	Content-Length:[23]
// 	Content-Type:[application/json]
// 	Date:[Sun, 19 Feb 2023 05:25:19 GMT]
// 	Server:[nginx/1.18.0]
// 	Strict-Transport-Security:[max-age=31536000; includeSubDomains; preload]]
// resp.Body {0xc0001fc480}

// Body
// 	{"ok":true,
// 	"result":[
// 		{"update_id":54696872,
// 		"message":
// 			{"message_id":147,
// 			"from":
// 				{"id":1700335247,
// 				"is_bot":false,
// 				"first_name":"Roman",
// 				"username":"BlackGooseOnDinner",
// 				"language_code":"en
// 				"},
// 			"chat":
// 				{"id":1700335247,
// 				"first_name":"Roman",
// 				"username":"BlackGooseOnDinner",
// 				"type":"private"
// 				},
// 			"date":1676784635,
// 			"text":"g"
// 			}
// 		}
// 			]
// 		}

// ////////////////////////////////////////////////////////////////////

// type Chat struct {
// 	ChatId int `json:"id"`
// }

// type Message struct {
// 	Chat Chat   `json:"chat"`
// 	Text string `json:"text"`
// }

// type Update struct {
// 	UpdateId int     `json:"update_id"`
// 	Message  Message `json:"message"`
// }

// type RestResponce struct {
// 	Result []Update `json:"result"`
// }

// type BotMessage struct {
// 	ChatId int    `json:"chat_id"`
// 	Text   string `json:"text"`
// }

//////////////////////////////////////////////////////////////////////

func main() {

	// алгоритм
	/*
		1 - с помощью GET запроса мы получаем свежие сообщения
			a получаем данные
			b читаем тело (заголовок читать не надо)
			c декодируем сообщение и парсим
			d
			i
		2 - с помощью POST отправляем данные

	*/

	// токен необходимо хранить в секрете
	botToken := "6161865465:AAHLcBcpSRWVDxIe1bK6R2l-feml36wHOAU"
	botApi := "https://api.telegram.org/bot"

	botUrl := botApi + botToken
	// идентификатор первого сообщения
	offset := 0
	for {
		// получаем обновления
		updates, err := getUpdates(botUrl, offset)

		if err != nil {
			fmt.Println("Something goes wrong: ", err.Error())
		}

		for _, update := range updates {
			err := respond(botUrl, update)
			offset = update.UpdateId + 1

			if err != nil {
				fmt.Println("Something goes wrong: ", err.Error())
			}
		}
		fmt.Println(updates)
		// fmt.Println("var = ", var)
	}

}

func getUpdates(botUrl string, offset int) ([]modules.Update, error) {

	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	fmt.Println("resp.Header", resp.Header)
	// resp.Body {0xc0000a4c00}
	fmt.Println("resp.Body", resp.Body)

	if err != nil {
		return nil, err

	}

	defer resp.Body.Close()

	// Body [123 34 111 107 34 58 116 114 117 101 44 34 114 101 115 117 108 116 34 58 91 93 125]
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Body", body)

	// как перевести байты в строку
	fmt.Println("Body", string(body))

	if err != nil {
		return nil, err
	}

	var restResponse modules.RestResponce
	// декодирование информации + парсинг входных данных
	err2 := json.Unmarshal(body, &restResponse)

	fmt.Println("restResponse", restResponse)
	fmt.Println("restResponse.Result", restResponse.Result)
	if err2 != nil {
		return nil, err
	}
	return restResponse.Result, nil

}

func respond(botUrl string, update modules.Update) error {
	var botMessage modules.BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text

	buf, err3 := json.Marshal(botMessage)
	if err3 != nil {
		return err3
	}

	res, err4 := http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	fmt.Println("my respond - ", res.Header)

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("my respond data - ", string(data))

	// my respond data -
	// {"ok":true,
	// "result":
	// 	{"message_id":222,
	// 	"from":
	// 		{"id":6161865465,
	// 		"is_bot":true,
	// 		"first_name":"go_bot",
	// 		"username":"go_special_bot"
	// 		},
	// 	"chat":
	// 		{"id":1700335247,
	// 		"first_name":"Roman",
	// 		"username":"BlackGooseOnDinner",
	// 		"type":"private"
	// 		},
	// 	"date":1676785665,
	// 	"text":"hs"}}

	if err4 != nil {
		return err4
	}
	return nil
}
