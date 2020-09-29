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

// An enumerated value that can link a `CatalogItemVariation` to an item option as one of its item option values.
type CatalogItemOptionValue struct {
	// Unique ID of the associated item option.
	ItemOptionId string `json:"item_option_id,omitempty"`
	// Name of this item option value. This is a searchable attribute for use in applicable query filters.
	Name string `json:"name,omitempty"`
	// A human-readable description for the option value. This is a searchable attribute for use in applicable query filters.
	Description string `json:"description,omitempty"`
	// The HTML-supported hex color for the item option (e.g., \"#ff8d4e85\"). Only displayed if `show_colors` is enabled on the parent `ItemOption`. When left unset, `color` defaults to white (\"#ffffff\") when `show_colors` is enabled on the parent `ItemOption`.
	Color string `json:"color,omitempty"`
	// Determines where this option value appears in a list of option values.
	Ordinal int32 `json:"ordinal,omitempty"`
	// The number of `CatalogItemVariation`s that currently use this item option value. Present only if `retrieve_counts` was specified on the request used to retrieve the parent item option of this value.
	ItemVariationCount int64 `json:"item_variation_count,omitempty"`
}