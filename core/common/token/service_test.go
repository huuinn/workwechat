package token

import (
	"testing"

	"github.com/go-redis/redis/v7"
)

func Test_service_Get(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	service := NewService("", "", client)

	t.Log(service.Get())
}

func Test_service_Delete(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	service := NewService("", "", client)

	t.Log(service.Delete(""))
}
