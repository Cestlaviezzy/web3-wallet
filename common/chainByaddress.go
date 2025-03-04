package common

import (
	"fmt"
	"regexp"
	"strings"
)

// 判断地址类型
func GetChainByAddress(address string) string {
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
