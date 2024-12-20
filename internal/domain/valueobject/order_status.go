package valueobject

// Order status
const (
	OrderStatusConfirmed    = "confirmed"
	OrderStatusNotConfirmed = "notconfirmed"
	OrderStatusApproved     = "approved"
	OrderStatusNotApproved  = "notapproved"
	OrderStatusPaid         = "paid"
	OrderStatusPreparing    = "preparing"
	OrderStatusDelivered    = "delivered"
	OrderStatusCanceled     = "canceled"
)

// Order Item status
const (
	OrderItemStatusConfirmed = "confirmed"
	OrderItemStatusCanceled  = "canceled"
)
