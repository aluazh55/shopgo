package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func OrderCost(reader *bufio.Reader) {
	fmt.Printf("\n--- Order cost calculation ---")
	fmt.Printf("Enter the cost of item")

	pStr, _ := reader.ReadString('\n')
	pStr = strings.TrimSpace(pStr)
	price, _ := strconv.Atoi(pStr)

	fmt.Printf("Enter the amount")
	qStr, _ := reader.ReadString('\n')
	qStr = strings.TrimSpace(qStr)
	qty, _ := strconv.Atoi(qStr)

	total := price * qty

	fmt.Printf("The total amount: %d tenge\n", total)

}
