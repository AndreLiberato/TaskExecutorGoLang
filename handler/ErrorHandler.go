package handler

// Check é uma função de auxílio para chegar erros
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
