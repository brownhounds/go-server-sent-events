package clients

import (
	"github.com/redis/go-redis/v9"
)

type connection struct {
	conn *redis.Client
}

var Redis *connection

func init() {
	Redis = &connection{
		conn: redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}),
	}
}

func (c *connection) Client() *redis.Client {
	return c.conn
}
