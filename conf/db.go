package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Database  string
	Charset   string
	IsRunning bool
}

var DbMasterList = []DbConfig{
	{
		Host:      "127.0.0.1",
		Port:      "3306",
		Username:  "root",
		Password:  "root123456",
		Database:  "lottery",
		Charset:   "utf8",
		IsRunning: true,
	},
}
var DbMaster DbConfig = DbMasterList[0]
