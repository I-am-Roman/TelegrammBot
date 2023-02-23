package modules

type Chat struct {
	ChatId int `json:"chat"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type RestResponce struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

/*
	RestResponce:
		Update:
			UpdateId
			Message:
				Text
				Chat:
					ChatId


*/
