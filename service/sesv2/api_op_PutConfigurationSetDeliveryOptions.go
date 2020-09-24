// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associate a configuration set with a dedicated IP pool. You can use dedicated IP
// pools to create groups of dedicated IP addresses for sending specific types of
// email.
func (c *Client) PutConfigurationSetDeliveryOptions(ctx context.Context, params *PutConfigurationSetDeliveryOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetDeliveryOptionsOutput, error) {
	stack := middleware.NewStack("PutConfigurationSetDeliveryOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutConfigurationSetDeliveryOptionsMiddlewares(stack)
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
	addOpPutConfigurationSetDeliveryOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetDeliveryOptions(options.Region), middleware.Before)
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
			OperationName: "PutConfigurationSetDeliveryOptions",
			Err:           err,
		}
	}
	out := result.(*PutConfigurationSetDeliveryOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to associate a configuration set with a dedicated IP pool.
type PutConfigurationSetDeliveryOptionsInput struct {
	// The name of the dedicated IP pool that you want to associate with the
	// configuration set.
	SendingPoolName *string
	// Specifies whether messages that use the configuration set are required to use
	// Transport Layer Security (TLS). If the value is Require, messages are only
	// delivered if a TLS connection can be established. If the value is Optional,
	// messages can be delivered in plain text if a TLS connection can't be
	// established.
	TlsPolicy types.TlsPolicy
	// The name of the configuration set that you want to associate with a dedicated IP
	// pool.
	ConfigurationSetName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetDeliveryOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutConfigurationSetDeliveryOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetDeliveryOptions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetDeliveryOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutConfigurationSetDeliveryOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetDeliveryOptions",
	}
}
