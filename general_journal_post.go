package myob

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-myob/utils"
)

func (c *Client) NewGeneralJournalPostRequest() GeneralJournalPostRequest {
	r := GeneralJournalPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GeneralJournalPostRequest struct {
	client      *Client
	queryParams *GeneralJournalPostRequestQueryParams
	pathParams  *GeneralJournalPostRequestPathParams
	method      string
	headers     http.Header
	requestBody GeneralJournalPostRequestBody
}

func (r GeneralJournalPostRequest) NewQueryParams() *GeneralJournalPostRequestQueryParams {
	return &GeneralJournalPostRequestQueryParams{}
}

type GeneralJournalPostRequestQueryParams struct {
}

func (p GeneralJournalPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GeneralJournalPostRequest) QueryParams() *GeneralJournalPostRequestQueryParams {
	return r.queryParams
}

func (r GeneralJournalPostRequest) NewPathParams() *GeneralJournalPostRequestPathParams {
	return &GeneralJournalPostRequestPathParams{}
}

type GeneralJournalPostRequestPathParams struct {
}

func (p *GeneralJournalPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GeneralJournalPostRequest) PathParams() *GeneralJournalPostRequestPathParams {
	return r.pathParams
}

func (r *GeneralJournalPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GeneralJournalPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *GeneralJournalPostRequest) Method() string {
	return r.method
}

func (r GeneralJournalPostRequest) NewRequestBody() GeneralJournalPostRequestBody {
	return GeneralJournalPostRequestBody{
		Lines: []GeneralJournalLine{},
	}
}

type GeneralJournalPostRequestBody GeneralJournal

func (r *GeneralJournalPostRequest) RequestBody() *GeneralJournalPostRequestBody {
	return nil
}

func (r *GeneralJournalPostRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GeneralJournalPostRequest) SetRequestBody(body GeneralJournalPostRequestBody) {
	r.requestBody = body
}

func (r *GeneralJournalPostRequest) NewResponseBody() *GeneralJournalPostResponseBody {
	return &GeneralJournalPostResponseBody{}
}

type GeneralJournalPostResponseBody struct {
	Items        Accounts   `json:"Items"`
	NextPageLink *utils.URL `json:"NextPageLink"`
	Count        int        `json:"Count"`
}

func (r *GeneralJournalPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/GeneralLedger/GeneralJournal", r.PathParams())
	return &u
}

func (r *GeneralJournalPostRequest) Do(ctx context.Context) (GeneralJournalPostResponseBody, error) {
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
