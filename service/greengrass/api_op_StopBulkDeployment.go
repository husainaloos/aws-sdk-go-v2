// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops the execution of a bulk deployment. This action returns a status of
// ''Stopping'' until the deployment is stopped. You cannot start a new bulk
// deployment while a previous deployment is in the ''Stopping'' state. This action
// doesn't rollback completed deployments or cancel pending deployments.
func (c *Client) StopBulkDeployment(ctx context.Context, params *StopBulkDeploymentInput, optFns ...func(*Options)) (*StopBulkDeploymentOutput, error) {
	stack := middleware.NewStack("StopBulkDeployment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStopBulkDeploymentMiddlewares(stack)
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
	addOpStopBulkDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopBulkDeployment(options.Region), middleware.Before)

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
			OperationName: "StopBulkDeployment",
			Err:           err,
		}
	}
	out := result.(*StopBulkDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopBulkDeploymentInput struct {
	// The ID of the bulk deployment.
	BulkDeploymentId *string
}

type StopBulkDeploymentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStopBulkDeploymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStopBulkDeployment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStopBulkDeployment{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopBulkDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "StopBulkDeployment",
	}
}