// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new query API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new query API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new query API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for query API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	QueryAllBalances(params *QueryAllBalancesParams, opts ...ClientOption) (*QueryAllBalancesOK, error)

	QueryBalance(params *QueryBalanceParams, opts ...ClientOption) (*QueryBalanceOK, error)

	QueryDenomMetadata(params *QueryDenomMetadataParams, opts ...ClientOption) (*QueryDenomMetadataOK, error)

	QueryDenomMetadataByQueryString(params *QueryDenomMetadataByQueryStringParams, opts ...ClientOption) (*QueryDenomMetadataByQueryStringOK, error)

	QueryDenomOwners(params *QueryDenomOwnersParams, opts ...ClientOption) (*QueryDenomOwnersOK, error)

	QueryDenomOwnersByQuery(params *QueryDenomOwnersByQueryParams, opts ...ClientOption) (*QueryDenomOwnersByQueryOK, error)

	QueryDenomsMetadata(params *QueryDenomsMetadataParams, opts ...ClientOption) (*QueryDenomsMetadataOK, error)

	QueryParams(params *QueryParamsParams, opts ...ClientOption) (*QueryParamsOK, error)

	QuerySendEnabled(params *QuerySendEnabledParams, opts ...ClientOption) (*QuerySendEnabledOK, error)

	QuerySpendableBalanceByDenom(params *QuerySpendableBalanceByDenomParams, opts ...ClientOption) (*QuerySpendableBalanceByDenomOK, error)

	QuerySpendableBalances(params *QuerySpendableBalancesParams, opts ...ClientOption) (*QuerySpendableBalancesOK, error)

	QuerySupplyOf(params *QuerySupplyOfParams, opts ...ClientOption) (*QuerySupplyOfOK, error)

	QueryTotalSupply(params *QueryTotalSupplyParams, opts ...ClientOption) (*QueryTotalSupplyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
QueryAllBalances query all balances API
*/
func (a *Client) QueryAllBalances(params *QueryAllBalancesParams, opts ...ClientOption) (*QueryAllBalancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryAllBalancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_AllBalances",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/balances/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryAllBalancesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryAllBalancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryAllBalancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryBalance query balance API
*/
func (a *Client) QueryBalance(params *QueryBalanceParams, opts ...ClientOption) (*QueryBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_Balance",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/balances/{address}/by_denom",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryBalanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryBalanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDenomMetadata query denom metadata API
*/
func (a *Client) QueryDenomMetadata(params *QueryDenomMetadataParams, opts ...ClientOption) (*QueryDenomMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDenomMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_DenomMetadata",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/denoms_metadata/{denom}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryDenomMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDenomMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDenomMetadataDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDenomMetadataByQueryString query denom metadata by query string API
*/
func (a *Client) QueryDenomMetadataByQueryString(params *QueryDenomMetadataByQueryStringParams, opts ...ClientOption) (*QueryDenomMetadataByQueryStringOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDenomMetadataByQueryStringParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_DenomMetadataByQueryString",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/denoms_metadata_by_query_string",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryDenomMetadataByQueryStringReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDenomMetadataByQueryStringOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDenomMetadataByQueryStringDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDenomOwners query denom owners API
*/
func (a *Client) QueryDenomOwners(params *QueryDenomOwnersParams, opts ...ClientOption) (*QueryDenomOwnersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDenomOwnersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_DenomOwners",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/denom_owners/{denom}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryDenomOwnersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDenomOwnersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDenomOwnersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDenomOwnersByQuery query denom owners by query API
*/
func (a *Client) QueryDenomOwnersByQuery(params *QueryDenomOwnersByQueryParams, opts ...ClientOption) (*QueryDenomOwnersByQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDenomOwnersByQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_DenomOwnersByQuery",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/denom_owners_by_query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryDenomOwnersByQueryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDenomOwnersByQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDenomOwnersByQueryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDenomsMetadata query denoms metadata API
*/
func (a *Client) QueryDenomsMetadata(params *QueryDenomsMetadataParams, opts ...ClientOption) (*QueryDenomsMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDenomsMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_DenomsMetadata",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/denoms_metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryDenomsMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDenomsMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDenomsMetadataDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryParams query params API
*/
func (a *Client) QueryParams(params *QueryParamsParams, opts ...ClientOption) (*QueryParamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryParamsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_Params",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/params",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryParamsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryParamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryParamsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QuerySendEnabled query send enabled API
*/
func (a *Client) QuerySendEnabled(params *QuerySendEnabledParams, opts ...ClientOption) (*QuerySendEnabledOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySendEnabledParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_SendEnabled",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/send_enabled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuerySendEnabledReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuerySendEnabledOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySendEnabledDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QuerySpendableBalanceByDenom query spendable balance by denom API
*/
func (a *Client) QuerySpendableBalanceByDenom(params *QuerySpendableBalanceByDenomParams, opts ...ClientOption) (*QuerySpendableBalanceByDenomOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySpendableBalanceByDenomParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_SpendableBalanceByDenom",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/spendable_balances/{address}/by_denom",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuerySpendableBalanceByDenomReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuerySpendableBalanceByDenomOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySpendableBalanceByDenomDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QuerySpendableBalances query spendable balances API
*/
func (a *Client) QuerySpendableBalances(params *QuerySpendableBalancesParams, opts ...ClientOption) (*QuerySpendableBalancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySpendableBalancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_SpendableBalances",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/spendable_balances/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuerySpendableBalancesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuerySpendableBalancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySpendableBalancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QuerySupplyOf query supply of API
*/
func (a *Client) QuerySupplyOf(params *QuerySupplyOfParams, opts ...ClientOption) (*QuerySupplyOfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySupplyOfParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_SupplyOf",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/supply/by_denom",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuerySupplyOfReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuerySupplyOfOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySupplyOfDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryTotalSupply query total supply API
*/
func (a *Client) QueryTotalSupply(params *QueryTotalSupplyParams, opts ...ClientOption) (*QueryTotalSupplyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryTotalSupplyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Query_TotalSupply",
		Method:             "GET",
		PathPattern:        "/cosmos/bank/v1beta1/supply",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryTotalSupplyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryTotalSupplyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryTotalSupplyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
