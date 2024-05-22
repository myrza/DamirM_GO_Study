package main

func fib(i int, a int, b int) {
	if i > 22 {
		return
	}
	i += 1

	println(i, a)
	a, b = b, a+b

	// рекурсия
	fib(i, a, b)
}
func main() {
	fib(0, 0, 1)
}
