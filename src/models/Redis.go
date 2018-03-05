package models

func SetRedisKey(key string, result string) {
	RedisDB.Set(key, result, 0)
}

func GetRedisKey(key string) (string, error) {
	return RedisDB.Get(key).Result()
}

func DeleteRedisKey(key string) {
	RedisDB.Del(key)
}
