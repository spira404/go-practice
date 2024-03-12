package main

import "fmt"

 // 1
type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Недостаточно средств")
	}
	w.Cash -= amount
	return nil
}

 // 2
type Card struct {
	Balance		int
	ValidUntil	string
	Cardholder	string
	CVV			string
	Number		string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Недостаточно средств")
	}
	c.Balance -= amount
	return nil
}

 // 3
type ApplePay struct {
	Money	int
	AppleID	string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Недостаточно средств")
	}
	a.Money -= amount
	return nil
}


type Payer interface {
	Pay(int) error
}


func Buy(p Payer) {

	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличными")
	case *Card:
		plasticCard := p.(*Card)
		fmt.Println("Встаувляйте карту", plasticCard.Cardholder)
	default:
		fmt.Println("Новый метод оплаты")
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка оплаты %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Оплата проведена через %T\n\n", p)
}


func main() {

	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "spira"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}
