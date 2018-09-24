package testify

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sfazilyesil/go-playground/mock-example"
	"github.com/stretchr/testify/mock"
	"github.com/sfazilyesil/go-playground/mock-example/testify-mock/mocks"
)

var _ = Describe("MockExample", func() {

	Context("When MakeInterview of SimpleInterviewMaker is called", func() {
		var (
			answerer = new(mocks.Answerer)

			questioner     = mockexample.NewStupidQuestioner()
			interviewMaker = mockexample.SimpleInterviewMaker{}

			interview *mockexample.Interview
		)

		BeforeEach(func() {
			answerer.On("Answer", mock.Anything).
				Return("What a stupid question.").
				Once()

			interview = interviewMaker.MakeInterview(answerer, questioner, 1)
		})

		It("should return interview", func() {
			Expect(interview).NotTo(BeNil())
		})

		It("should return interview with a content containing one qa pair", func() {
			Expect(interview.Content).To(HaveLen(1))
		})

		Specify("Answerer answers only ones.", func() {
			answerer.AssertExpectations(GinkgoT())
		})
	})

	Context("When Answer of mock answerer is called with expected param", func() {
		var (
			answerer = new(mocks.Answerer)
			answer   string
		)

		BeforeEach(func() {
			answerer.On("Answer", "Doktor bu ne?").
				Return("Elinin körü").
				Once()

			answer = answerer.Answer(1, "Doktor bu ne?")
		})

		It("should return expected answer", func() {
			Expect(answer).To(Equal("Elinin körü"))
		})

		It("should be called once with correct params", func() {
			answerer.AssertExpectations(GinkgoT())
		})
	})

	Context("When Answer of mock answerer is called with unexpected param", func() {
		var (
			answerer = new(mocks.Answerer)
			answer   string
		)

		BeforeEach(func() {
			answerer.On("Answer", "Doktor bu ne?").
				Return("Elinin körü")

			answerer.On("Answer", "Sen de kimsin?").
				Return("Elinin körü")

			answer = answerer.Answer(1, "Sen de kimsin?")
		})

		It("should fail", func() {
			answerer.AssertExpectations(GinkgoT())
		})
	})

	Context("When MakeInterview of mock interview maker is called", func() {

	})
})
