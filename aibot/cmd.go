package aibot

import (
	"encoding/json/jsontext"

	"github.com/google/uuid"
)

const (
	CMD_PING         = "ping"
	CMD_MSG_CALLBACK = "aibot_msg_callback"
	CMD_SUBSCRIBE    = "aibot_subscribe"
	CMD_RESPOND_MSG  = "aibot_respond_msg"
)

type Headers struct {
	ReqID string `json:"req_id"`
}

type SubscribeBody struct {
	BotID  string `json:"bot_id"`
	Secret string `json:"secret"`
}
type SubscribeReq struct {
	Cmd     string        `json:"cmd"`
	Headers Headers       `json:"headers"`
	Body    SubscribeBody `json:"body"`
}

type PingReq struct {
	Cmd     string  `json:"cmd"`
	Headers Headers `json:"headers"`
}

type Response struct {
	Cmd     string  `json:"cmd,omitempty"`
	Headers Headers `json:"headers,omitzero"`

	Body jsontext.Value `json:"body,omitempty"`

	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

func (r Response) IsCmd() bool {
	return r.Cmd != ""
}

type MsgCallBackBody struct {
	MsgID    string `json:"msgid"`
	AiBotID  string `json:"aibotid"`
	ChatID   string `json:"chatid"`
	ChatType string `json:"chattype"`
	From     struct {
		UserID string `json:"userid"`
	} `json:"from"`
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type TextMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type MarkdownMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type RespondMsg[T any] struct {
	Cmd     string  `json:"cmd,omitempty"`
	Headers Headers `json:"headers,omitzero"`
	Body    T       `json:"body"`
}

func ReqID(cmd string) string {
	return cmd + "#" + uuid.NewString()
}
