package controller

import (
	"go-fund-science/utils"
	"reflect"

	"github.com/cdfmlr/crud/log"
	"github.com/gin-gonic/gin"
)

var logger = log.ZoneLogger("crud/config")

type Pagination struct {
	PageIndex int
	PageSize  int
	Total     int
}

type PageQuery struct {
	PageIndex   int `form:"pageIndex"`
	PageSize    int `form:"pageSize"`
	QueryString string
}

func PaginationHandler[T any](c *gin.Context) PageQuery {
	var pageQuery PageQuery
	if err := c.ShouldBindQuery(&pageQuery); err != nil {
		logger.WithContext(c).WithError(err).
			Warn("PaginationHandler: bind request failed")
		ResponseError(c, CodeBadRequest, err)
		return PageQuery{}
	}
	var model T
	if err := c.ShouldBindQuery(&model); err != nil {
		logger.WithContext(c).WithError(err).
			Warn("PaginationHandler: bind request failed")
		ResponseError(c, CodeBadRequest, err)
		return PageQuery{}
	}

	pageQuery.QueryString = AccessStructField(model)

	return pageQuery
}

func AccessStructField(a interface{}) string {

	clause := utils.Clause{
		Conditions: "",
		Variables:  []interface{}{},
	}

	//获取reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取到该结构体有几个字段
	num := val.NumField()
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i).Tag.Get("form")

		v := val.Field(i)

		clause.CreateCondition(tagVal, utils.Equal, []string{string(v.String())})
	}

	return clause.BuildSQLStatement()
}
