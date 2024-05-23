package enums

type TransactionStatusEnum string

var (
	WaitingForConfirmationStatus TransactionStatusEnum = "waiting_for_confirmation"
	WaitingForMOU                TransactionStatusEnum = "waiting_for_mou"
	WaitingForMOUConfirmation    TransactionStatusEnum = "waiting_for_mou_confirmation"
	Confirmed                    TransactionStatusEnum = "confirmed"
	WaitingForInitialPayment     TransactionStatusEnum = "waiting_for_initial_payment"
	Ongoing                      TransactionStatusEnum = "ongoing"
	WaitingForFurtherPayments    TransactionStatusEnum = "waiting_for_further_payments"
	Suspended                    TransactionStatusEnum = "suspended"
	Done                         TransactionStatusEnum = "done"
)
