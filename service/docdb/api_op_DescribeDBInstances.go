// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about provisioned Amazon DocumentDB instances. This API
// supports pagination.
func (c *Client) DescribeDBInstances(ctx context.Context, params *DescribeDBInstancesInput, optFns ...func(*Options)) (*DescribeDBInstancesOutput, error) {
	if params == nil {
		params = &DescribeDBInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBInstances", params, optFns, addOperationDescribeDBInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBInstances ().
type DescribeDBInstancesInput struct {

	// The user-provided instance identifier. If this parameter is specified,
	// information from only the specific instance is returned. This parameter isn't
	// case sensitive. Constraints:
	//
	//     * If provided, must match the identifier of an
	// existing DBInstance.
	DBInstanceIdentifier *string

	// A filter that specifies one or more instances to describe. Supported filters:
	//
	//
	// * db-cluster-id - Accepts cluster identifiers and cluster Amazon Resource Names
	// (ARNs). The results list includes only the information about the instances that
	// are associated with the clusters that are identified by these ARNs.
	//
	//     *
	// db-instance-id - Accepts instance identifiers and instance ARNs. The results
	// list includes only the information about the instances that are identified by
	// these ARNs.
	Filters []*types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Represents the output of DescribeDBInstances ().
type DescribeDBInstancesOutput struct {

	// Detailed information about one or more instances.
	DBInstances []*types.DBInstance

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBInstances{}, middleware.After)
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
	addOpDescribeDBInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBInstances",
	}
}
