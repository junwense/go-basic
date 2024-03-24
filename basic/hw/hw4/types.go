package hw4

import (
	"context"
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

var ErrorKeyNotFound = errors.New("key未找到")
var ErrorCacheHasClose = errors.New("缓存已经关闭")

type Cache interface {
	Get(ctx context.Context, key string) (any, error)

	Set(ctx context.Context, key string, val any, timeout time.Duration) error
}

type Close interface {
	Close(ctx context.Context) error
}

type LocalCache struct {
	// 并发安全 ，
	// 不支持范型，
	// 支持争对于key设置过期时间
	// api丰富
	// 支持过期删除的事件
	// 不支持close方法,只能把cache里所有元素刷新掉,但是这样也解决了close重复调用的问题
	// 删除元素方式包含2种，一种是get的时候惰性删除，一种是自己定时删除
	c     cache.Cache
	close bool
}

func (l *LocalCache) Close(ctx context.Context) error {
	l.close = true
	l.c.Flush()
	return nil
}

func (l *LocalCache) Get(ctx context.Context, key string) (any, error) {

	if l.close {
		return nil, ErrorCacheHasClose
	}

	get, b := l.c.Get(key)
	if !b {
		return nil, ErrorKeyNotFound
	}

	return get, nil
}

func (l *LocalCache) Set(ctx context.Context, key string, val any, timeout time.Duration) error {
	if l.close {
		return ErrorCacheHasClose
	}
	l.c.Set(key, val, timeout)
	return nil
}

type RedisCache struct {
}

func (r *RedisCache) Get(ctx context.Context, key string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) Set(ctx context.Context, key string, val any, timeout time.Duration) error {
	//TODO implement me
	panic("implement me")
}
