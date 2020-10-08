package main

import (
	"fmt"

	"github.com/tonyoditanto/latihan_go/bagian2/package_1"

	"github.com/tonyoditanto/latihan_go/bagian2/package_2"
)

func main() {

	package_1.PrintPackage1()
	package_2.PrintPackage2()

	var x, y int

	nama1, nama2, nama3 := "tony", "varian", "yoditanto"

	x = 10
	y = 5

	result1 := package_1.Add(x, y)
	fmt.Println(fmt.Sprintf("Hasil penjumlahannya adalah %d", result1))

	fmt.Println(package_1.CustomPrint(nama1))
	fmt.Println(package_1.CustomPrint(nama2))
	fmt.Println(package_1.CustomPrint(nama3))

	result2, flag2 := package_1.AddDoubleReturn(x, y)

	fmt.Println(fmt.Sprintf("Hasil penjumlahannya adalah %d , %t", result2, flag2))

	package_1.PrintPackage1()
	package_2.PrintPackage2()
}
