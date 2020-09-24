// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a notification rule for a resource. You can change the events that
// trigger the notification rule, the status of the rule, and the targets that
// receive the notifications. To add or remove tags for a notification rule, you
// must use TagResource () and UntagResource ().
func (c *Client) UpdateNotificationRule(ctx context.Context, params *UpdateNotificationRuleInput, optFns ...func(*Options)) (*UpdateNotificationRuleOutput, error) {
	stack := middleware.NewStack("UpdateNotificationRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateNotificationRuleMiddlewares(stack)
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
	addOpUpdateNotificationRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotificationRule(options.Region), middleware.Before)
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
			OperationName: "UpdateNotificationRule",
			Err:           err,
		}
	}
	out := result.(*UpdateNotificationRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotificationRuleInput struct {
	// The address and type of the targets to receive notifications from this
	// notification rule.
	Targets []*types.Target
	// The level of detail to include in the notifications for this resource. BASIC
	// will include only the contents of the event as it would appear in AWS
	// CloudWatch. FULL will include any supplemental information provided by AWS
	// CodeStar Notifications and/or the service for the resource for which the
	// notification is created.
	DetailType types.DetailType
	// A list of event types associated with this notification rule.
	EventTypeIds []*string
	// The status of the notification rule. Valid statuses include enabled (sending
	// notifications) or disabled (not sending notifications).
	Status types.NotificationRuleStatus
	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string
	// The name of the notification rule.
	Name *string
}

type UpdateNotificationRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateNotificationRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNotificationRule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNotificationRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateNotificationRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "UpdateNotificationRule",
	}
}
