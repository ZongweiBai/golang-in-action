package config

import (
	"context"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"strconv"
	"sync"
	"time"
)

var redisOnce sync.Once

// InitRedis 初始化Redis连接
func InitRedis() {
	if CONFIG.Redis.Enabled == false {
		return
	}
	redisOnce.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     CONFIG.Redis.Host + ":" + strconv.Itoa(CONFIG.Redis.Port),
			Password: CONFIG.Redis.Password,
			DB:       CONFIG.Redis.DB,
			PoolSize: CONFIG.Redis.PoolSize,
		})

		// 哨兵模式
		// rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		// 	MasterName:    "master",
		// 	SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
		// })

		// 集群模式
		// rdb := redis.NewClusterClient(&redis.ClusterOptions{
		// 	Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
		// })

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			LOG.Errorf("初始化Redis失败:%v", err)
			panic(err)
		}

		RDB = rdb
	})
}
