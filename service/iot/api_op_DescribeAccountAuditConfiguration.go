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
)

// Gets information about the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks are
// enabled or disabled.
func (c *Client) DescribeAccountAuditConfiguration(ctx context.Context, params *DescribeAccountAuditConfigurationInput, optFns ...func(*Options)) (*DescribeAccountAuditConfigurationOutput, error) {
	stack := middleware.NewStack("DescribeAccountAuditConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAccountAuditConfigurationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAuditConfiguration(options.Region), middleware.Before)

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
			OperationName: "DescribeAccountAuditConfiguration",
			Err:           err,
		}
	}
	out := result.(*DescribeAccountAuditConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAuditConfigurationInput struct {
}

type DescribeAccountAuditConfigurationOutput struct {
	// The ARN of the role that grants permission to AWS IoT to access information
	// about your devices, policies, certificates, and other items as required when
	// performing an audit. On the first call to UpdateAccountAuditConfiguration, this
	// parameter is required.
	RoleArn *string
	// Which audit checks are enabled and disabled for this account.
	AuditCheckConfigurations map[string]*types.AuditCheckConfiguration
	// Information about the targets to which audit notifications are sent for this
	// account.
	AuditNotificationTargetConfigurations map[string]*types.AuditNotificationTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAccountAuditConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountAuditConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountAuditConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccountAuditConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeAccountAuditConfiguration",
	}
}