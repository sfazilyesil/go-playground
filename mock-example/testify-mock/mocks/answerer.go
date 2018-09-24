package mocks

import "github.com/stretchr/testify/mock"

type Answerer struct {
	mock.Mock
}

func (m *Answerer) Answer(no int, question string) string {
	args := m.Called(no, question)
	return args.String(0)
}
