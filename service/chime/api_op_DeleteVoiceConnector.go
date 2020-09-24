// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified Amazon Chime Voice Connector. Any phone numbers associated
// with the Amazon Chime Voice Connector must be disassociated from it before it
// can be deleted.
func (c *Client) DeleteVoiceConnector(ctx context.Context, params *DeleteVoiceConnectorInput, optFns ...func(*Options)) (*DeleteVoiceConnectorOutput, error) {
	stack := middleware.NewStack("DeleteVoiceConnector", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteVoiceConnectorMiddlewares(stack)
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
	addOpDeleteVoiceConnectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVoiceConnector(options.Region), middleware.Before)
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
			OperationName: "DeleteVoiceConnector",
			Err:           err,
		}
	}
	out := result.(*DeleteVoiceConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVoiceConnectorInput struct {
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

type DeleteVoiceConnectorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteVoiceConnectorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVoiceConnector{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVoiceConnector{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVoiceConnector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteVoiceConnector",
	}
}
