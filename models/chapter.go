package models

type Chapter struct {
	Id             int    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	CourseId       int    `xorm:"INTEGER" json:"courseId"`     // 对应的课程
	ChapterIndex   int    `xorm:"INTEGER" json:"chapterIndex"` // 第几章节
	ChapterTitle   string `xorm:"varchar(100)" json:"chapterTitle"`
	ChapterContent string `xorm:"text" json:"chapterContent"`
	Status         string `xorm:"varchar(50)" json:"status"`
}
