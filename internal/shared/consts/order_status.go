package sharedconsts

// Order status
const (
	OrderStatusNotConfirmed      = "order-not-confirmed"  // status inicial da ordem
	OrderStatusConfirmed         = "order-confirmed"      // order confirmed by the customer
	OrderStatusCheckoutConfirmed = "checkout-confirmed"   // checkout confirmado e aguardando pagamento
	OrderStatusPaymentApproved   = "payment-approved"     // payment approved by the payment gateway
	OrderStatusNotApproved       = "payment-not-approved" // em caso de recusa do pagamento pelo gateway
)

// Order Item status
const (
	OrderItemStatusConfirmed = "item-confirmed"
	OrderItemStatusCanceled  = "item-canceled"
)

// Kitchen
const (
	KitchenReceived  = "received"
	KitchenPreparing = "preparing"
	KitchenReady     = "ready"
	KitchenDelivered = "delivered"
)
