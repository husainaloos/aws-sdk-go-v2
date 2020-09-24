// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of image recipes.
func (c *Client) ListImageRecipes(ctx context.Context, params *ListImageRecipesInput, optFns ...func(*Options)) (*ListImageRecipesOutput, error) {
	stack := middleware.NewStack("ListImageRecipes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListImageRecipesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListImageRecipes(options.Region), middleware.Before)
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
			OperationName: "ListImageRecipes",
			Err:           err,
		}
	}
	out := result.(*ListImageRecipesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImageRecipesInput struct {
	// The maximum items to return in a request.
	MaxResults *int32
	// The owner defines which image recipes you want to list. By default, this request
	// will only show image recipes owned by your account. You can use this field to
	// specify if you want to view image recipes owned by yourself, by Amazon, or those
	// image recipes that have been shared with you by other customers.
	Owner types.Ownership
	// The filters.
	Filters []*types.Filter
	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string
}

type ListImageRecipesOutput struct {
	// The request ID that uniquely identifies this request.
	RequestId *string
	// The next token used for paginated responses. When this is not empty, there are
	// additional elements that the service has not included in this request. Use this
	// token with the next request to retrieve additional objects.
	NextToken *string
	// The list of image pipelines.
	ImageRecipeSummaryList []*types.ImageRecipeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListImageRecipesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListImageRecipes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListImageRecipes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListImageRecipes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImageRecipes",
	}
}
