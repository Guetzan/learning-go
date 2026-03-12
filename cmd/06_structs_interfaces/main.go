package main

import "fmt"

type engine interface {
	kmsLeft() uint16
}

type gasEngine struct {
	kpl    uint8 //kilometers per liter
	liters uint8 
	manufacturer
}

type electricEngine struct {
	kpkwh uint8 //kilometers per kWh
	kwh uint8
}

func (engine gasEngine) kmsLeft() uint16 {
	return uint16(engine.kpl *  engine.liters)
}

func (engine electricEngine) kmsLeft() uint16 {
	return uint16(engine.kpkwh * engine.kwh)
}

func canMakeIt(engine engine, kilometers uint16) {
	if engine.kmsLeft() < kilometers {
		fmt.Println("There's no way you're making it there")
	} else {
		fmt.Println("You're able to make it there!")
	}
} 

type manufacturer struct {
	name string 
	foundation uint16
}

func main() {
	var engine gasEngine = gasEngine{15, 50, manufacturer{"Peugeot", 1810}}
	fmt.Println(engine.kpl, engine.liters, engine.name)
	fmt.Println(engine.kmsLeft())
	canMakeIt(engine, 239)

	var engine2 electricEngine = electricEngine{22, 76}
	fmt.Println(engine2.kpkwh, engine2.kwh)
	fmt.Println(engine2.kmsLeft())
	canMakeIt(engine2, 239)
}