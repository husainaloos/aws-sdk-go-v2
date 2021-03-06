// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a notification.
func (c *Client) UpdateNotification(ctx context.Context, params *UpdateNotificationInput, optFns ...func(*Options)) (*UpdateNotificationOutput, error) {
	if params == nil {
		params = &UpdateNotificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNotification", params, optFns, addOperationUpdateNotificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of UpdateNotification
type UpdateNotificationInput struct {

	// The accountId that is associated with the budget whose notification you want to
	// update.
	//
	// This member is required.
	AccountId *string

	// The name of the budget whose notification you want to update.
	//
	// This member is required.
	BudgetName *string

	// The updated notification to be associated with a budget.
	//
	// This member is required.
	NewNotification *types.Notification

	// The previous notification that is associated with a budget.
	//
	// This member is required.
	OldNotification *types.Notification
}

// Response of UpdateNotification
type UpdateNotificationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateNotificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNotification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNotification{}, middleware.After)
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
	addOpUpdateNotificationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotification(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateNotification(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "UpdateNotification",
	}
}
