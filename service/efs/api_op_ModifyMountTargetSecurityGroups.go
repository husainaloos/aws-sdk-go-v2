// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the set of security groups in effect for a mount target. When you
// create a mount target, Amazon EFS also creates a new network interface. For more
// information, see CreateMountTarget (). This operation replaces the security
// groups in effect for the network interface associated with a mount target, with
// the SecurityGroups provided in the request. This operation requires that the
// network interface of the mount target has been created and the lifecycle state
// of the mount target is not deleted. The operation requires permissions for the
// following actions:
//
//     * elasticfilesystem:ModifyMountTargetSecurityGroups
// action on the mount target's file system.
//
//     *
// ec2:ModifyNetworkInterfaceAttribute action on the mount target's network
// interface.
func (c *Client) ModifyMountTargetSecurityGroups(ctx context.Context, params *ModifyMountTargetSecurityGroupsInput, optFns ...func(*Options)) (*ModifyMountTargetSecurityGroupsOutput, error) {
	stack := middleware.NewStack("ModifyMountTargetSecurityGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpModifyMountTargetSecurityGroupsMiddlewares(stack)
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
	addOpModifyMountTargetSecurityGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyMountTargetSecurityGroups(options.Region), middleware.Before)
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
			OperationName: "ModifyMountTargetSecurityGroups",
			Err:           err,
		}
	}
	out := result.(*ModifyMountTargetSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyMountTargetSecurityGroupsInput struct {
	// An array of up to five VPC security group IDs.
	SecurityGroups []*string
	// The ID of the mount target whose security groups you want to modify.
	MountTargetId *string
}

type ModifyMountTargetSecurityGroupsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpModifyMountTargetSecurityGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpModifyMountTargetSecurityGroups{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpModifyMountTargetSecurityGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyMountTargetSecurityGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "ModifyMountTargetSecurityGroups",
	}
}
