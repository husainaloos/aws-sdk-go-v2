// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified scaling plan. Deleting a scaling plan deletes the
// underlying ScalingInstruction () for all of the scalable resources that are
// covered by the plan. If the plan has launched resources or has scaling
// activities in progress, you must delete those resources separately.
func (c *Client) DeleteScalingPlan(ctx context.Context, params *DeleteScalingPlanInput, optFns ...func(*Options)) (*DeleteScalingPlanOutput, error) {
	if params == nil {
		params = &DeleteScalingPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteScalingPlan", params, optFns, addOperationDeleteScalingPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteScalingPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteScalingPlanInput struct {

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan.
	//
	// This member is required.
	ScalingPlanVersion *int64
}

type DeleteScalingPlanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteScalingPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteScalingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteScalingPlan{}, middleware.After)
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
	addOpDeleteScalingPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteScalingPlan(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteScalingPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "DeleteScalingPlan",
	}
}
