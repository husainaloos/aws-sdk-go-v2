// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers an AWS Batch job definition.
func (c *Client) RegisterJobDefinition(ctx context.Context, params *RegisterJobDefinitionInput, optFns ...func(*Options)) (*RegisterJobDefinitionOutput, error) {
	stack := middleware.NewStack("RegisterJobDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRegisterJobDefinitionMiddlewares(stack)
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
	addOpRegisterJobDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterJobDefinition(options.Region), middleware.Before)
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
			OperationName: "RegisterJobDefinition",
			Err:           err,
		}
	}
	out := result.(*RegisterJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterJobDefinitionInput struct {
	// An object with various properties specific to multi-node parallel jobs. If you
	// specify node properties for a job, it becomes a multi-node parallel job. For
	// more information, see Multi-node Parallel Jobs
	// (https://docs.aws.amazon.com/batch/latest/userguide/multi-node-parallel-jobs.html)
	// in the AWS Batch User Guide. If the job definition's type parameter is
	// container, then you must specify either containerProperties or nodeProperties.
	NodeProperties *types.NodeProperties
	// An object with various properties specific to single-node container-based jobs.
	// If the job definition's type parameter is container, then you must specify
	// either containerProperties or nodeProperties.
	ContainerProperties *types.ContainerProperties
	// The retry strategy to use for failed jobs that are submitted with this job
	// definition. Any retry strategy that is specified during a SubmitJob () operation
	// overrides the retry strategy defined here. If a job is terminated due to a
	// timeout, it is not retried.
	RetryStrategy *types.RetryStrategy
	// The type of job definition.
	Type types.JobDefinitionType
	// Default parameter substitution placeholders to set in the job definition.
	// Parameters are specified as a key-value pair mapping. Parameters in a SubmitJob
	// request override any corresponding parameter defaults from the job definition.
	Parameters map[string]*string
	// The name of the job definition to register. Up to 128 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	JobDefinitionName *string
	// The timeout configuration for jobs that are submitted with this job definition,
	// after which AWS Batch terminates your jobs if they have not finished. If a job
	// is terminated due to a timeout, it is not retried. The minimum value for the
	// timeout is 60 seconds. Any timeout configuration that is specified during a
	// SubmitJob () operation overrides the timeout configuration defined here. For
	// more information, see Job Timeouts
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/job_timeouts.html)
	// in the Amazon Elastic Container Service Developer Guide.
	Timeout *types.JobTimeout
}

type RegisterJobDefinitionOutput struct {
	// The revision of the job definition.
	Revision *int32
	// The name of the job definition.
	JobDefinitionName *string
	// The Amazon Resource Name (ARN) of the job definition.
	JobDefinitionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRegisterJobDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRegisterJobDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterJobDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterJobDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "RegisterJobDefinition",
	}
}
