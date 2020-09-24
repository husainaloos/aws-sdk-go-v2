// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates metadata about an application.
func (c *Client) UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) {
	stack := middleware.NewStack("UpdateApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateApplicationMiddlewares(stack)
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
	addOpUpdateApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplication(options.Region), middleware.Before)
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
			OperationName: "UpdateApplication",
			Err:           err,
		}
	}
	out := result.(*UpdateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationInput struct {
	// New name of the application to be updated.
	Name *string
	// New description of the application to be updated.
	Description *string
	// Configuration ID of the application to be updated.
	ConfigurationId *string
}

type UpdateApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApplication{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "UpdateApplication",
	}
}
