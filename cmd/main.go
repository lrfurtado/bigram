package main

import "github.com/lrfurtado/bigram"
import "flag"
import "fmt"

func main() {
	input := flag.String("input", "", "string to be broken into bigrams ")

	flag.Parse()

	if *input == "" {
		panic("Please provide an inputi using -input switch")
	}

	fmt.Println(bigram.Parse(*input))
	fmt.Println(bigram.Histogram(*input))

}
