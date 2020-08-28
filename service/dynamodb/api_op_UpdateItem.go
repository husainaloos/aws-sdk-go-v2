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

// Edits an existing item's attributes, or adds a new item to the table if it does
// not already exist. You can put, delete, or add attribute values. You can also
// perform a conditional update on an existing item (insert a new attribute
// name-value pair if it doesn't exist, or replace an existing name-value pair if
// it has certain expected attribute values). You can also return the item's
// attribute values in the same UpdateItem operation using the ReturnValues
// parameter.
func (c *Client) UpdateItem(ctx context.Context, params *UpdateItemInput, optFns ...func(*Options)) (*UpdateItemOutput, error) {
	stack := middleware.NewStack("UpdateItem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpUpdateItemMiddlewares(stack)
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
	addOpUpdateItemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateItem(options.Region), middleware.Before)

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
			OperationName: "UpdateItem",
			Err:           err,
		}
	}
	out := result.(*UpdateItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdateItem operation.
type UpdateItemInput struct {
	// This is a legacy parameter. Use UpdateExpression instead. For more information,
	// see AttributeUpdates
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributeUpdates.html)
	// in the Amazon DynamoDB Developer Guide.
	AttributeUpdates map[string]*types.AttributeValueUpdate
	// A condition that must be satisfied in order for a conditional update to succeed.
	// An expression can contain any of the following:
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
	// For more information about condition expressions, see Specifying
	// Conditions
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
	// in the Amazon DynamoDB Developer Guide.) To work around this, you could specify
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
	// which are placeholders for the actual value at runtime. For more information
	// about expression attribute names, see Specifying Item Attributes
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
	// The primary key of the item to be updated. Each element consists of an attribute
	// name and a value for that attribute. For the primary key, you must provide all
	// of the attributes. For example, with a simple primary key, you only need to
	// provide a value for the partition key. For a composite primary key, you must
	// provide values for both the partition key and the sort key.
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
	// Use ReturnValues if you want to get the item attributes as they appear before or
	// after they are updated. For UpdateItem, the valid values are:
	//
	//     * NONE - If
	// ReturnValues is not specified, or if its value is NONE, then nothing is
	// returned. (This setting is the default for ReturnValues.)
	//
	//     * ALL_OLD -
	// Returns all of the attributes of the item, as they appeared before the
	// UpdateItem operation.
	//
	//     * UPDATED_OLD - Returns only the updated attributes,
	// as they appeared before the UpdateItem operation.
	//
	//     * ALL_NEW - Returns all
	// of the attributes of the item, as they appear after the UpdateItem operation.
	//
	//
	// * UPDATED_NEW - Returns only the updated attributes, as they appear after the
	// UpdateItem operation.
	//
	// There is no additional cost associated with requesting a
	// return value aside from the small network and processing overhead of receiving a
	// larger response. No read capacity units are consumed. The values returned are
	// strongly consistent.
	ReturnValues types.ReturnValue
	// The name of the table containing the item to update.
	TableName *string
	// An expression that defines one or more attributes to be updated, the action to
	// be performed on them, and new values for them. The following action values are
	// available for UpdateExpression.
	//
	//     * SET - Adds one or more attributes and
	// values to an item. If any of these attributes already exist, they are replaced
	// by the new values. You can also use SET to add or subtract from an attribute
	// that is of type Number. For example: SET myNum = myNum + :valSET supports the
	// following functions:
	//
	//         * if_not_exists (path, operand) - if the item does
	// not contain an attribute at the specified path, then if_not_exists evaluates to
	// operand; otherwise, it evaluates to path. You can use this function to avoid
	// overwriting an attribute that may already be present in the item.
	//
	//         *
	// list_append (operand, operand) - evaluates to a list with a new element added to
	// it. You can append the new element to the start or the end of the list by
	// reversing the order of the operands.
	//
	//     These function names are
	// case-sensitive.
	//
	//     * REMOVE - Removes one or more attributes from an item.
	//
	//
	// * ADD - Adds the specified value to the item, if the attribute does not already
	// exist. If the attribute does exist, then the behavior of ADD depends on the data
	// type of the attribute:
	//
	//         * If the existing attribute is a number, and if
	// Value is also a number, then Value is mathematically added to the existing
	// attribute. If Value is a negative number, then it is subtracted from the
	// existing attribute. If you use ADD to increment or decrement a number value for
	// an item that doesn't exist before the update, DynamoDB uses 0 as the initial
	// value. Similarly, if you use ADD for an existing item to increment or decrement
	// an attribute value that doesn't exist before the update, DynamoDB uses 0 as the
	// initial value. For example, suppose that the item you want to update doesn't
	// have an attribute named itemcount, but you decide to ADD the number 3 to this
	// attribute anyway. DynamoDB will create the itemcount attribute, set its initial
	// value to 0, and finally add 3 to it. The result will be a new itemcount
	// attribute in the item, with a value of 3.
	//
	//         * If the existing data type
	// is a set and if Value is also a set, then Value is added to the existing set.
	// For example, if the attribute value is the set [1,2], and the ADD action
	// specified [3], then the final attribute value is [1,2,3]. An error occurs if an
	// ADD action is specified for a set attribute and the attribute type specified
	// does not match the existing set type. Both sets must have the same primitive
	// data type. For example, if the existing data type is a set of strings, the Value
	// must also be a set of strings.
	//
	//     The ADD action only supports Number and set
	// data types. In addition, ADD can only be used on top-level attributes, not
	// nested attributes.
	//
	//     * DELETE - Deletes an element from a set. If a set of
	// values is specified, then those values are subtracted from the old set. For
	// example, if the attribute value was the set [a,b,c] and the DELETE action
	// specifies [a,c], then the final attribute value is [b]. Specifying an empty set
	// is an error. The DELETE action only supports set data types. In addition, DELETE
	// can only be used on top-level attributes, not nested attributes.  </li> </ul>
	// <p>You can have many actions in a single expression, such as the following:
	// <code>SET a=:value1, b=:value2 DELETE :value3, :value4, :value5</code> </p>
	// <p>For more information on update expressions, see <a
	// href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.Modifying.html">Modifying
	// Items and Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	UpdateExpression *string
}

// Represents the output of an UpdateItem operation.
type UpdateItemOutput struct {
	// A map of attribute values as they appear before or after the UpdateItem
	// operation, as determined by the ReturnValues parameter. The Attributes map is
	// only present if ReturnValues was specified as something other than NONE in the
	// request. Each element represents one attribute.
	Attributes map[string]*types.AttributeValue
	// The capacity units consumed by the UpdateItem operation. The data returned
	// includes the total provisioned throughput consumed, along with statistics for
	// the table and any indexes involved in the operation. ConsumedCapacity is only
	// returned if the ReturnConsumedCapacity parameter was specified. For more
	// information, see Provisioned Throughput
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *types.ConsumedCapacity
	// Information about item collections, if any, that were affected by the UpdateItem
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

func addawsAwsjson10_serdeOpUpdateItemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateItem{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateItem{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateItem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateItem",
	}
}
