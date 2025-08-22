package myob

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-myob/odata"
	"github.com/omniboost/go-myob/utils"
)

func (c *Client) NewGeneralJournalsGetRequest() GeneralJournalsGetRequest {
	r := GeneralJournalsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GeneralJournalsGetRequest struct {
	client      *Client
	queryParams *GeneralJournalsGetRequestQueryParams
	pathParams  *GeneralJournalsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody GeneralJournalsGetRequestBody
}

func (r GeneralJournalsGetRequest) NewQueryParams() *GeneralJournalsGetRequestQueryParams {
	return &GeneralJournalsGetRequestQueryParams{
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(),
	}
}

type GeneralJournalsGetRequestQueryParams struct {
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}

func (p GeneralJournalsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GeneralJournalsGetRequest) QueryParams() *GeneralJournalsGetRequestQueryParams {
	return r.queryParams
}

func (r GeneralJournalsGetRequest) NewPathParams() *GeneralJournalsGetRequestPathParams {
	return &GeneralJournalsGetRequestPathParams{}
}

type GeneralJournalsGetRequestPathParams struct {
}

func (p *GeneralJournalsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GeneralJournalsGetRequest) PathParams() *GeneralJournalsGetRequestPathParams {
	return r.pathParams
}

func (r *GeneralJournalsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GeneralJournalsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *GeneralJournalsGetRequest) Method() string {
	return r.method
}

func (r GeneralJournalsGetRequest) NewRequestBody() GeneralJournalsGetRequestBody {
	return GeneralJournalsGetRequestBody{}
}

type GeneralJournalsGetRequestBody struct {
}

func (r *GeneralJournalsGetRequest) RequestBody() *GeneralJournalsGetRequestBody {
	return nil
}

func (r *GeneralJournalsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GeneralJournalsGetRequest) SetRequestBody(body GeneralJournalsGetRequestBody) {
	r.requestBody = body
}

func (r *GeneralJournalsGetRequest) NewResponseBody() *GeneralJournalsGetResponseBody {
	return &GeneralJournalsGetResponseBody{}
}

type GeneralJournalsGetResponseBody struct {
	Items        GeneralJournals `json:"Items"`
	NextPageLink *utils.URL      `json:"NextPageLink"`
	Count        int             `json:"Count"`
}

func (r *GeneralJournalsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/GeneralLedger/GeneralJournal", r.PathParams())
	return &u
}

func (r *GeneralJournalsGetRequest) Do(ctx context.Context) (GeneralJournalsGetResponseBody, error) {
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

func (r *GeneralJournalsGetRequest) All(ctx context.Context) (GeneralJournalsGetResponseBody, error) {
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
