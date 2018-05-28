package consts

const (
	Success = 0 // 成功

	Status = "status"
	Data   = "data"

	ErrParam    = "Invalid parameter" // 参数错误
	ErrValidate = "Validate against"  // 参数验证失败

	ErrAdmin         = "Invalid admin"                  // 用户不存在
	ErrLogin         = "Incorrect username or password" // 密码错误
	ErrLoginRequired = "Login required"                 // 未登录
	ErrPerm          = "Invalid permission"             // 权限错误

	ErrSession = "Session error" // Session 错误

	ErrNoFound = "Not found"     // 未查询到数据
	ErrMongo   = "Mongodb error" // MySQL 错误
)
