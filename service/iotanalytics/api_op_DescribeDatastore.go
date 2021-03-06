// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a data store.
func (c *Client) DescribeDatastore(ctx context.Context, params *DescribeDatastoreInput, optFns ...func(*Options)) (*DescribeDatastoreOutput, error) {
	if params == nil {
		params = &DescribeDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDatastore", params, optFns, addOperationDescribeDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatastoreInput struct {

	// The name of the data store
	//
	// This member is required.
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

func addOperationDescribeDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDatastore{}, middleware.After)
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
	addOpDescribeDatastoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatastore(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDatastore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "DescribeDatastore",
	}
}
