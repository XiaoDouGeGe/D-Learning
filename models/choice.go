package models

type Choice struct {
	Id         int    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	QuestionId int    `xorm:"INTEGER" json:"questionId"`          // 对应的问题
	AContent   string `xorm:"text default 'AAA'" json:"aContent"` // A选项内容
	BContent   string `xorm:"text default 'BBB'" json:"bContent"` // B选项内容
	CContent   string `xorm:"text default 'CCC'" json:"cContent"` // C选项内容
	DContent   string `xorm:"text default 'DDD'" json:"dContent"` // D选项内容
	AShow      int    `xorm:"INTEGER default 1" json:"aShow"`     // A选项是否显示：1显示，0不显示
	BShow      int    `xorm:"INTEGER default 1" json:"bShow"`     // B选项是否显示：1显示，0不显示
	CShow      int    `xorm:"INTEGER default 1" json:"cShow"`     // C选项是否显示：1显示，0不显示
	DShow      int    `xorm:"INTEGER default 1" json:"dShow"`     // D选项是否显示：1显示，0不显示
}
