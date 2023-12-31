package request

type Request interface {
	GetURL() string
	GetMethod() string
	GetHeaders() map[string]string
	GetData() interface{}
	GetVersion() string
}

// BaseRequest is the base struct of service requests
type BaseRequest struct {
	URL     string // resource url, i.e. /regions/${regionId}/elasticIps/${elasticIpId}
	Method  string
	Header  map[string]string
	Data    interface{}
	Version string
}

func (r *BaseRequest) GetURL() string {
	return r.URL
}

func (r *BaseRequest) GetMethod() string {
	return r.Method
}

func (r *BaseRequest) GetVersion() string {
	return r.Version
}

func (r *BaseRequest) GetData() interface{} {
	return r
}

func (r *BaseRequest) GetHeaders() map[string]string {
	return r.Header
}

// AddHeader only adds pin or erp, they will be encoded to base64 code
func (r *BaseRequest) AddHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}
