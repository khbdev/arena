package redis

import (
	"encoding/json"
	"errors"
	"time"
	"user-service/internal/config"

	"github.com/redis/go-redis/v9"
)

// ======================
//       WRITE-THROUGH
// ======================
func Set(key string, value interface{}, ttl time.Duration) error {
	// value ni JSON ga aylantiramiz
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Redis'ga yozish
	return config.RedisClient.Set(config.Ctx, key, data, ttl).Err()
}

// ======================
//       READ-THROUGH
// ======================
func Get(key string, dest interface{}, fallback func() (interface{}, error), ttl time.Duration) error {
	// 1️⃣ Redis'dan olish
	data, err := config.RedisClient.Get(config.Ctx, key).Result()
	if err == nil {
		// topildi → unmarshal qilamiz
		return json.Unmarshal([]byte(data), dest)
	}

	// 2️⃣ Redis'da topilmasa
	if !errors.Is(err, redis.Nil) && err != nil {
		return err
	}

	// fallback (DB dan olish)
	dbData, err := fallback()
	if err != nil {
		return err
	}

	// Redis'ga yozamiz (write-through)
	if err := Set(key, dbData, ttl); err != nil {
		return err
	}

	// dest ga copy
	bytes, _ := json.Marshal(dbData)
	return json.Unmarshal(bytes, dest)
}

// ======================
//        DELETE
// ======================
func Delete(key string) error {
	return config.RedisClient.Del(config.Ctx, key).Err()
}
