// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a deployment configuration.
func (c *Client) CreateDeploymentConfig(ctx context.Context, params *CreateDeploymentConfigInput, optFns ...func(*Options)) (*CreateDeploymentConfigOutput, error) {
	stack := middleware.NewStack("CreateDeploymentConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDeploymentConfigMiddlewares(stack)
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
	addOpCreateDeploymentConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeploymentConfig(options.Region), middleware.Before)
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
			OperationName: "CreateDeploymentConfig",
			Err:           err,
		}
	}
	out := result.(*CreateDeploymentConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateDeploymentConfig operation.
type CreateDeploymentConfigInput struct {
	// The name of the deployment configuration to create.
	DeploymentConfigName *string
	// The minimum number of healthy instances that should be available at any time
	// during the deployment. There are two parameters expected in the input: type and
	// value. The type parameter takes either of the following values:
	//
	//     *
	// HOST_COUNT: The value parameter represents the minimum number of healthy
	// instances as an absolute value.
	//
	//     * FLEET_PERCENT: The value parameter
	// represents the minimum number of healthy instances as a percentage of the total
	// number of instances in the deployment. If you specify FLEET_PERCENT, at the
	// start of the deployment, AWS CodeDeploy converts the percentage to the
	// equivalent number of instances and rounds up fractional instances.
	//
	// The value
	// parameter takes an integer. For example, to set a minimum of 95% healthy
	// instance, specify a type of FLEET_PERCENT and a value of 95.
	MinimumHealthyHosts *types.MinimumHealthyHosts
	// The configuration that specifies how the deployment traffic is routed.
	TrafficRoutingConfig *types.TrafficRoutingConfig
	// The destination platform type for the deployment (Lambda, Server, or ECS).
	ComputePlatform types.ComputePlatform
}

// Represents the output of a CreateDeploymentConfig operation.
type CreateDeploymentConfigOutput struct {
	// A unique deployment configuration ID.
	DeploymentConfigId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDeploymentConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDeploymentConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDeploymentConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDeploymentConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "CreateDeploymentConfig",
	}
}
