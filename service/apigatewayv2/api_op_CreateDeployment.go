// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a Deployment for an API.
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	stack := middleware.NewStack("CreateDeployment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDeploymentMiddlewares(stack)
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
	addOpCreateDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeployment(options.Region), middleware.Before)
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
			OperationName: "CreateDeployment",
			Err:           err,
		}
	}
	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new Deployment resource to represent a deployment.
type CreateDeploymentInput struct {
	// The name of the Stage resource for the Deployment resource to create.
	StageName *string
	// The description for the deployment resource.
	Description *string
	// The API identifier.
	ApiId *string
}

type CreateDeploymentOutput struct {
	// Specifies whether a deployment was automatically released.
	AutoDeployed *bool
	// The date and time when the Deployment resource was created.
	CreatedDate *time.Time
	// May contain additional feedback on the status of an API deployment.
	DeploymentStatusMessage *string
	// The identifier for the deployment.
	DeploymentId *string
	// The status of the deployment: PENDING, FAILED, or SUCCEEDED.
	DeploymentStatus types.DeploymentStatus
	// The description for the deployment.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDeploymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeployment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeployment{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateDeployment",
	}
}
