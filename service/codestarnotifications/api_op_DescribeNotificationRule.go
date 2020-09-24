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
	"time"
)

// Returns information about a specified notification rule.
func (c *Client) DescribeNotificationRule(ctx context.Context, params *DescribeNotificationRuleInput, optFns ...func(*Options)) (*DescribeNotificationRuleOutput, error) {
	stack := middleware.NewStack("DescribeNotificationRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeNotificationRuleMiddlewares(stack)
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
	addOpDescribeNotificationRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotificationRule(options.Region), middleware.Before)
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
			OperationName: "DescribeNotificationRule",
			Err:           err,
		}
	}
	out := result.(*DescribeNotificationRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotificationRuleInput struct {
	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string
}

type DescribeNotificationRuleOutput struct {
	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string
	// The date and time the notification rule was most recently updated, in timestamp
	// format.
	LastModifiedTimestamp *time.Time
	// The level of detail included in the notifications for this resource. BASIC will
	// include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar
	// Notifications and/or the service for the resource for which the notification is
	// created.
	DetailType types.DetailType
	// The status of the notification rule. Valid statuses are on (sending
	// notifications) or off (not sending notifications).
	Status types.NotificationRuleStatus
	// A list of the event types associated with the notification rule.
	EventTypes []*types.EventTypeSummary
	// The Amazon Resource Name (ARN) of the resource associated with the notification
	// rule.
	Resource *string
	// The name of the notification rule.
	Name *string
	// The tags associated with the notification rule.
	Tags map[string]*string
	// The date and time the notification rule was created, in timestamp format.
	CreatedTimestamp *time.Time
	// A list of the SNS topics associated with the notification rule.
	Targets []*types.TargetSummary
	// The name or email alias of the person who created the notification rule.
	CreatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeNotificationRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNotificationRule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNotificationRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNotificationRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "DescribeNotificationRule",
	}
}
