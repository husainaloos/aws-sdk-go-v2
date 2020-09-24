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

// Starts a job run using a job definition.
func (c *Client) StartJobRun(ctx context.Context, params *StartJobRunInput, optFns ...func(*Options)) (*StartJobRunOutput, error) {
	stack := middleware.NewStack("StartJobRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartJobRunMiddlewares(stack)
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
	addOpStartJobRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartJobRun(options.Region), middleware.Before)
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
			OperationName: "StartJobRun",
			Err:           err,
		}
	}
	out := result.(*StartJobRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartJobRunInput struct {
	// The name of the job definition to use.
	JobName *string
	// The ID of a previous JobRun to retry.
	JobRunId *string
	// The name of the SecurityConfiguration structure to be used with this job run.
	SecurityConfiguration *string
	// The JobRun timeout in minutes. This is the maximum time that a job run can
	// consume resources before it is terminated and enters TIMEOUT status. The default
	// is 2,880 minutes (48 hours). This overrides the timeout value set in the parent
	// job.
	Timeout *int32
	// This field is deprecated. Use MaxCapacity instead.  <p>The number of AWS Glue
	// data processing units (DPUs) to allocate to this JobRun. From 2 to 100 DPUs can
	// be allocated; the default is 10. A DPU is a relative measure of processing power
	// that consists of 4 vCPUs of compute capacity and 16 GB of memory. For more
	// information, see the <a
	// href="https://docs.aws.amazon.com/https:/aws.amazon.com/glue/pricing/">AWS Glue
	// pricing page</a>.</p>
	AllocatedCapacity *int32
	// The job arguments specifically for this run. For this job run, they replace the
	// default arguments set in the job definition itself. You can specify arguments
	// here that your own job-execution script consumes, as well as arguments that AWS
	// Glue itself consumes. For information about how to specify and consume your own
	// Job arguments, see the Calling AWS Glue APIs in Python
	// (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html)
	// topic in the developer guide. For information about the key-value pairs that AWS
	// Glue consumes to set up your job, see the Special Parameters Used by AWS Glue
	// (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	// topic in the developer guide.
	Arguments map[string]*string
	// The number of AWS Glue data processing units (DPUs) that can be allocated when
	// this job runs. A DPU is a relative measure of processing power that consists of
	// 4 vCPUs of compute capacity and 16 GB of memory. For more information, see the
	// AWS Glue pricing page
	// (https://docs.aws.amazon.com/https:/aws.amazon.com/glue/pricing/).  <p>Do not
	// set <code>Max Capacity</code> if using <code>WorkerType</code> and
	// <code>NumberOfWorkers</code>.</p> <p>The value that can be allocated for
	// <code>MaxCapacity</code> depends on whether you are running a Python shell job,
	// or an Apache Spark ETL job:</p> <ul> <li> <p>When you specify a Python shell job
	// (<code>JobCommand.Name</code>="pythonshell"), you can allocate either 0.0625 or
	// 1 DPU. The default is 0.0625 DPU.</p> </li> <li> <p>When you specify an Apache
	// Spark ETL job (<code>JobCommand.Name</code>="glueetl"), you can allocate from 2
	// to 100 DPUs. The default is 10 DPUs. This job type cannot have a fractional DPU
	// allocation.</p> </li> </ul>
	MaxCapacity *float64
	// Specifies configuration properties of a job run notification.
	NotificationProperty *types.NotificationProperty
	// The type of predefined worker that is allocated when a job runs. Accepts a value
	// of Standard, G.1X, or G.2X.
	//
	//     * For the Standard worker type, each worker
	// provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	//
	//
	// * For the G.1X worker type, each worker provides 4 vCPU, 16 GB of memory and a
	// 64GB disk, and 1 executor per worker.
	//
	//     * For the G.2X worker type, each
	// worker provides 8 vCPU, 32 GB of memory and a 128GB disk, and 1 executor per
	// worker.
	WorkerType types.WorkerType
	// The number of workers of a defined workerType that are allocated when a job
	// runs.  <p>The maximum number of workers you can define are 299 for
	// <code>G.1X</code>, and 149 for <code>G.2X</code>. </p>
	NumberOfWorkers *int32
}

type StartJobRunOutput struct {
	// The ID assigned to this job run.
	JobRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartJobRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartJobRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartJobRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartJobRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartJobRun",
	}
}
