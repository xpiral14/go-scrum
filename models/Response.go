package models

type Response struct {
	Command int `json:"command"`
	Payload interface{}
}

func NewResponse(command int, payload interface{}) *Response {
	return &Response{
		command,
		payload,
	}
}
