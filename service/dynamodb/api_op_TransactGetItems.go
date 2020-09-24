// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// TransactGetItems is a synchronous operation that atomically retrieves multiple
// items from one or more tables (but not from indexes) in a single account and
// Region. A TransactGetItems call can contain up to 25 TransactGetItem objects,
// each of which contains a Get structure that specifies an item to retrieve from a
// table in the account and Region. A call to TransactGetItems cannot retrieve
// items from tables in more than one AWS account or Region. The aggregate size of
// the items in the transaction cannot exceed 4 MB. DynamoDB rejects the entire
// TransactGetItems request if any of the following is true:
//
//     * A conflicting
// operation is in the process of updating an item to be read.
//
//     * There is
// insufficient provisioned capacity for the transaction to be completed.
//
//     *
// There is a user error, such as an invalid data format.
//
//     * The aggregate size
// of the items in the transaction cannot exceed 4 MB.
func (c *Client) TransactGetItems(ctx context.Context, params *TransactGetItemsInput, optFns ...func(*Options)) (*TransactGetItemsOutput, error) {
	stack := middleware.NewStack("TransactGetItems", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpTransactGetItemsMiddlewares(stack)
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
	addOpTransactGetItemsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTransactGetItems(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)

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
			OperationName: "TransactGetItems",
			Err:           err,
		}
	}
	out := result.(*TransactGetItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TransactGetItemsInput struct {
	// An ordered array of up to 25 TransactGetItem objects, each of which contains a
	// Get structure.
	TransactItems []*types.TransactGetItem
	// A value of TOTAL causes consumed capacity information to be returned, and a
	// value of NONE prevents that information from being returned. No other value is
	// valid.
	ReturnConsumedCapacity types.ReturnConsumedCapacity
}

type TransactGetItemsOutput struct {
	// If the ReturnConsumedCapacity value was TOTAL, this is an array of
	// ConsumedCapacity objects, one for each table addressed by TransactGetItem
	// objects in the TransactItems parameter. These ConsumedCapacity objects report
	// the read-capacity units consumed by the TransactGetItems call in that table.
	ConsumedCapacity []*types.ConsumedCapacity
	// An ordered array of up to 25 ItemResponse objects, each of which corresponds to
	// the TransactGetItem object in the same position in the TransactItems array. Each
	// ItemResponse object contains a Map of the name-value pairs that are the
	// projected attributes of the requested item. If a requested item could not be
	// retrieved, the corresponding ItemResponse object is Null, or if the requested
	// item has no projected attributes, the corresponding ItemResponse object is an
	// empty Map.
	Responses []*types.ItemResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpTransactGetItemsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpTransactGetItems{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpTransactGetItems{}, middleware.After)
}

func newServiceMetadataMiddleware_opTransactGetItems(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "TransactGetItems",
	}
}
