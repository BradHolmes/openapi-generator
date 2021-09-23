/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

func inArray(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

type FakeClassnameTags123Api interface {

	/*
	TestClassname To test class name in snake case

	To test class name in snake case

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiTestClassnameRequest
	*/
	TestClassname(ctx _context.Context) ApiTestClassnameRequest

	// TestClassnameExecute executes the request
	//  @return Client
	TestClassnameExecute(r ApiTestClassnameRequest) (Client, *_nethttp.Response, error)
}

// FakeClassnameTags123ApiService FakeClassnameTags123Api service
type FakeClassnameTags123ApiService service

type ApiTestClassnameRequest struct {
	ctx _context.Context
	ApiService FakeClassnameTags123Api
	body *Client
	accept *string
}

// client model
func (r *ApiTestClassnameRequest) Body(body Client) *ApiTestClassnameRequest {
	r.body = &body
	return r
}

func (r ApiTestClassnameRequest) Execute() (Client, *_nethttp.Response, error) {
	return r.ApiService.TestClassnameExecute(r)
}

/*
TestClassname To test class name in snake case

To test class name in snake case
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTestClassnameRequest
*/
func (a *FakeClassnameTags123ApiService) TestClassname(ctx _context.Context, body *Client, accept string) ApiTestClassnameRequest {
	return ApiTestClassnameRequest{
		ApiService: a,
		ctx: ctx,
		body: body,
		accept: accept
	}
}

// Execute executes the request
//  @return Client
func (a *FakeClassnameTags123ApiService) TestClassnameExecute(r ApiTestClassnameRequest) (Client, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FakeClassnameTags123ApiService.TestClassname")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fake_classname_test"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}
	if inArray(r.accept, localVarHTTPHeaderAccepts) {
		localVarHTTPHeaderAccepts := []string {r.accept}
	}	
	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_query"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key_query", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
