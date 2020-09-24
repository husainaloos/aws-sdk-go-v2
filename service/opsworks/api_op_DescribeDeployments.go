// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Requests a description of a specified set of deployments. This call accepts only
// one resource-identifying parameter. Required Permissions: To use this action, an
// IAM user must have a Show, Deploy, or Manage permissions level for the stack, or
// an attached policy that explicitly grants permissions. For more information
// about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeDeployments(ctx context.Context, params *DescribeDeploymentsInput, optFns ...func(*Options)) (*DescribeDeploymentsOutput, error) {
	stack := middleware.NewStack("DescribeDeployments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDeploymentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeployments(options.Region), middleware.Before)
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
			OperationName: "DescribeDeployments",
			Err:           err,
		}
	}
	out := result.(*DescribeDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeploymentsInput struct {
	// The app ID. If you include this parameter, the command returns a description of
	// the commands associated with the specified app.
	AppId *string
	// The stack ID. If you include this parameter, the command returns a description
	// of the commands associated with the specified stack.
	StackId *string
	// An array of deployment IDs to be described. If you include this parameter, the
	// command returns a description of the specified deployments. Otherwise, it
	// returns a description of every deployment.
	DeploymentIds []*string
}

// Contains the response to a DescribeDeployments request.
type DescribeDeploymentsOutput struct {
	// An array of Deployment objects that describe the deployments.
	Deployments []*types.Deployment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDeploymentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDeployments{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDeployments{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDeployments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeDeployments",
	}
}
