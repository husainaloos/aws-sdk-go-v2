// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables lifecycle management by creating a new LifecycleConfiguration object. A
// LifecycleConfiguration object defines when files in an Amazon EFS file system
// are automatically transitioned to the lower-cost EFS Infrequent Access (IA)
// storage class. A LifecycleConfiguration applies to all files in a file system.
// Each Amazon EFS file system supports one lifecycle configuration, which applies
// to all files in the file system. If a LifecycleConfiguration object already
// exists for the specified file system, a PutLifecycleConfiguration call modifies
// the existing configuration. A PutLifecycleConfiguration call with an empty
// LifecyclePolicies array in the request body deletes any existing
// LifecycleConfiguration and disables lifecycle management.  <p>In the request,
// specify the following: </p> <ul> <li> <p>The ID for the file system for which
// you are enabling, disabling, or modifying lifecycle management.</p> </li> <li>
// <p>A <code>LifecyclePolicies</code> array of <code>LifecyclePolicy</code>
// objects that define when files are moved to the IA storage class. The array can
// contain only one <code>LifecyclePolicy</code> item.</p> </li> </ul> <p>This
// operation requires permissions for the
// <code>elasticfilesystem:PutLifecycleConfiguration</code> operation.</p> <p>To
// apply a <code>LifecycleConfiguration</code> object to an encrypted file system,
// you need the same AWS Key Management Service (AWS KMS) permissions as when you
// created the encrypted file system. </p>
func (c *Client) PutLifecycleConfiguration(ctx context.Context, params *PutLifecycleConfigurationInput, optFns ...func(*Options)) (*PutLifecycleConfigurationOutput, error) {
	stack := middleware.NewStack("PutLifecycleConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutLifecycleConfigurationMiddlewares(stack)
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
	addOpPutLifecycleConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutLifecycleConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleConfigurationInput struct {
	// The ID of the file system for which you are creating the LifecycleConfiguration
	// object (String).
	FileSystemId *string
	// An array of LifecyclePolicy objects that define the file system's
	// LifecycleConfiguration object. A LifecycleConfiguration object tells lifecycle
	// management when to transition files from the Standard storage class to the
	// Infrequent Access storage class.
	LifecyclePolicies []*types.LifecyclePolicy
}

type PutLifecycleConfigurationOutput struct {
	// An array of lifecycle management policies. Currently, EFS supports a maximum of
	// one policy per file system.
	LifecyclePolicies []*types.LifecyclePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutLifecycleConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutLifecycleConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLifecycleConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutLifecycleConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "PutLifecycleConfiguration",
	}
}
