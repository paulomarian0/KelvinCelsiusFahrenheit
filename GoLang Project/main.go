package main

import "fmt"

//Projeto simples: Pegar temperaturas e mostrar em kelvin, fahreint e celsius

func main() {

	var entradaTemp float64

	var entradaSwitch int

	for {

		fmt.Println("Escolha '1' para entrar com a temperatura em Kelvin")
		fmt.Println("Escolha '2' para entrar com a temperatura em Fahrenheit")
		fmt.Println("Escolha '3' para entrar com a temperatura em Celsius")
		fmt.Println("Escolha '4' para finalizar o programa")
		fmt.Scan(&entradaSwitch)

		// FINALIZAR O PROGRAMA ANTES DE ENTRAR NO SWITCH
		if entradaSwitch == 4 {
			fmt.Println("Finalizando o programa!!!!!!")
			break
		}
		fmt.Println("Digite a temperatura:")

		fmt.Scan(&entradaTemp)

		switch entradaSwitch {

		case 1:

			fmt.Println("Você digitou a temperatura:", entradaTemp, "em graus Kelvin")
			fmt.Println("Temperatura em Celsius:", kelvinToCelsius(entradaTemp))
			fmt.Println("Temperatura em Fahrenheit:", kelvinToFahnreint(entradaTemp))
			fmt.Println("---------------------------------------------------------------")

		case 2:
			fmt.Println("Você digitou a temperatura:", entradaTemp, "em graus Celsius")
			fmt.Println("Temperatura em Kelvin:", celsiusToKelvin(entradaTemp))
			fmt.Println("Temperatura em Fahrenheit:", celsiusToFahrenheit(entradaTemp))
			fmt.Println("---------------------------------------------------------------")

		case 3:
			fmt.Println("Você digitou a temperatura:", entradaTemp, "em graus Fahrenheit")

			fmt.Println("Temperatura em Celsius:", fahrenheitToCelsius(entradaTemp))
			fmt.Println("Temperatura em Kelvin:", fahrenheitToKelvin(entradaTemp))
			fmt.Println("---------------------------------------------------------------")

		default:
			fmt.Println("Entrada inválida")

		}

	}
}

//	ENTRADA EM KELVIN

func kelvinToFahnreint(entrada float64) float64 {

	celsius := entrada - 273

	fahrenheit := (1.8 * celsius) + 32

	return fahrenheit

}

func kelvinToCelsius(entrada float64) float64 {

	celsius := entrada - 273

	return celsius

}

// ENTRADA EM CELSIUS

func celsiusToKelvin(entrada float64) float64 {

	kelvin := entrada + 273

	return kelvin
}

func celsiusToFahrenheit(entrada float64) float64 {

	fahrenheit := 1.8*entrada + 32

	return fahrenheit

}

// ENTRADA EM FAHRENTOIDNTND

func fahrenheitToCelsius(entrada float64) float64 {

	celsius := ((entrada - 32) * 5) / 9

	return celsius
}

func fahrenheitToKelvin(entrada float64) float64 {

	celsius := ((entrada - 32) * 5) / 9

	kelvin := celsius + 273

	return kelvin
}
