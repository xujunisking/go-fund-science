package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Model is the interface for all models.
// It only requires an Identity() method to return the primary key field
// name and value.
type Model interface {
	// Identity returns the primary key field of the model.
	// A very common case is that the primary key field is ID.
	Identity() (fieldName string, value any)
}

// BasicModel implements Model interface with an auto increment primary key ID.
//
// BasicModel is actually the gorm.Model struct which contains the following
// fields:
//
//	ID, CreatedAt, UpdatedAt, DeletedAt
//
// It is a good idea to embed this struct as the base struct for all models:
//
//	type User struct {
//	  orm.BasicModel
//	}
type BaseModel struct {
	//ID      uuid.UUID  `gorm:"primarykey;column:id" form:"id" json:"id"`
	ID      uuid.UUID  `gorm:"primarykey;column:id" form:"id" json:"id"`
	Created *LocalTime `gorm:"type:datetime;column:created" form:"created" json:"created"`
	Updated *LocalTime `gorm:"type:datetime;column:updated" form:"updated" json:"updated"`
}

// type BasicModel gorm.Model
type BasicModel BaseModel

func (m BasicModel) Identity() (fieldName string, value any) {
	return "ID", m.ID
}

type LocalTime time.Time

// MarshalJSON 序列化为JSON
func (t LocalTime) MarshalJSON() ([]byte, error) {
	v := time.Time(t)
	if v.IsZero() {
		return []byte(""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", v.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化为JSON
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	v, err := time.Parse("2006-01-02 15:04:05", string(data)[1:20])
	*t = LocalTime(v)
	return err
}

// String 重写String方法
func (t *LocalTime) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// SetRaw 读取数据库值
func (t *LocalTime) SetRaw(value interface{}) error {
	switch value.(type) {
	case time.Time:
		*t = LocalTime(value.(time.Time))
	}
	return nil
}

// RawValue 写入数据库
func (t *LocalTime) RawValue() interface{} {
	v := time.Time(*t)
	str := v.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return nil
	}
	return str
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 实现 driver.Valuer 接口
func (t LocalTime) Value() (driver.Value, error) {
	v := time.Time(t)
	return v, nil
}

// // region 日期转json代码
// type Time struct {
// 	MTime time.Time
// }

// // MarshalJSON 序列化为JSON
// func (t Time) MarshalJSON() ([]byte, error) {
// 	if t.MTime.IsZero() {
// 		return []byte("\"\""), nil
// 	}
// 	stamp := fmt.Sprintf("\"%s\"", t.MTime.Format("2006-01-02 15:04:05"))
// 	return []byte(stamp), nil
// }

// // UnmarshalJSON 反序列化为JSON
// func (t *Time) UnmarshalJSON(data []byte) error {
// 	var err error
// 	t.MTime, err = time.Parse("2006-01-02 15:04:05", string(data)[1:20])
// 	return err
// }

// // String 重写String方法
// func (t *Time) String() string {
// 	data, _ := json.Marshal(t)
// 	return string(data)
// }

// // SetRaw 读取数据库值
// func (t *Time) SetRaw(value interface{}) error {
// 	switch value.(type) {
// 	case time.Time:
// 		t.MTime = value.(time.Time)
// 	}
// 	return nil
// }

// // RawValue 写入数据库
// func (t *Time) RawValue() interface{} {
// 	str := t.MTime.Format("2006-01-02 15:04:05")
// 	if str == "0001-01-01 00:00:00" {
// 		return nil
// 	}
// 	return str
// }

// func (t *Time) Scan(v interface{}) error {
// 	value, ok := v.(time.Time)
// 	if ok {
// 		*t = Time{MTime: value}
// 		return nil
// 	}
// 	return fmt.Errorf("can not convert %v to timestamp", v)
// }
