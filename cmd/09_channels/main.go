package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_DDR5_RAM_PRICE float32 = 500
var MAX_BOMB_APU_PRICE float32 = 80

func main() {
	startedAt := time.Now()

	var ramChannel = make(chan string)
	var apuChannel = make(chan string)
	var websites = []string{"kabum.com.br", "amazon.com.br", "terabyteshop.com.br"}

	for _, website := range websites {
		go checkRamPrices(website, ramChannel)
		go checkApuPrices(website, apuChannel)
	}

	notify(ramChannel, apuChannel)
	fmt.Println(time.Since(startedAt))
}

func checkRamPrices(website string, ramChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var price float32 = rand.Float32() * 10000

		if price <= MAX_DDR5_RAM_PRICE {
			ramChannel <- website
			break
		}
	}
}

func checkApuPrices(website string, apuChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var price float32 = rand.Float32()*1000

		if price <= MAX_BOMB_APU_PRICE {
			apuChannel <- website
			break
		}
	}
}

func notify(ramChannel chan string, apuChannel chan string) {
	select {
		case website := <- ramChannel:
			fmt.Println("A RAM deal was found at", website)
		case website := <- apuChannel:
			fmt.Println("A APU deal was found at", website)
	}
}
