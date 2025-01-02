package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	words:=strings.Fields(corpus)

	query:=os.Args[1:]

queries:
	for _,q:=range query{
		for i,w:=range words{
			// if strings.ToLower(q)==strings.ToLower(w){
			if strings.EqualFold(q, w) {
				fmt.Printf("#%-2d: %q\n",i+1,w)
				break queries
			}
		}
	}
}