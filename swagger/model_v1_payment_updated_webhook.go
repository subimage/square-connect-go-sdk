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

// Published when a charge is made or refunded through the Square Point of Sale app or the Transactions API.
type V1PaymentUpdatedWebhook struct {
	// The ID of the target merchant associated with the event.
	MerchantId string `json:"merchant_id,omitempty"`
	// The ID of the target merchant associated with the event.
	LocationId string `json:"location_id,omitempty"`
	// The type of event this represents, i.e. `PAYMENT_UPDATED`.
	EventType string `json:"event_type,omitempty"`
	// The ID of the updated V1 Payment.
	EntityId string `json:"entity_id,omitempty"`
}