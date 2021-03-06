// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a task to estimate the quality of the transform.  <p>When you provide
// label sets as examples of truth, AWS Glue machine learning uses some of those
// examples to learn from them. The rest of the labels are used as a test to
// estimate quality.</p> <p>Returns a unique identifier for the run. You can call
// <code>GetMLTaskRun</code> to get more information about the stats of the
// <code>EvaluationTaskRun</code>.</p>
func (c *Client) StartMLEvaluationTaskRun(ctx context.Context, params *StartMLEvaluationTaskRunInput, optFns ...func(*Options)) (*StartMLEvaluationTaskRunOutput, error) {
	if params == nil {
		params = &StartMLEvaluationTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMLEvaluationTaskRun", params, optFns, addOperationStartMLEvaluationTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMLEvaluationTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLEvaluationTaskRunInput struct {

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string
}

type StartMLEvaluationTaskRunOutput struct {

	// The unique identifier associated with this run.
	TaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartMLEvaluationTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMLEvaluationTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMLEvaluationTaskRun{}, middleware.After)
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
	addOpStartMLEvaluationTaskRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLEvaluationTaskRun(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartMLEvaluationTaskRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartMLEvaluationTaskRun",
	}
}
