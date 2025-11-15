package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AddToCart(reader *bufio.Reader) {
	fmt.Println("\n--- Adding order to busket ---")

	fmt.Print("Enter the title of item: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Enter the price of item")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	price, _ := strconv.Atoi(priceStr)

	fmt.Printf("Enter the amount of items")
	qtyStr, _ := reader.ReadString('\n')
	qtyStr = strings.TrimSpace(qtyStr)
	qty, _ := strconv.Atoi(qtyStr)

	fmt.Printf("You added to busket: %s, quantity: %d, price per item: %d tenge",
		title, price, qty)

}
