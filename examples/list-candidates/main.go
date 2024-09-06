package main

import (
	"context"
	"fmt"
	"os"

	"github.com/quantumsheep/go-hackerrank"
)

var (
	hackerRankApiKey = os.Getenv("HACKERRANK_API_KEY")
)

func main() {
	api := hackerrank.NewClient(hackerRankApiKey)

	candidates, err := hackerrank.Paginate(func(offset int) ([]hackerrank.TestCandidateIndex, error) {
		res, err := api.V3GetTestByTestIdCandidates(context.Background(), &hackerrank.V3GetTestByTestIdCandidatesParams{
			TestId: "<YOUR TEST ID>",
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
