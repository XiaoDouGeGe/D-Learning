package models

type Exercise struct {
	Id        int    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	UserId    int    `xorm:"INTEGER" json:"userId"`
	ChapterId int    `xorm:"INTEGER" json:"chapterId"`
	Content   string `xorm:"text default '[]'" json:"content"`                       // 答题情况，json格式，[{qId:题目序号,uAnswer:用户答案}]
	Score     int    `xorm:"INTEGER" json:"score"`                                   // 得分
	ETime     string `xorm:"varchar(20) default '2025-01-01 12:00:00'" json:"eTime"` // 考试时间 2025-01-01 12:00:00
}
