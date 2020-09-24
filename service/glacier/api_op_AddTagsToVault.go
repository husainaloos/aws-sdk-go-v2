// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation adds the specified tags to a vault. Each tag is composed of a key
// and a value. Each vault can have up to 10 tags. If your request would cause the
// tag limit for the vault to be exceeded, the operation throws the
// LimitExceededException error. If a tag already exists on the vault under a
// specified key, the existing key value will be overwritten. For more information
// about tags, see Tagging Amazon S3 Glacier Resources
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/tagging.html).
func (c *Client) AddTagsToVault(ctx context.Context, params *AddTagsToVaultInput, optFns ...func(*Options)) (*AddTagsToVaultOutput, error) {
	stack := middleware.NewStack("AddTagsToVault", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAddTagsToVaultMiddlewares(stack)
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
	addOpAddTagsToVaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToVault(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)

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
			OperationName: "AddTagsToVault",
			Err:           err,
		}
	}
	out := result.(*AddTagsToVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input values for AddTagsToVault.
type AddTagsToVaultInput struct {
	// The tags to add to the vault. Each tag is composed of a key and a value. The
	// value can be an empty string.
	Tags map[string]*string
	// The name of the vault.
	VaultName *string
	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	AccountId *string
}

type AddTagsToVaultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAddTagsToVaultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAddTagsToVault{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAddTagsToVault{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddTagsToVault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "AddTagsToVault",
	}
}
