// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Glue machine learning transform. This operation creates the
// transform and all the necessary parameters to train it.  <p>Call this operation
// as the first step in the process of using a machine learning transform (such as
// the <code>FindMatches</code> transform) for deduplicating data. You can provide
// an optional <code>Description</code>, in addition to the parameters that you
// want to use for your algorithm.</p> <p>You must also specify certain parameters
// for the tasks that AWS Glue runs on your behalf as part of learning from your
// data and creating a high-quality machine learning transform. These parameters
// include <code>Role</code>, and optionally, <code>AllocatedCapacity</code>,
// <code>Timeout</code>, and <code>MaxRetries</code>. For more information, see <a
// href="https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html">Jobs</a>.</p>
func (c *Client) CreateMLTransform(ctx context.Context, params *CreateMLTransformInput, optFns ...func(*Options)) (*CreateMLTransformOutput, error) {
	stack := middleware.NewStack("CreateMLTransform", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateMLTransformMiddlewares(stack)
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
	addOpCreateMLTransformValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMLTransform(options.Region), middleware.Before)
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
			OperationName: "CreateMLTransform",
			Err:           err,
		}
	}
	out := result.(*CreateMLTransformOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMLTransformInput struct {
	// The maximum number of times to retry a task for this transform after a task run
	// fails.
	MaxRetries *int32
	// The name or Amazon Resource Name (ARN) of the IAM role with the required
	// permissions. The required permissions include both AWS Glue service role
	// permissions to AWS Glue resources, and Amazon S3 permissions required by the
	// transform.  <ul> <li> <p>This role needs AWS Glue service role permissions to
	// allow access to resources in AWS Glue. See <a
	// href="https://docs.aws.amazon.com/glue/latest/dg/attach-policy-iam-user.html">Attach
	// a Policy to IAM Users That Access AWS Glue</a>.</p> </li> <li> <p>This role
	// needs permission to your Amazon Simple Storage Service (Amazon S3) sources,
	// targets, temporary directory, scripts, and any libraries used by the task run
	// for this transform.</p> </li> </ul>
	Role *string
	// The type of predefined worker that is allocated when this task runs. Accepts a
	// value of Standard, G.1X, or G.2X.
	//
	//     * For the Standard worker type, each
	// worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per
	// worker.
	//
	//     * For the G.1X worker type, each worker provides 4 vCPU, 16 GB of
	// memory and a 64GB disk, and 1 executor per worker.
	//
	//     * For the G.2X worker
	// type, each worker provides 8 vCPU, 32 GB of memory and a 128GB disk, and 1
	// executor per worker.
	//
	//     <p> <code>MaxCapacity</code> is a mutually exclusive
	// option with <code>NumberOfWorkers</code> and <code>WorkerType</code>.</p> <ul>
	// <li> <p>If either <code>NumberOfWorkers</code> or <code>WorkerType</code> is
	// set, then <code>MaxCapacity</code> cannot be set.</p> </li> <li> <p>If
	// <code>MaxCapacity</code> is set then neither <code>NumberOfWorkers</code> or
	// <code>WorkerType</code> can be set.</p> </li> <li> <p>If <code>WorkerType</code>
	// is set, then <code>NumberOfWorkers</code> is required (and vice versa).</p>
	// </li> <li> <p> <code>MaxCapacity</code> and <code>NumberOfWorkers</code> must
	// both be at least 1.</p> </li> </ul>
	WorkerType types.WorkerType
	// A description of the machine learning transform that is being defined. The
	// default is an empty string.
	Description *string
	// The algorithmic parameters that are specific to the transform type used.
	// Conditionally dependent on the transform type.
	Parameters *types.TransformParameters
	// The number of workers of a defined workerType that are allocated when this task
	// runs.  <p>If <code>WorkerType</code> is set, then <code>NumberOfWorkers</code>
	// is required (and vice versa).</p>
	NumberOfWorkers *int32
	// A list of AWS Glue table definitions used by the transform.
	InputRecordTables []*types.GlueTable
	// The number of AWS Glue data processing units (DPUs) that are allocated to task
	// runs for this transform. You can allocate from 2 to 100 DPUs; the default is 10.
	// A DPU is a relative measure of processing power that consists of 4 vCPUs of
	// compute capacity and 16 GB of memory. For more information, see the AWS Glue
	// pricing page (https://aws.amazon.com/glue/pricing/).  <p>
	// <code>MaxCapacity</code> is a mutually exclusive option with
	// <code>NumberOfWorkers</code> and <code>WorkerType</code>.</p> <ul> <li> <p>If
	// either <code>NumberOfWorkers</code> or <code>WorkerType</code> is set, then
	// <code>MaxCapacity</code> cannot be set.</p> </li> <li> <p>If
	// <code>MaxCapacity</code> is set then neither <code>NumberOfWorkers</code> or
	// <code>WorkerType</code> can be set.</p> </li> <li> <p>If <code>WorkerType</code>
	// is set, then <code>NumberOfWorkers</code> is required (and vice versa).</p>
	// </li> <li> <p> <code>MaxCapacity</code> and <code>NumberOfWorkers</code> must
	// both be at least 1.</p> </li> </ul> <p>When the <code>WorkerType</code> field is
	// set to a value other than <code>Standard</code>, the <code>MaxCapacity</code>
	// field is set automatically and becomes read-only.</p> <p>When the
	// <code>WorkerType</code> field is set to a value other than
	// <code>Standard</code>, the <code>MaxCapacity</code> field is set automatically
	// and becomes read-only.</p>
	MaxCapacity *float64
	// The tags to use with this machine learning transform. You may use tags to limit
	// access to the machine learning transform. For more information about tags in AWS
	// Glue, see AWS Tags in AWS Glue
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the developer
	// guide.
	Tags map[string]*string
	// The unique name that you give the transform when you create it.
	Name *string
	// The timeout of the task run for this transform in minutes. This is the maximum
	// time that a task run for this transform can consume resources before it is
	// terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout *int32
	// This value determines which version of AWS Glue this machine learning transform
	// is compatible with. Glue 1.0 is recommended for most customers. If the value is
	// not set, the Glue compatibility defaults to Glue 0.9. For more information, see
	// AWS Glue Versions
	// (https://docs.aws.amazon.com/glue/latest/dg/release-notes.html#release-notes-versions)
	// in the developer guide.
	GlueVersion *string
}

type CreateMLTransformOutput struct {
	// A unique identifier that is generated for the transform.
	TransformId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateMLTransformMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMLTransform{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMLTransform{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateMLTransform(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateMLTransform",
	}
}
