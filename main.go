package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 判断地址类型
func getChainByAddress(address string) string {
	address = strings.TrimSpace(address)

	if len(address) == 0 {
		return "Invalid address"
	}

	if strings.HasPrefix(address, "1") || strings.HasPrefix(address, "3") || strings.HasPrefix(address, "bc1") {
		return "BTC"
	}

	if strings.HasPrefix(address, "0x") && len(address) == 42 {
		return "ETH/BSC"
	}

	if len(address) == 44 && isBase58(address) {
		return "SOL"
	}

	if strings.HasPrefix(address, "T") && len(address) == 34 {
		return "TRX"
	}

	if strings.HasPrefix(address, "f1") || strings.HasPrefix(address, "f3") || strings.HasPrefix(address, "f5") {
		return "FIL"
	}
	return "Unknown address"
}

func isBase58(s string) bool {
	regex := "^[1-9A-HJ-NP-Za-km-z]+$"
	//re := regexp.MustCompile(regex)
	re, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println("Invalid base58", err)
		return false
	}
	return re.MatchString(s)
}

func main() {
	addresses := []string{
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",           // BTC
		"0x742d35Cc6634C0532925a3b844Bc454e4438f44e",   // ETH
		"0x2b6e2b1f9c254e22a457ff636c5cc51556d30a5f",   // BSC
		"5hNtdHzv9f2PHcNjw5h3M3rRS7Emfym7N9u9ysRffvqD", // Solana
		"TJ56bNUfD5ZZoXr9KbgaJ4yWkdyfgQT2kX",           // TRX
		"f1v9X5a8yauWVYgyu7Ym2ws5w8YTKThmyXtHk",        // Filecoin
	}

	for _, address := range addresses {
		chain := getChainByAddress(address)
		fmt.Println(address, chain)
	}
}
