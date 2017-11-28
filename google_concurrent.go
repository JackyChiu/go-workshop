// The comments in these files for the presentation tool used to generate the slides.
// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

// START1 OMIT
func Google(query string) []Result {
	var results []Result
	resultsChan := make(chan Result)

	go func() { resultsChan <- Web(query) }()
	go func() { resultsChan <- Image(query) }()
	go func() { resultsChan <- Video(query) }()

	for i := 0; i < 3; i++ {
		results = append(results, <-resultsChan)
	}

	return results
}

// STOP1 OMIT

// START2 OMIT
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result // HL

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// STOP2 OMIT

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang") // HL
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)

}
