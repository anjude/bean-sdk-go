package response

import (
	"encoding/json"
	"fmt"
	"github.com/anjude/bcore/pkg/berr"
	"io"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type HttpResponse struct {
	ErrCode int32       `json:"err_code"`
	Msg     string      `json:"msg"`
	Detail  string      `json:"detail"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code      int64  `json:"code"`
	Message   string `json:"message,omitempty"`
	RequestID string `json:"requestID,omitempty"`
}

type BaseResponse struct {
	ErrorResponse
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) error {
	if err := json.Unmarshal(body, r); err != nil {
		return err
	}
	if r.Code > 0 {
		return berr.NewBizError(r.Code, r.Message)
	}
	return nil
}

func ParseFromHttpResponse(rawResponse *http.Response, response Response) error {
	defer rawResponse.Body.Close()
	body, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return err
	}
	if rawResponse.StatusCode != 200 {
		return fmt.Errorf("request fail with status: %s, with body: %s", rawResponse.Status, body)
	}

	if err := response.ParseErrorFromHTTPResponse(body); err != nil {
		return err
	}

	return json.Unmarshal(body, &response)
}
