package myob

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-myob/odata"
	"github.com/omniboost/go-myob/utils"
)

func (c *Client) NewTaxCodesGetRequest() TaxCodesGetRequest {
	r := TaxCodesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxCodesGetRequest struct {
	client      *Client
	queryParams *TaxCodesGetRequestQueryParams
	pathParams  *TaxCodesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxCodesGetRequestBody
}

func (r TaxCodesGetRequest) NewQueryParams() *TaxCodesGetRequestQueryParams {
	return &TaxCodesGetRequestQueryParams{
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(),
	}
}

type TaxCodesGetRequestQueryParams struct {
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}

func (p TaxCodesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxCodesGetRequest) QueryParams() *TaxCodesGetRequestQueryParams {
	return r.queryParams
}

func (r TaxCodesGetRequest) NewPathParams() *TaxCodesGetRequestPathParams {
	return &TaxCodesGetRequestPathParams{}
}

type TaxCodesGetRequestPathParams struct {
}

func (p *TaxCodesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxCodesGetRequest) PathParams() *TaxCodesGetRequestPathParams {
	return r.pathParams
}

func (r *TaxCodesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxCodesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxCodesGetRequest) Method() string {
	return r.method
}

func (r TaxCodesGetRequest) NewRequestBody() TaxCodesGetRequestBody {
	return TaxCodesGetRequestBody{}
}

type TaxCodesGetRequestBody struct {
}

func (r *TaxCodesGetRequest) RequestBody() *TaxCodesGetRequestBody {
	return nil
}

func (r *TaxCodesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxCodesGetRequest) SetRequestBody(body TaxCodesGetRequestBody) {
	r.requestBody = body
}

func (r *TaxCodesGetRequest) NewResponseBody() *TaxCodesGetResponseBody {
	return &TaxCodesGetResponseBody{}
}

type TaxCodesGetResponseBody struct {
	Items        TaxCodes   `json:"Items"`
	NextPageLink *utils.URL `json:"NextPageLink"`
	Count        int        `json:"Count"`
}

func (r *TaxCodesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/GeneralLedger/TaxCode", r.PathParams())
	return &u
}

func (r *TaxCodesGetRequest) Do(ctx context.Context) (TaxCodesGetResponseBody, error) {
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

func (r *TaxCodesGetRequest) All(ctx context.Context) (TaxCodesGetResponseBody, error) {
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
