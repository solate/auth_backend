package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type PermissionCache struct {
	rdb *redis.Redis
}

func NewPermissionCache(rdb *redis.Redis) *PermissionCache {
	return &PermissionCache{rdb: rdb}
}

func (c *PermissionCache) GetUserPermissions(ctx context.Context, userId int64) ([]string, error) {
	key := fmt.Sprintf("user:perm:%d", userId)

	// 从缓存获取
	val, err := c.rdb.Get(key)
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var permissions []string
	err = json.Unmarshal([]byte(val), &permissions)
	return permissions, err
}

func (c *PermissionCache) SetUserPermissions(ctx context.Context, userId int64, permissions []string) error {
	key := fmt.Sprintf("user:perm:%d", userId)

	val, err := json.Marshal(permissions)
	if err != nil {
		return err
	}

	return c.rdb.Set(key, string(val))
}

func (c *PermissionCache) DeleteUserPermissions(ctx context.Context, userId int64) error {
	key := fmt.Sprintf("user:perm:%d", userId)
	_, err := c.rdb.Del(key)
	return err
}
