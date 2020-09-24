// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves details for the specified Amazon Chime Voice Connector group, such as
// timestamps, name, and associated VoiceConnectorItems.
func (c *Client) GetVoiceConnectorGroup(ctx context.Context, params *GetVoiceConnectorGroupInput, optFns ...func(*Options)) (*GetVoiceConnectorGroupOutput, error) {
	stack := middleware.NewStack("GetVoiceConnectorGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetVoiceConnectorGroupMiddlewares(stack)
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
	addOpGetVoiceConnectorGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorGroup(options.Region), middleware.Before)
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
			OperationName: "GetVoiceConnectorGroup",
			Err:           err,
		}
	}
	out := result.(*GetVoiceConnectorGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorGroupInput struct {
	// The Amazon Chime Voice Connector group ID.
	VoiceConnectorGroupId *string
}

type GetVoiceConnectorGroupOutput struct {
	// The Amazon Chime Voice Connector group details.
	VoiceConnectorGroup *types.VoiceConnectorGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetVoiceConnectorGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVoiceConnectorGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetVoiceConnectorGroup",
	}
}
