package dao

import (
	"go-fund-science/database"
	"go-fund-science/logger"
	"go-fund-science/models"
	"go-fund-science/utils"
	"time"

	"github.com/google/uuid"
)

const tableName = "com_person"

// 新增人员
func Insert(person models.ComPerson) (string, error) {
	person.ID = uuid.New()
	t := models.LocalTime(time.Now())
	person.Created = &t
	person.Updated = &t
	tx := database.DB.Table(tableName).Create(&person)
	if tx.Error != nil {
		logger.Info("insert person in database.com_person fail:" + tx.Error.Error())
		return "", tx.Error
	}
	return person.ID.String(), nil
}

// 删除人员
func Delete(id string) (string, error) {
	tx := database.DB.Table(tableName).Delete(&models.ComPerson{}, "id = ?", id)
	if tx.Error != nil {
		logger.Info("delete person in database.com_person fail:" + tx.Error.Error())
		return "", tx.Error
	}
	return id, nil
}

// 更新人员所有信息
func Update(person models.ComPerson) (string, error) {
	tx := database.DB.Table(tableName).Save(&person)

	if tx.Error != nil {
		logger.Info("update person in database.com_person fail:" + tx.Error.Error())
		return "", tx.Error
	}
	return person.ID.String(), nil
}

// 根据人员ID更新人员基础信息(姓名、性别和手机号)
func UpdatePersonBaseInfo(person models.ComPerson) (string, error) {
	tx := database.DB.Table(tableName).Updates(models.ComPerson{PersonName: person.PersonName, Sex: person.Sex, Mobile: person.Mobile})

	if tx.Error != nil {
		logger.Info("update PersonBaseInfo in database.com_person fail:" + tx.Error.Error())
		return "", tx.Error
	}

	return person.ID.String(), nil
}

// 根据人员ID获取人员信息
func GetPersonByID(id string) (models.ComPerson, error) {
	var person models.ComPerson
	tx := database.DB.Table(tableName).Where("id = ?", id).First(&person)
	if tx.Error != nil {
		logger.Info("GetPersonByID fail:" + tx.Error.Error())
		return person, tx.Error
	}
	return person, nil
}

// 根据人员姓名获取人员信息
func GetPersonByPersonName(personName string) (models.ComPerson, error) {
	var person models.ComPerson
	clause := utils.NewClause()
	clause.CreateCondition("person_name", utils.LIKE, []string{personName})
	strSQL := clause.BuildSQLStatement()
	tx := database.DB.Table(tableName).Where(strSQL).First(&person)
	if tx.Error != nil {
		logger.Info("GetPersonByPersonName fail:" + tx.Error.Error())
		return person, tx.Error
	}
	return person, nil
}

// 根据身份证号获取人员信息
func GetPersonByCertID(CertID string) (models.ComPerson, error) {
	var person models.ComPerson
	tx := database.DB.Table(tableName).Where("cert_id = ?", CertID).First(&person)
	if tx.Error != nil {
		logger.Info("GetPersonByCertID fail:" + tx.Error.Error())
		return person, tx.Error
	}
	return person, nil
}

// 分页查询符合条件的人员信息列表
func GetDataListByPage(strWhere string, orderString string, pageIndex int, pageSize int, rowCount *int64) ([]models.ComPerson, error) {
	var persons []models.ComPerson
	offset := (pageIndex - 1) * pageSize

	// 分页查询
	tx := database.DB.Table(tableName).Debug().Where(strWhere).Count(rowCount).Offset(offset).Limit(pageSize).Find(&persons)
	if tx.Error != nil {
		logger.Info("GetDataListByPage query fail:" + tx.Error.Error())
		return persons, tx.Error
	}

	return persons, nil
}

// var persons []models.ComPerson
// 	page := 1
// 	pageSize := 10
// 	offset := (page - 1) * pageSize
// 	// 分页查询
// 	orm.DB.Table("com_person").Debug().Where("person_name like '%徐%'").Offset(offset).Limit(pageSize).Find(&persons)

// 	for _, person := range persons {
// 		fmt.Printf("%v\n", *person.PersonName)
// 	}

// 	orm.DB.Raw("SELECT * FROM com_person WHERE name = ?", "john").Scan(&persons)
