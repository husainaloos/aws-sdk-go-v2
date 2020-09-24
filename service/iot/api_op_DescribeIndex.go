// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a search index.
func (c *Client) DescribeIndex(ctx context.Context, params *DescribeIndexInput, optFns ...func(*Options)) (*DescribeIndexOutput, error) {
	stack := middleware.NewStack("DescribeIndex", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeIndexMiddlewares(stack)
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
	addOpDescribeIndexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIndex(options.Region), middleware.Before)
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
			OperationName: "DescribeIndex",
			Err:           err,
		}
	}
	out := result.(*DescribeIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIndexInput struct {
	// The index name.
	IndexName *string
}

type DescribeIndexOutput struct {
	// The index status.
	IndexStatus types.IndexStatus
	// The index name.
	IndexName *string
	// Contains a value that specifies the type of indexing performed. Valid values
	// are:
	//
	//     * REGISTRY – Your thing index contains only registry data.
	//
	//     *
	// REGISTRY_AND_SHADOW - Your thing index contains registry data and shadow data.
	//
	//
	// * REGISTRY_AND_CONNECTIVITY_STATUS - Your thing index contains registry data and
	// thing connectivity status data.
	//
	//     *
	// REGISTRY_AND_SHADOW_AND_CONNECTIVITY_STATUS - Your thing index contains registry
	// data, shadow data, and thing connectivity status data.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeIndexMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIndex{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIndex{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeIndex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeIndex",
	}
}
