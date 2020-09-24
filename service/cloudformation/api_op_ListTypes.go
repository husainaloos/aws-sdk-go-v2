// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns summary information about types that have been registered with
// CloudFormation.
func (c *Client) ListTypes(ctx context.Context, params *ListTypesInput, optFns ...func(*Options)) (*ListTypesOutput, error) {
	stack := middleware.NewStack("ListTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListTypesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTypes(options.Region), middleware.Before)
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
			OperationName: "ListTypes",
			Err:           err,
		}
	}
	out := result.(*ListTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypesInput struct {
	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
	// The provisioning behavior of the type. AWS CloudFormation determines the
	// provisioning type during registration, based on the types of handlers in the
	// schema handler package submitted. Valid values include:
	//
	//     * FULLY_MUTABLE:
	// The type includes an update handler to process updates to the type during stack
	// update operations.
	//
	//     * IMMUTABLE: The type does not include an update
	// handler, so the type cannot be updated and must instead be replaced during stack
	// update operations.
	//
	//     * NON_PROVISIONABLE: The type does not include create,
	// read, and delete handlers, and therefore cannot actually be provisioned.
	ProvisioningType types.ProvisioningType
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32
	// The deprecation status of the types that you want to get summary information
	// about. Valid values include:
	//
	//     * LIVE: The type is registered for use in
	// CloudFormation operations.
	//
	//     * DEPRECATED: The type has been deregistered and
	// can no longer be used in CloudFormation operations.
	DeprecatedStatus types.DeprecatedStatus
	// The scope at which the type is visible and usable in CloudFormation operations.
	// Valid values include:
	//
	//     * PRIVATE: The type is only visible and usable within
	// the account in which it is registered. Currently, AWS CloudFormation marks any
	// types you create as PRIVATE.
	//
	//     * PUBLIC: The type is publically visible and
	// usable within any Amazon account.
	//
	// The default is PRIVATE.
	Visibility types.Visibility
}

type ListTypesOutput struct {
	// A list of TypeSummary structures that contain information about the specified
	// types.
	TypeSummaries []*types.TypeSummary
	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call this action again and assign
	// that token to the request object's NextToken parameter. If the request returns
	// all results, NextToken is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListTypes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListTypes",
	}
}
