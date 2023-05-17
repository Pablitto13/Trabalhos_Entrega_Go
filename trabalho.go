package main

import "fmt"

const EbuliKelvin = 373.2

func main() {
	kelvin := EbuliKelvin
	celsius := EbuliKelvin - 273

	fmt.Printf("O ponto de ebulição em Kelvin é: %g e o ponto de ebulição em Celsius é: %g", kelvin, celsius)

}
