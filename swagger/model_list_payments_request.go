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

// Describes a request to list payments using  [ListPayments](api-endpoint:Payments-ListPayments).  The maximum results per page is 100.
type ListPaymentsRequest struct {
	// The timestamp for the beginning of the reporting period, in RFC 3339 format. Inclusive. Default: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`
	// The timestamp for the end of the reporting period, in RFC 3339 format.  Default: The current time.
	EndTime string `json:"end_time,omitempty"`
	// The order in which results are listed: - `ASC` - Oldest to newest. - `DESC` - Newest to oldest (default).
	SortOrder string `json:"sort_order,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Limit results to the location supplied. By default, results are returned for the default (main) location associated with the seller.
	LocationId string `json:"location_id,omitempty"`
	// The exact amount in the `total_money` for a payment.
	Total int64 `json:"total,omitempty"`
	// The last four digits of a payment card.
	Last4 string `json:"last_4,omitempty"`
	// The brand of the payment card (for example, VISA).
	CardBrand string `json:"card_brand,omitempty"`
	// The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  The default value of 100 is also the maximum allowed value. If the provided value is  greater than 100, it is ignored and the default value is used instead.  Default: `100`
	Limit int32 `json:"limit,omitempty"`
}
