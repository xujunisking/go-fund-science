package redisManager

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisConfig struct {
	Timeout  int    //连接超时时间(单位：秒)
	Host     string //服务器地址
	Port     int    //服务器端口
	Database int    //数据库
	Password string //密码
	Pool     Pool   //redis连接池配置
}

type Pool struct {
	MaxIdle         int  // 最大空闲连接
	MaxActive       int  // 最大连接数
	IdleTimeout     int  // 最大空闲时间(单位：秒)
	Wait            bool // 设置为true，则客户端会等待;设置为false，则客户端会立即返回一个错误
	MaxConnLifetime int  // 连接最大生命周期(单位：秒)
}

var RedisConn *redis.Pool

func NewRedisConn(config *RedisConfig) error {
	address := fmt.Sprintf("%s:%s", config.Host, string(config.Port))
	RedisConn = &redis.Pool{
		// Maximum number of idle connections in the pool.(连接池idle连接数)
		MaxIdle: config.Pool.MaxIdle,
		// Maximum number of connections allocated by the pool at a given time.(连接池在给定时间内分配的最大连接数)
		// When zero, there is no limit on the number of connections in the pool(当连接池中的连接数为0时，没有限制)
		MaxActive: config.Pool.MaxActive,

		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: time.Duration(config.Pool.IdleTimeout) * time.Second,

		// If Wait is true and the pool is at the MaxActive limit, then Get() waits
		// for a connection to be returned to the pool before returning.
		Wait: config.Pool.Wait,

		// Close connections older than this duration. If the value is zero, then
		// the pool does not close connections based on age.
		MaxConnLifetime: time.Duration(config.Pool.MaxConnLifetime) * time.Second,

		//创建和配置链接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, redis.DialDatabase(config.Database), redis.DialConnectTimeout(time.Duration(config.Timeout)*time.Second))
			if err != nil {
				return nil, err
			}
			if config.Password != "" && config.Password != "none" {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}
