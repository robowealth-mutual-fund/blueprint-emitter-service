package codec

import (
	"encoding/json"
	"github.com/lovoo/goka"
)

type UserCodec struct {
	From    string
	To      string
	Content string
}

func (u *UserCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (u *UserCodec) Decode(data []byte) (interface{}, error) {
	var msg []UserCodec
	return msg, json.Unmarshal(data, &msg)
}

func NewCodec() goka.Codec {
	return &UserCodec{}
}
