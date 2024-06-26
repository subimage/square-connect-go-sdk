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

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type OAuthApiService service

/*
OAuthApiService Authorize
As part of a URL sent to a seller to authorize permissions for  the developer, &#x60;Authorize&#x60; displays an authorization page and a  list of requested permissions. This is not a callable API endpoint.  The completed URL looks similar to the following example: https://connect.squareup.com/oauth2/authorize?client_id&#x3D;{YOUR_APP_ID}&amp;scope&#x3D;CUSTOMERS_WRITE+CUSTOMERS_READ&amp;session&#x3D;False&amp;state&#x3D;82201dd8d83d23cc8a48caf52b  The seller can approve or deny the permissions. If approved,&#x60; Authorize&#x60;  returns an &#x60;AuthorizeResponse&#x60; that is sent to the redirect URL and includes  a state string and an authorization code. The code is used in the &#x60;ObtainToken&#x60;  call to obtain an access token and a refresh token that the developer uses  to manage resources on behalf of the seller.  __Important:__ The &#x60;AuthorizeResponse&#x60; is sent to the redirect URL that you set on  the OAuth page of your application in the Developer Dashboard.  If an error occurs or the seller denies the request, &#x60;Authorize&#x60; returns an  error response that includes &#x60;error&#x60; and &#x60;error_description&#x60; values. If the  error is due to the seller denying the request, the error value is &#x60;access_denied&#x60;  and the &#x60;error_description&#x60; is &#x60;user_denied&#x60;.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientId The Square-issued ID for your application, available from  the OAuth page for your application on the Developer Dashboard.
 * @param optional nil or *OAuthApiAuthorizeOpts - Optional Parameters:
     * @param "Scope" (optional.Interface of OAuthPermission) -  A space-separated list of the permissions that the application is requesting. Default: \&quot;&#x60;MERCHANT_PROFILE_READ PAYMENTS_READ SETTLEMENTS_READ BANK_ACCOUNTS_READ&#x60;\&quot;
     * @param "Locale" (optional.String) -  The locale to present the permission request form in. Square detects the appropriate locale automatically. Only provide this value if the application can definitively determine the preferred locale.  Currently supported values: &#x60;en-IE&#x60;, &#x60;en-US&#x60;, &#x60;en-CA&#x60;, &#x60;es-US&#x60;, &#x60;fr-CA&#x60;, and &#x60;ja-JP&#x60;.
     * @param "Session" (optional.Bool) -  If &#x60;false&#x60;, the user must log in to their Square account to view the Permission Request form, even if they already have a valid user session. This value has no effect in Sandbox. Default: &#x60;true&#x60;
     * @param "State" (optional.String) -  When provided, &#x60;state&#x60; is passed to the configured redirect URL after the Permission Request form is submitted. You can include &#x60;state&#x60; and verify its value to help protect against cross-site request forgery.
     * @param "CodeChallenge" (optional.String) -  When provided, the oauth flow will use PKCE to authorize. The &#x60;code_challenge&#x60; will be associated with the authorization_code and a &#x60;code_verifier&#x60; will need to passed in to obtain the access token.
@return AuthorizeResponse
*/

type OAuthApiAuthorizeOpts struct {
	Scope         optional.Interface
	Locale        optional.String
	Session       optional.Bool
	State         optional.String
	CodeChallenge optional.String
}

func (a *OAuthApiService) Authorize(ctx context.Context, clientId string, localVarOptionals *OAuthApiAuthorizeOpts) (AuthorizeResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue AuthorizeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("client_id", parameterToString(clientId, ""))
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarQueryParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Locale.IsSet() {
		localVarQueryParams.Add("locale", parameterToString(localVarOptionals.Locale.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Session.IsSet() {
		localVarQueryParams.Add("session", parameterToString(localVarOptionals.Session.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarQueryParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallenge.IsSet() {
		localVarQueryParams.Add("code_challenge", parameterToString(localVarOptionals.CodeChallenge.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v AuthorizeResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OAuthApiService ObtainToken
Returns an OAuth access token and a refresh token unless the  &#x60;short_lived&#x60; parameter is set to &#x60;true&#x60;, in which case the endpoint  returns only an access token.  The &#x60;grant_type&#x60; parameter specifies the type of OAuth request. If  &#x60;grant_type&#x60; is &#x60;authorization_code&#x60;, you must include the authorization  code you received when a seller granted you authorization. If &#x60;grant_type&#x60;  is &#x60;refresh_token&#x60;, you must provide a valid refresh token. If you are using  an old version of the Square APIs (prior to March 13, 2019), &#x60;grant_type&#x60;  can be &#x60;migration_token&#x60; and you must provide a valid migration token.  You can use the &#x60;scopes&#x60; parameter to limit the set of permissions granted  to the access token and refresh token. You can use the &#x60;short_lived&#x60; parameter  to create an access token that expires in 24 hours.  __Note:__ OAuth tokens should be encrypted and stored on a secure server.  Application clients should never interact directly with OAuth tokens.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body An object containing the fields to POST for the request.

See the corresponding object definition for field details.
@return ObtainTokenResponse
*/
func (a *OAuthApiService) ObtainToken(ctx context.Context, body ObtainTokenRequest) (ObtainTokenResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ObtainTokenResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ObtainTokenResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OAuthApiService RenewToken
&#x60;RenewToken&#x60; is deprecated. For information about refreshing OAuth access tokens, see [Migrate from Renew to Refresh OAuth Tokens](https://developer.squareup.com/docs/oauth-api/migrate-to-refresh-tokens).  Renews an OAuth access token before it expires.  OAuth access tokens besides your application&#x27;s personal access token expire after 30 days. You can also renew expired tokens within 15 days of their expiration. You cannot renew an access token that has been expired for more than 15 days. Instead, the associated user must recomplete the OAuth flow from the beginning.  __Important:__ The &#x60;Authorization&#x60; header for this endpoint must have the following format:  &#x60;&#x60;&#x60; Authorization: Client APPLICATION_SECRET &#x60;&#x60;&#x60;  Replace &#x60;APPLICATION_SECRET&#x60; with the application secret on the Credentials page in the [Developer Dashboard](https://developer.squareup.com/apps).
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body An object containing the fields to POST for the request.

See the corresponding object definition for field details.
  - @param clientId Your application ID, which is available in the OAuth page in the [Developer Dashboard](https://developer.squareup.com/apps).

@return RenewTokenResponse
*/
func (a *OAuthApiService) RenewToken(ctx context.Context, body RenewTokenRequest, clientId string) (RenewTokenResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RenewTokenResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/clients/{client_id}/access-token/renew"
	localVarPath = strings.Replace(localVarPath, "{"+"client_id"+"}", fmt.Sprintf("%v", clientId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v RenewTokenResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
OAuthApiService RevokeToken
Revokes an access token generated with the OAuth flow.  If an account has more than one OAuth access token for your application, this endpoint revokes all of them, regardless of which token you specify. When an OAuth access token is revoked, all of the active subscriptions associated with that OAuth token are canceled immediately.  __Important:__ The &#x60;Authorization&#x60; header for this endpoint must have the following format:  &#x60;&#x60;&#x60; Authorization: Client APPLICATION_SECRET &#x60;&#x60;&#x60;  Replace &#x60;APPLICATION_SECRET&#x60; with the application secret on the OAuth page for your application on the Developer Dashboard.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body An object containing the fields to POST for the request.

See the corresponding object definition for field details.
@return RevokeTokenResponse
*/
func (a *OAuthApiService) RevokeToken(ctx context.Context, body RevokeTokenRequest) (RevokeTokenResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RevokeTokenResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/revoke"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v RevokeTokenResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
