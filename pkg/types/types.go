package types

//Money представляют собой денежную сумму в минимальных единицах (центы, копейки, дирамы и тд)
type Money int64

//Currency представляют код валюты
type Currency string

//коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет номер карты
type PAN string

//Card представлет информацию о платежной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

//Category представляет собой категорию в который был совершен платеж (авир, аптеки, рестораны и тд)
type Category string

//Category представляет собой статус платежа
type Status string

//
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment представляет информацию о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

type PaymentSource struct {
	Type    string //'card'
	Number  string //номер вида '5058 xxxx xxxx 8888'
	Balance Money  //баланс в дирамах
}
