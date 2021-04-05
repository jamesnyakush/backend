package helpers

import (
	"github.com/jamesnyakush/pesaswap/pesaswap/api"
	"log"
)

const (
	consumerKey = "dMMcTvqhxl"
	appSecret   = "wjqRX5tTSzFyJbwW33ANzwh9V"
)

func CreateUserPesaSwap(firstname, lastname, email, phone string) {

	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

	res, err := svc.NewUser(
		firstname,
		lastname,
		email,
		phone,
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

func UpdateUserPesaSwap() {

}

func CardPaymentPesaSwap(currency, amount string) {

	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

	res, err := svc.CardPayment(
		currency,
		amount,
		"012022",
		"111",
		"4242424242424242",
		"AC312E60-BA4B-11EA-A700-4903572028c",
		"AC312E60-BA4B-11EA-A700-4903572028FB",
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
