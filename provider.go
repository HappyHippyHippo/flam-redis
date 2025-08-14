package redis

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
)

type provider struct{}

func NewProvider() flam.Provider {
	return &provider{}
}

func (provider) Id() string {
	return providerId
}

func (provider) Register(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	var e error
	provide := func(constructor any, opts ...dig.ProvideOption) bool {
		e = container.Provide(constructor, opts...)
		return e == nil
	}

	_ = provide(newDefaultConnectionCreator, dig.Group(ConnectionCreatorGroup)) &&
		provide(newConnectionFactory) &&
		provide(newFacade)

	return e
}

func (provider) Boot(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		configFacade config.Facade,
	) error {
		DefaultHost = configFacade.String(PathDefaultHost, DefaultHost)
		DefaultPort = configFacade.Int(PathDefaultPort, DefaultPort)
		DefaultDatabase = configFacade.Int(PathDefaultDatabase, DefaultDatabase)

		return nil
	})
}

func (provider) Close(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		connectionFactory connectionFactory,
	) error {
		if e := connectionFactory.Close(); e != nil {
			return e
		}

		return nil
	})
}
