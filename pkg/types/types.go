package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платежной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

// Category представляет собой категорию, в которой был совершен платеж (авто, аптеки, рестораны и т.д.).
type Category string

// Payment предстваляет информацию о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
