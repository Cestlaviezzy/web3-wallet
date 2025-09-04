package config

import (
	"fmt"
	"os"
)

// Config API Gateway配置
type Config struct {
	Server   ServerConfig   `json:"server"`
	Services ServicesConfig `json:"services"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

// ServicesConfig 后端服务配置
type ServicesConfig struct {
	WalletAddr string `json:"wallet_addr"`
	ChainAddr  string `json:"chain_addr"`
	TxAddr     string `json:"tx_addr"`
	AdminAddr  string `json:"admin_addr"`
	NotifyAddr string `json:"notify_addr"`
}

// Address 返回服务器监听地址
func (s ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

// Load 加载配置
func Load() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Port: getEnvOrDefault("PORT", "3050"),
			Host: getEnvOrDefault("HOST", "0.0.0.0"),
		},
		Services: ServicesConfig{
			WalletAddr: getEnvOrDefault("WALLET_SERVICE_ADDR", "localhost:50001"),
			ChainAddr:  getEnvOrDefault("CHAIN_SERVICE_ADDR", "localhost:50002"),
			TxAddr:     getEnvOrDefault("TRANSACTION_SERVICE_ADDR", "localhost:50003"),
			AdminAddr:  getEnvOrDefault("ADMIN_SERVICE_ADDR", "localhost:50004"),
			NotifyAddr: getEnvOrDefault("NOTIFY_SERVICE_ADDR", "localhost:50005"),
		},
	}

	return config, nil
}

// getEnvOrDefault 获取环境变量或默认值
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
