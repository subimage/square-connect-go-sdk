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

// Contains the measurement unit for a quantity and a precision which specifies the number of digits after the decimal point for decimal quantities.
type OrderQuantityUnit struct {
	MeasurementUnit *MeasurementUnit `json:"measurement_unit,omitempty"`
	// For non-integer quantities, represents the number of digits after the decimal point that are recorded for this quantity.  For example, a precision of 1 allows quantities like `\"1.0\"` and `\"1.1\"`, but not `\"1.01\"`.  Min: 0. Max: 5.
	Precision int32 `json:"precision,omitempty"`
}