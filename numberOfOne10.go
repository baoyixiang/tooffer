package main

func NumberOf1(n int) int {
	num := 0
	addNum := 0
	if n < 0 {
		n = -n
		addNum = 1
	}

	for n != 0 {
		if n & 1 == 1 {
			num ++
		}
		n = n >> 1
	}

	return num + addNum
}

func NumberOf1Better(n int) int {
	tmpCnt := 1
	num := 0
	for tmpCnt != 0 {
		if n & tmpCnt == 1 {
			num ++
		}
		tmpCnt = tmpCnt << 1
	}
	return num
}

func NumberOf1Best(n int) int {
	num := 0
	for n != 0 {
		num ++
		n = n & (n - 1)
	}
	return num
}