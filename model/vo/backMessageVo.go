package vo

import (
	"encoding/json"
)

type BackMessageVo struct {
	Id      string
	Key     string
	Value   string
	GuildId string
}

// implements encoding.BinaryMarshaler
func (e *BackMessageVo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(e)
}

// implements encoding.BinaryUnmarshaler
func (e *BackMessageVo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}
