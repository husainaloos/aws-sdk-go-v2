// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows the destination domain owner to delete an existing inbound cross-cluster
// search connection.
func (c *Client) DeleteInboundCrossClusterSearchConnection(ctx context.Context, params *DeleteInboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*DeleteInboundCrossClusterSearchConnectionOutput, error) {
	stack := middleware.NewStack("DeleteInboundCrossClusterSearchConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteInboundCrossClusterSearchConnectionMiddlewares(stack)
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
	addOpDeleteInboundCrossClusterSearchConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInboundCrossClusterSearchConnection(options.Region), middleware.Before)
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
			OperationName: "DeleteInboundCrossClusterSearchConnection",
			Err:           err,
		}
	}
	out := result.(*DeleteInboundCrossClusterSearchConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteInboundCrossClusterSearchConnection ()
// operation.
type DeleteInboundCrossClusterSearchConnectionInput struct {
	// The id of the inbound connection that you want to permanently delete.
	CrossClusterSearchConnectionId *string
}

// The result of a DeleteInboundCrossClusterSearchConnection () operation. Contains
// details of deleted inbound connection.
type DeleteInboundCrossClusterSearchConnectionOutput struct {
	// Specifies the InboundCrossClusterSearchConnection () of deleted inbound
	// connection.
	CrossClusterSearchConnection *types.InboundCrossClusterSearchConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteInboundCrossClusterSearchConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteInboundCrossClusterSearchConnection{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteInboundCrossClusterSearchConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteInboundCrossClusterSearchConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DeleteInboundCrossClusterSearchConnection",
	}
}
