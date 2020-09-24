// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the resolvers that are associated with a specific function.
func (c *Client) ListResolversByFunction(ctx context.Context, params *ListResolversByFunctionInput, optFns ...func(*Options)) (*ListResolversByFunctionOutput, error) {
	stack := middleware.NewStack("ListResolversByFunction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListResolversByFunctionMiddlewares(stack)
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
	addOpListResolversByFunctionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResolversByFunction(options.Region), middleware.Before)
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
			OperationName: "ListResolversByFunction",
			Err:           err,
		}
	}
	out := result.(*ListResolversByFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolversByFunctionInput struct {
	// The Function ID.
	FunctionId *string
	// The API ID.
	ApiId *string
	// An identifier that was returned from the previous call to this operation, which
	// you can use to return the next set of items in the list.
	NextToken *string
	// The maximum number of results you want the request to return.
	MaxResults *int32
}

type ListResolversByFunctionOutput struct {
	// The list of resolvers.
	Resolvers []*types.Resolver
	// An identifier that can be used to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListResolversByFunctionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListResolversByFunction{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListResolversByFunction{}, middleware.After)
}

func newServiceMetadataMiddleware_opListResolversByFunction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "ListResolversByFunction",
	}
}
