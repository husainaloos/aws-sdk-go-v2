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

// Returns status information for sources within an aggregator. The status includes
// information about the last time AWS Config verified authorization between the
// source account and an aggregator account. In case of a failure, the status
// contains the related error code or message.
func (c *Client) DescribeConfigurationAggregatorSourcesStatus(ctx context.Context, params *DescribeConfigurationAggregatorSourcesStatusInput, optFns ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	stack := middleware.NewStack("DescribeConfigurationAggregatorSourcesStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeConfigurationAggregatorSourcesStatusMiddlewares(stack)
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
	addOpDescribeConfigurationAggregatorSourcesStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(options.Region), middleware.Before)
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
			OperationName: "DescribeConfigurationAggregatorSourcesStatus",
			Err:           err,
		}
	}
	out := result.(*DescribeConfigurationAggregatorSourcesStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationAggregatorSourcesStatusInput struct {
	// Filters the status type.
	//
	//     * Valid value FAILED indicates errors while moving
	// data.
	//
	//     * Valid value SUCCEEDED indicates the data was successfully moved.
	//
	//
	// * Valid value OUTDATED indicates the data is not the most recent.
	UpdateStatus []types.AggregatedSourceStatusType
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The name of the configuration aggregator.
	ConfigurationAggregatorName *string
	// The maximum number of AggregatorSourceStatus returned on each page. The default
	// is maximum. If you specify 0, AWS Config uses the default.
	Limit *int32
}

type DescribeConfigurationAggregatorSourcesStatusOutput struct {
	// Returns an AggregatedSourceStatus object.
	AggregatedSourceStatusList []*types.AggregatedSourceStatus
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeConfigurationAggregatorSourcesStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigurationAggregatorSourcesStatus",
	}
}
