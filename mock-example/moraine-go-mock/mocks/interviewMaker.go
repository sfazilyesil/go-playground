package mocks

import (
	"github.com/maraino/go-mock"
	"github.com/sfazilyesil/go-playground/mock-example"
)

type InterviewMaker struct {
	mock.Mock
}

func (m *InterviewMaker) MakeInterview(a mockexample.Answerer, q mockexample.Questioner, length int) *mockexample.Interview {
	args := m.Called(a, q, length)
	return args.Get(0).(*mockexample.Interview)
}
