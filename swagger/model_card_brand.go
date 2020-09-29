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

// CardBrand : Indicates a card's brand, such as `VISA` or `MASTERCARD`.
type CardBrand string

// List of CardBrand
const (
	OTHER_BRAND_CardBrand         CardBrand = "OTHER_BRAND"
	VISA_CardBrand                CardBrand = "VISA"
	MASTERCARD_CardBrand          CardBrand = "MASTERCARD"
	AMERICAN_EXPRESS_CardBrand    CardBrand = "AMERICAN_EXPRESS"
	DISCOVER_CardBrand            CardBrand = "DISCOVER"
	DISCOVER_DINERS_CardBrand     CardBrand = "DISCOVER_DINERS"
	JCB_CardBrand                 CardBrand = "JCB"
	CHINA_UNIONPAY_CardBrand      CardBrand = "CHINA_UNIONPAY"
	SQUARE_GIFT_CARD_CardBrand    CardBrand = "SQUARE_GIFT_CARD"
	SQUARE_CAPITAL_CARD_CardBrand CardBrand = "SQUARE_CAPITAL_CARD"
	INTERAC_CardBrand             CardBrand = "INTERAC"
	EFTPOS_CardBrand              CardBrand = "EFTPOS"
	FELICA_CardBrand              CardBrand = "FELICA"
)