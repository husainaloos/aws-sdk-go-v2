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
)

// Starts a model compilation job. After the model has been compiled, Amazon
// SageMaker saves the resulting model artifacts to an Amazon Simple Storage
// Service (Amazon S3) bucket that you specify. If you choose to host your model
// using Amazon SageMaker hosting services, you can use the resulting model
// artifacts as part of the model. You can also use the artifacts with AWS IoT
// Greengrass. In that case, deploy them as an ML resource. In the request body,
// you provide the following:
//
//     * A name for the compilation job
//
//     *
// Information about the input model artifacts
//
//     * The output location for the
// compiled model and the device (target) that the model runs on
//
//     * The Amazon
// Resource Name (ARN) of the IAM role that Amazon SageMaker assumes to perform the
// model compilation job.
//
// You can also provide a Tag to track the model
// compilation job's resource use and costs. The response body contains the
// CompilationJobArn for the compiled job. To stop a model compilation job, use
// StopCompilationJob (). To get information about a particular model compilation
// job, use DescribeCompilationJob (). To get information about multiple model
// compilation jobs, use ListCompilationJobs ().
func (c *Client) CreateCompilationJob(ctx context.Context, params *CreateCompilationJobInput, optFns ...func(*Options)) (*CreateCompilationJobOutput, error) {
	stack := middleware.NewStack("CreateCompilationJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateCompilationJobMiddlewares(stack)
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
	addOpCreateCompilationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCompilationJob(options.Region), middleware.Before)
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
			OperationName: "CreateCompilationJob",
			Err:           err,
		}
	}
	out := result.(*CreateCompilationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCompilationJobInput struct {
	// Provides information about the output location for the compiled model and the
	// target device the model runs on.
	OutputConfig *types.OutputConfig
	// Specifies a limit to how long a model compilation job can run. When the job
	// reaches the time limit, Amazon SageMaker ends the compilation job. Use this API
	// to cap model training costs.
	StoppingCondition *types.StoppingCondition
	// Provides information about the location of input model artifacts, the name and
	// shape of the expected data inputs, and the framework in which the model was
	// trained.
	InputConfig *types.InputConfig
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to
	// perform tasks on your behalf. During model compilation, Amazon SageMaker needs
	// your permission to:
	//
	//     * Read input data from an S3 bucket
	//
	//     * Write model
	// artifacts to an S3 bucket
	//
	//     * Write logs to Amazon CloudWatch Logs
	//
	//     *
	// Publish metrics to Amazon CloudWatch
	//
	// You grant permissions for all of these
	// tasks to an IAM role. To pass this role to Amazon SageMaker, the caller of this
	// API must have the iam:PassRole permission. For more information, see Amazon
	// SageMaker Roles.
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html)
	RoleArn *string
	// A name for the model compilation job. The name must be unique within the AWS
	// Region and within your AWS account.
	CompilationJobName *string
}

type CreateCompilationJobOutput struct {
	// If the action is successful, the service sends back an HTTP 200 response. Amazon
	// SageMaker returns the following data in JSON format:
	//
	//     * CompilationJobArn:
	// The Amazon Resource Name (ARN) of the compiled job.
	CompilationJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateCompilationJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCompilationJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCompilationJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCompilationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateCompilationJob",
	}
}
