package models

import "time"

type Chat struct {
	Id        int       `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	UserId    int       `xorm:"INTEGER" json:"userId"`
	CourseId  int       `xorm:"INTEGER" json:"courseId"`
	Role      string    `xorm:"varchar(10) not null" json:"role"`
	Content   string    `xorm:"text" json:"content"`
	ChatTime  time.Time `json:"-"`
	ChatTimeF string    `xorm:"-" json:"time,omitempty"` // 格式化后的对话时间，便于显示
	Next      int       `xorm:"INTEGER" json:"-"`        // 下一句话 允许为空
}
