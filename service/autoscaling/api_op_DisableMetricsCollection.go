// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables group metrics for the specified Auto Scaling group.
func (c *Client) DisableMetricsCollection(ctx context.Context, params *DisableMetricsCollectionInput, optFns ...func(*Options)) (*DisableMetricsCollectionOutput, error) {
	if params == nil {
		params = &DisableMetricsCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableMetricsCollection", params, optFns, addOperationDisableMetricsCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableMetricsCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableMetricsCollectionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Specifies one or more of the following metrics:
	//
	//     * GroupMinSize
	//
	//     *
	// GroupMaxSize
	//
	//     * GroupDesiredCapacity
	//
	//     * GroupInServiceInstances
	//
	//     *
	// GroupPendingInstances
	//
	//     * GroupStandbyInstances
	//
	//     *
	// GroupTerminatingInstances
	//
	//     * GroupTotalInstances
	//
	//     *
	// GroupInServiceCapacity
	//
	//     * GroupPendingCapacity
	//
	//     * GroupStandbyCapacity
	//
	//
	// * GroupTerminatingCapacity
	//
	//     * GroupTotalCapacity
	//
	// If you omit this
	// parameter, all metrics are disabled.
	Metrics []*string
}

type DisableMetricsCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableMetricsCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableMetricsCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableMetricsCollection{}, middleware.After)
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
	addOpDisableMetricsCollectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableMetricsCollection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisableMetricsCollection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DisableMetricsCollection",
	}
}
