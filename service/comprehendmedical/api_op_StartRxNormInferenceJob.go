// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an asynchronous job to detect medication entities and link them to the
// RxNorm ontology. Use the DescribeRxNormInferenceJob operation to track the
// status of a job.
func (c *Client) StartRxNormInferenceJob(ctx context.Context, params *StartRxNormInferenceJobInput, optFns ...func(*Options)) (*StartRxNormInferenceJobOutput, error) {
	stack := middleware.NewStack("StartRxNormInferenceJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartRxNormInferenceJobMiddlewares(stack)
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
	addIdempotencyToken_opStartRxNormInferenceJobMiddleware(stack, options)
	addOpStartRxNormInferenceJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartRxNormInferenceJob(options.Region), middleware.Before)
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
			OperationName: "StartRxNormInferenceJob",
			Err:           err,
		}
	}
	out := result.(*StartRxNormInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRxNormInferenceJobInput struct {
	// Specifies where to send the output files.
	OutputDataConfig *types.OutputDataConfig
	// Specifies the format and location of the input data for the job.
	InputDataConfig *types.InputDataConfig
	// An AWS Key Management Service key to encrypt your output files. If you do not
	// specify a key, the files are written in plain text.
	KMSKey *string
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend Medical read access to your input data. For
	// more information, see  Role-Based Permissions Required for Asynchronous
	// Operations
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions-med.html#auth-role-permissions-med).
	DataAccessRoleArn *string
	// The language of the input documents. All documents must be in the same language.
	LanguageCode types.LanguageCode
	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend Medical generates one.
	ClientRequestToken *string
	// The identifier of the job.
	JobName *string
}

type StartRxNormInferenceJobOutput struct {
	// The identifier of the job.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartRxNormInferenceJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartRxNormInferenceJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartRxNormInferenceJob{}, middleware.After)
}

type idempotencyToken_initializeOpStartRxNormInferenceJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartRxNormInferenceJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartRxNormInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartRxNormInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartRxNormInferenceJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartRxNormInferenceJobMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartRxNormInferenceJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartRxNormInferenceJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "StartRxNormInferenceJob",
	}
}
