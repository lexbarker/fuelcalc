package fuelcalc

import "fmt"

//import "fmt"

func (fc *Falc) Consumption() {
	mpl := fc.Miles / fc.Litres
	fmt.Printf("you achieved %f mile per litre\n", mpl)
	mpg := mpl * L2gal
	fmt.Printf("mpg  = %f\n", mpg)
}

func Testmeout() {}
