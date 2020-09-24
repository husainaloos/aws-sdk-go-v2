// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns a description of a processing job.
func (c *Client) DescribeProcessingJob(ctx context.Context, params *DescribeProcessingJobInput, optFns ...func(*Options)) (*DescribeProcessingJobOutput, error) {
	stack := middleware.NewStack("DescribeProcessingJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeProcessingJobMiddlewares(stack)
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
	addOpDescribeProcessingJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProcessingJob(options.Region), middleware.Before)
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
			OperationName: "DescribeProcessingJob",
			Err:           err,
		}
	}
	out := result.(*DescribeProcessingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProcessingJobInput struct {
	// The name of the processing job. The name must be unique within an AWS Region in
	// the AWS account.
	ProcessingJobName *string
}

type DescribeProcessingJobOutput struct {
	// The ARN of a monitoring schedule for an endpoint associated with this processing
	// job.
	MonitoringScheduleArn *string
	// The time at which the processing job started.
	ProcessingStartTime *time.Time
	// The configuration information used to create an experiment.
	ExperimentConfig *types.ExperimentConfig
	// The time limit for how long the processing job is allowed to run.
	StoppingCondition *types.ProcessingStoppingCondition
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy
	// for a processing job. In distributed training, you specify more than one
	// instance.
	ProcessingResources *types.ProcessingResources
	// Networking options for a processing job.
	NetworkConfig *types.NetworkConfig
	// The ARN of an AutoML job associated with this processing job.
	AutoMLJobArn *string
	// A string, up to one KB in size, that contains the reason a processing job
	// failed, if it failed.
	FailureReason *string
	// The ARN of a training job associated with this processing job.
	TrainingJobArn *string
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	RoleArn *string
	// The time at which the processing job completed.
	ProcessingEndTime *time.Time
	// The inputs for a processing job.
	ProcessingInputs []*types.ProcessingInput
	// The name of the processing job. The name must be unique within an AWS Region in
	// the AWS account.
	ProcessingJobName *string
	// The Amazon Resource Name (ARN) of the processing job.
	ProcessingJobArn *string
	// An optional string, up to one KB in size, that contains metadata from the
	// processing container when the processing job exits.
	ExitMessage *string
	// Provides the status of a processing job.
	ProcessingJobStatus types.ProcessingJobStatus
	// Configures the processing job to run a specified container image.
	AppSpecification *types.AppSpecification
	// Output configuration for the processing job.
	ProcessingOutputConfig *types.ProcessingOutputConfig
	// The time at which the processing job was created.
	CreationTime *time.Time
	// The environment variables set in the Docker container.
	Environment map[string]*string
	// The time at which the processing job was last modified.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeProcessingJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProcessingJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProcessingJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProcessingJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeProcessingJob",
	}
}
