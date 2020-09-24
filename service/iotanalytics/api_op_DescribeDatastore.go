// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a data store.
func (c *Client) DescribeDatastore(ctx context.Context, params *DescribeDatastoreInput, optFns ...func(*Options)) (*DescribeDatastoreOutput, error) {
	stack := middleware.NewStack("DescribeDatastore", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeDatastoreMiddlewares(stack)
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
	addOpDescribeDatastoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatastore(options.Region), middleware.Before)
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
			OperationName: "DescribeDatastore",
			Err:           err,
		}
	}
	out := result.(*DescribeDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatastoreInput struct {
	// The name of the data store
	DatastoreName *string
	// If true, additional statistical information about the data store is included in
	// the response. This feature cannot be used with a data store whose S3 storage is
	// customer-managed.
	IncludeStatistics *bool
}

type DescribeDatastoreOutput struct {
	// Information about the data store.
	Datastore *types.Datastore
	// Additional statistical information about the data store. Included if the
	// 'includeStatistics' parameter is set to true in the request.
	Statistics *types.DatastoreStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeDatastoreMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDatastore{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDatastore{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDatastore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "DescribeDatastore",
	}
}
