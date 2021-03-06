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

// Updates the status of a model version. You can perform the following status
// updates:
//
//     * Change the TRAINING_COMPLETE status to ACTIVE.
//
//     * Change
// ACTIVEto INACTIVE.
func (c *Client) UpdateModelVersionStatus(ctx context.Context, params *UpdateModelVersionStatusInput, optFns ...func(*Options)) (*UpdateModelVersionStatusOutput, error) {
	if params == nil {
		params = &UpdateModelVersionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModelVersionStatus", params, optFns, addOperationUpdateModelVersionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelVersionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateModelVersionStatusInput struct {

	// The model ID of the model version to update.
	//
	// This member is required.
	ModelId *string

	// The model type.
	//
	// This member is required.
	ModelType types.ModelTypeEnum

	// The model version number.
	//
	// This member is required.
	ModelVersionNumber *string

	// The model version status.
	//
	// This member is required.
	Status types.ModelVersionStatus
}

type UpdateModelVersionStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateModelVersionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateModelVersionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateModelVersionStatus{}, middleware.After)
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
	addOpUpdateModelVersionStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateModelVersionStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateModelVersionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "UpdateModelVersionStatus",
	}
}
