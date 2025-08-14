package redis

type Facade interface {
	HasConnection(id string) bool
	ListConnections() []string
	GetConnection(id string) (Connection, error)
	AddConnection(id string, connection Connection) error
}

type facade struct {
	connectionFactory connectionFactory
}

func newFacade(
	connectionFactory connectionFactory,
) *facade {
	return &facade{
		connectionFactory: connectionFactory,
	}
}

func (facade facade) HasConnection(
	id string,
) bool {
	return facade.connectionFactory.Has(id)
}

func (facade facade) ListConnections() []string {
	return facade.connectionFactory.List()
}

func (facade facade) GetConnection(
	id string,
) (Connection, error) {
	return facade.connectionFactory.Get(id)
}

func (facade facade) AddConnection(
	id string,
	connection Connection,
) error {
	return facade.connectionFactory.Add(id, connection)
}
