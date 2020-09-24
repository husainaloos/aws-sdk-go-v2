// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata such as the name, the network interfaces, and the status (that
// is, whether the agent is running or not) for an agent. To specify which agent to
// describe, use the Amazon Resource Name (ARN) of the agent in your request.
func (c *Client) DescribeAgent(ctx context.Context, params *DescribeAgentInput, optFns ...func(*Options)) (*DescribeAgentOutput, error) {
	stack := middleware.NewStack("DescribeAgent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAgentMiddlewares(stack)
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
	addOpDescribeAgentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAgent(options.Region), middleware.Before)
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
			OperationName: "DescribeAgent",
			Err:           err,
		}
	}
	out := result.(*DescribeAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeAgent
type DescribeAgentInput struct {
	// The Amazon Resource Name (ARN) of the agent to describe.
	AgentArn *string
}

// DescribeAgentResponse
type DescribeAgentOutput struct {
	// The time that the agent was activated (that is, created in your account).
	CreationTime *time.Time
	// The name of the agent.
	Name *string
	// The type of endpoint that your agent is connected to. If the endpoint is a VPC
	// endpoint, the agent is not accessible over the public internet.
	EndpointType types.EndpointType
	// The time that the agent last connected to DataSyc.
	LastConnectionTime *time.Time
	// The status of the agent. If the status is ONLINE, then the agent is configured
	// properly and is available to use. The Running status is the normal running
	// status for an agent. If the status is OFFLINE, the agent's VM is turned off or
	// the agent is in an unhealthy state. When the issue that caused the unhealthy
	// state is resolved, the agent returns to ONLINE status.
	Status types.AgentStatus
	// The subnet and the security group that DataSync used to access a VPC endpoint.
	PrivateLinkConfig *types.PrivateLinkConfig
	// The Amazon Resource Name (ARN) of the agent.
	AgentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAgentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAgent{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAgent{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAgent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeAgent",
	}
}
