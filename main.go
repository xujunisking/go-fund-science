package main

import (
	"fmt"
	conf "go-fund-science/config"
	"go-fund-science/database"
	"go-fund-science/docs"
	"go-fund-science/middleware"
	"go-fund-science/router"
	redisManager "go-fund-science/utils/redis"

	"github.com/cdfmlr/crud/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type MyConfig struct {
	Name              string
	config.BaseConfig `mapstructure:",squash"`
	database.DBConn
	Redis redisManager.RedisConfig
}

var myConfig MyConfig

// 运行环境
const currentEvnMode = conf.Dev

// 读取配置文件
func ReadConfig() error {
	err := config.Init(&myConfig, conf.ReadConfigInfo(currentEvnMode))
	if err != nil {
		fmt.Printf("read load file [application_%v.yaml] error = %v", currentEvnMode, err)
		return err
	}
	return nil
}

// 连接数据库
func DataBaseConnection() error {
	_, err := database.ConnectDB(database.DBDriverMySQL, myConfig.DB.DSN)
	if err != nil {
		fmt.Printf("connect database error = %v", err)
		return err
	}
	return nil
}

// 初始化
func init() {

	err := ReadConfig()
	if err != nil {
		fmt.Printf("read Config error: %v", err)
		return
	}

	myConfig.DB.DSN = database.ConnString(&myConfig.DBConn)

	err = DataBaseConnection()
	if err != nil {
		fmt.Printf("database connettion error: %v", err)
		return
	}

	err = redisManager.NewRedisConn(&myConfig.Redis)
	if err != nil {
		fmt.Printf("reids connettion error: %v", err)
		return
	}

	fmt.Println(myConfig)
	//注册模型
	//orm.RegisterModel(Todo{}, Project{})
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	this is go-gin-gorm example.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost" + myConfig.HTTP.Addr
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	route := gin.Default()
	//注册swagger api相关路由
	route.Use(middleware.Cors()) // 跨域
	//port := fmt.Sprintf(":%d", config.GetInt("application.http-server-port"))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NewRouter(route)

	route.Run(myConfig.HTTP.Addr)
}
