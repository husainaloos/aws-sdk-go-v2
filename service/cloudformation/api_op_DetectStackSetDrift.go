// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detect drift on a stack set. When CloudFormation performs drift detection on a
// stack set, it performs drift detection on the stack associated with each stack
// instance in the stack set. For more information, see How CloudFormation Performs
// Drift Detection on a Stack Set
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-drift.html).
// DetectStackSetDrift returns the OperationId of the stack set drift detection
// operation. Use this operation id with DescribeStackSetOperation () to monitor
// the progress of the drift detection operation. The drift detection operation may
// take some time, depending on the number of stack instances included in the stack
// set, as well as the number of resources included in each stack. Once the
// operation has completed, use the following actions to return drift
// information:
//
//     * Use DescribeStackSet () to return detailed informaiton about
// the stack set, including detailed information about the last completed drift
// operation performed on the stack set. (Information about drift operations that
// are in progress is not included.)
//
//     * Use ListStackInstances () to return a
// list of stack instances belonging to the stack set, including the drift status
// and last drift time checked of each instance.
//
//     * Use DescribeStackInstance
// () to return detailed information about a specific stack instance, including its
// drift status and last drift time checked.
//
// For more information on performing a
// drift detection operation on a stack set, see Detecting Unmanaged Changes in
// Stack Sets
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-drift.html).
// You can only run a single drift detection operation on a given stack set at one
// time. To stop a drift detection stack set operation, use StopStackSetOperation
// ().
func (c *Client) DetectStackSetDrift(ctx context.Context, params *DetectStackSetDriftInput, optFns ...func(*Options)) (*DetectStackSetDriftOutput, error) {
	stack := middleware.NewStack("DetectStackSetDrift", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDetectStackSetDriftMiddlewares(stack)
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
	addIdempotencyToken_opDetectStackSetDriftMiddleware(stack, options)
	addOpDetectStackSetDriftValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectStackSetDrift(options.Region), middleware.Before)
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
			OperationName: "DetectStackSetDrift",
			Err:           err,
		}
	}
	out := result.(*DetectStackSetDriftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectStackSetDriftInput struct {
	// The name of the stack set on which to perform the drift detection operation.
	StackSetName *string
	// The ID of the stack set operation.
	OperationId *string
	// The user-specified preferences for how AWS CloudFormation performs a stack set
	// operation. For more information on maximum concurrent accounts and failure
	// tolerance, see Stack set operation options
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stackset-ops-options).
	OperationPreferences *types.StackSetOperationPreferences
}

type DetectStackSetDriftOutput struct {
	// The ID of the drift detection stack set operation. you can use this operation id
	// with DescribeStackSetOperation () to monitor the progress of the drift detection
	// operation.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDetectStackSetDriftMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDetectStackSetDrift{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDetectStackSetDrift{}, middleware.After)
}

type idempotencyToken_initializeOpDetectStackSetDrift struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDetectStackSetDrift) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDetectStackSetDrift) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DetectStackSetDriftInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DetectStackSetDriftInput ")
	}

	if input.OperationId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.OperationId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDetectStackSetDriftMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpDetectStackSetDrift{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDetectStackSetDrift(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DetectStackSetDrift",
	}
}
