// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the encryption key if a Key Management Service (KMS) customer master key
// (CMK) has been specified to be used to encrypt content in Amazon Fraud Detector.
func (c *Client) GetKMSEncryptionKey(ctx context.Context, params *GetKMSEncryptionKeyInput, optFns ...func(*Options)) (*GetKMSEncryptionKeyOutput, error) {
	if params == nil {
		params = &GetKMSEncryptionKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKMSEncryptionKey", params, optFns, addOperationGetKMSEncryptionKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKMSEncryptionKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKMSEncryptionKeyInput struct {
}

type GetKMSEncryptionKeyOutput struct {

	// The KMS encryption key.
	KmsKey *types.KMSKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetKMSEncryptionKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetKMSEncryptionKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetKMSEncryptionKey{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetKMSEncryptionKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetKMSEncryptionKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetKMSEncryptionKey",
	}
}
