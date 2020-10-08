package main

import "fmt"

type Geo struct {
	Lat  float64
	Long float64
}

func main() {
	fmt.Println("Cetak Array: ")
	array()

	fmt.Println("\nCetak Slice: ")
	slices()

	fmt.Println("\nCetak Map: ")
	mapExample()
}

func array() {
	//array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {

	names := []string{
		"Tony",
		"Varian",
		"Yoditanto",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "--"
	fmt.Println(a, b)
	fmt.Println(names)
}

func mapExample() {
	mapGeo := make(map[string]interface{})
	//mapPeserta := make(map[string][]string)
	mapGeo["Malang"] = Geo{
		Lat:  7.7956,
		Long: 110.3695,
	}

	mapGeo["Silicon Valley"] = Geo{
		Lat:  6.2088,
		Long: 106.8456,
	}

	//print map key and value
	for i, v := range mapGeo {
		fmt.Println(fmt.Sprintf("Koordinat %s  : %v", i, v))
	}
}
