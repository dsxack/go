package sugar

// Must panics if err is not nil, otherwise returns t.
// It is useful for simplifying error handling.
// Example:
//
//	file, err := os.Open("file.txt")
//	if err != nil {
//	  panic(err)
//	}
//	defer file.Close()
//	// use file
//
// can be simplified to:
//
//	file := sugar.Must(os.Open("file.txt")).(*os.File)
//	defer file.Close()
//	// use file
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
