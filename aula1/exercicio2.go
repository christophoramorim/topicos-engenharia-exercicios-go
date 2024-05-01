package main

import (
	"errors"
	"fmt"
)

func getInfo(input string) (int, error) {
	info := len(input)
	if info == 0 {
		return 0, errors.New("a string está vazia")
	}
	return info, nil
}

func main() {
	vlr_ini := "Olá, mundo!"
	size, err := getInfo(vlr_ini)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Printf("Tamanho da string: %d\n", size)
}
