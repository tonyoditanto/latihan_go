package package_1

import "fmt"

func PrintPackage1() {
	fmt.Println("Print dari package_1")
}

func CustomPrint(str string) string {
	return fmt.Sprintf("Halo, nama saya %s", str)
}

func add(x int, y int) int {
	return x + y
}

func Add(x int, y int) int {
	return x + y
}

func AddDoubleReturn(x int, y int) (int, bool) {
	return x + y, true
}
