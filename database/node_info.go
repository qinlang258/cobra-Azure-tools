package database

type NodeInfo struct {
	ID             int     `gorm:"column:id;primary_key;auto_increment;comment:ID"`
	IP             string  `gorm:"column:ip;type:varchar(255);comment:IP"`
	Datetime       string  `gorm:"column:datetime;type:varchar(255);comment:服务器时间"`
	Status         string  `gorm:"column:status;type:varchar(255);comment:服务器状态"`
	StartedTime    string  `gorm:"column:started_time;type:varchar(255);comment:开机时间"` // Pointed type for nullable datetime
	UserNumber     int     `gorm:"column:user_number;type:int;comment:运行的用户数量 "`
	LoadAverage1m  float64 `gorm:"column:loadaverage_1m;type:decimal(10,2);comment:1分钟负载"`
	LoadAverage5m  float64 `gorm:"column:loadaverage_5m;type:decimal(10,2);comment:5分钟负载"`   // Changed type to float64 for decimal representation
	LoadAverage15m float64 `gorm:"column:loadaverage_15m;type:decimal(10,2);comment:15分钟负载"` // Changed type to float64 for decimal representation
}

func (NodeInfo) TableName() string {
	return "node_info"
}
