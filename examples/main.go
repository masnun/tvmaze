package main

import "github.com/masnun/tvmaze"
import "fmt"

func main() {
	results := tvmaze.Search("doctor who")
	for _, result := range results {
		fmt.Println(result.Show.Name, "-", result.Score)
	}
}
