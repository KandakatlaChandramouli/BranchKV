package state_machine

import (
	"sync"
)

type StateMachine struct {
	state string
	mu    sync.RWMutex
}

func NewStateMachine() *StateMachine {

	return &StateMachine{
		state: "init",
	}
}

func (s *StateMachine) Transition(
	state string,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.state = state
}

func (s *StateMachine) State() string {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.state
}
