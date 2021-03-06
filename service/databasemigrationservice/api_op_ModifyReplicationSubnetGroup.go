// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the settings for the specified replication subnet group.
func (c *Client) ModifyReplicationSubnetGroup(ctx context.Context, params *ModifyReplicationSubnetGroupInput, optFns ...func(*Options)) (*ModifyReplicationSubnetGroupOutput, error) {
	if params == nil {
		params = &ModifyReplicationSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReplicationSubnetGroup", params, optFns, addOperationModifyReplicationSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReplicationSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyReplicationSubnetGroupInput struct {

	// The name of the replication instance subnet group.
	//
	// This member is required.
	ReplicationSubnetGroupIdentifier *string

	// A list of subnet IDs.
	//
	// This member is required.
	SubnetIds []*string

	// A description for the replication instance subnet group.
	ReplicationSubnetGroupDescription *string
}

//
type ModifyReplicationSubnetGroupOutput struct {

	// The modified replication subnet group.
	ReplicationSubnetGroup *types.ReplicationSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyReplicationSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyReplicationSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyReplicationSubnetGroup{}, middleware.After)
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
	addOpModifyReplicationSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationSubnetGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyReplicationSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ModifyReplicationSubnetGroup",
	}
}
