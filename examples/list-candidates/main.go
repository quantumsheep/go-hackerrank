package main

import (
	"fmt"

	"github.com/quantumsheep/go-hackerrank"
)

func main() {
	api := hackerrank.NewClient("<YOUR API KEY>")

	candidates, err := hackerrank.Paginate(func(offset int) ([]hackerrank.TestCandidateIndex, error) {
		res, err := api.V3GetTestByTestIdCandidates(&hackerrank.V3GetTestByTestIdCandidatesParams{
			TestId: "1706056",
			Limit:  100,
			Offset: offset,
		})
		return res.Data, err
	})
	if err != nil {
		panic(err)
	}

	for _, candidate := range candidates {
		fmt.Printf("%s %f\n", candidate.Email, candidate.Score)
	}
}
