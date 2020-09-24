// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details for the specified configuration recorders. If the
// configuration recorder is not specified, this action returns the details for all
// configuration recorders associated with the account. Currently, you can specify
// only one configuration recorder per region in your account.
func (c *Client) DescribeConfigurationRecorders(ctx context.Context, params *DescribeConfigurationRecordersInput, optFns ...func(*Options)) (*DescribeConfigurationRecordersOutput, error) {
	stack := middleware.NewStack("DescribeConfigurationRecorders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeConfigurationRecordersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationRecorders(options.Region), middleware.Before)
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
			OperationName: "DescribeConfigurationRecorders",
			Err:           err,
		}
	}
	out := result.(*DescribeConfigurationRecordersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeConfigurationRecorders () action.
type DescribeConfigurationRecordersInput struct {
	// A list of configuration recorder names.
	ConfigurationRecorderNames []*string
}

// The output for the DescribeConfigurationRecorders () action.
type DescribeConfigurationRecordersOutput struct {
	// A list that contains the descriptions of the specified configuration recorders.
	ConfigurationRecorders []*types.ConfigurationRecorder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeConfigurationRecordersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationRecorders{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationRecorders{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeConfigurationRecorders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigurationRecorders",
	}
}
