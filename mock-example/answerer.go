package mockexample


type Answerer interface {
	Answer(no int, question string) string
}
