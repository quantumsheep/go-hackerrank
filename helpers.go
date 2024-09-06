package hackerrank

func Paginate[Result any](f func(offset int) ([]Result, error)) ([]Result, error) {
	var results []Result

	for {
		res, err := f(len(results))
		if err != nil {
			return nil, err
		}
		if len(res) == 0 {
			break
		}

		results = append(results, res...)
	}

	return results, nil
}
