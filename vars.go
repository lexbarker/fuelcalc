// vars.go provides the variables, contstants structs and interfaces for the fuelcalc package
package fuelcalc

//mpg struct, provides miles and litres to determine MPG
type Mpg struct {
	Miles  float64
	Litres float64
}

//constant for litre to gallon conversion
const L2gal = 4.53

//Provides Miles and mpg to determine fuel used
type Fuelused struct {
	Miles  float64
	Mpg    float64
	Litres float64
}

const Gal2l = 4.54
