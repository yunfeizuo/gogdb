package main

type Q interface {
	// myQuestions(id int) []Question
	// askQuestion(QuestionInput) Question
}

// type Question struct {
//     GetImages(size int) []string
// }

type Vertex interface {
	GetId() string
}

type Node struct {
	Id string
}

func (n *Node) GetId() string {
	return n.Id
}
