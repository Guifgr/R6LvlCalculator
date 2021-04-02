package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func calcularXpNivel(nivel int) int {
	result := 0
	if nivel >= 2 {
		result = (nivel-2)*500 + 5000
	}
	return result
}

func somaXpAcumulado(nivel int) int {
	acumulado := 0
	for i := 0; i <= nivel; i++ {
		acumulado += calcularXpNivel(i)
	}
	return acumulado
}

func calcularXp(nivelAtual, nivelDesejado int) int {
	totalXpAtual := somaXpAcumulado(nivelAtual)
	totalXpDesejado := somaXpAcumulado(nivelDesejado)
	return totalXpDesejado - totalXpAtual
}

func nextStep() {
	fmt.Print("\nPressione 'Enter' para continuar!")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	CallClear() //function to clear the console or terminal
}

func main() {
	CallClear() //function to clear the console or terminal
	var nivelAtual, nivelDesejado, valorVenda, nivelDeVenda int
	fmt.Printf("Seja bem vindo!\n\nSe quiser descobrir quantos XP precisa para subir até determinado nível\n")
	fmt.Printf("Vamos apenas precisar que nos próximos passos você digite seu nível atual e em que nível quer chegar\n")
	nextStep()
	fmt.Printf("Agora digite em que nível você está!\n\nDigite o nível:  ")
	fmt.Scan(&nivelAtual)
	if nivelAtual < 0 {
		nivelAtual = 0
	}
	CallClear() //function to clear the console or terminal
	fmt.Printf("Agora digite em que nível quer chegar!\n\nDigite o nível:  ")
	fmt.Scan(&nivelDesejado)
	if nivelDesejado < 0 {
		nivelDesejado = 0
	}
	xpNecessario := calcularXp(int(nivelAtual), int(nivelDesejado))
	CallClear() //function to clear the console or terminal
	fmt.Println("Você precisa upar do nível", nivelAtual, "ao", nivelDesejado, "=", nivelDesejado-nivelAtual)
	fmt.Println("e precisa de", xpNecessario, "xp para alcançar seu objetivo")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	nextStep()
	fmt.Println("Vamos calcular quanto vale o que você upou da conta")
	fmt.Println("Primeiro vamos precisar saber por quanto a conta vai ser vendida e em que nível")
	nextStep()
	fmt.Printf("Qual o valor de venda da conta?\n")
	fmt.Scan(&valorVenda)
	if valorVenda < 0 {
		valorVenda = 0
	}
	CallClear()
	fmt.Printf("Em qual nivel a conta vai ser vendida?\n")
	fmt.Scan(&nivelDeVenda)
	if nivelDeVenda < 0 {
		nivelDeVenda = 0
	}
	valorPorXp := float32(valorVenda) / float32(somaXpAcumulado(int(nivelDeVenda)))
	CallClear()
	fmt.Println("Você pegou a conta em que nível?")
	fmt.Scan(&nivelAtual)
	if nivelAtual < 0 {
		nivelAtual = 0
	}
	CallClear()
	fmt.Println("Você upou até de que nivel?")
	fmt.Scan(&nivelDesejado)
	if nivelDesejado < 0 {
		nivelDesejado = 0
	}
	if nivelDesejado > nivelDeVenda {
		nivelDesejado = nivelDeVenda
	}
	CallClear()
	var valorSerPago float32
	if nivelAtual >= 0 && nivelDesejado >= 0 && valorPorXp > 0 {
		valorSerPago = float32(calcularXp(int(nivelAtual), int(nivelDesejado))) * valorPorXp
	} else {
		valorSerPago = 0
	}
	fmt.Printf("O valor justo para você receber pelo que você upou é de R$%.2f\n", valorSerPago)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	nextStep()
}
