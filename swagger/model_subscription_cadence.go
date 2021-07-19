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

// SubscriptionCadence : Determines the billing cadence of a [Subscription](entity:Subscription)
type SubscriptionCadence string

// List of SubscriptionCadence
const (
	SUBSCRIPTION_CADENCE_DO_NOT_USE_SubscriptionCadence SubscriptionCadence = "SUBSCRIPTION_CADENCE_DO_NOT_USE"
	DAILY_SubscriptionCadence                           SubscriptionCadence = "DAILY"
	WEEKLY_SubscriptionCadence                          SubscriptionCadence = "WEEKLY"
	EVERY_TWO_WEEKS_SubscriptionCadence                 SubscriptionCadence = "EVERY_TWO_WEEKS"
	THIRTY_DAYS_SubscriptionCadence                     SubscriptionCadence = "THIRTY_DAYS"
	SIXTY_DAYS_SubscriptionCadence                      SubscriptionCadence = "SIXTY_DAYS"
	NINETY_DAYS_SubscriptionCadence                     SubscriptionCadence = "NINETY_DAYS"
	MONTHLY_SubscriptionCadence                         SubscriptionCadence = "MONTHLY"
	EVERY_TWO_MONTHS_SubscriptionCadence                SubscriptionCadence = "EVERY_TWO_MONTHS"
	QUARTERLY_SubscriptionCadence                       SubscriptionCadence = "QUARTERLY"
	EVERY_FOUR_MONTHS_SubscriptionCadence               SubscriptionCadence = "EVERY_FOUR_MONTHS"
	EVERY_SIX_MONTHS_SubscriptionCadence                SubscriptionCadence = "EVERY_SIX_MONTHS"
	ANNUAL_SubscriptionCadence                          SubscriptionCadence = "ANNUAL"
	EVERY_TWO_YEARS_SubscriptionCadence                 SubscriptionCadence = "EVERY_TWO_YEARS"
)
