// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for DeleteCustomerGateway.
type DeleteCustomerGatewayInput struct {
	_ struct{} `type:"structure"`

	// The ID of the customer gateway.
	//
	// CustomerGatewayId is a required field
	CustomerGatewayId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s DeleteCustomerGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCustomerGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCustomerGatewayInput"}

	if s.CustomerGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomerGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCustomerGatewayOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCustomerGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCustomerGateway = "DeleteCustomerGateway"

// DeleteCustomerGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified customer gateway. You must delete the VPN connection
// before you can delete the customer gateway.
//
//    // Example sending a request using DeleteCustomerGatewayRequest.
//    req := client.DeleteCustomerGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteCustomerGateway
func (c *Client) DeleteCustomerGatewayRequest(input *DeleteCustomerGatewayInput) DeleteCustomerGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteCustomerGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCustomerGatewayInput{}
	}

	req := c.newRequest(op, input, &DeleteCustomerGatewayOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteCustomerGatewayRequest{Request: req, Input: input, Copy: c.DeleteCustomerGatewayRequest}
}

// DeleteCustomerGatewayRequest is the request type for the
// DeleteCustomerGateway API operation.
type DeleteCustomerGatewayRequest struct {
	*aws.Request
	Input *DeleteCustomerGatewayInput
	Copy  func(*DeleteCustomerGatewayInput) DeleteCustomerGatewayRequest
}

// Send marshals and sends the DeleteCustomerGateway API request.
func (r DeleteCustomerGatewayRequest) Send(ctx context.Context) (*DeleteCustomerGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCustomerGatewayResponse{
		DeleteCustomerGatewayOutput: r.Request.Data.(*DeleteCustomerGatewayOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCustomerGatewayResponse is the response type for the
// DeleteCustomerGateway API operation.
type DeleteCustomerGatewayResponse struct {
	*DeleteCustomerGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCustomerGateway request.
func (r *DeleteCustomerGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
