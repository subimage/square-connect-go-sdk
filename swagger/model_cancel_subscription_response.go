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

// Defines fields that are included in a [CancelSubscription](api-endpoint:Subscriptions-CancelSubscription) response.
type CancelSubscriptionResponse struct {
	// Information about errors encountered during the request.
	Errors       []ModelError  `json:"errors,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}
