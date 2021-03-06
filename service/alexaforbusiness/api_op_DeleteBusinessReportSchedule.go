// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the recurring report delivery schedule with the specified schedule ARN.
func (c *Client) DeleteBusinessReportSchedule(ctx context.Context, params *DeleteBusinessReportScheduleInput, optFns ...func(*Options)) (*DeleteBusinessReportScheduleOutput, error) {
	if params == nil {
		params = &DeleteBusinessReportScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBusinessReportSchedule", params, optFns, addOperationDeleteBusinessReportScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBusinessReportScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBusinessReportScheduleInput struct {

	// The ARN of the business report schedule.
	//
	// This member is required.
	ScheduleArn *string
}

type DeleteBusinessReportScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBusinessReportScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBusinessReportSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBusinessReportSchedule{}, middleware.After)
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
	addOpDeleteBusinessReportScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBusinessReportSchedule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBusinessReportSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "DeleteBusinessReportSchedule",
	}
}
