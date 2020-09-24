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

// Deletes a notebook instance lifecycle configuration.
func (c *Client) DeleteNotebookInstanceLifecycleConfig(ctx context.Context, params *DeleteNotebookInstanceLifecycleConfigInput, optFns ...func(*Options)) (*DeleteNotebookInstanceLifecycleConfigOutput, error) {
	stack := middleware.NewStack("DeleteNotebookInstanceLifecycleConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteNotebookInstanceLifecycleConfigMiddlewares(stack)
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
	addOpDeleteNotebookInstanceLifecycleConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteNotebookInstanceLifecycleConfig(options.Region), middleware.Before)
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
			OperationName: "DeleteNotebookInstanceLifecycleConfig",
			Err:           err,
		}
	}
	out := result.(*DeleteNotebookInstanceLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteNotebookInstanceLifecycleConfigInput struct {
	// The name of the lifecycle configuration to delete.
	NotebookInstanceLifecycleConfigName *string
}

type DeleteNotebookInstanceLifecycleConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteNotebookInstanceLifecycleConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteNotebookInstanceLifecycleConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteNotebookInstanceLifecycleConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteNotebookInstanceLifecycleConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteNotebookInstanceLifecycleConfig",
	}
}
