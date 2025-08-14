package redis

import (
	flam "github.com/happyhippyhippo/flam"
)

type ConnectionCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (Connection, error)
}
