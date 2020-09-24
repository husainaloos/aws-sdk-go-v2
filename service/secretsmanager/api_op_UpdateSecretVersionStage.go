// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the staging labels attached to a version of a secret. Staging labels
// are used to track a version as it progresses through the secret rotation
// process. You can attach a staging label to only one version of a secret at a
// time. If a staging label to be added is already attached to another version,
// then it is moved--removed from the other version first and then attached to this
// one. For more information about staging labels, see Staging Labels
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/terms-concepts.html#term_staging-label)
// in the AWS Secrets Manager User Guide. The staging labels that you specify in
// the VersionStage parameter are added to the existing list of staging
// labels--they don't replace it. You can move the AWSCURRENT staging label to this
// version by including it in this call. Whenever you move AWSCURRENT, Secrets
// Manager automatically moves the label AWSPREVIOUS to the version that AWSCURRENT
// was removed from. If this action results in the last label being removed from a
// version, then the version is considered to be 'deprecated' and can be deleted by
// Secrets Manager. Minimum permissions To run this command, you must have the
// following permissions:
//
//     * secretsmanager:UpdateSecretVersionStage
//
// Related
// operations
//
//     * To get the list of staging labels that are currently
// associated with a version of a secret, use DescribeSecret () and examine the
// SecretVersionsToStages response value.
func (c *Client) UpdateSecretVersionStage(ctx context.Context, params *UpdateSecretVersionStageInput, optFns ...func(*Options)) (*UpdateSecretVersionStageOutput, error) {
	stack := middleware.NewStack("UpdateSecretVersionStage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateSecretVersionStageMiddlewares(stack)
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
	addOpUpdateSecretVersionStageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecretVersionStage(options.Region), middleware.Before)
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
			OperationName: "UpdateSecretVersionStage",
			Err:           err,
		}
	}
	out := result.(*UpdateSecretVersionStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecretVersionStageInput struct {
	// (Optional) The secret version ID that you want to add the staging label. If you
	// want to remove a label from a version, then do not specify this parameter. If
	// the staging label is already attached to a different version of the secret, then
	// you must also specify the RemoveFromVersionId parameter.
	MoveToVersionId *string
	// Specifies the secret with the version with the list of staging labels you want
	// to modify. You can specify either the Amazon Resource Name (ARN) or the friendly
	// name of the secret. If you specify an ARN, we generally recommend that you
	// specify a complete ARN. You can specify a partial ARN too—for example, if you
	// don’t include the final hyphen and six random characters that Secrets Manager
	// adds at the end of the ARN when you created the secret. A partial ARN match can
	// work as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as a
	// partial ARN, then those characters cause Secrets Manager to assume that you’re
	// specifying a complete ARN. This confusion can cause unexpected results. To avoid
	// this situation, we recommend that you don’t create secret names ending with a
	// hyphen followed by six characters. If you specify an incomplete ARN without the
	// random suffix, and instead provide the 'friendly name', you must not include the
	// random suffix. If you do include the random suffix added by Secrets Manager, you
	// receive either a ResourceNotFoundException or an AccessDeniedException error,
	// depending on your permissions.
	SecretId *string
	// Specifies the secret version ID of the version that the staging label is to be
	// removed from. If the staging label you are trying to attach to one version is
	// already attached to a different version, then you must include this parameter
	// and specify the version that the label is to be removed from. If the label is
	// attached and you either do not specify this parameter, or the version ID does
	// not match, then the operation fails.
	RemoveFromVersionId *string
	// The staging label to add to this version.
	VersionStage *string
}

type UpdateSecretVersionStageOutput struct {
	// The ARN of the secret with the modified staging label.
	ARN *string
	// The friendly name of the secret with the modified staging label.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateSecretVersionStageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSecretVersionStage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSecretVersionStage{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSecretVersionStage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "UpdateSecretVersionStage",
	}
}
