package types

// RelayerResponse 统一响应格式 (兼容MetaSignX-Wallet)
type RelayerResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// 状态码常量定义
const (
	StatusSuccess               = 200
	StatusBadRequest            = 400
	StatusNotFound              = 404
	StatusUnsupportedMediaType  = 415
	StatusUnprocessableEntity   = 422
	StatusInternalError         = 500
	StatusTransactionValidating = 1001
	StatusContractNotDeployed   = 1002
	StatusChainContractError    = 1003
	StatusThirdError            = 2000
	StatusOther                 = 3000
)

// Success 创建成功响应
func Success(data interface{}) RelayerResponse {
	return RelayerResponse{
		StatusCode: StatusSuccess,
		Message:    "200 OK",
		Data:       data,
	}
}

// Error 创建错误响应
func Error(code int, message string) RelayerResponse {
	return RelayerResponse{
		StatusCode: code,
		Message:    message,
		Data:       nil,
	}
}

// SubmitterStatus 提交者状态 (兼容MetaSignX-Wallet)
type SubmitterStatus struct {
	Address              string `json:"address"`
	Nonce                int64  `json:"nonce"`
	Balance              string `json:"balance"`
	IsSendingTransaction bool   `json:"isSendingTransaction"`
	IsBlocking           bool   `json:"isBlocking"`
}

// SubmittersResponse 提交者状态响应
type SubmittersResponse struct {
	Status      []SubmitterStatus `json:"status"`
	FeeReceiver string            `json:"feeReceiver"`
}
