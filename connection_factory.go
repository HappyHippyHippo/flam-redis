package redis

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
)

type connectionFactory flam.Factory[Connection]

type connectionFactoryArgs struct {
	dig.In

	Creators      []ConnectionCreator `group:"flam.redis.connection.creator"`
	FactoryConfig flam.FactoryConfig
}

func newConnectionFactory(
	args connectionFactoryArgs,
) (connectionFactory, error) {
	var creators []flam.ResourceCreator[Connection]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathConnections,
		args.FactoryConfig,
		nil,
	)
}
