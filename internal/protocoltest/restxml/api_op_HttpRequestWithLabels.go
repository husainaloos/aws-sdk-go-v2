// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests are serialized when there's no input payload but
// there are HTTP labels.
func (c *Client) HttpRequestWithLabels(ctx context.Context, params *HttpRequestWithLabelsInput, optFns ...func(*Options)) (*HttpRequestWithLabelsOutput, error) {
	stack := middleware.NewStack("HttpRequestWithLabels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpHttpRequestWithLabelsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpHttpRequestWithLabelsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpRequestWithLabels(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "HttpRequestWithLabels",
			Err:           err,
		}
	}
	out := result.(*HttpRequestWithLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsInput struct {
	String_ *string
	Short   *int16
	Integer *int32
	Long    *int64
	Float   *float32
	Double  *float64
	// Serialized in the path as true or false.
	Boolean *bool
	// Note that this member has no format, so it's serialized as an RFC 3399
	// date-time.
	Timestamp *time.Time
}

type HttpRequestWithLabelsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHttpRequestWithLabelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHttpRequestWithLabels{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHttpRequestWithLabels{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpRequestWithLabels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithLabels",
	}
}
