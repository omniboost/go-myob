package myob

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-myob/odata"
	"github.com/omniboost/go-myob/utils"
)

func (c *Client) NewJobsGetRequest() JobsGetRequest {
	r := JobsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JobsGetRequest struct {
	client      *Client
	queryParams *JobsGetRequestQueryParams
	pathParams  *JobsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JobsGetRequestBody
}

func (r JobsGetRequest) NewQueryParams() *JobsGetRequestQueryParams {
	return &JobsGetRequestQueryParams{
		Filter:  odata.NewFilter(),
		Top:     odata.NewTop(),
		Skip:    odata.NewSkip(),
		OrderBy: odata.NewOrderBy(),
	}
}

type JobsGetRequestQueryParams struct {
	Filter  *odata.Filter  `schema:"$filter,omitempty"`
	Top     *odata.Top     `schema:"$top,omitempty"`
	Skip    *odata.Skip    `schema:"$skip,omitempty"`
	OrderBy *odata.OrderBy `schema:"$orderby,omitempty"`
}

func (p JobsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JobsGetRequest) QueryParams() *JobsGetRequestQueryParams {
	return r.queryParams
}

func (r JobsGetRequest) NewPathParams() *JobsGetRequestPathParams {
	return &JobsGetRequestPathParams{}
}

type JobsGetRequestPathParams struct {
}

func (p *JobsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JobsGetRequest) PathParams() *JobsGetRequestPathParams {
	return r.pathParams
}

func (r *JobsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JobsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JobsGetRequest) Method() string {
	return r.method
}

func (r JobsGetRequest) NewRequestBody() JobsGetRequestBody {
	return JobsGetRequestBody{}
}

type JobsGetRequestBody struct {
}

func (r *JobsGetRequest) RequestBody() *JobsGetRequestBody {
	return nil
}

func (r *JobsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JobsGetRequest) SetRequestBody(body JobsGetRequestBody) {
	r.requestBody = body
}

func (r *JobsGetRequest) NewResponseBody() *JobsGetResponseBody {
	return &JobsGetResponseBody{}
}

type JobsGetResponseBody struct {
	Items        Jobs       `json:"Items"`
	NextPageLink *utils.URL `json:"NextPageLink"`
	Count        int        `json:"Count"`
}

func (r *JobsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/GeneralLedger/Job", r.PathParams())
	return &u
}

func (r *JobsGetRequest) Do(ctx context.Context) (JobsGetResponseBody, error) {
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

func (r *JobsGetRequest) All(ctx context.Context) (JobsGetResponseBody, error) {
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
