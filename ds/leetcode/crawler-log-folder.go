package main

import "fmt"

func crawlerLogfolder(logs []string) (val int) {
	stack := []string{}
	for _, v := range logs {
		switch v {
		case "../":
			if len(stack) == 0 {
				continue
			}
			stack = stack[:len(stack)-1]
		case "./":
			continue
		default:
			stack = append(stack, v)
		}
	}
	return len(stack)
}

func main() {
	fmt.Println(crawlerLogfolder([]string{"d1/", "d2/", "../"}))
}
