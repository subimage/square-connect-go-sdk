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

// The custom attribute filters in a set of [customer filters](entity:CustomerFilter) used in a search query. Use this filter to search based on [custom attributes](entity:CustomAttribute) that are assigned to customer profiles. For more information, see [Search by custom attribute](https://developer.squareup.com/docs/customers-api/use-the-api/search-customers#search-by-custom-attribute).
type CustomerCustomAttributeFilters struct {
	// The custom attribute filters. Each filter must specify `key` and include the `filter` field with a type-specific filter, the `updated_at` field, or both. The provided keys must be unique within the list of custom attribute filters.
	Filters []CustomerCustomAttributeFilter `json:"filters,omitempty"`
}