// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds, updates, or removes the description of an experiment. Updates the display
// name of an experiment.
func (c *Client) UpdateExperiment(ctx context.Context, params *UpdateExperimentInput, optFns ...func(*Options)) (*UpdateExperimentOutput, error) {
	stack := middleware.NewStack("UpdateExperiment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateExperimentMiddlewares(stack)
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
	addOpUpdateExperimentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateExperiment(options.Region), middleware.Before)
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
			OperationName: "UpdateExperiment",
			Err:           err,
		}
	}
	out := result.(*UpdateExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateExperimentInput struct {
	// The description of the experiment.
	Description *string
	// The name of the experiment to update.
	ExperimentName *string
	// The name of the experiment as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, ExperimentName is displayed.
	DisplayName *string
}

type UpdateExperimentOutput struct {
	// The Amazon Resource Name (ARN) of the experiment.
	ExperimentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateExperimentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateExperiment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateExperiment{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateExperiment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateExperiment",
	}
}
