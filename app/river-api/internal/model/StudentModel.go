package model

type StudentModel struct {
	//jsonName
	Id         int64  `gorm:"column:id;" json:"id"`
	StudName   string `gorm:"column:stud_name;default:''" json:"studName"`
	StudAge    int32  `gorm:"column:stud_age;default:'0'" json:"studAge"`
	StudSex    string `gorm:"column:stud_sex;default:''" json:"studSex"`
	CreateTime string `gorm:"column:create_time;default:'1970-01-01 00:00:00'" json:"createTime"`
	UpdateTime string `gorm:"column:update_time;default:'1970-01-01 00:00:00'" json:"updateTime"`
}
