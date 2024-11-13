package models

import (
	"github.com/google/uuid"
)

type ComPerson struct {
	BaseModel   `mapstructure:",squash"`
	CertId      string     `gorm:"unique;not null;column:cert_id" form:"cert_id" json:"cert_id"`
	PersonName  *string    `gorm:"not null;default:1;column:person_name" form:"person_name" json:"person_name"`
	Sex         string     `gorm:"column:sex" form:"sex" json:"sex"`
	Birthday    *LocalTime `gorm:"type:datetime;column:birthday" form:"birthday" json:"birthday"`
	UnitId      uuid.UUID  `gorm:"column:unit_id" form:"unit_id" json:"unit_id"`
	PersonState int32      `gorm:"column:person_state" form:"person_state" json:"person_state"`
	Education   uuid.UUID  `gorm:"column:education" form:"education" json:"education"`
	Nation      uuid.UUID  `gorm:"column:nation" form:"nation" json:"nation"`
	Title       uuid.UUID  `gorm:"column:title" form:"title" json:"title"`
	Spec        uuid.UUID  `gorm:"column:spec" form:"spec" json:"spec"`
	Email       string     `gorm:"type:varchar(100)column:email" form:"email" json:"email"`
	Mobile      string     `gorm:"type:varchar(100)column:mobile" form:"mobile" json:"mobile"`
	Remark      string     `gorm:"type:varchar(200)column:remark" form:"remark" json:"remark"`
	//Email       sql.NullString `gorm:"type:varchar(100)column:email" form:"email" json:"email"`
}

//数据库中保存为空字符串
//user :=User{Name:new(string)，Age:18))} // 字段类型为*string
//user :=User{Name:sql.NullString{"", true}, Age:18} //字段类型为sql.NullString
