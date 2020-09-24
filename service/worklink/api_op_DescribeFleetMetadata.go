// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Provides basic information for the specified fleet, excluding identity provider,
// networking, and device configuration details.
func (c *Client) DescribeFleetMetadata(ctx context.Context, params *DescribeFleetMetadataInput, optFns ...func(*Options)) (*DescribeFleetMetadataOutput, error) {
	stack := middleware.NewStack("DescribeFleetMetadata", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeFleetMetadataMiddlewares(stack)
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
	addOpDescribeFleetMetadataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetMetadata(options.Region), middleware.Before)
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
			OperationName: "DescribeFleetMetadata",
			Err:           err,
		}
	}
	out := result.(*DescribeFleetMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetMetadataInput struct {
	// The Amazon Resource Name (ARN) of the fleet.
	FleetArn *string
}

type DescribeFleetMetadataOutput struct {
	// The time that the fleet was created.
	CreatedTime *time.Time
	// The name of the fleet.
	FleetName *string
	// The current state of the fleet.
	FleetStatus types.FleetStatus
	// The time that the fleet was last updated.
	LastUpdatedTime *time.Time
	// The option to optimize for better performance by routing traffic through the
	// closest AWS Region to users, which may be outside of your home Region.
	OptimizeForEndUserLocation *bool
	// The name to display.
	DisplayName *string
	// The tags attached to the resource. A tag is a key-value pair.
	Tags map[string]*string
	// The identifier used by users to sign in to the Amazon WorkLink app.
	CompanyCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeFleetMetadataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFleetMetadata{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFleetMetadata{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFleetMetadata(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DescribeFleetMetadata",
	}
}
