package enum

type PayStatus int

const (
	UnPay PayStatus = 0
	Pay PayStatus = 1
)

func (p PayStatus) String() string {
	switch p {
	case UnPay:
		return "未支付"
	case Pay:
		return "已支付"
	default:
		return "UNKNOWN"
	}
}