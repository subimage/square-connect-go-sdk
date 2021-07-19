/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Present only when `GiftCardActivityType` is LOAD.
type GiftCardActivityLoad struct {
	AmountMoney *Money `json:"amount_money,omitempty"`
	// The `order_id` of the order associated with the activity. It is populated along with `line_item_uid` and is required if using the Square Orders API.
	OrderId string `json:"order_id,omitempty"`
	// The `line_item_uid` of the gift card’s line item in the order associated with the activity. It is populated along with `order_id` and is required if using the Square Orders API.
	LineItemUid string `json:"line_item_uid,omitempty"`
	// A client-specified ID to associate an entity, in another system, with this gift card activity. This can be used to track the order or payment related information when the Square Orders API is not being used.
	ReferenceId string `json:"reference_id,omitempty"`
	// If you are not using the Orders API, this field is required because it is used to identify a buyer  to perform compliance checks.
	BuyerPaymentInstrumentIds []string `json:"buyer_payment_instrument_ids,omitempty"`
}
