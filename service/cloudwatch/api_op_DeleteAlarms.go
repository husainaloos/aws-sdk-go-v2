// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified alarms. You can delete up to 100 alarms in one operation.
// However, this total can include no more than one composite alarm. For example,
// you could delete 99 metric alarms and one composite alarms with one operation,
// but you can't delete two composite alarms with one operation. In the event of an
// error, no alarms are deleted. It is possible to create a loop or cycle of
// composite alarms, where composite alarm A depends on composite alarm B, and
// composite alarm B also depends on composite alarm A. In this scenario, you can't
// delete any composite alarm that is part of the cycle because there is always
// still a composite alarm that depends on that alarm that you want to delete. To
// get out of such a situation, you must break the cycle by changing the rule of
// one of the composite alarms in the cycle to remove a dependency that creates the
// cycle. The simplest change to make to break a cycle is to change the AlarmRule
// of one of the alarms to False. Additionally, the evaluation of composite alarms
// stops if CloudWatch detects a cycle in the evaluation path.
func (c *Client) DeleteAlarms(ctx context.Context, params *DeleteAlarmsInput, optFns ...func(*Options)) (*DeleteAlarmsOutput, error) {
	stack := middleware.NewStack("DeleteAlarms", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteAlarmsMiddlewares(stack)
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
	addOpDeleteAlarmsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAlarms(options.Region), middleware.Before)
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
			OperationName: "DeleteAlarms",
			Err:           err,
		}
	}
	out := result.(*DeleteAlarmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAlarmsInput struct {
	// The alarms to be deleted.
	AlarmNames []*string
}

type DeleteAlarmsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteAlarmsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAlarms{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAlarms{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAlarms(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DeleteAlarms",
	}
}
