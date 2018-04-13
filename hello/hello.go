
package main

const SIX = 6

func main() {
	println("Hello world from Go!")
	println("The answer is:", calculateAnswer())
	println("5 ** 2 =", square(5))
	println("3 + 12 =", add(3, 12))
}

func calculateAnswer() int {
	seven := 7
	return SIX * seven
}

func square(n int) int {
	return n * n
}

func add(a, b int) int {
	return a + b
}
