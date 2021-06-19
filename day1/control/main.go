package main

func exampleIf() {
	age := 11
	if age == 19 {
		println("age == 19")
	} else if age > 19 {
		println("age >19")
	} else {
		println("unknown")
	}
}

func exampleFor() {
	for i := 1; i <= 10; i++ {
		println(i)
	}
}

func exampleSwitch() {
	switch n := 6; n {
	case 1, 3, 5, 7, 9:
		println("单数")
	case 2, 4, 6, 8:
		println("双数")
	default:
		println(n)
	}
}

func main() {
	exampleSwitch()
}
