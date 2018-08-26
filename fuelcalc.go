package main

import "fmt"

//import "fmt"

type Falc struct {
	miles  float64
	litres float64
	//l2gal  float64
}

const l2gal = 4.53

func (fc *Falc) Consumption() {
	mpl := fc.miles / fc.litres
	fmt.Printf("you achieved %f mile per litre\n", mpl)
	mpg := mpl * l2gal
	fmt.Printf("mpg  = %f\n", mpg)
}

func main() {
	var fc Falc
	fc.miles = 300
	//var litres float64
	//var litre float64

	fmt.Scan(&fc.litres)
	fc.Consumption()
	//mpl := float64(miles) / float64(litres)
	//fmt.Printf("you achieved %f mile per litre\n", mpl)
	//mpg := mpl * l2gal
	//fmt.Printf("mpg  = %f\n", mpg)

}
