// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the specified configuration.
func (c *Client) DescribeConfiguration(ctx context.Context, params *DescribeConfigurationInput, optFns ...func(*Options)) (*DescribeConfigurationOutput, error) {
	stack := middleware.NewStack("DescribeConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeConfigurationMiddlewares(stack)
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
	addOpDescribeConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfiguration(options.Region), middleware.Before)
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
			OperationName: "DescribeConfiguration",
			Err:           err,
		}
	}
	out := result.(*DescribeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationInput struct {
	// The unique ID that Amazon MQ generates for the configuration.
	ConfigurationId *string
}

type DescribeConfigurationOutput struct {
	// Required. The name of the configuration. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~).
	// This value must be 1-150 characters long.
	Name *string
	// The list of all tags associated with this configuration.
	Tags map[string]*string
	// Required. The type of broker engine. Note: Currently, Amazon MQ supports only
	// ACTIVEMQ.
	EngineType types.EngineType
	// The authentication strategy associated with the configuration.
	AuthenticationStrategy types.AuthenticationStrategy
	// Required. The latest revision of the configuration.
	LatestRevision *types.ConfigurationRevision
	// Required. The ARN of the configuration.
	Arn *string
	// Required. The date and time of the configuration revision.
	Created *time.Time
	// Required. The version of the broker engine. For a list of supported engine
	// versions, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string
	// Required. The description of the configuration.
	Description *string
	// Required. The unique ID that Amazon MQ generates for the configuration.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeConfiguration",
	}
}
