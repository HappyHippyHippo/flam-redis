package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	flam "github.com/happyhippyhippo/flam"
)

type defaultConnectionCreator struct{}

func newDefaultConnectionCreator() *defaultConnectionCreator {
	return &defaultConnectionCreator{}
}

func (defaultConnectionCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == ConnectionDriverDefault
}

func (defaultConnectionCreator) Create(
	config flam.Bag,
) (Connection, error) {
	address := fmt.Sprintf(
		"%s:%d",
		config.String("host", DefaultHost),
		config.Int("port", DefaultPort),
	)

	return redis.NewClient(&redis.Options{
		Addr: address,
		DB:   config.Int("db", DefaultDatabase),
	}), nil
}
