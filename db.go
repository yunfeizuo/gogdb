package main

import (
	"fmt"
	"reflect"
)

/*

  type Query {
    labels: [Label]
    myQuestions(userId: Int): [Question]
    questionsOpen2Me(userId: Int): [Question]
  }

  type Mutation {
    addQuestion(userId: Int, content: String, labelId: Int, images: [String], makePublic: Boolean, one2one: Boolean): Int
    offer(questionId: Int!, userId: Int!, price: Int!): String
    chat(questionId: Int!, userId: Int!, textMsg: String, voiceMsg: String, imgMsg: String): DialogMessage
  }

// query

root {									---> 	Vertex {

	myQuestions(id) {
		id
		text
		images(small)
		answeredBy {
			id
			name
			email
			wechat {
				openid
				profile
			}
		}
	}
}
*/

type Question struct {
	Content string
	Length  int
}

func main() {

	var root = QueryEdge{}
	root.Name = "root"
	var i int
	contentProp := QueryEdge{Name: "Content", Target: QueryVertex{Kind: Scalar}, Type: reflect.TypeOf(root.Name)}
	lengthProp := QueryEdge{Name: "Length", Target: QueryVertex{Kind: Scalar}, Type: reflect.TypeOf(i)}
	edges := []QueryEdge{contentProp, lengthProp}
	root.Target = QueryVertex{Id: 23, Edges: edges, Kind: Object}
	root.Type = reflect.TypeOf(Question{})

	v := ResolveEdge(1, root)
	fmt.Println("%v", v)
}
