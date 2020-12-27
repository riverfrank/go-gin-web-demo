package model

type TeacherModel struct {
	//jsonName
	Id          int64  `gorm:"column:id;" json:"id"`
	TeacherName string `gorm:"column:teacher_name;default:''" json:"teacherName"`
	CreateTime  string `gorm:"column:create_time;default:'1970-01-01 00:00:00'" json:"createTime"`
	UpdateTime  string `gorm:"column:update_time;default:'1970-01-01 00:00:00'" json:"updateTime"`
}
