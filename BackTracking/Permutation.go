package main

import "fmt"

func main() {
	sample := "roc"
	sampleRune := []rune(sample)
	generatePermutation(sampleRune, 0, len(sampleRune)-1)
}

func generatePermutation(sampleRune []rune, left, right int) {
	if left == right {
		fmt.Println(string(sampleRune))
	} else {
		for i := left; i <= right; i++ {
			//do
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
			//recur
			generatePermutation(sampleRune, left+1, right)
			//undo
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
		}
	}
}
