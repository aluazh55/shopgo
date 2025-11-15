package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aluazh55/shopgo/shop"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1) Add item to busket")
	fmt.Println("2) Calculate the price of order")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	ch, _ := strconv.Atoi(input)

	if ch == 0 {
		shop.AddToCart(reader)

	} else {
		shop.OrderCost(reader)
	}

}
