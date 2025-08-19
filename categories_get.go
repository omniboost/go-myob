package myob

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-myob/odata"
	"github.com/omniboost/go-myob/utils"
)

func (c *Client) NewCategoriesGetRequest() CategoriesGetRequest {
	r := CategoriesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CategoriesGetRequest struct {
	client      *Client
	queryParams *CategoriesGetRequestQueryParams
	pathParams  *CategoriesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CategoriesGetRequestBody
}

func (r CategoriesGetRequest) NewQueryParams() *CategoriesGetRequestQueryParams {
	return &CategoriesGetRequestQueryParams{
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(),
	}
}

type CategoriesGetRequestQueryParams struct {
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}

func (p CategoriesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CategoriesGetRequest) QueryParams() *CategoriesGetRequestQueryParams {
	return r.queryParams
}

func (r CategoriesGetRequest) NewPathParams() *CategoriesGetRequestPathParams {
	return &CategoriesGetRequestPathParams{}
}

type CategoriesGetRequestPathParams struct {
}

func (p *CategoriesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CategoriesGetRequest) PathParams() *CategoriesGetRequestPathParams {
	return r.pathParams
}

func (r *CategoriesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CategoriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CategoriesGetRequest) Method() string {
	return r.method
}

func (r CategoriesGetRequest) NewRequestBody() CategoriesGetRequestBody {
	return CategoriesGetRequestBody{}
}

type CategoriesGetRequestBody struct {
}

func (r *CategoriesGetRequest) RequestBody() *CategoriesGetRequestBody {
	return nil
}

func (r *CategoriesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CategoriesGetRequest) SetRequestBody(body CategoriesGetRequestBody) {
	r.requestBody = body
}

func (r *CategoriesGetRequest) NewResponseBody() *CategoriesGetResponseBody {
	return &CategoriesGetResponseBody{}
}

type CategoriesGetResponseBody struct {
	Items        Categories `json:"Items"`
	NextPageLink *utils.URL `json:"NextPageLink"`
	Count        int        `json:"Count"`
}

func (r *CategoriesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/GeneralLedger/Category", r.PathParams())
	return &u
}

func (r *CategoriesGetRequest) Do(ctx context.Context) (CategoriesGetResponseBody, error) {
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

func (r *CategoriesGetRequest) All(ctx context.Context) (CategoriesGetResponseBody, error) {
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
