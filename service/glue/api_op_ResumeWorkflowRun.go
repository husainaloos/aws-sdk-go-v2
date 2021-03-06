// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Restarts any completed nodes in a workflow run and resumes the run execution.
func (c *Client) ResumeWorkflowRun(ctx context.Context, params *ResumeWorkflowRunInput, optFns ...func(*Options)) (*ResumeWorkflowRunOutput, error) {
	if params == nil {
		params = &ResumeWorkflowRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResumeWorkflowRun", params, optFns, addOperationResumeWorkflowRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResumeWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeWorkflowRunInput struct {

	// The name of the workflow to resume.
	//
	// This member is required.
	Name *string

	// A list of the node IDs for the nodes you want to restart. The nodes that are to
	// be restarted must have an execution attempt in the original run.
	//
	// This member is required.
	NodeIds []*string

	// The ID of the workflow run to resume.
	//
	// This member is required.
	RunId *string
}

type ResumeWorkflowRunOutput struct {

	// A list of the node IDs for the nodes that were actually restarted.
	NodeIds []*string

	// The new ID assigned to the resumed workflow run. Each resume of a workflow run
	// will have a new run ID.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResumeWorkflowRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResumeWorkflowRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResumeWorkflowRun{}, middleware.After)
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
	addOpResumeWorkflowRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResumeWorkflowRun(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResumeWorkflowRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ResumeWorkflowRun",
	}
}
