// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a global table from an existing table. A global table creates a
// replication relationship between two or more DynamoDB tables with the same table
// name in the provided Regions. This operation only applies to Version 2017.11.29
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
// of global tables.  <p>If you want to add a new replica table to a global table,
// each of the following conditions  must be true:
//
//     * The table must have the
// same primary key as all of the other replicas.
//
//     * The table must have the
// same name as all of the other replicas.
//
//     * The table must have DynamoDB
// Streams enabled, with the stream containing both the new and the old images of
// the item.
//
//     * None of the replica tables in the global table can contain any
// data.
//
// If global secondary indexes are specified, then the following conditions
// must also be met:
//
//     * The global secondary indexes must have the same name.
//
//
// * The global secondary indexes must have the same hash key and sort key (if
// present).
//
// If local secondary indexes are specified, then the following
// conditions must also be met:
//
//     * The local secondary indexes must have the
// same name.
//
//     * The local secondary indexes must have the same hash key and
// sort key (if present).
//
//     <important> <p> Write capacity settings should be
// set consistently across your replica tables and secondary indexes. DynamoDB
// strongly recommends enabling auto scaling to manage the write capacity settings
// for all of your global tables replicas and indexes. </p> <p> If you prefer to
// manage write capacity settings manually, you should provision equal replicated
// write capacity units to your replica tables. You should also provision equal
// replicated write capacity units to matching secondary indexes across your global
// table. </p> </important>
func (c *Client) CreateGlobalTable(ctx context.Context, params *CreateGlobalTableInput, optFns ...func(*Options)) (*CreateGlobalTableOutput, error) {
	if params == nil {
		params = &CreateGlobalTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGlobalTable", params, optFns, addOperationCreateGlobalTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGlobalTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGlobalTableInput struct {

	// The global table name.
	//
	// This member is required.
	GlobalTableName *string

	// The Regions where the global table needs to be created.
	//
	// This member is required.
	ReplicationGroup []*types.Replica
}

type CreateGlobalTableOutput struct {

	// Contains the details of the global table.
	GlobalTableDescription *types.GlobalTableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateGlobalTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateGlobalTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateGlobalTable{}, middleware.After)
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
	addOpCreateGlobalTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGlobalTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opCreateGlobalTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "CreateGlobalTable",
	}
}
