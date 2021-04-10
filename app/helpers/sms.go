package helpers

import (
	"flag"
	"fmt"
	sms "github.com/jamesnyakush/AT-SMS"
	"log"
)

var (
	apiKey    = "8ce04b6c81f08cd40ce918f9b01bf979bf8d0052ac4e5203e39f4726d059aacf"
	username  = "mlaru"
	shortcode = ""
)

func Sms(recipient, message string) {
	//recipient := flag.String("r", "254746445198", "The phone number of the recipient of the message")
	//message := flag.String("m", "Karibu kenya ndugu yangu", "The message to be sent")
	env := flag.String("e", "Prod", "The environment of the api")

	flag.Parse()

	if recipient == "" || message == "" {
		log.Fatal("please enter all required arguments. see --help")
	}

	if apiKey == "" || username == "" {
		log.Fatal("missing required environment variables: AT_APIKEY, AT_USERNAME")
	}

	smsService := sms.NewService(username, apiKey, *env)

	sendResponse, err := smsService.Send(shortcode, recipient, message)

	if err != nil {
		fmt.Printf("Failed to send sms: %v", err)
	}
	fmt.Printf("SMS Send reponse: %v\n", sendResponse)

}
