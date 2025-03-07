package cache

import (
	"errors"
	"github.com/gookit/color"
	"github.com/redis/go-redis/v9"
	"time"
)

type Queue struct {
	QueueName string
	C         *redis.Client
}

// NewQueue 创建队列
func NewQueue(queueName string, r *redis.Client) *Queue {
	return &Queue{QueueName: queueName, C: r}
}

// Push 推送消息
func (q *Queue) Push(message string) {
	_, err := q.C.RPush(ctx, q.QueueName, message).Result()
	if err != nil {
		color.Redln("Push err: ", err)
	}
}

// RPop 消费消息
func (q *Queue) RPop(handle func(message string) error) error {
	message, err := q.C.RPop(ctx, q.QueueName).Result()
	if errors.Is(err, redis.Nil) {
		time.Sleep(1 * time.Second)
		return nil
	} else if err != nil {
		color.Redln("Pop err: ", err)
		return err
	}
	err = handle(message)
	if err != nil {
		return err
	}
	return nil
}

// LPop 消费消息
func (q *Queue) LPop(handle func(message string) error) error {
	message, err := q.C.LPop(ctx, q.QueueName).Result()
	if errors.Is(err, redis.Nil) {
		time.Sleep(1 * time.Second)
		return nil
	} else if err != nil {
		color.Redln("Pop err: ", err)
		return err
	}
	err = handle(message)
	if err != nil {
		return err
	}
	return nil

}

// Len 队列长度
func (q *Queue) Len() int64 {
	l, err := q.C.LLen(ctx, q.QueueName).Result()
	if err != nil {
		color.Redln("Len err: ", err)
	}
	return l
}

// Clear 清空队列
func (q *Queue) Clear() {
	_, err := q.C.Del(ctx, q.QueueName).Result()
	if err != nil {
		color.Redln("Clear err: ", err)
	}
}

// Exists 判断key是否存在
func Exists(key string) bool {
	_, err := Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return false
	}
	return true
}

// Set 设置key
func Set(key, value string, expiration time.Duration) {
	_, err := Redis.Set(ctx, key, value, expiration).Result()
	if err != nil {
		color.Redln("Set err: ", err)
	}
}

// Get 获取key
func Get(key string) string {
	value, err := Redis.Get(ctx, key).Result()
	if err != nil {
		color.Redln("Get err: ", err)
	}
	return value
}

// Del 删除key
func Del(key string) {
	_, err := Redis.Del(ctx, key).Result()
	if err != nil {
		color.Redln("Del err: ", err)
	}
}

// Expire 设置key过期时间
func Expire(key string, expiration time.Duration) {
	_, err := Redis.Expire(ctx, key, expiration).Result()
	if err != nil {
		color.Redln("Expire err: ", err)
	}
}

// HSet 设置hash key
func HSet(key, field, value string) {
	_, err := Redis.HSet(ctx, key, field, value).Result()
	if err != nil {
		color.Redln("HSet err: ", err)
	}
}
