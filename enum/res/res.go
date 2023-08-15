package res

const (
	OK = iota
	FAIL
	WHAT
	EXPIRED
)

var res = map[int]string{
	OK:      "OK 啦",
	FAIL:    "????",
	WHAT:    "沒有這種東西",
	EXPIRED: "已過期",
}

func GetMsg(e int) string {
	return res[e]
}
