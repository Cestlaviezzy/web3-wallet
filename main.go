package main

import (
	"fmt"
	"web3-wallet/common"
)

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
		chain := common.GetChainByAddress(address)
		fmt.Println(address, chain)
	}
}
