package valueobject

// Order status
const (
	OrderStatusConfirmed    = "confirmed"
	OrderStatusNotConfirmed = "notconfirmed"
	OrderStatusPaid         = "paid"
	OrderStatusPreparing    = "preparing"
	OrderStatusDelivered    = "delivered"
	OrderStatusCanceled     = "canceled"
	OrderStatusApproved     = "approved"
	OrderStatusNotApproved  = "notapproved"
)

// Order Item status
const (
	OrderItemStatusConfirmed = "confirmed"
	OrderItemStatusCanceled  = "canceled"
)
