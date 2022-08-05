package user

import (
	"encoding/json"

	"github.com/lovoo/goka"
)

type Login struct {
	UserID    string
	Timestamp int64
}

func NewCodec() goka.Codec {
	return &Login{}
}

func (l *Login) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (l *Login) Decode(data []byte) (interface{}, error) {
	return &l, json.Unmarshal(data, &l)
}
