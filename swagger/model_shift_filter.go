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

// Defines a filter used in a search for `Shift` records. `AND` logic is used by Square's servers to apply each filter property specified.
type ShiftFilter struct {
	// Fetch shifts for the specified location.
	LocationIds []string `json:"location_ids"`
	// Fetch shifts for the specified employees. DEPRECATED at version 2020-08-26. Use `team_member_ids` instead
	EmployeeIds []string           `json:"employee_ids,omitempty"`
	Status      *ShiftFilterStatus `json:"status,omitempty"`
	Start       *TimeRange         `json:"start,omitempty"`
	End         *TimeRange         `json:"end,omitempty"`
	Workday     *ShiftWorkday      `json:"workday,omitempty"`
	// Fetch shifts for the specified team members. Replaced `employee_ids` at version \"2020-08-26\"
	TeamMemberIds []string `json:"team_member_ids"`
}