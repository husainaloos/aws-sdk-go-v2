// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// SetTerminationProtection locks a cluster (job flow) so the EC2 instances in the
// cluster cannot be terminated by user intervention, an API call, or in the event
// of a job-flow error. The cluster still terminates upon successful completion of
// the job flow. Calling SetTerminationProtection on a cluster is similar to
// calling the Amazon EC2 DisableAPITermination API on all EC2 instances in a
// cluster. SetTerminationProtection is used to prevent accidental termination of a
// cluster and to ensure that in the event of an error, the instances persist so
// that you can recover any data stored in their ephemeral instance storage.  <p>
// To terminate a cluster that has been locked by setting
// <code>SetTerminationProtection</code> to <code>true</code>, you must first
// unlock the job flow by a subsequent call to
// <code>SetTerminationProtection</code> in which you set the value to
// <code>false</code>. </p> <p> For more information, see<a
// href="https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html">Managing
// Cluster Termination</a> in the <i>Amazon EMR Management Guide</i>. </p>
func (c *Client) SetTerminationProtection(ctx context.Context, params *SetTerminationProtectionInput, optFns ...func(*Options)) (*SetTerminationProtectionOutput, error) {
	stack := middleware.NewStack("SetTerminationProtection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSetTerminationProtectionMiddlewares(stack)
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
	addOpSetTerminationProtectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetTerminationProtection(options.Region), middleware.Before)
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
			OperationName: "SetTerminationProtection",
			Err:           err,
		}
	}
	out := result.(*SetTerminationProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input argument to the TerminationProtection () operation.
type SetTerminationProtectionInput struct {
	// A Boolean that indicates whether to protect the cluster and prevent the Amazon
	// EC2 instances in the cluster from shutting down due to API calls, user
	// intervention, or job-flow error.
	TerminationProtected *bool
	// A list of strings that uniquely identify the clusters to protect. This
	// identifier is returned by RunJobFlow () and can also be obtained from
	// DescribeJobFlows () .
	JobFlowIds []*string
}

type SetTerminationProtectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSetTerminationProtectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSetTerminationProtection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetTerminationProtection{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetTerminationProtection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "SetTerminationProtection",
	}
}
