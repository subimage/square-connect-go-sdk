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

// Describes a gift card activity of the ACTIVATE type.
type GiftCardActivityActivate struct {
	AmountMoney *Money `json:"amount_money,omitempty"`
	// The ID of the order associated with the activity.  This is required if your application uses the Square Orders API.
	OrderId string `json:"order_id,omitempty"`
	// The `line_item_uid` of the gift card line item in an order.  This is required if your application uses the Square Orders API.
	LineItemUid string `json:"line_item_uid,omitempty"`
	// If your application does not use the Square Orders API, you can optionally use this field  to associate the gift card activity with a client-side entity.
	ReferenceId string `json:"reference_id,omitempty"`
	// Required if your application does not use the Square Orders API.  This is a list of client-provided payment instrument IDs.  Square uses this information to perform compliance checks. If you use the Square Orders API, Square has the necessary instrument IDs to perform necessary  compliance checks.
	BuyerPaymentInstrumentIds []string `json:"buyer_payment_instrument_ids,omitempty"`
}
