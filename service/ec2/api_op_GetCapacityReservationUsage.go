// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets usage information about a Capacity Reservation. If the Capacity Reservation
// is shared, it shows usage information for the Capacity Reservation owner and
// each AWS account that is currently using the shared capacity. If the Capacity
// Reservation is not shared, it shows only the Capacity Reservation owner's usage.
func (c *Client) GetCapacityReservationUsage(ctx context.Context, params *GetCapacityReservationUsageInput, optFns ...func(*Options)) (*GetCapacityReservationUsageOutput, error) {
	stack := middleware.NewStack("GetCapacityReservationUsage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpGetCapacityReservationUsageMiddlewares(stack)
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
	addOpGetCapacityReservationUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCapacityReservationUsage(options.Region), middleware.Before)
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
			OperationName: "GetCapacityReservationUsage",
			Err:           err,
		}
	}
	out := result.(*GetCapacityReservationUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCapacityReservationUsageInput struct {
	// The ID of the Capacity Reservation.
	CapacityReservationId *string
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error. Valid range: Minimum value of 1.
	// Maximum value of 1000.
	MaxResults *int32
	// The token to use to retrieve the next page of results.
	NextToken *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type GetCapacityReservationUsageOutput struct {
	// The number of instances for which the Capacity Reservation reserves capacity.
	TotalInstanceCount *int32
	// The ID of the Capacity Reservation.
	CapacityReservationId *string
	// The type of instance for which the Capacity Reservation reserves capacity.
	InstanceType *string
	// The current state of the Capacity Reservation. A Capacity Reservation can be in
	// one of the following states:
	//
	//     * active - The Capacity Reservation is active
	// and the capacity is available for your use.
	//
	//     * expired - The Capacity
	// Reservation expired automatically at the date and time specified in your
	// request. The reserved capacity is no longer available for your use.
	//
	//     *
	// cancelled - The Capacity Reservation was manually cancelled. The reserved
	// capacity is no longer available for your use.
	//
	//     * pending - The Capacity
	// Reservation request was successful but the capacity provisioning is still
	// pending.
	//
	//     * failed - The Capacity Reservation request has failed. A request
	// might fail due to invalid request parameters, capacity constraints, or instance
	// limit constraints. Failed requests are retained for 60 minutes.
	State types.CapacityReservationState
	// The remaining capacity. Indicates the number of instances that can be launched
	// in the Capacity Reservation.
	AvailableInstanceCount *int32
	// Information about the Capacity Reservation usage.
	InstanceUsages []*types.InstanceUsage
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpGetCapacityReservationUsageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpGetCapacityReservationUsage{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpGetCapacityReservationUsage{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCapacityReservationUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetCapacityReservationUsage",
	}
}
