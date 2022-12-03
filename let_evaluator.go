package main

import(
	"strconv"
)

type Binding struct {
	varname string
	value int
}

func Lookup(varname string, e []Binding string) Binding.value {
	for i := range e{
		if e[i].varname == varname{
			return e[i].value
		}
	}
}

func evaluate(localRoot astNode, e []Binding string) int{
	switch localRoot.contents
		case "let":
			varname := localRoot.children[0].contents
			exp1Val := evaluate(*localRoot.children[1], e)
			var newBindingList []Binding = []Binding{Binding{varname, exp1Val}}
			e = append(newBindingList, e...)
			return evaluate(*localRoot.children[2], e)
		case "identifier":
			varname := localRoot.children[0].contents
			return Lookup(localRoot.contents, e)
		case "if":
			zerocheck := evaluate(*localRoot.children[0], e)
			if localRoot.children[0].ttype == "iszero" && zerocheck == 0{
				return evaluate(*localRoot.children[1], e)
			}
			else{
				return evaluate(*localRoot.children[2], e)
			}
		case "iszero":
			return evaluate(*localRoot.children[0], e)
		case "minus":
			return (evaluate(*localRoot.children[0], e) - evaluate(*localRoot.children[1], e))
		case "integer":
			return strconv.Atoi(localRoot.contents)
}