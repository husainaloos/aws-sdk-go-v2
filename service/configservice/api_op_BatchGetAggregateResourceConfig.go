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

// Returns the current configuration items for resources that are present in your
// AWS Config aggregator. The operation also returns a list of resources that are
// not processed in the current request. If there are no unprocessed resources, the
// operation returns an empty unprocessedResourceIdentifiers list.  <note> <ul>
// <li> <p>The API does not return results for deleted resources.</p> </li> <li>
// <p> The API does not return tags and relationships.</p> </li> </ul> </note>
func (c *Client) BatchGetAggregateResourceConfig(ctx context.Context, params *BatchGetAggregateResourceConfigInput, optFns ...func(*Options)) (*BatchGetAggregateResourceConfigOutput, error) {
	stack := middleware.NewStack("BatchGetAggregateResourceConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetAggregateResourceConfigMiddlewares(stack)
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
	addOpBatchGetAggregateResourceConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetAggregateResourceConfig(options.Region), middleware.Before)
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
			OperationName: "BatchGetAggregateResourceConfig",
			Err:           err,
		}
	}
	out := result.(*BatchGetAggregateResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetAggregateResourceConfigInput struct {
	// A list of aggregate ResourceIdentifiers objects.
	ResourceIdentifiers []*types.AggregateResourceIdentifier
	// The name of the configuration aggregator.
	ConfigurationAggregatorName *string
}

type BatchGetAggregateResourceConfigOutput struct {
	// A list of resource identifiers that were not processed with current scope. The
	// list is empty if all the resources are processed.
	UnprocessedResourceIdentifiers []*types.AggregateResourceIdentifier
	// A list that contains the current configuration of one or more resources.
	BaseConfigurationItems []*types.BaseConfigurationItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetAggregateResourceConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetAggregateResourceConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetAggregateResourceConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetAggregateResourceConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "BatchGetAggregateResourceConfig",
	}
}
