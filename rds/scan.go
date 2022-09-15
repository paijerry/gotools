package rds

import (
	"github.com/gomodule/redigo/redis"
)

// 參考：https://stackoverflow.com/questions/31498296/how-to-scan-keys-from-redis-using-golang-using-scan-not-keys
// 取代 redis KEYS command
func scanKeys(conn redis.Conn, match string) (keys []string, err error) {
	iter := 0 // scan 游標

	for {
		// 開始游標為 0
		var arr []interface{}
		var err2 error
		if match != "" {
			if arr, err2 = redis.Values(conn.Do("SCAN", iter, "MATCH", match)); err2 != nil {
				return keys, err2
			}
		} else {
			if arr, err2 = redis.Values(conn.Do("SCAN", iter)); err2 != nil {
				return keys, err2
			}
		}

		iter, _ = redis.Int(arr[0], nil)
		ks, _ := redis.Strings(arr[1], nil)
		keys = append(keys, ks...)

		// 游標返回 0 則停止查詢
		if iter == 0 {
			break
		}
	}

	return keys, nil
}
