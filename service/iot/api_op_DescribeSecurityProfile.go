// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets information about a Device Defender security profile.
func (c *Client) DescribeSecurityProfile(ctx context.Context, params *DescribeSecurityProfileInput, optFns ...func(*Options)) (*DescribeSecurityProfileOutput, error) {
	if params == nil {
		params = &DescribeSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSecurityProfile", params, optFns, addOperationDescribeSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityProfileInput struct {

	// The name of the security profile whose information you want to get.
	//
	// This member is required.
	SecurityProfileName *string
}

type DescribeSecurityProfileOutput struct {

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here. Note: This API field is deprecated. Please use
	// DescribeSecurityProfileResponse$additionalMetricsToRetainV2 () instead.
	AdditionalMetricsToRetain []*string

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here.
	AdditionalMetricsToRetainV2 []*types.MetricToRetain

	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]*types.AlertTarget

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []*types.Behavior

	// The time the security profile was created.
	CreationDate *time.Time

	// The time the security profile was last modified.
	LastModifiedDate *time.Time

	// The ARN of the security profile.
	SecurityProfileArn *string

	// A description of the security profile (associated with the security profile when
	// it was created or updated).
	SecurityProfileDescription *string

	// The name of the security profile.
	SecurityProfileName *string

	// The version of the security profile. A new version is generated whenever the
	// security profile is updated.
	Version *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSecurityProfile{}, middleware.After)
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
	addOpDescribeSecurityProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSecurityProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeSecurityProfile",
	}
}
