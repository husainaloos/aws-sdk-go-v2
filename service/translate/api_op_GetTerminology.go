// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a custom terminology.
func (c *Client) GetTerminology(ctx context.Context, params *GetTerminologyInput, optFns ...func(*Options)) (*GetTerminologyOutput, error) {
	stack := middleware.NewStack("GetTerminology", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTerminologyMiddlewares(stack)
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
	addOpGetTerminologyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTerminology(options.Region), middleware.Before)
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
			OperationName: "GetTerminology",
			Err:           err,
		}
	}
	out := result.(*GetTerminologyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTerminologyInput struct {
	// The data format of the custom terminology being retrieved, either CSV or TMX.
	TerminologyDataFormat types.TerminologyDataFormat
	// The name of the custom terminology being retrieved.
	Name *string
}

type GetTerminologyOutput struct {
	// The properties of the custom terminology being retrieved.
	TerminologyProperties *types.TerminologyProperties
	// The data location of the custom terminology being retrieved. The custom
	// terminology file is returned in a presigned url that has a 30 minute expiration.
	TerminologyDataLocation *types.TerminologyDataLocation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTerminologyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTerminology{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTerminology{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTerminology(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "GetTerminology",
	}
}
