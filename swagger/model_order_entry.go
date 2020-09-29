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

// A lightweight description of an [Order](#type-order) that is returned when `returned_entries` is true on a [SearchOrderRequest](#type-searchorderrequest)
type OrderEntry struct {
	// The id of the Order
	OrderId string `json:"order_id,omitempty"`
	// Version number which is incremented each time an update is committed to the order. Orders that were not created through the API will not include a version and thus cannot be updated.  [Read more about working with versions](https://developer.squareup.com/docs/orders-api/manage-orders#update-orders).
	Version int32 `json:"version,omitempty"`
	// The location id the Order belongs to.
	LocationId string `json:"location_id,omitempty"`
}