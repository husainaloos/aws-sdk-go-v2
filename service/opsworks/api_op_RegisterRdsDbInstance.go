// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers an Amazon RDS instance with a stack. Required Permissions: To use this
// action, an IAM user must have a Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) RegisterRdsDbInstance(ctx context.Context, params *RegisterRdsDbInstanceInput, optFns ...func(*Options)) (*RegisterRdsDbInstanceOutput, error) {
	if params == nil {
		params = &RegisterRdsDbInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterRdsDbInstance", params, optFns, addOperationRegisterRdsDbInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterRdsDbInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterRdsDbInstanceInput struct {

	// The database password.
	//
	// This member is required.
	DbPassword *string

	// The database's master user name.
	//
	// This member is required.
	DbUser *string

	// The Amazon RDS instance's ARN.
	//
	// This member is required.
	RdsDbInstanceArn *string

	// The stack ID.
	//
	// This member is required.
	StackId *string
}

type RegisterRdsDbInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterRdsDbInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterRdsDbInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterRdsDbInstance{}, middleware.After)
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
	addOpRegisterRdsDbInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterRdsDbInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterRdsDbInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "RegisterRdsDbInstance",
	}
}
