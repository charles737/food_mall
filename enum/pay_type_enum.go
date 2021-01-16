package enum

type PayType int

const (
	Bank PayType = 0
	WeChat PayType = 1
	AliPay PayType = 2
	PayPal PayType = 3
)

func (p PayType) String() string {
	switch p {
	case Bank:
		return "银行卡"
	case WeChat:
		return "微信"
	case AliPay:
		return "支付宝"
	case PayPal:
		return "PayPal"
	default:
		return "UNKNOWN"
	}
}