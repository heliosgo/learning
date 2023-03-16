package errx

var (
	message = map[uint32]string{
		OK:                "SUCCESS",
		ServerCommonError: "服务器开小差啦，请稍后再来试一试",
		RequestParamError: "请求参数错误",
		DBError:           "数据库错误",
		TokenExpireError:  "Token 过期",
	}
)

func mapErrMsg(code uint32) string {
	if msg, ok := message[code]; ok {
		return msg
	}

	return "服务器开小差啦，请稍后再来试一试"
}
