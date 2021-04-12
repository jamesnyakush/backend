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

func UpdateUserPesaSwap(firstname, lastname, email, phone, address1, address2, country, external string) {
	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

	res, err := svc.UserUpdate(
		firstname,
		lastname,
		email,
		phone,
		address1,
		address2,
		country,
		external,
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

func CardPaymentPesaSwap(currency, amount, expiry, security, number, transactionId, customerId string) {

	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

	res, err := svc.CardPayment(
		currency,
		amount,
		expiry,
		security,
		number,
		transactionId,
		customerId,
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

func PesaSwapAccountBalance(account string) {
	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

	res, err := svc.AccountBalance(
		account,
		"",
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

func PesaSwapFundTransfer(source, dest, sourceNarration, destNarration, amount, currency, bankCode, environment, customerId, transactionId string) {
	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

	res, err := svc.FundTransfer(
		source,
		dest,
		sourceNarration,
		destNarration,
		amount,
		currency,
		bankCode,
		environment,
		customerId,
		transactionId,
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)

}

func PesaSwapSendToMpesa(source, dest, sourceNarration, destNarration, amount, currency, environment, customerId, transactionId string) {
	svc, err := api.New(consumerKey, appSecret, api.SANDBOX)

	if err != nil {
		panic(err)
	}

	res, err := svc.SendToMpesa(
		source,
		dest,
		sourceNarration,
		destNarration,
		amount,
		currency,
		environment,
		customerId,
		transactionId,
	)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
