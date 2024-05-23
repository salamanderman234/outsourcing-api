package enums

type MidtransStatusEnum string

var (
	CaptureMidtranstatus        MidtransStatusEnum = "capture"
	SettlementMidtransStatus    MidtransStatusEnum = "settlement"
	PendingMidtransStatus       MidtransStatusEnum = "pending"
	DenyMidtransStatus          MidtransStatusEnum = "deny"
	CancelMidtransStatus        MidtransStatusEnum = "cancel"
	ExpireMidtransStatus        MidtransStatusEnum = "expire"
	FailureMidtransStatus       MidtransStatusEnum = "failure"
	RefundMidtransaStatus       MidtransStatusEnum = "refund"
	PartialRefundMidtransStatus MidtransStatusEnum = "partial_refund"
	AuthorizeMidtransStatus     MidtransStatusEnum = "authorize"
)
