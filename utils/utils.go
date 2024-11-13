package utils

import (
	"go-fund-science/models"
	"go-fund-science/utils/result"
	"time"

	"github.com/cdfmlr/crud/log"

	"github.com/google/uuid"
)

var logger = log.ZoneLogger("crud/config")

// 将字符串转换为UUID
func UUIDParse(uid string) uuid.UUID {
	ID, err := uuid.Parse(uid)
	if err != nil {
		logger.WithError(err).
			Error("IUUIDParse error: readFromFile error")
		return uuid.Nil
	}
	return ID
}

// 返回全零值UUID
func ZeroUUID() uuid.UUID {
	ID, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")

	return ID
}

// 字符串转换为时间
func DateTimeParse(timeStr string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		logger.WithError(err).
			Error("DateTimeParse error: readFromFile error")
		return time.Time{} //时间的零值
	}
	return t
}

func ModelTimeParse(timeStr string) models.LocalTime {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		logger.WithError(err).
			Error("DateTimeParse error: readFromFile error")
		return models.LocalTime(time.Time{}) //时间的零值
	}
	return models.LocalTime(t)
}

// func ModelTimeParse(timeStr string) models.Time {
// 	layout := "2006-01-02 15:04:05"
// 	t, err := time.Parse(layout, timeStr)
// 	if err != nil {
// 		logger.WithError(err).
// 			Error("DateTimeParse error: readFromFile error")
// 		return models.Time{MTime: time.Time{}} //时间的零值
// 	}
// 	return models.Time{MTime: t}
// }

func Result(code int, data interface{}, msg string) result.Result {
	if msg == "" {
		msg = result.ResultMsg(code)
	}

	return result.Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
