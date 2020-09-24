// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables or disables the publishing of reputation metrics for emails sent using a
// specific configuration set in a given AWS Region. Reputation metrics include
// bounce and complaint rates. These metrics are published to Amazon CloudWatch. By
// using CloudWatch, you can create alarms when bounce or complaint rates exceed
// certain thresholds. You can execute this operation no more than once per second.
func (c *Client) UpdateConfigurationSetReputationMetricsEnabled(ctx context.Context, params *UpdateConfigurationSetReputationMetricsEnabledInput, optFns ...func(*Options)) (*UpdateConfigurationSetReputationMetricsEnabledOutput, error) {
	stack := middleware.NewStack("UpdateConfigurationSetReputationMetricsEnabled", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateConfigurationSetReputationMetricsEnabledMiddlewares(stack)
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
	addOpUpdateConfigurationSetReputationMetricsEnabledValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfigurationSetReputationMetricsEnabled(options.Region), middleware.Before)
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
			OperationName: "UpdateConfigurationSetReputationMetricsEnabled",
			Err:           err,
		}
	}
	out := result.(*UpdateConfigurationSetReputationMetricsEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to modify the reputation metric publishing settings for a
// configuration set.
type UpdateConfigurationSetReputationMetricsEnabledInput struct {
	// The name of the configuration set that you want to update.
	ConfigurationSetName *string
	// Describes whether or not Amazon SES will publish reputation metrics for the
	// configuration set, such as bounce and complaint rates, to Amazon CloudWatch.
	Enabled *bool
}

type UpdateConfigurationSetReputationMetricsEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateConfigurationSetReputationMetricsEnabledMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateConfigurationSetReputationMetricsEnabled{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateConfigurationSetReputationMetricsEnabled{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateConfigurationSetReputationMetricsEnabled(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "UpdateConfigurationSetReputationMetricsEnabled",
	}
}
