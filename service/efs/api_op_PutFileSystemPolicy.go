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

// Applies an Amazon EFS FileSystemPolicy to an Amazon EFS file system. A file
// system policy is an IAM resource-based policy and can contain multiple policy
// statements. A file system always has exactly one file system policy, which can
// be the default policy or an explicit policy set or updated using this API
// operation. When an explicit policy is set, it overrides the default policy. For
// more information about the default file system policy, see Default EFS File
// System Policy
// (https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html#default-filesystempolicy).
// This operation requires permissions for the
// elasticfilesystem:PutFileSystemPolicy action.
func (c *Client) PutFileSystemPolicy(ctx context.Context, params *PutFileSystemPolicyInput, optFns ...func(*Options)) (*PutFileSystemPolicyOutput, error) {
	stack := middleware.NewStack("PutFileSystemPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutFileSystemPolicyMiddlewares(stack)
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
	addOpPutFileSystemPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutFileSystemPolicy(options.Region), middleware.Before)
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
			OperationName: "PutFileSystemPolicy",
			Err:           err,
		}
	}
	out := result.(*PutFileSystemPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutFileSystemPolicyInput struct {
	// The ID of the EFS file system that you want to create or update the
	// FileSystemPolicy for.
	FileSystemId *string
	// (Optional) A flag to indicate whether to bypass the FileSystemPolicy lockout
	// safety check. The policy lockout safety check determines whether the policy in
	// the request will prevent the principal making the request will be locked out
	// from making future PutFileSystemPolicy requests on the file system. Set
	// BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the
	// principal that is making the request from making a subsequent
	// PutFileSystemPolicy request on the file system. The default value is False.
	BypassPolicyLockoutSafetyCheck *bool
	// The FileSystemPolicy that you're creating. Accepts a JSON formatted policy
	// definition. To find out more about the elements that make up a file system
	// policy, see EFS Resource-based Policies
	// (https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies).
	Policy *string
}

type PutFileSystemPolicyOutput struct {
	// Specifies the EFS file system to which the FileSystemPolicy applies.
	FileSystemId *string
	// The JSON formatted FileSystemPolicy for the EFS file system.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutFileSystemPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutFileSystemPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutFileSystemPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutFileSystemPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "PutFileSystemPolicy",
	}
}
