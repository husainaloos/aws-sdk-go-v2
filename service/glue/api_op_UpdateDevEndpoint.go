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

// Updates a specified development endpoint.
func (c *Client) UpdateDevEndpoint(ctx context.Context, params *UpdateDevEndpointInput, optFns ...func(*Options)) (*UpdateDevEndpointOutput, error) {
	stack := middleware.NewStack("UpdateDevEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDevEndpointMiddlewares(stack)
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
	addOpUpdateDevEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDevEndpoint(options.Region), middleware.Before)
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
			OperationName: "UpdateDevEndpoint",
			Err:           err,
		}
	}
	out := result.(*UpdateDevEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDevEndpointInput struct {
	// The name of the DevEndpoint to be updated.
	EndpointName *string
	// The list of public keys to be deleted from the DevEndpoint.
	DeletePublicKeys []*string
	// The list of argument keys to be deleted from the map of arguments used to
	// configure the DevEndpoint.
	DeleteArguments []*string
	// The list of public keys for the DevEndpoint to use.
	AddPublicKeys []*string
	// The public key for the DevEndpoint to use.
	PublicKey *string
	// The map of arguments to add the map of arguments used to configure the
	// DevEndpoint.  <p>Valid arguments are:</p> <ul> <li> <p>
	// <code>"--enable-glue-datacatalog": ""</code> </p> </li> <li> <p>
	// <code>"GLUE_PYTHON_VERSION": "3"</code> </p> </li> <li> <p>
	// <code>"GLUE_PYTHON_VERSION": "2"</code> </p> </li> </ul> <p>You can specify a
	// version of Python support for development endpoints by using the
	// <code>Arguments</code> parameter in the <code>CreateDevEndpoint</code> or
	// <code>UpdateDevEndpoint</code> APIs. If no arguments are provided, the version
	// defaults to Python 2.</p>
	AddArguments map[string]*string
	// Custom Python or Java libraries to be loaded in the DevEndpoint.
	CustomLibraries *types.DevEndpointCustomLibraries
	// True if the list of custom libraries to be loaded in the development endpoint
	// needs to be updated, or False if otherwise.
	UpdateEtlLibraries *bool
}

type UpdateDevEndpointOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDevEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDevEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDevEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDevEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateDevEndpoint",
	}
}
