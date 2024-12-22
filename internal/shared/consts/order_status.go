package sharedconsts

// Order status
const (
	OrderStatusConfirmed       = "order-confirmed" // order confirmed by the customer
	OrderStatusNotConfirmed    = "order-not-confirmed"
	OrderStatusPaymentApproved = "payment-approved" // payment approved by the payment gateway
	OrderStatusNotApproved     = "payment-not-approved"
	// OrderStatusPreparing       = "preparing"
	// OrderStatusDelivered       = "delivered"
	// OrderStatusCanceled        = "canceled"
)

// Order Item status
const (
	OrderItemStatusConfirmed = "item-confirmed"
	OrderItemStatusCanceled  = "item-canceled"
)
