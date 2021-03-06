// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops an Amazon RDS DB instance. When you stop a DB instance, Amazon RDS retains
// the DB instance's metadata, including its endpoint, DB parameter group, and
// option group membership. Amazon RDS also retains the transaction logs so you can
// do a point-in-time restore if necessary.  <p>For more information, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StopInstance.html">
// Stopping an Amazon RDS DB Instance Temporarily</a> in the <i>Amazon RDS User
// Guide.</i> </p> <note> <p> This command doesn't apply to Aurora MySQL and Aurora
// PostgreSQL. For Aurora clusters, use <code>StopDBCluster</code> instead. </p>
// </note>
func (c *Client) StopDBInstance(ctx context.Context, params *StopDBInstanceInput, optFns ...func(*Options)) (*StopDBInstanceOutput, error) {
	if params == nil {
		params = &StopDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDBInstance", params, optFns, addOperationStopDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDBInstanceInput struct {

	// The user-supplied instance identifier.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The user-supplied instance identifier of the DB Snapshot created immediately
	// before the DB instance is stopped.
	DBSnapshotIdentifier *string
}

type StopDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStopDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStopDBInstance{}, middleware.After)
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
	addOpStopDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDBInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StopDBInstance",
	}
}
