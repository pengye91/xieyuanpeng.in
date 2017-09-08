package configs

var (
	AccessLogPath    = BACKEND_ROOT + "/logs/gin_access.log"
	ErrorLogPath     = BACKEND_ROOT + "/logs/gin_error.log"
	AllToOneFilePath = BACKEND_ROOT + "/logs/gin_all_to_one.log"
)

//func init() {
//	os.Setenv("GIN_ACCESS_LOG", os.Getenv("MYGOPATH")+"/xieyuanpeng.in/logs/gin_access.log")
//	os.Setenv("GIN_ERROR_LOG", os.Getenv("MYGOPATH")+"/xieyuanpeng.in/logs/gin_error.log")
//}
