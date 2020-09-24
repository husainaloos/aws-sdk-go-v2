// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastoredata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets the headers for an object at the specified path.
func (c *Client) DescribeObject(ctx context.Context, params *DescribeObjectInput, optFns ...func(*Options)) (*DescribeObjectOutput, error) {
	stack := middleware.NewStack("DescribeObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeObjectMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeObject(options.Region), middleware.Before)
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
			OperationName: "DescribeObject",
			Err:           err,
		}
	}
	out := result.(*DescribeObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeObjectInput struct {
	// The path (including the file name) where the object is stored in the container.
	// Format: //
	Path *string
}

type DescribeObjectOutput struct {
	// The date and time that the object was last modified.
	LastModified *time.Time
	// The content type of the object.
	ContentType *string
	// An optional CacheControl header that allows the caller to control the object's
	// cache behavior. Headers can be passed in as specified in the HTTP at
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
	// (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9). Headers with
	// a custom user-defined value are also accepted.
	CacheControl *string
	// The ETag that represents a unique instance of the object.
	ETag *string
	// The length of the object in bytes.
	ContentLength *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "DescribeObject",
	}
}
