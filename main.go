// C00407978

package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var tok Token
var tokenQueue []Token



func printAST(root astNode) {
	root.printTreeWork(0)
}

func (nd astNode) printTreeWork(indentLevel int) {
	outString := ""
	for i := 0; i < indentLevel; i++ {
		outString += "    "
	}
	fmt.Printf("%s", outString)
	fmt.Printf("%s \n", nd.contents)
	if nd.termsym == false && len(nd.children) > 0 {
		for i := 0; i < len(nd.children); i++ {
			nd.children[i].printTreeWork(indentLevel+1)
		}
	}
}

func main(){

	tokenQueue = []Token{}
	file, error := os.Open("./example.txt")
	if error != nil{
		log.Fatal(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := let_scanner.scanner.New(line)
		
		for tok := s.NextToken(); tok.Type != token.EOF; tok = s.NextToken() {
			tokenQueue = append(tokenQueue, tok)
		}
	}
	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	tree := parse()

	printAST(tree)
}
