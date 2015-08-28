package event

type Event interface {
	SetValue(interface{}) (bool, error)
	GetValue() (interface{}, error)
	GetLTime() uint64
}

type BaseEvent struct {
	ltime uint64
	value interface{}
}

func (e *BaseEvent) SetValue(v interface{}) (isSuccess bool, err error) {

	e.value = v;

	return true, nil
}

func (e *BaseEvent) GetValue() (interface{}, error) {

	return e.value, nil
}

func (e *BaseEvent) GetLTime() uint64 {

	return e.ltime
}

type NormalEvent struct {
	Event
	BaseEvent
}

type TaggedEvent struct {
	Event
	BaseEvent
	tags map[string]interface{}
}
