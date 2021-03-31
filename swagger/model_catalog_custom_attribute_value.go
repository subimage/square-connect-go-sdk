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

// An instance of a custom attribute. Custom attributes can be defined and added to `ITEM` and `ITEM_VARIATION` type catalog objects. [Read more about custom attributes](https://developer.squareup.com/docs/catalog-api/add-custom-attributes).
type CatalogCustomAttributeValue struct {
	// The name of the custom attribute.
	Name string `json:"name,omitempty"`
	// The string value of the custom attribute.  Populated if `type` = `STRING`.
	StringValue *string `json:"string_value,omitempty"`
	// __Read-only.__ The id of the [CatalogCustomAttributeDefinition](#type-CatalogCustomAttributeDefinition) this value belongs to.
	CustomAttributeDefinitionId string                                `json:"custom_attribute_definition_id,omitempty"`
	Type_                       *CatalogCustomAttributeDefinitionType `json:"type,omitempty"`
	// Populated if `type` = `NUMBER`. Contains a string representation of a decimal number, using a `.` as the decimal separator.
	NumberValue *string `json:"number_value,omitempty"`
	// A `true` or `false` value. Populated if `type` = `BOOLEAN`.
	BooleanValue *bool `json:"boolean_value,omitempty"`
	// One or more choices from `allowed_selections`. Populated if `type` = `SELECTION`.
	SelectionUidValues *[]string `json:"selection_uid_values,omitempty"`
	// __Read-only.__ A copy of key from the associated `CatalogCustomAttributeDefinition`.
	Key string `json:"key,omitempty"`
}
