package models

type Question struct {
	Id        int    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	ChapterId int    `xorm:"INTEGER" json:"chapterId"`  // 对应的章节
	Qcontent  string `xorm:"text" json:"qContent"`      // 问题
	Qtype     string `xorm:"varchar(50)" json:"qType"`  // 类型：单选single-choice、多选multiple-choice、判断judgment
	Answer    string `xorm:"varchar(2)" json:"qAnswer"` // 答案
	Analysis  string `xorm:"text" json:"qAnalysis"`     // 解析
	Mark      int    `xorm:"INTEGER" json:"mark"`       // 分数
}
