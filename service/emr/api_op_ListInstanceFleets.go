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

// Lists all available details about the instance fleets in a cluster. The instance
// fleet configuration is available only in Amazon EMR versions 4.8.0 and later,
// excluding 5.0.x versions.
func (c *Client) ListInstanceFleets(ctx context.Context, params *ListInstanceFleetsInput, optFns ...func(*Options)) (*ListInstanceFleetsOutput, error) {
	if params == nil {
		params = &ListInstanceFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstanceFleets", params, optFns, addOperationListInstanceFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstanceFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstanceFleetsInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string
}

type ListInstanceFleetsOutput struct {

	// The list of instance fleets for the cluster and given filters.
	InstanceFleets []*types.InstanceFleet

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInstanceFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInstanceFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInstanceFleets{}, middleware.After)
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
	addOpListInstanceFleetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceFleets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListInstanceFleets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListInstanceFleets",
	}
}
