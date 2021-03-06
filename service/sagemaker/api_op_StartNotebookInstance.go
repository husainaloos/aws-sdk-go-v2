// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Launches an ML compute instance with the latest version of the libraries and
// attaches your ML storage volume. After configuring the notebook instance, Amazon
// SageMaker sets the notebook instance status to InService. A notebook instance's
// status must be InService before you can connect to your Jupyter notebook.
func (c *Client) StartNotebookInstance(ctx context.Context, params *StartNotebookInstanceInput, optFns ...func(*Options)) (*StartNotebookInstanceOutput, error) {
	if params == nil {
		params = &StartNotebookInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartNotebookInstance", params, optFns, addOperationStartNotebookInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartNotebookInstanceInput struct {

	// The name of the notebook instance to start.
	//
	// This member is required.
	NotebookInstanceName *string
}

type StartNotebookInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartNotebookInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartNotebookInstance{}, middleware.After)
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
	addOpStartNotebookInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartNotebookInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartNotebookInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StartNotebookInstance",
	}
}
