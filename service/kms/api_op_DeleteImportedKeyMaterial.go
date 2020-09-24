// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes key material that you previously imported. This operation makes the
// specified customer master key (CMK) unusable. For more information about
// importing key material into AWS KMS, see Importing Key Material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in
// the AWS Key Management Service Developer Guide. You cannot perform this
// operation on a CMK in a different AWS account. When the specified CMK is in the
// PendingDeletion state, this operation does not change the CMK's state.
// Otherwise, it changes the CMK's state to PendingImport. After you delete key
// material, you can use ImportKeyMaterial () to reimport the same key material
// into the CMK. The CMK that you use for this operation must be in a compatible
// key state. For details, see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) DeleteImportedKeyMaterial(ctx context.Context, params *DeleteImportedKeyMaterialInput, optFns ...func(*Options)) (*DeleteImportedKeyMaterialOutput, error) {
	stack := middleware.NewStack("DeleteImportedKeyMaterial", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteImportedKeyMaterialMiddlewares(stack)
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
	addOpDeleteImportedKeyMaterialValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteImportedKeyMaterial(options.Region), middleware.Before)
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
			OperationName: "DeleteImportedKeyMaterial",
			Err:           err,
		}
	}
	out := result.(*DeleteImportedKeyMaterialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteImportedKeyMaterialInput struct {
	// Identifies the CMK from which you are deleting imported key material. The Origin
	// of the CMK must be EXTERNAL. Specify the key ID or the Amazon Resource Name
	// (ARN) of the CMK. For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys () or DescribeKey ().
	KeyId *string
}

type DeleteImportedKeyMaterialOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteImportedKeyMaterialMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteImportedKeyMaterial{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteImportedKeyMaterial{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteImportedKeyMaterial(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DeleteImportedKeyMaterial",
	}
}
