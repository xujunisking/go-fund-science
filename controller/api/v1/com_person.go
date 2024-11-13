package controller

import (
	con "go-fund-science/controller"
	"go-fund-science/dao"
	"go-fund-science/models"
	"go-fund-science/utils"
	"go-fund-science/utils/result"
	"net/http"

	"github.com/cdfmlr/crud/log"
	"github.com/gin-gonic/gin"
)

var logger = log.ZoneLogger("crud/config")

func PersonInsert(c *gin.Context) {
	var person models.ComPerson
	err := c.ShouldBindQuery(&person)
	if err != nil {
		logger.WithContext(c).WithError(err).
			Warn("ShouldBindQuery Person error: bind request failed")
		c.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	ID, err := dao.Insert(person)
	if err != nil {
		logger.WithContext(c).WithError(err).
			Warn("Insert Person error: bind request failed")
		c.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Result(result.OK, ID, "人员新增成功！"))
}

func PersonDelete(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusOK, utils.Result(result.OK, nil, "参数缺少人员id"))
		return
	}
	ID, err := dao.Delete(id)
	if err != nil {
		logger.WithContext(c).WithError(err).
			Warn("Delete Person error: bind request failed")
		c.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Result(result.OK, ID, "人员删除成功！"))
}

func PersonUpdate(c *gin.Context) {
	var person models.ComPerson
	err := c.ShouldBindQuery(&person)
	if err != nil {
		logger.WithContext(c).WithError(err).
			Warn("ShouldBindQuery Person error: bind request failed")
		c.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}
	ID, err := dao.Update(person)
	if err != nil {
		logger.WithContext(c).WithError(err).
			Warn("Update Person error: bind request failed")
		c.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Result(result.OK, ID, "人员信息修改成功！"))
}

// 根据人员ID获取人员信息
// GetPersonByID godoc
// @Summary      显示人员信息
// @Description  根据人员ID获取人员信息
// @Tags         Persons
// @Accept       json
// @Produce      json
// @Param        id   query      string  true  "ID"
// @Success      200  {object}  models.ComPerson
// @Failure      400  {string}  string "{"msg": "error info"}"
// @Router       /persons/GetPersonByID [get]
func GetPersonByID(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		id = utils.ZeroUUID().String()
	}

	person, err := dao.GetPersonByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.Result(result.OK, person, ""))
}

// 根据人员姓名获取人员信息
// GetPersonByPersonName godoc
// @Summary      显示人员信息
// @Description  根据人员姓名获取人员信息
// @Tags         Persons
// @Accept       json
// @Produce      json
// @Param        personName   query      string  true  "personName"
// @Success      200  {object}  models.ComPerson
// @Failure      400  {string}  string "{"msg": "error info"}"
// @Router       /persons/GetPersonByPersonName [get]
func GetPersonByPersonName(ctx *gin.Context) {
	personName, ok := ctx.GetQuery("personName")
	if !ok {
		personName = "somebody"
	}

	person, err := dao.GetPersonByPersonName(personName)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.Result(result.OK, person, ""))
}

// 根据身份证号获取人员信息
// GetPersonByCertID godoc
// @Summary      显示人员信息
// @Description  根据身份证号获取人员信息
// @Tags         Persons
// @Accept       json
// @Produce      json
// @Param        certID   query      string  true  "certID"
// @Success      200  {object}  models.ComPerson
// @Failure      400  {string}  string "{"msg": "error info"}"
// @Router       /persons/GetPersonByCertID [get]
func GetPersonByCertID(ctx *gin.Context) {
	certID, ok := ctx.GetQuery("certID")
	if !ok {
		certID = "somebody"
	}

	person, err := dao.GetPersonByCertID(certID)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.Result(result.OK, person, ""))
}

// 分页查询人员信息
func GetDataListByPage(c *gin.Context) {
	p := con.PaginationHandler[models.ComPerson](c)
	var rowCount int64
	persons, err := dao.GetDataListByPage(p.QueryString, "cert_id asc", p.PageIndex, p.PageSize, &rowCount)
	if err != nil {
		logger.WithContext(c).WithError(err).
			Warn("GetDataListByPage: bind request failed")
		c.JSON(http.StatusOK, utils.Result(result.OK, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Result(result.OK, persons, ""))
}
