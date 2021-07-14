package main

func main() {
	a := "abb"
	bytes := []byte(a)
	replaceBlank(bytes)
	println(string(bytes[:10]))
}
