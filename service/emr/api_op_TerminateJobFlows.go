// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// TerminateJobFlows shuts a list of clusters (job flows) down. When a job flow is
// shut down, any step not yet completed is canceled and the EC2 instances on which
// the cluster is running are stopped. Any log files not already saved are uploaded
// to Amazon S3 if a LogUri was specified when the cluster was created. The maximum
// number of clusters allowed is 10. The call to TerminateJobFlows is asynchronous.
// Depending on the configuration of the cluster, it may take up to 1-5 minutes for
// the cluster to completely terminate and release allocated resources, such as
// Amazon EC2 instances.
func (c *Client) TerminateJobFlows(ctx context.Context, params *TerminateJobFlowsInput, optFns ...func(*Options)) (*TerminateJobFlowsOutput, error) {
	if params == nil {
		params = &TerminateJobFlowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateJobFlows", params, optFns, addOperationTerminateJobFlowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateJobFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the TerminateJobFlows () operation.
type TerminateJobFlowsInput struct {

	// A list of job flows to be shutdown.
	//
	// This member is required.
	JobFlowIds []*string
}

type TerminateJobFlowsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTerminateJobFlowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateJobFlows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateJobFlows{}, middleware.After)
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
	addOpTerminateJobFlowsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateJobFlows(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTerminateJobFlows(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "TerminateJobFlows",
	}
}
