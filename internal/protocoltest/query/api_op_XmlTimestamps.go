// Code generated by smithy-go-codegen DO NOT EDIT.
package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This tests how timestamps are serialized, including using the default format of
// date-time and various @timestampFormat trait values.
func (c *Client) XmlTimestamps(ctx context.Context, params *XmlTimestampsInput, optFns ...func(*Options)) (*XmlTimestampsOutput, error) {
	stack := middleware.NewStack("XmlTimestamps", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "XmlTimestamps",
	}, middleware.Before)
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "XmlTimestamps",
			Err:           err,
		}
	}
	out := result.(*XmlTimestampsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlTimestampsInput struct {
}

type XmlTimestampsOutput struct {
	Normal       *time.Time
	DateTime     *time.Time
	EpochSeconds *time.Time
	HttpDate     *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
