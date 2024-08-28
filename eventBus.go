package ebus

type EBus struct {
	eventCloset map[*Signal][]SignalHandler
}

func NewEBus() (*EBus, error) {
	return &EBus{
		eventCloset: make(map[*Signal][]SignalHandler),
	}, nil
}

func (e *EBus) CreateSignal(name string) (*Signal, error) {
	return newSignal(e, name)
}

func (e *EBus) dispatch(signal *Signal) {
	for _, handler := range e.eventCloset[signal] {
		handler()
	}
}

func (e *EBus) subscribe(signal *Signal, handler SignalHandler) {
	e.eventCloset[signal] = append(e.eventCloset[signal], handler)
}
