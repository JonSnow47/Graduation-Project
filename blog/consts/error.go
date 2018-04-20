package consts

const (
	Failure = iota // 失败
	Success        // 成功

	ErrParam = "Invalid parameter"  // 参数错误
	ErrPerm  = "Invalid permission" // 权限错误

	ErrAdmin         = "Invalid admin"                  // 用户不存在
	ErrLogin         = "Incorrect username or password" // 密码错误
	ErrLoginRequired = "Login required"                 // 未登录

	ErrSession = "Session error" // Session 错误

	ErrNoFound = "Not found"     // 未查询到数据
	ErrMongo   = "Mongodb error" // MySQL 错误
)
