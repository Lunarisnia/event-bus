package ebus

type Signal struct {
	Name string
	bus  *EBus
}

type SignalHandler func()

func newSignal(bus *EBus, name string) (*Signal, error) {
	return &Signal{
		bus:  bus,
		Name: name,
	}, nil
}

func (s *Signal) Emit() {
	s.bus.dispatch(s)
}

func (s *Signal) Connect(handler SignalHandler) {
	s.bus.subscribe(s, handler)
}
