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

// LoyaltyProgramRewardDefinitionScope : Indicates the scope of the reward tier.
type LoyaltyProgramRewardDefinitionScope string

// List of LoyaltyProgramRewardDefinitionScope
const (
	ORDER_LoyaltyProgramRewardDefinitionScope          LoyaltyProgramRewardDefinitionScope = "ORDER"
	ITEM_VARIATION_LoyaltyProgramRewardDefinitionScope LoyaltyProgramRewardDefinitionScope = "ITEM_VARIATION"
	CATEGORY_LoyaltyProgramRewardDefinitionScope       LoyaltyProgramRewardDefinitionScope = "CATEGORY"
)