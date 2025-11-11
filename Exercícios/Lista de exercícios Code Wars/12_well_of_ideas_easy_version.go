// For every good kata idea there seem to be quite a few bad ones!

// In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. 
// If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. 
// If there are no good ideas, as is often the case, return 'Fail!'.

//go:build ignore

package main

import "fmt"

func Well(x []string) string {
	countGoodIdeas := 0

	for _, idea := range(x){
		if idea == "good" {
			countGoodIdeas += 1
		}
	}
	if countGoodIdeas == 0 {
		return "Fail!"
	} else if countGoodIdeas > 2 {
		return "I smell a series!"
	} 
	return "Publish!"
}

func main() {
	fmt.Println(Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}))
	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}))
	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"}))
}