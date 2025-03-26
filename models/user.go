package models

type User struct {
	Id       int    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	Username string `xorm:"varchar(30) not null" json:"username"`
	Password string `xorm:"varchar(128) not null" json:"-"`
	Secret   string `xorm:"varchar(30)" json:"secret"`
	Status   string `xorm:"varchar(10) not null" json:"status"`
}
