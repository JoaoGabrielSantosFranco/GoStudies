package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	originUnit := os.Args[len(os.Args)-1]

	finalUnit := converter(originUnit)

	originValues := os.Args[1 : len(os.Args)-1]

	fmt.Println("Valores originais:", originValues)
	fmt.Println("Unidade original:", originUnit)
	fmt.Println("Unidade convertida:", finalUnit)
	converterValues(originValues, originUnit, finalUnit)

}

func converter(originUnit string) string {
	if originUnit == "celsius" {
		return "fahrenheit"
	} else if originUnit == "quilometros" {
		return "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!\n", originUnit)
		os.Exit(1)
	}
	return ""
}

func converterValues(originValues []string, originUnit string, finalUnit string) {
	for _, valStr := range originValues {
		valFloat, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Printf("Valor inválido: %s\n", valStr)
			continue
		}
		converted := converterValue(float32(valFloat), originUnit)
		fmt.Printf("%f %s = %f %s\n", valFloat, originUnit, converted, finalUnit)
	}
}

func converterValue(originValue float32, originUnit string) float32 {
	if originUnit == "celsius" {
		return originValue*1.8 + 32
	} else if originUnit == "quilometros" {
		return originValue * 0.621371
	}
	return originValue
}
