// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to list distributions that are associated with a specified AWS
// WAF web ACL.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListDistributionsByWebACLIdRequest
type ListDistributionsByWebACLIdInput struct {
	_ struct{} `type:"structure"`

	// Use Marker and MaxItems to control pagination of results. If you have more
	// than MaxItems distributions that satisfy the request, the response includes
	// a NextMarker element. To get the next page of results, submit another request.
	// For the value of Marker, specify the value of NextMarker from the last response.
	// (For the first request, omit Marker.)
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of distributions that you want CloudFront to return in
	// the response body. The maximum and default values are both 100.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`

	// The ID of the AWS WAF web ACL that you want to list the associated distributions.
	// If you specify "null" for the ID, the request returns a list of the distributions
	// that aren't associated with a web ACL.
	//
	// WebACLId is a required field
	WebACLId *string `location:"uri" locationName:"WebACLId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListDistributionsByWebACLIdInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDistributionsByWebACLIdInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDistributionsByWebACLIdInput"}

	if s.WebACLId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsByWebACLIdInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.WebACLId != nil {
		v := *s.WebACLId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "WebACLId", protocol.StringValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The response to a request to list the distributions that are associated with
// a specified AWS WAF web ACL.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListDistributionsByWebACLIdResult
type ListDistributionsByWebACLIdOutput struct {
	_ struct{} `type:"structure" payload:"DistributionList"`

	// The DistributionList type.
	DistributionList *DistributionList `type:"structure"`
}

// String returns the string representation
func (s ListDistributionsByWebACLIdOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsByWebACLIdOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DistributionList != nil {
		v := s.DistributionList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "DistributionList", v, metadata)
	}
	return nil
}

// UnmarshalAWSXML decodes the AWS API shape using the passed in *xml.Decoder.
func (s *ListDistributionsByWebACLIdOutput) UnmarshalAWSXML(d *xml.Decoder) (err error) {
	defer func() {
		if err != nil {
			*s = ListDistributionsByWebACLIdOutput{}
		}
	}()
	for {
		tok, err := d.Token()
		if tok == nil || err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("fail to UnmarshalAWSXML ListDistributionsByWebACLIdOutput, %s", err)
		}
		start, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		if s.DistributionList == nil {
			s.DistributionList = &DistributionList{}
		}
		err = s.DistributionList.unmarshalAWSXML(d, start)
		if err != nil {
			return fmt.Errorf("fail to UnmarshalAWSXML ListDistributionsByWebACLIdOutput, %s", err)
		}
		return nil
	}
}

const opListDistributionsByWebACLId = "ListDistributionsByWebACLId2019_03_26"

// ListDistributionsByWebACLIdRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// List the distributions that are associated with a specified AWS WAF web ACL.
//
//    // Example sending a request using ListDistributionsByWebACLIdRequest.
//    req := client.ListDistributionsByWebACLIdRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListDistributionsByWebACLId
func (c *Client) ListDistributionsByWebACLIdRequest(input *ListDistributionsByWebACLIdInput) ListDistributionsByWebACLIdRequest {
	op := &aws.Operation{
		Name:       opListDistributionsByWebACLId,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/distributionsByWebACLId/{WebACLId}",
	}

	if input == nil {
		input = &ListDistributionsByWebACLIdInput{}
	}

	req := c.newRequest(op, input, &ListDistributionsByWebACLIdOutput{})
	return ListDistributionsByWebACLIdRequest{Request: req, Input: input, Copy: c.ListDistributionsByWebACLIdRequest}
}

// ListDistributionsByWebACLIdRequest is the request type for the
// ListDistributionsByWebACLId API operation.
type ListDistributionsByWebACLIdRequest struct {
	*aws.Request
	Input *ListDistributionsByWebACLIdInput
	Copy  func(*ListDistributionsByWebACLIdInput) ListDistributionsByWebACLIdRequest
}

// Send marshals and sends the ListDistributionsByWebACLId API request.
func (r ListDistributionsByWebACLIdRequest) Send(ctx context.Context) (*ListDistributionsByWebACLIdResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDistributionsByWebACLIdResponse{
		ListDistributionsByWebACLIdOutput: r.Request.Data.(*ListDistributionsByWebACLIdOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDistributionsByWebACLIdResponse is the response type for the
// ListDistributionsByWebACLId API operation.
type ListDistributionsByWebACLIdResponse struct {
	*ListDistributionsByWebACLIdOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDistributionsByWebACLId request.
func (r *ListDistributionsByWebACLIdResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
