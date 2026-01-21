
1. Take a char , and tell whether the char is vovel, consonent or a special char


2. The same program with empty switch, and also check whether a char is a vovel or a consonent or a special char

char := 'a'

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is vovel")
	default:
		println(string(char), "either consonent or a special char")
	}

