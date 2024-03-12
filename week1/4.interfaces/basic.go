package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не достаточно средств")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main(){
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
}

// you can embed interfaces
// you can define empty example i-face then example.(youri-face).(field)
