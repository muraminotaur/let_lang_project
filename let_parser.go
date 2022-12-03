package main

import (
	"os"
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

func checkToken(character string) {
	if tokenQueue[0] != character {
		os.Exit()
	}
	advanceToken()
}

func parse() astNode{
	root := astNode{}

switch tokenQueue[0].tokenType
	// 1
	case "integer":
		initTreeNd(&root, true) //dequeue
	// 2
	case "minus":
		initTreeNd(&root, false) //dequeue
		checkToken(tokenList, "(", "Error in parse: lparen not found in case minus") //dequeue
		child1 := parse()
		checkToken(tokenList, ",", "Error in parse: comma not found in case minus")
		child2 := parse()
		checkToken(tokenList, ")", "Error in parse: rparen not found in case minus")
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
	// 3
	case "iszero":
		initTreeNd(&root, false) //dequeue
		checkToken(tokenList, "(", "Error in parse: lparen not found in case iszero") //dequeue
		child1 := parse()
		checkToken(tokenList, ")", "Error in parse: rparen not found in case iszero")
		root.children = append(root.children, &child1)
	case "if":
		initTreeNd(&root, false) //dequeue
		child1 := parse()
		checkToken(tokenList, "then", "Error in parse: then not found in case if")
		child2 := parse()
		checkToken(tokenList, ")", "Error in parse: else not found in case if")
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
		checkToken(tokenList, "=", "Error in parse: equal not found in case let")
		child2 := parse()
		checkToken(tokenList, "in", "Error in parse: in not found in case let")
		child3 := parse()
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)

	return root
}

func initTreeNd(nd *astNode, isterm bool){
	nd.termsym = isterm
	nd.children = make([]*astNode, 0, 5)
	nd.ttype = tokenList[0].tokenType
	nd.contents = tokenList[0].tokenValue
	advanceToken() // dequeue
}

