package mockexample

type Questioner interface {
	Ask() string
}

type stupidQuestioner struct{}

func NewStupidQuestioner() *stupidQuestioner {
	return &stupidQuestioner{}
}

func (q *stupidQuestioner) Ask() string {
	return "Kim daha çok bilir: çok gezen mi çok okuyan mı?"
}
