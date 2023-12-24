package domains

type RoleEnum string

const (
	AdminRole       RoleEnum = "admin"
	EmployeeRole    RoleEnum = "employee"
	ServiceUserRole RoleEnum = "service_user"
	SupervisorRole  RoleEnum = "supervisor"
)

type OrderStatusEnum string

const (
	WaitingForPaymentOrderStatus      OrderStatusEnum = "waiting_for_payment"
	WaitingForConfirmationOrderStatus OrderStatusEnum = "waiting_for_confirmation"
	ProcessedOrderStatus              OrderStatusEnum = "processed"
	WaitingForFurtherPaymentStatus    OrderStatusEnum = "waiting_for_further_payment"
	OngoingOrderStatus                OrderStatusEnum = "ongoing"
	CompletedOrderStatus              OrderStatusEnum = "completed"
	CancelledOrderStatus              OrderStatusEnum = "cancelled"
)

type PaymentTypeEnum string

const (
	FullPaymentType        PaymentTypeEnum = "full"
	DownPaymentPaymentType PaymentTypeEnum = "dp"
	ThreeTerminPaymentType PaymentTypeEnum = "3_termin"
)
