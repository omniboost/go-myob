package myob

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-myob/odata"
	"github.com/omniboost/go-myob/utils"
)

func (c *Client) NewAccountsGetRequest() AccountsGetRequest {
	r := AccountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountsGetRequest struct {
	client      *Client
	queryParams *AccountsGetRequestQueryParams
	pathParams  *AccountsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountsGetRequestBody
}

func (r AccountsGetRequest) NewQueryParams() *AccountsGetRequestQueryParams {
	return &AccountsGetRequestQueryParams{
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(),
	}
}

type AccountsGetRequestQueryParams struct {
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}

func (p AccountsGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountsGetRequest) QueryParams() *AccountsGetRequestQueryParams {
	return r.queryParams
}

func (r AccountsGetRequest) NewPathParams() *AccountsGetRequestPathParams {
	return &AccountsGetRequestPathParams{}
}

type AccountsGetRequestPathParams struct {
}

func (p *AccountsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountsGetRequest) PathParams() *AccountsGetRequestPathParams {
	return r.pathParams
}

func (r *AccountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountsGetRequest) Method() string {
	return r.method
}

func (r AccountsGetRequest) NewRequestBody() AccountsGetRequestBody {
	return AccountsGetRequestBody{}
}

type AccountsGetRequestBody struct {
}

func (r *AccountsGetRequest) RequestBody() *AccountsGetRequestBody {
	return nil
}

func (r *AccountsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountsGetRequest) SetRequestBody(body AccountsGetRequestBody) {
	r.requestBody = body
}

func (r *AccountsGetRequest) NewResponseBody() *AccountsGetResponseBody {
	return &AccountsGetResponseBody{}
}

type AccountsGetResponseBody struct {
	Items        Accounts   `json:"Items"`
	NextPageLink *utils.URL `json:"NextPageLink"`
	Count        int        `json:"Count"`
}

func (r *AccountsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/GeneralLedger/Account", r.PathParams())
	return &u
}

func (r *AccountsGetRequest) Do(ctx context.Context) (AccountsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *AccountsGetRequest) All(ctx context.Context) (AccountsGetResponseBody, error) {
	responseBody := r.NewResponseBody()

	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return *responseBody, err
		}

		responseBody.Items = append(responseBody.Items, resp.Items...)
		responseBody.Count = resp.Count

		if resp.NextPageLink == nil {
			break
		}

		skip := resp.NextPageLink.Query().Get("$skip")
		skipInt, err := strconv.Atoi(skip)
		if err != nil {
			break
		}
		r.queryParams.Skip.Set(skipInt)

		top := resp.NextPageLink.Query().Get("$top")
		topInt, err := strconv.Atoi(top)
		if err != nil {
			break
		}
		r.queryParams.Top.Set(topInt)
	}

	return *responseBody, nil
}
