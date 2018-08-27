// fuelcalc provides function to caluclate mile per gallon
package fuelcalc

func (fc *Mpg) Consumption() (mpg float64) {
	mpl := fc.Miles / fc.Litres
	mpg = mpl * L2gal
	return
}

func (fc *Fuelused) Fuelrequired() {
	fused := fc.Miles / fc.Mpg
	fc.Litres = fused * Gal2l
	return
}

func Testmeout() {}
