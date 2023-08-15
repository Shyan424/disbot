package res

type Res int

const (
	OK Res = iota
	FAIL
	WHAT
	EXPIRED
)

var mapToString = map[Res]string{
	OK:      "OK 啦",
	FAIL:    "????",
	WHAT:    "沒有這種東西",
	EXPIRED: "已過期",
}

func (r Res) GetMsg() string {
	return mapToString[r]
}
