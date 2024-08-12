package main

import "fmt"

const ebulicaoF = 212.0

func main() {
	var F = ebulicaoF
	var C = (F - 32) * 5 / 9

	fmt.Println("Ponto de ebulição da ága em ºF é: ", F)
	fmt.Println("Ponto de ebulição da água em ºC é: ", C)
}    