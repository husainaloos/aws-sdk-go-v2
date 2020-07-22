// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbgov2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a single item in a table by primary key. You can perform a conditional
// delete operation that deletes the item if it exists, or if it has an expected
// attribute value. In addition to deleting an item, you can also return the item's
// attribute values in the same operation, using the ReturnValues parameter. Unless
// you specify conditions, the DeleteItem is an idempotent operation; running it
// multiple times on the same item or attribute does not result in an error
// response. Conditional deletes are useful for deleting items only if specific
// conditions are met. If those conditions are met, DynamoDB performs the delete.
// Otherwise, the item is not deleted.
func (c *Client) DeleteItem(ctx context.Context, params *DeleteItemInput, optFns ...func(*Options)) (*DeleteItemOutput, error) {
	stack := middleware.NewStack("DeleteItem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addOpDeleteItemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteItem(options.Region), middleware.Before)
	addawsAwsjson10_serdeOpDeleteItemMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "DeleteItem",
			Err:           err,
		}
	}
	out := result.(*DeleteItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteItem operation.
type DeleteItemInput struct {
	// A condition that must be satisfied in order for a conditional DeleteItem to
	// succeed. An expression can contain any of the following:
	//
	//     * Functions:
	// attribute_exists | attribute_not_exists | attribute_type | contains |
	// begins_with | size These function names are case-sensitive.
	//
	//     * Comparison
	// operators: = | <> | < | > | <= | >= | BETWEEN | IN
	//
	//     * Logical operators: AND
	// | OR | NOT
	//
	// For more information about condition expressions, see Condition
	// Expressions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionExpression *string
	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see ConditionalOperator
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionalOperator types.ConditionalOperator
	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see Expected
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html)
	// in the Amazon DynamoDB Developer Guide.
	Expected map[string]*types.ExpectedAttributeValue
	// One or more substitution tokens for attribute names in an expression. The
	// following are some use cases for using ExpressionAttributeNames:
	//
	//     * To
	// access an attribute whose name conflicts with a DynamoDB reserved word.
	//
	//     *
	// To create a placeholder for repeating occurrences of an attribute name in an
	// expression.
	//
	//     * To prevent special characters in an attribute name from being
	// misinterpreted in an expression.
	//
	// Use the # character in an expression to
	// dereference an attribute name. For example, consider the following attribute
	// name:
	//
	//     * Percentile
	//
	// The name of this attribute conflicts with a reserved
	// word, so it cannot be used directly in an expression. (For the complete list of
	// reserved words, see Reserved Words
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide). To work around this, you could specify
	// the following for ExpressionAttributeNames:
	//
	//     * {"#P":"Percentile"}
	//
	// You
	// could then use this substitution in an expression, as in this example:
	//
	//     * #P
	// = :val
	//
	// Tokens that begin with the : character are expression attribute values,
	// which are placeholders for the actual value at runtime. For more information on
	// expression attribute names, see Specifying Item Attributes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeNames map[string]*string
	// One or more values that can be substituted in an expression. Use the : (colon)
	// character in an expression to dereference an attribute value. For example,
	// suppose that you wanted to check whether the value of the ProductStatus
	// attribute was one of the following: Available | Backordered | Discontinued You
	// would first need to specify ExpressionAttributeValues as follows: {
	// ":avail":{"S":"Available"}, ":back":{"S":"Backordered"},
	// ":disc":{"S":"Discontinued"} } You could then use these values in an expression,
	// such as this: ProductStatus IN (:avail, :back, :disc) For more information on
	// expression attribute values, see Condition Expressions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeValues map[string]*types.AttributeValue
	// A map of attribute names to AttributeValue objects, representing the primary key
	// of the item to delete. For the primary key, you must provide all of the
	// attributes. For example, with a simple primary key, you only need to provide a
	// value for the partition key. For a composite primary key, you must provide
	// values for both the partition key and the sort key.
	Key map[string]*types.AttributeValue
	// Determines the level of detail about provisioned throughput consumption that is
	// returned in the response:
	//
	//     * INDEXES - The response includes the aggregate
	// ConsumedCapacity for the operation, together with ConsumedCapacity for each
	// table and secondary index that was accessed. Note that some operations, such as
	// GetItem and BatchGetItem, do not access any indexes at all. In these cases,
	// specifying INDEXES will only return ConsumedCapacity information for table(s).
	//
	//
	// * TOTAL - The response includes only the aggregate ConsumedCapacity for the
	// operation.
	//
	//     * NONE - No ConsumedCapacity details are included in the
	// response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity
	// Determines whether item collection metrics are returned. If set to SIZE, the
	// response includes statistics about item collections, if any, that were modified
	// during the operation are returned in the response. If set to NONE (the default),
	// no statistics are returned.
	ReturnItemCollectionMetrics types.ReturnItemCollectionMetrics
	// Use ReturnValues if you want to get the item attributes as they appeared before
	// they were deleted. For DeleteItem, the valid values are:
	//
	//     * NONE - If
	// ReturnValues is not specified, or if its value is NONE, then nothing is
	// returned. (This setting is the default for ReturnValues.)
	//
	//     * ALL_OLD - The
	// content of the old item is returned.
	//
	// The ReturnValues parameter is used by
	// several DynamoDB operations; however, DeleteItem does not recognize any values
	// other than NONE or ALL_OLD.
	ReturnValues types.ReturnValue
	// The name of the table from which to delete the item.
	TableName *string
}

// Represents the output of a DeleteItem operation.
type DeleteItemOutput struct {
	// A map of attribute names to AttributeValue objects, representing the item as it
	// appeared before the DeleteItem operation. This map appears in the response only
	// if ReturnValues was specified as ALL_OLD in the request.
	Attributes map[string]*types.AttributeValue
	// The capacity units consumed by the DeleteItem operation. The data returned
	// includes the total provisioned throughput consumed, along with statistics for
	// the table and any indexes involved in the operation. ConsumedCapacity is only
	// returned if the ReturnConsumedCapacity parameter was specified. For more
	// information, see Provisioned Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *types.ConsumedCapacity
	// Information about item collections, if any, that were affected by the DeleteItem
	// operation. ItemCollectionMetrics is only returned if the
	// ReturnItemCollectionMetrics parameter was specified. If the table does not have
	// any local secondary indexes, this information is not returned in the response.
	// Each ItemCollectionMetrics element consists of:
	//
	//     * ItemCollectionKey - The
	// partition key value of the item collection. This is the same as the partition
	// key value of the item itself.
	//
	//     * SizeEstimateRangeGB - An estimate of item
	// collection size, in gigabytes. This value is a two-element array containing a
	// lower bound and an upper bound for the estimate. The estimate includes the size
	// of all the items in the table, plus the size of all attributes projected into
	// all of the local secondary indexes on that table. Use this estimate to measure
	// whether a local secondary index is approaching its size limit. The estimate is
	// subject to change over time; therefore, do not rely on the precision or accuracy
	// of the estimate.
	ItemCollectionMetrics *types.ItemCollectionMetrics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDeleteItemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteItem{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteItem{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteItem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "DeleteItem",
	}
}
