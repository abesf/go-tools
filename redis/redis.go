package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)
var once sync.Once
var redisClients map[string]*redis.Client
//单例模式 初始化redis链接
func GetRedisClient(addr,pass string) *redis.Client {
	// 使用sync.Once确保每个地址只被初始化一次
	once.Do(func() {
		redisClients = make(map[string]*redis.Client)
	})

	// 检查是否已经有一个Redis客户端针对这个地址
	client, exists := redisClients[addr]
	if exists {
		return client
	}

	// 如果没有，则创建一个新的Redis客户端
	client = redis.NewClient(&redis.Options{
		Addr:     addr, // 使用传入的地址配置
		Password: pass,   // Redis密码
		DB:       0,    // use default DB
		PoolSize: 100,
	})

	// 存储在map中，以便下次重用
	redisClients[addr] = client

	return client
}
//资源加锁
func AcquireLock(conn *redis.Client,key string,value string,t time.Duration) (ok bool, err error) {
	ok,err=conn.SetNX(context.Background(),key,value,t).Result()
	//fmt.Printf("AcquireLock key=%+v  ok=%+v ,err=%+v ",key,ok,err)
	return
}
//资源解锁
func ReleaseLock(conn *redis.Client,key string) (ok int64, err error) {
	ok,err=conn.Del(context.Background(),key).Result()
	//fmt.Printf("ok=%+v\n,err=%+v\n",ok,err)
	return
}