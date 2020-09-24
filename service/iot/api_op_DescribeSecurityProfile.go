// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets information about a Device Defender security profile.
func (c *Client) DescribeSecurityProfile(ctx context.Context, params *DescribeSecurityProfileInput, optFns ...func(*Options)) (*DescribeSecurityProfileOutput, error) {
	stack := middleware.NewStack("DescribeSecurityProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeSecurityProfileMiddlewares(stack)
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
	addOpDescribeSecurityProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityProfile(options.Region), middleware.Before)
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
			OperationName: "DescribeSecurityProfile",
			Err:           err,
		}
	}
	out := result.(*DescribeSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityProfileInput struct {
	// The name of the security profile whose information you want to get.
	SecurityProfileName *string
}

type DescribeSecurityProfileOutput struct {
	// The name of the security profile.
	SecurityProfileName *string
	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here. Note: This API field is deprecated. Please use
	// DescribeSecurityProfileResponse$additionalMetricsToRetainV2 () instead.
	AdditionalMetricsToRetain []*string
	// A description of the security profile (associated with the security profile when
	// it was created or updated).
	SecurityProfileDescription *string
	// The version of the security profile. A new version is generated whenever the
	// security profile is updated.
	Version *int64
	// The time the security profile was created.
	CreationDate *time.Time
	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]*types.AlertTarget
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []*types.Behavior
	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here.
	AdditionalMetricsToRetainV2 []*types.MetricToRetain
	// The time the security profile was last modified.
	LastModifiedDate *time.Time
	// The ARN of the security profile.
	SecurityProfileArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeSecurityProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSecurityProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSecurityProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSecurityProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeSecurityProfile",
	}
}
