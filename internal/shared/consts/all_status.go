package sharedconsts

// Order status
const (
	OrderStatusNotConfirmed      = "order-not-confirmed"  // status inicial da ordem
	OrderStatusConfirmed         = "order-confirmed"      // order confirmed by the customer
	OrderStatusCheckoutConfirmed = "checkout-confirmed"   // checkout confirmado e aguardando pagamento
	OrderStatusPaymentApproved   = "payment-approved"     // payment approved by the payment gateway
	OrderStatusNotApproved       = "payment-not-approved" // em caso de recusa do pagamento pelo gateway

	OrderReceivedByKitchen      = "order-received-by-kitchen"       // pedido recebido pela cozinha (recebido)
	OrderInPreparationByKitchen = "order-in-preparation-by-kitchen" // pedido em preparo na cozinha (em preparo)
	OrderReadyByKitchen         = "order-ready-by-kitchen"          // pedido pronto na cozinha (pronto)
	OrderFinalizedByKitchen     = "order-finalized-by-kitchen"      // pedido finalizado na cozinha (finalizado)
)

// Order Item status
const (
	OrderItemStatusConfirmed = "item-confirmed"
	OrderItemStatusCanceled  = "item-canceled"
)
