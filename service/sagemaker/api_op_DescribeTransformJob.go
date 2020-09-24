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

// Returns information about a transform job.
func (c *Client) DescribeTransformJob(ctx context.Context, params *DescribeTransformJobInput, optFns ...func(*Options)) (*DescribeTransformJobOutput, error) {
	stack := middleware.NewStack("DescribeTransformJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTransformJobMiddlewares(stack)
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
	addOpDescribeTransformJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransformJob(options.Region), middleware.Before)
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
			OperationName: "DescribeTransformJob",
			Err:           err,
		}
	}
	out := result.(*DescribeTransformJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransformJobInput struct {
	// The name of the transform job that you want to view details of.
	TransformJobName *string
}

type DescribeTransformJobOutput struct {
	// The name of the model used in the transform job.
	ModelName *string
	// A timestamp that shows when the transform Job was created.
	CreationTime *time.Time
	// Specifies the number of records to include in a mini-batch for an HTTP inference
	// request. A record is a single unit of input data that inference can be made on.
	// For example, a single line in a CSV file is a record. To enable the batch
	// strategy, you must set SplitType to Line, RecordIO, or TFRecord.
	BatchStrategy types.BatchStrategy
	// The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling job
	// that created the transform or training job.
	LabelingJobArn *string
	// The name of the transform job.
	TransformJobName *string
	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	//     * CreateProcessingJob ()
	//
	//     *
	// CreateTrainingJob ()
	//
	//     * CreateTransformJob ()
	ExperimentConfig *types.ExperimentConfig
	// The maximum payload size, in MB, used in the transform job.
	MaxPayloadInMB *int32
	// If the transform job failed, FailureReason describes why it failed. A transform
	// job creates a log file, which includes error messages, and stores it as an
	// Amazon S3 object. For more information, see Log Amazon SageMaker Events with
	// Amazon CloudWatch
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/logging-cloudwatch.html).
	FailureReason *string
	// The Amazon Resource Name (ARN) of the AutoML transform job.
	AutoMLJobArn *string
	// The status of the transform job. If the transform job failed, the reason is
	// returned in the FailureReason field.
	TransformJobStatus types.TransformJobStatus
	// Indicates when the transform job has been  completed, or has stopped or failed.
	// You are billed for the time interval between this time and the value of
	// <code>TransformStartTime</code>.</p>
	TransformEndTime *time.Time
	// The data structure used to specify the data to be used for inference in a batch
	// transform job and to associate the data that is relevant to the prediction
	// results in the output. The input filter provided allows you to exclude input
	// data that is not needed for inference in a batch transform job. The output
	// filter provided allows you to include input data relevant to interpreting the
	// predictions in the output from the job. For more information, see Associate
	// Prediction Results with their Corresponding Input Records
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html).
	DataProcessing *types.DataProcessing
	// The maximum number of parallel requests on each instance node that can be
	// launched in a transform job. The default value is 1.
	MaxConcurrentTransforms *int32
	// Describes the resources, including ML instance types and ML instance count, to
	// use for the transform job.
	TransformResources *types.TransformResources
	// Describes the dataset to be transformed and the Amazon S3 location where it is
	// stored.
	TransformInput *types.TransformInput
	// Indicates when the transform job starts on ML instances. You are billed for the
	// time interval between this time and the value of TransformEndTime.
	TransformStartTime *time.Time
	// The Amazon Resource Name (ARN) of the transform job.
	TransformJobArn *string
	// The timeout and maximum number of retries for processing a transform job
	// invocation.
	ModelClientConfig *types.ModelClientConfig
	// The environment variables to set in the Docker container. We support up to 16
	// key and values entries in the map.
	Environment map[string]*string
	// Identifies the Amazon S3 location where you want Amazon SageMaker to save the
	// results from the transform job.
	TransformOutput *types.TransformOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTransformJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTransformJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTransformJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTransformJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeTransformJob",
	}
}
