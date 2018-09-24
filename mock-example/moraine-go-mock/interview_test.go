package morainegomock

import (
	"github.com/maraino/go-mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sfazilyesil/go-playground/mock-example"
	"github.com/sfazilyesil/go-playground/mock-example/moraine-go-mock/mocks"
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
			answerer.When("Answer", 1, mock.Any).
				Return("Ne aptalca bir soru!").
				Times(1)

			interview = interviewMaker.MakeInterview(answerer, questioner, 1)
		})

		It("should return interview", func() {
			Expect(interview).NotTo(BeNil())
		})

		It("should return interview with a content containing one qa pair", func() {
			Expect(interview.Content).To(HaveLen(1))
		})

		It("should return interview with the expected content", func() {
			Expect(interview.Content[0]).To(Equal(mockexample.QAPair{
				Question: "Kim daha çok bilir: çok gezen mi çok okuyan mı?",
				Answer: "Ne aptalca bir soru!",
			}))
		})

		Specify("Answerer answers only ones.", func() {
			answerer.Verify()
		})
	})

	Context("When Answer of mock answerer is called with expected param", func() {
		var (
			answerer = new(mocks.Answerer)
			answer   string
		)

		BeforeEach(func() {
			answerer.When("Answer", 1, "Doktor bu ne?").
				Return("Elinin körü").
				Times(1)

			answer = answerer.Answer(1, "Doktor bu ne?")
		})

		It("should return expected answer", func() {
			Expect(answer).To(Equal("Elinin körü"))
		})

		It("should be called once with correct params", func() {
			answerer.Verify()
		})
	})

	Context("When Answer of mock answerer is called with unexpected param", func() {
		var (
			answerer = new(mocks.Answerer)
			answer   string
		)

		BeforeEach(func() {
			answerer.When("Answer", 1, "Doktor bu ne?").
				Return("Elinin körü").
				Times(0)

			answerer.When("Answer", 1, "Sen de kimsin?").
				Return("Elinin körü").
				AtLeast(1)

			answer = answerer.Answer(1, "Sen de kimsin?")
		})

		It("should fail", func() {
			Expect(answerer.Verify()).To(BeTrue())
		})
		It("blah blah hhh", func() {
			Expect(1).To(Equal(1))
		})
	})

	Context("When MakeInterview of mock interview maker is called", func() {

	})
})
