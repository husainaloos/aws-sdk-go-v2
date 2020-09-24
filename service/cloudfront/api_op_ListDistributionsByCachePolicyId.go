// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of distribution IDs for distributions that have a cache behavior
// that’s associated with the specified cache policy. You can optionally specify
// the maximum number of items to receive in the response. If the total number of
// items in the list exceeds the maximum that you specify, or the default maximum,
// the response is paginated. To get the next page of items, send a subsequent
// request that specifies the NextMarker value from the current response as the
// Marker value in the subsequent request.
func (c *Client) ListDistributionsByCachePolicyId(ctx context.Context, params *ListDistributionsByCachePolicyIdInput, optFns ...func(*Options)) (*ListDistributionsByCachePolicyIdOutput, error) {
	stack := middleware.NewStack("ListDistributionsByCachePolicyId", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListDistributionsByCachePolicyIdMiddlewares(stack)
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
	addOpListDistributionsByCachePolicyIdValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributionsByCachePolicyId(options.Region), middleware.Before)
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
			OperationName: "ListDistributionsByCachePolicyId",
			Err:           err,
		}
	}
	out := result.(*ListDistributionsByCachePolicyIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDistributionsByCachePolicyIdInput struct {
	// Use this field when paginating results to indicate where to begin in your list
	// of distribution IDs. The response includes distribution IDs in the list that
	// occur after the marker. To get the next page of the list, set this field’s value
	// to the value of NextMarker from the current page’s response.
	Marker *string
	// The maximum number of distribution IDs that you want in the response.
	MaxItems *string
	// The ID of the cache policy whose associated distribution IDs you want to list.
	CachePolicyId *string
}

type ListDistributionsByCachePolicyIdOutput struct {
	// A list of distribution IDs.
	DistributionIdList *types.DistributionIdList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListDistributionsByCachePolicyIdMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListDistributionsByCachePolicyId{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListDistributionsByCachePolicyId{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDistributionsByCachePolicyId(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListDistributionsByCachePolicyId",
	}
}
