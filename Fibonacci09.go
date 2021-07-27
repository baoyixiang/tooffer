package main

func Fibonacci(n int) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

func FibonacciBetter(n int) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	firstNum := int64(0)
	secondNum := int64(1)
	thirdNum := firstNum + secondNum

	for i := 2; i <= n; i++ {
		thirdNum = firstNum + secondNum
		firstNum = secondNum
		secondNum = thirdNum
	}

	return thirdNum
}
