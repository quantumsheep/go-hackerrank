# HackerRank API client for Go

This is a Go client for the HackerRank API.

It works by parsing [HackerRank's OpenAPI specification](https://www.hackerrank.com/apidoc) and generating Go code to interact with the API. A custom program needed to be written to generate the Go code because the HackerRank's OpenAPI specification is not compatible with existing Go OpenAPI client generators.

See the [examples](examples) directory for usage examples.
