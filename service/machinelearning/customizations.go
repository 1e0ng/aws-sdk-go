package machinelearning

import (
	"net/url"

	"github.com/awslabs/aws-sdk-go/aws"
)

func init() {
	initRequest = func(r *aws.Request) {
		r.Handlers.Build.PushBack(updatePredictEndpoint)
	}
}

// updatePredictEndpoint rewrites the request endpoint to use the
// "PredictEndpoint" parameter of the Predict operation.
func updatePredictEndpoint(r *aws.Request) {
	if !r.ParamsFilled() {
		return
	}

	r.Endpoint = *r.Params.(*PredictInput).PredictEndpoint

	uri, err := url.Parse(r.Endpoint)
	if err != nil {
		r.Error = err
		return
	}
	r.HTTPRequest.URL = uri
}