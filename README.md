# go-fund-science
## GORM(Object Relational Mapping 对象关系映射)
## 1.开启go modules功能
### 1)GO111MODULE=off，关闭go modules功能，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
### 2)GO111MODULE=on，开启go modules功能，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
### 3)GO111MODULE=auto，默认值，go命令会根据当前目录中是否有go.mod文件来决定是否启用module功能。这种情况下可以分为两种情形：
#### A.当项目路径在GOPATH目录外部时， 设置为GO111MODULE = on
#### B.当项目路径位于GOPATH内部时，即使存在go.mod, 设置为GO111MODULE = off
go env -w GO111MODULE=auto
## 2.生成go.mod文件
go mod init go-fund-science
## 3.下载并安装 gin
go get -u github.com/gin-gonic/gin
go get github.com/google/uuid
go get github.com/sirupsen/logrus
go get -u github.com/lestrrat-go/file-rotatelogs
go get github.com/cdfmlr/crud/orm
go get github.com/fsnotify/fsnotify
go get -u https://github.com/spf13/viper
## 4.代码完成后安装swagger
go get -u github.com/swaggo/swag/cmd/swag
## 确保$GOPATH/bin已经添加到你的环境变量$PATH中，这样你就可以在任何地方运行swag命令。(确保swag.exe文件在go的安装目录下的bin文件夹中)
npm install -g swag
go install github.com/swaggo/swag/cmd/swag@latest
swag init

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
## 移除未使用的‌依赖项‌\更新依赖项的版本‌\保持go.mod文件的同步
go mod tidy


## swagger注释标签
### main函数标签如下：
// @title Swagger Example API
// @version 1.0
// @description this is a sample server celler server
// @termsOfService https://www.swagger.io/terms/
 
// @contact.name dracula
// @contact.url http://www.swagger.io/support
// @contact.email abc.xyz@qq.com
 
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host 127.0.0.1:8080

### controller层函数标签如下：
/* @Param name query string true "Name"
name：参数名
query：参数位置，可以是query（查询参数）、path（路径参数）、header（头部参数）、body（请求体参数）等
string：参数类型
true：是否必填
"Name"：参数描述
*/
// @Summary 测试sayHello
// @Description 向你说hello
// @Tags 测试
// @Accept json
// @Param name query string true "人名"
// @Success 200 {string} string "{"msg": "hello wy"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hello [get]