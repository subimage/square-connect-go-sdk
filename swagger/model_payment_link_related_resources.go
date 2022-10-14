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

type PaymentLinkRelatedResources struct {
	// The order associated with the payment link.
	Orders []Order `json:"orders,omitempty"`
	// The subscription plan associated with the payment link.
	SubscriptionPlans []CatalogObject `json:"subscription_plans,omitempty"`
}