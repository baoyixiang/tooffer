package main

func replaceBlank(replaceBytes []byte) {
	blankCnt := 0
	bytesLen := len(replaceBytes)
	for _, ch := range replaceBytes {
		// 空格得ascii码为 32
		if ch == 32 {
			blankCnt ++
		}
	}

	if blankCnt == 0 {
		return
	}

	// 扩容，备好新空间
	for i := 0; i < blankCnt; i++ {
		replaceBytes = append(replaceBytes, 32, 32)
	}

	newIndex := len(replaceBytes) - 1
	for i := bytesLen - 1; i >= 0 ; i-- {
		if replaceBytes[i] == ' ' {
			replaceBytes[newIndex] = 51
			newIndex--
			replaceBytes[newIndex] = 50
			newIndex--
			replaceBytes[newIndex] = 37
			newIndex--

			continue
		}

		replaceBytes[newIndex] = replaceBytes[i]
		newIndex--
	}

	println(replaceBytes)
}