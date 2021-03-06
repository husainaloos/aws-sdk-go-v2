// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the published version of a dashboard.
func (c *Client) UpdateDashboardPublishedVersion(ctx context.Context, params *UpdateDashboardPublishedVersionInput, optFns ...func(*Options)) (*UpdateDashboardPublishedVersionOutput, error) {
	if params == nil {
		params = &UpdateDashboardPublishedVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDashboardPublishedVersion", params, optFns, addOperationUpdateDashboardPublishedVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDashboardPublishedVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDashboardPublishedVersionInput struct {

	// The ID of the AWS account that contains the dashboard that you're updating.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard.
	//
	// This member is required.
	DashboardId *string

	// The version number of the dashboard.
	//
	// This member is required.
	VersionNumber *int64
}

type UpdateDashboardPublishedVersionOutput struct {

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string

	// The ID for the dashboard.
	DashboardId *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDashboardPublishedVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDashboardPublishedVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDashboardPublishedVersion{}, middleware.After)
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
	addOpUpdateDashboardPublishedVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDashboardPublishedVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDashboardPublishedVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDashboardPublishedVersion",
	}
}
