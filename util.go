package zkteco

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// GetRedisConn gets the Redis connection.
func GetRedisConn(redisAddr, redisPassword string) (c redis.Conn, err error) {
	pongStr := ""

	if c, err = redis.Dial("tcp", redisAddr); err != nil {
		goto end
	}

	if len(redisPassword) != 0 {
		if _, err = c.Do("AUTH", redisPassword); err != nil {
			goto end
		}
	}

	if pongStr, err = redis.String(c.Do("PING")); err != nil {
		goto end
	}

	if pongStr != "PONG" {
		err = fmt.Errorf("Redis PING != PONG(%v)", pongStr)
		goto end
	}
end:
	if err != nil {
		debugPrintf("GetRedisConn() error: %v\n", err)
		return c, err
	}
	return c, nil
}
