package main

import "fmt"

func main() {
	x := 0

	// if
	if x == 0 {
		fmt.Println("X vale zero")
	} else if x == 1 {
		fmt.Println("X vale 1")
	} else {
		fmt.Println("X nao vale nem 0 nem 1")
	}

	// switch case
	y := 5

	switch  {
	case y < 5:
		fmt.Println("Y é menor que 5")
		break
	case y > 5:
		fmt.Println("Y é maior que 5")
		break
	default:
		fmt.Println("Y vale 5")
	}

	quemTaNoEscritorioHoje := "Lucas"

	switch quemTaNoEscritorioHoje {
	case "Nathália", "Lucas":
		fmt.Println("Nathália e lucas estão no escritório")
		fallthrough // executa a próxima etapa também
	case "Miu":
		fmt.Println("meus papais vieram eu também vim")
	case "Deli", "Rebz":
		fmt.Println("O deli e a rebecca estão pintando o sete no escritório")
	default:
		fmt.Println("Ningúem veio :-(")
	}

	var ola interface{}
	ola = true

	switch ola.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float")
		
	}
}
