package main

import (
	"os"
    "fmt"
)

type astNode struct {
	parent *astNode
	ttype string
	termsym bool
	contents string
	children []*astNode
}

func advanceToken() {
	tokenQueue = tokenQueue[1:]
}

func checkToken(character string, err string) {
	if tokenQueue[0].Literal != character {
		os.Exit(0)
	} else{
      fmt.Print(err)      
    }
	advanceToken()
}

func parse() astNode{
	root := astNode{}

switch tokenQueue[0].Type {
	// 1
	case "integer":
		initTreeNd(&root, true) //dequeue
	// 2
	case "minus":
		initTreeNd(&root, false) //dequeue
		checkToken("(", "Error in parse: lparen not found in case minus") //dequeue
		child1 := parse()
		checkToken(",", "Error in parse: comma not found in case minus")
		child2 := parse()
		checkToken(")", "Error in parse: rparen not found in case minus")
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
	// 3
	case "iszero":
		initTreeNd(&root, false) //dequeue
		checkToken("(", "Error in parse: lparen not found in case iszero") //dequeue
		child1 := parse()
		checkToken(")", "Error in parse: rparen not found in case iszero")
		root.children = append(root.children, &child1)
	case "if":
		initTreeNd(&root, false) //dequeue
		child1 := parse()
		checkToken("then", "Error in parse: then not found in case if")
		child2 := parse()
		checkToken(")", "Error in parse: else not found in case if")
		child3 := parse()
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)
	// 5
	case "identifier":
		initTreeNd(&root, true) //dequeue
	// 6
	case "let":
		initTreeNd(&root, false) //dequeue
		child1 := parse()
		checkToken("=", "Error in parse: equal not found in case let")
		child2 := parse()
		checkToken("in", "Error in parse: in not found in case let")
		child3 := parse()
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)

    }
	return root
}

func initTreeNd(nd *astNode, isterm bool){
	nd.termsym = isterm
	nd.children = make([]*astNode, 0, 5)

	nd.ttype = tokenQueue[0].Type
	nd.contents = tokenQueue[0].Literal
	advanceToken() // dequeue
}

