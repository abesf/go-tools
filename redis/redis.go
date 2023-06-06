package redis

import (
	"context"
	_ "context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	once     sync.Once     // 用于确保单例模式的once对象
	instance *redis2.Client // Redis客户端对象
)

func NewRedis(addr,auth string) (r *redis2.Client, cf func(), err error) {

	once.Do(func(){
		instance= redis2.NewClient(&redis2.Options{
			Addr:    addr,
			Password: auth, // no password set
			DB:       0,  // use default DB
			PoolSize: 10,
		})
	})
	r=instance
	//r = redis.NewRedis(&cfg)
	cf = func() { r.Close() }
	return
}
//分布式锁 加锁
func AcquireLock(conn *redis2.Client,key string,value string,t time.Duration) (ok bool, err error) {
	//ok,err=conn.SetNX(context.Background(),key,1,time.Second*10).Result()
	ok,err=conn.SetNX(context.Background(),key,value,t).Result()
	fmt.Printf("ok=%+v\n,err=%+v\n",ok,err)
	return
}
//分布式锁 解锁
func ReleaseLock(conn *redis2.Client,key string)  {
	ok,err:=conn.Del(context.Background(),key).Result()
	fmt.Printf("ok=%+v\n,err=%+v\n",ok,err)
}
//获取指定key 的 value
func Get(conn *redis2.Client,key string) (ok string, err error) {
	ok,err=conn.Get(context.Background(),key).Result()
	fmt.Printf("ok=%+v\n,err=%+v\n",ok,err)
	return
}