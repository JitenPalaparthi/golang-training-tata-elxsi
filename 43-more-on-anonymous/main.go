package main

func main() {

	r1 := calc(10, 20, func(i1, i2 int) int {
		return i1 + i2
	})

	println(r1)

	r2 := calc(20, 10, sub)
	println(r2)

	//handler("user/add", "user:1")

}

func calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func sub(i, j int) int {
	return i - j
}

func handler(pattern string, data any, handle func() string) string {
	return handle()
}

// tasks
// calc function create this like
/*
func calc(a, b any, fn func(any, any) any) any {
	return fn(a, b)
}

*/
