// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a target network from the specified Client VPN endpoint. When you
// disassociate the last target network from a Client VPN, the following happens:
//
//
// * The route that was automatically added for the VPC is deleted
//
//     * All
// active client connections are terminated
//
//     * New client connections are
// disallowed
//
//     * The Client VPN endpoint's status changes to pending-associate
func (c *Client) DisassociateClientVpnTargetNetwork(ctx context.Context, params *DisassociateClientVpnTargetNetworkInput, optFns ...func(*Options)) (*DisassociateClientVpnTargetNetworkOutput, error) {
	stack := middleware.NewStack("DisassociateClientVpnTargetNetwork", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisassociateClientVpnTargetNetworkMiddlewares(stack)
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
	addOpDisassociateClientVpnTargetNetworkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateClientVpnTargetNetwork(options.Region), middleware.Before)
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
			OperationName: "DisassociateClientVpnTargetNetwork",
			Err:           err,
		}
	}
	out := result.(*DisassociateClientVpnTargetNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateClientVpnTargetNetworkInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the target network association.
	AssociationId *string
	// The ID of the Client VPN endpoint from which to disassociate the target network.
	ClientVpnEndpointId *string
}

type DisassociateClientVpnTargetNetworkOutput struct {
	// The ID of the target network association.
	AssociationId *string
	// The current state of the target network association.
	Status *types.AssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisassociateClientVpnTargetNetworkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisassociateClientVpnTargetNetwork{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateClientVpnTargetNetwork{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateClientVpnTargetNetwork(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateClientVpnTargetNetwork",
	}
}
