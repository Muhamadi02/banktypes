package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string

// Category представляет собой категорию, в которой был совершен платеж (авто, аптекы, рестораны и т.д.).
type Category string

// Status представляет собой статус платежа.
type Status string

// Предопределенные статусы платежей.
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

// Category представляет собой категорию, в которой был совершен платёж (авто, аптеки, рестораны и т.д.).
type Category string

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
