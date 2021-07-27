package main

func Power(base float32, exp int) float32 {
	if floatEquals(base, 0) {
		return 0
	}

	if exp == 0 {
		return 1
	}

	positiveExp := true
	if exp < 0 {
		exp = - exp
		positiveExp = false
	}

	oddExp := false
	if exp & 1 == 1 {
		oddExp = true
	}

	retResult := base
	for exp > 1 {
		retResult = retResult * retResult
		exp = exp >> 1
	}

	if oddExp {
		retResult = retResult * base
	}
	if !positiveExp {
		retResult = 1 / retResult
	}
	return retResult
}

func floatEquals(num1, num2 float32) bool {
	if (num1 - num2) < 0.0000001 && (num1 - num2) > -0.0000001 {
		return true
	}
	return false
}