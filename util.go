package main

func unwrap[T any](output T, err error) T {
	if err != nil {
		panic(err)
	}
	return output
}
