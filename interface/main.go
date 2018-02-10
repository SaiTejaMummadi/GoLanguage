package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"xen", "sam", "andy", "lilly", "sony"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	s:=[]string{"wen","oka","qwd","aew"}
	sort.Strings(s)
	fmt.Println(s)
	n:=[]int{1,9,6,3,2,6}
	sort.Ints(n)
	fmt.Println(n)
}
