package mockexample

import "fmt"

type QAPair struct {
	Question string
	Answer   string
}

type Interview struct {
	Content []QAPair
}

type InterviewMaker interface {
	MakeInterview(answerer Answerer, questioner Questioner, length int) *Interview
}

type SimpleInterviewMaker struct{}

func (s *SimpleInterviewMaker) MakeInterview(a Answerer, q Questioner, length int) *Interview {
	qaPairs := make([]QAPair, length)

	for i := 0; i < length; i++ {
		question := q.Ask()
		fmt.Println("Question:", question)

		answer := a.Answer(i+1, question)
		fmt.Println("Answer: ", answer)

		qaPairs[i] = QAPair{Question: question, Answer: answer}
	}

	return &Interview{Content: qaPairs}
}
