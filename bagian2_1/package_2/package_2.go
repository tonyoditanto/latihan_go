package package_2

import (
	"fmt"

	"github.com/tonyoditanto/latihan_go/bagian2_1/package_1"
)

func PrintPackage2() {
	fmt.Println("Print dari package_2")
}

func ImportDariPackage1() {
	package_1.PrintPackage1()
}
