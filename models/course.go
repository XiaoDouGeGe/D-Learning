package models

type Course struct {
	Id         int    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	CourseName string `xorm:"VARCHAR(30)" json:"courseName"`
	CourseDesc string `xorm:"VARCHAR(100)" json:"courseDesc"`
	CourseImg  string `xorm:"VARCHAR(100)" json:"courseImg"`
	Status     string `xorm:"VARCHAR(50)" json:"status"`
}
