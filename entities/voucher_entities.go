package entities

type VoucherEntity struct {
	Id      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Expired string `json:"expired_date"`
}
