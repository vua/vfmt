package internal

import (
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`@\[(.*?)\]`)

func Parse(text string) string {
	container := make([]string, 0)

	prev := 0
	for _, idx := range r.FindAllStringSubmatchIndex(text, -1) {
		if prev < idx[0] {
			container = append(container, text[prev:idx[0]])
		}
		// text[idx[0]:idx[1]]:@[***:***]
		// text[idx[2]:idx[3]]:***:***
		token := text[idx[2]:idx[3]]
		if !strings.Contains(token, "::") {
			container = append(container, token)
		} else {
			// seg[0]:console output
			// seg[1]:style
			seg := strings.Split(token, "::")
			if len(seg) != 2 {
				panic(Err(SyntaxERR))
			}
			token:=seg[0]
			for _,style:=range strings.Split(seg[1],"|"){
				if style=="" {
					continue
				}
				temp,err:=ApplyStyle(token,style)
				if err!=nil{
					panic(err)
				}
				token=temp
			}
			container = append(container, token)
		}
		prev = idx[1]
	}
	//tail
	if prev < len(text) {
		container = append(container, text[prev:])
	}

	return strings.Join(container,"")
}
