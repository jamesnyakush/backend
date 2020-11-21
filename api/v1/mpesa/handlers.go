package mpesa

type mpesaC2BRequest struct {
	TransactionType   string `json:"transaction_type"`
	TransactionID     string `json:"transaction_id"`
	TransactionTime   string `json:"transaction_time"`
	TransactionAmount string `json:"transaction_amount"`
	MSISDN            string `json:"msisdn"`
	ShortCode         string `json:"short_code"`
	OrgAccountBalance string `json:"org_account_balance"`
	BillRefNumber     string `json:"bill_ref_number"`
	InvoiceNumber     string `json:"invoice_number"`
	FirstName         string `json:"first_name"`
	MiddleName        string `json:"middle_name"`
	LastName          string `json:"last_name"`
}

type mpesaB2CRequest struct {
	ResultCode                          string `json:"result_code"`
	ResultDesc                          string `json:"result_desc"`
	OriginatorConversationID            string `json:"originator_conversation_id"`
	ConversationID                      string `json:"conversation_id"`
	TransactionID                       string `json:"transaction_id"`
	Msisdn                              string `json:"msisdn"`
	B2CWorkingAccountAvailableFunds     string `json:"b2c_working_account_available_funds"`
	B2CUtilityAccountAvailableFunds     string `json:"b2c_utility_account_available_funds"`
	Amount                              string `json:"amount"`
	TransactionCompletedDateTime        string `json:"transaction_completed_date_time"`
	B2CChargesPaidAccountAvailableFunds string `json:"b2c_charges_paid_account_available_funds"`
	ReceiverPartyPublicName             string `json:"receiver_party_public_name"`
	B2CRecipientIsRegisteredCustomer    string `json:"b2c_recipient_is_registered_customer"`
}
