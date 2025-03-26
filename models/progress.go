package models

type Progress struct {
	Id         int    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	UserId     int    `xorm:"INTEGER" json:"userId"`
	ChapterId  int    `xorm:"INTEGER" json:"chapterId"`
	Status     string `xorm:"varchar(10) not null" json:"status"` // Not started、In progress、Completed
	Percentage int    `xorm:"INTEGER" json:"percentage"`          // 0-100
}
