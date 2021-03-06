// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides information about the bootstrap actions associated with a cluster.
func (c *Client) ListBootstrapActions(ctx context.Context, params *ListBootstrapActionsInput, optFns ...func(*Options)) (*ListBootstrapActionsOutput, error) {
	if params == nil {
		params = &ListBootstrapActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBootstrapActions", params, optFns, addOperationListBootstrapActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBootstrapActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which bootstrap actions to retrieve.
type ListBootstrapActionsInput struct {

	// The cluster identifier for the bootstrap actions to list.
	//
	// This member is required.
	ClusterId *string

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string
}

// This output contains the bootstrap actions detail.
type ListBootstrapActionsOutput struct {

	// The bootstrap actions associated with the cluster.
	BootstrapActions []*types.Command

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBootstrapActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBootstrapActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBootstrapActions{}, middleware.After)
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
	addOpListBootstrapActionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBootstrapActions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListBootstrapActions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListBootstrapActions",
	}
}
