// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of tags for the specified AWS CloudHSM cluster. This is a paginated
// operation, which means that each response might contain only a subset of all the
// tags. When the response contains only a subset of tags, it includes a NextToken
// value. Use this value in a subsequent ListTags request to get more tags. When
// you receive a response with no NextToken (or an empty or null value), that means
// there are no more tags to get.
func (c *Client) ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) {
	if params == nil {
		params = &ListTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTags", params, optFns, addOperationListTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsInput struct {

	// The cluster identifier (ID) for the cluster whose tags you are getting. To find
	// the cluster ID, use DescribeClusters ().
	//
	// This member is required.
	ResourceId *string

	// The maximum number of tags to return in the response. When there are more tags
	// than the number you specify, the response contains a NextToken value.
	MaxResults *int32

	// The NextToken value that you received in the previous response. Use this value
	// to get more tags.
	NextToken *string
}

type ListTagsOutput struct {

	// A list of tags.
	//
	// This member is required.
	TagList []*types.Tag

	// An opaque string that indicates that the response contains only a subset of
	// tags. Use this value in a subsequent ListTags request to get more tags.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTags{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTags(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "ListTags",
	}
}
