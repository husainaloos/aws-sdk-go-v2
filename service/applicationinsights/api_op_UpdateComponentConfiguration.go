// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the monitoring configurations for the component. The configuration input
// parameter is an escaped JSON of the configuration and should match the schema of
// what is returned by DescribeComponentConfigurationRecommendation.
func (c *Client) UpdateComponentConfiguration(ctx context.Context, params *UpdateComponentConfigurationInput, optFns ...func(*Options)) (*UpdateComponentConfigurationOutput, error) {
	stack := middleware.NewStack("UpdateComponentConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateComponentConfigurationMiddlewares(stack)
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
	addOpUpdateComponentConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComponentConfiguration(options.Region), middleware.Before)
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
			OperationName: "UpdateComponentConfiguration",
			Err:           err,
		}
	}
	out := result.(*UpdateComponentConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateComponentConfigurationInput struct {
	// The tier of the application component. Supported tiers include DOT_NET_WORKER,
	// DOT_NET_WEB, DOT_NET_CORE, SQL_SERVER, and DEFAULT.
	Tier types.Tier
	// The name of the component.
	ComponentName *string
	// The name of the resource group.
	ResourceGroupName *string
	// Indicates whether the application component is monitored.
	Monitor *bool
	// The configuration settings of the component. The value is the escaped JSON of
	// the configuration. For more information about the JSON format, see Working with
	// JSON
	// (https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/working-with-json.html).
	// You can send a request to DescribeComponentConfigurationRecommendation to see
	// the recommended configuration for a component. For the complete format of the
	// component configuration file, see Component Configuration
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config.html).
	ComponentConfiguration *string
}

type UpdateComponentConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateComponentConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateComponentConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateComponentConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateComponentConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "UpdateComponentConfiguration",
	}
}
