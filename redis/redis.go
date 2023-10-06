package redis

import (
	"github.com/go-redis/redis/v8"
	"sync"
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