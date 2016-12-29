package manta

const (
	entityCreated = iota
	entityUpdated
	entityDeleted
	entityEntered
	entityLeft
)

type entity struct {
	index  int32
	serial int32
	class  *class
	values []interface{}
}

func newEntity(index, serial int32, class *class, values []interface{}) *entity {
	return &entity{
		index:  index,
		serial: serial,
		class:  class,
		values: values,
	}
}
