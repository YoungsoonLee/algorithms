package main

import "fmt"

type address struct {
	detail [3]map[string]string
}

// type detail struct {
// 	detail map[string]string
// }

func main() {

	//sd := &detail{detail: d}

	a := [3]map[string]string{}
	for i := 0; i < 3; i++ {
		d := make(map[string]string, 8)
		//sd := &detail{detail: d}
		a[i] = d
	}

	//address := &address{index: a}
	a[0]["1"] = "1"
	a[0]["2"] = "1"
	a[0]["3"] = "1"
	a[0]["4"] = "1"
	a[0]["5"] = "1"
	a[0]["6"] = "1"
	a[0]["7"] = "1"
	a[0]["8"] = "1"
	fmt.Println(a[0])

	// a[0].detail["line1"] = "1"
	// a[0].detail["line2"] = "1"
	// address.index[0].detail["line3"] = "1"

	// address.index[1].detail["line1"] = "2"
	// address.index[1].detail["line2"] = "2"
	// address.index[1].detail["line3"] = "2"

	// address.index[2].detail["line1"] = "3"
	// address.index[2].detail["line2"] = "3"
	// address.index[2].detail["line3"] = "3"

	// fmt.Println(address.index[0])
	// fmt.Println(address.index[1])
	// fmt.Println(address.index[2])
}
