package errx

const (
	OK uint32 = 200

	// 前 3 位代表业务 后3 位代表具体错误

	// 全局错误码
	// 服务器开小差
	ServerCommonError uint32 = 100001
	// 请求参数错误
	RequestParamError uint32 = 100002
	// 数据库错误
	DBError uint32 = 100003
	// token 过期
	TokenExpireError uint32 = 100004
)
