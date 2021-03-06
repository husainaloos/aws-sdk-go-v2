// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Selenium testing project. Projects are used to track TestGridSession
// () instances.
func (c *Client) CreateTestGridProject(ctx context.Context, params *CreateTestGridProjectInput, optFns ...func(*Options)) (*CreateTestGridProjectOutput, error) {
	if params == nil {
		params = &CreateTestGridProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTestGridProject", params, optFns, addOperationCreateTestGridProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTestGridProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTestGridProjectInput struct {

	// Human-readable name of the Selenium testing project.
	//
	// This member is required.
	Name *string

	// Human-readable description of the project.
	Description *string
}

type CreateTestGridProjectOutput struct {

	// ARN of the Selenium testing project that was created.
	TestGridProject *types.TestGridProject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTestGridProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTestGridProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTestGridProject{}, middleware.After)
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
	addOpCreateTestGridProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTestGridProject(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateTestGridProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateTestGridProject",
	}
}
