package enums

type PaymentEnum string

var (
	FullPayment PaymentEnum = "full"
	DpPayment   PaymentEnum = "dp"
	ThreeTermin PaymentEnum = "3_termin"
)

type PaymentSubTypeEnum string

var (
	DPSubType                PaymentSubTypeEnum = "dp_first"
	ThreeTerminFirstSubType  PaymentSubTypeEnum = "termin_first"
	ThreeTerminSecondSubType PaymentSubTypeEnum = "termin_second"
)

type DPStatus string

var (
	WaitingForDP          DPStatus = "waiting_for_dp"
	DPCompleted           DPStatus = "dp_completed"
	WaitingForRemainingDP DPStatus = "waiting_for_remaining_dp"
	DPRemainingCompleted  DPStatus = "remaining_dp_completed"
)

type TerminStatus string

var (
	WaitingForFirstTermin  DPStatus = "waiting_for_first_termin"
	FirstTerminCompleted   DPStatus = "first_termin_completed"
	WaitingForSecondTermin DPStatus = "waiting_for_second_termin"
	SecondTerminCompleted  DPStatus = "second_termin_completed"
	WaitingForThirdTermin  DPStatus = "waiting_for_third_termin"
	ThirdTerminCompleted   DPStatus = "third_termin_completed"
)
