package main

import(
	"strconv"
)

type Binding struct {
	varname string
	value int
}

func Lookup(varname string, e []Binding) int {
	for i := range e{
		if e[i].varname == varname{
			return e[i].value
		}            
	}
    return 0
}

func evaluate(localRoot astNode, e []Binding) int{
	switch localRoot.contents {
		case "let":
			varname := localRoot.children[0].contents
			exp1Val := evaluate(*localRoot.children[1], e)
			var newBindingList []Binding = []Binding{Binding{varname, exp1Val}}
			e = append(newBindingList, e...)
			return evaluate(*localRoot.children[2], e)
		case "identifier":
			return Lookup(localRoot.contents, e)
		case "if":
			zerocheck := evaluate(*localRoot.children[0], e)
			if localRoot.children[0].ttype == "iszero" && zerocheck == 0{
				return evaluate(*localRoot.children[1], e)
			} else{
				return evaluate(*localRoot.children[2], e)
			}
		case "iszero":
			return evaluate(*localRoot.children[0], e)
		case "minus":
			return (evaluate(*localRoot.children[0], e) - evaluate(*localRoot.children[1], e))
		case "integer":
            exp1Val, _ := strconv.Atoi(localRoot.contents)
			return exp1Val
    }
    return 0
}
