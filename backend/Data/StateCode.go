package Data

// 不太清楚具体该怎么定义，总之先用类似HTTP头的状态码
// TODO: Goland叫我改成驼峰拼法
const (
	LOGIN_SUCCESS    = 2000
	REGISTER_SUCCESS = 2001
	VALIDATE_SUCCESS = 2002

	INVALID_PARAMS = 4000

	ERROR_TOKEN_CHECK_FAIL = 4010
	ERROR_TOKEN_EXPIRED    = 4011
	ERROR_WRONG_PASSWORD   = 4012
	ERROR_NONEXISTENT_USER = 4013
	ERROR_EXISTED_USER     = 4014
)

// 向前端发送的数据格式
// "code"    : 处理的状态
// "message" : 附加的消息
// "account" : 用户的账户
// "token"   : Token，只在登录的时候发

// 前端发回的数据格式
// "token"   : Token
// "data"    : 其他的数据，之后再增加
