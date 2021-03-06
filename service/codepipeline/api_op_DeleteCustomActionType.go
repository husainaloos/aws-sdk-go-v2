// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Marks a custom action as deleted. PollForJobs for the custom action fails after
// the action is marked for deletion. Used for custom actions only. To re-create a
// custom action after it has been deleted you must use a string in the version
// field that has never been used before. This string can be an incremented version
// number, for example. To restore a deleted custom action, use a JSON file that is
// identical to the deleted action, including the original string in the version
// field.
func (c *Client) DeleteCustomActionType(ctx context.Context, params *DeleteCustomActionTypeInput, optFns ...func(*Options)) (*DeleteCustomActionTypeOutput, error) {
	if params == nil {
		params = &DeleteCustomActionTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomActionType", params, optFns, addOperationDeleteCustomActionTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomActionTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteCustomActionType operation. The custom action
// will be marked as deleted.
type DeleteCustomActionTypeInput struct {

	// The category of the custom action that you want to delete, such as source or
	// deploy.
	//
	// This member is required.
	Category types.ActionCategory

	// The provider of the service used in the custom action, such as AWS CodeDeploy.
	//
	// This member is required.
	Provider *string

	// The version of the custom action to delete.
	//
	// This member is required.
	Version *string
}

type DeleteCustomActionTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCustomActionTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCustomActionType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCustomActionType{}, middleware.After)
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
	addOpDeleteCustomActionTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomActionType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCustomActionType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "DeleteCustomActionType",
	}
}
