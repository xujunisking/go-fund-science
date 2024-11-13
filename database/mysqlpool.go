package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logs "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	conn *gorm.DB
)

type DBConn struct {
	UserName string
	Password string
	Host     string
	Port     int
	Dbname   string
	Timeout  string
}

// 配置MySQL连接参数
// username := "root"  //账号
// password := "123456" //密码
// host := "127.0.0.1" //数据库地址，可以是Ip或者域名
// port := 3306 //数据库端口
// Dbname := "tizi365" //数据库名
// timeout := "10s" //连接超时，10秒
func ConnDB(username, password, host string, port int, Dbname, timeout string) *gorm.DB {
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return _db
}

// 配置MySQL连接参数
// username := "root"  //账号
// password := "123456" //密码
// host := "127.0.0.1" //数据库地址，可以是Ip或者域名
// port := 3306 //数据库端口
// Dbname := "tizi365" //数据库名
// timeout := "10s" //连接超时，10秒
func ConnString(connConfig *DBConn) string {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		connConfig.UserName, connConfig.Password, connConfig.Host, connConfig.Port, connConfig.Dbname, connConfig.Timeout)

	return dsn
}

// 定义自定义的NamingStrategy
var namingStrategy = schema.NamingStrategy{
	TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`(表前缀)
	SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled(单数表名)
	NoLowerCase:   true, // skip the snake_casing of names(设置为false时，表名会被转换为小写，而设置为true则会保持表名的原始大小写)
	//NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name(字段名替换器)
}

var newLogger = logs.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logs.Config{
		SlowThreshold:             time.Second, // 慢查询的阈值，当查询超出这个时间阈值时，Gorm会记录该查询为慢查询。
		LogLevel:                  logs.Silent, // 日志级别
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger(确保在查询未找到记录时不会返回错误)
		ParameterizedQueries:      true,        // Don't include params in the SQL log(控制 SQL 日志中是否包含参数)
		Colorful:                  true,        // Disable color(true时，日志将以彩色方式输出)
	},
)

func ConnectionDB(connConfig *DBConn) (conn *gorm.DB, err error) {
	dsn := ConnString(connConfig)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//NamingStrategy: namingStrategy,
		Logger: newLogger,
	})
	if err != nil {
		panic("mysql database connection failed, error=" + err.Error())
	}

	conn = _db

	sqlDB, err := _db.DB()
	if err != nil {
		panic("sql.DB() failed, error=" + err.Error())
	}

	//设置连接池的最大闲置连接数
	sqlDB.SetConnMaxIdleTime(10)
	//设置连接池中的最大连接数量
	sqlDB.SetMaxOpenConns(100)
	//设置连接的最大复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return conn, err
}

func GetConntion() *gorm.DB {
	return conn
}

func InitDb(connConfig *DBConn) (*gorm.DB, error) {
	dsn := ConnString(connConfig)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	conn = db

	sqlDB, err := db.DB()
	if err != nil {
		panic("sql.DB() failed, error=" + err.Error())
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(100)

	//db.SingularTable(true)//这个函数已经被取消
	//根据对应结构体初始化数据库表
	//db.AutoMigrate(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))
	return db, nil
}
