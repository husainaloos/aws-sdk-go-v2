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

// Accepts a request to attach a VPC to a transit gateway. The VPC attachment must
// be in the pendingAcceptance state. Use DescribeTransitGatewayVpcAttachments ()
// to view your pending VPC attachment requests. Use
// RejectTransitGatewayVpcAttachment () to reject a VPC attachment request.
func (c *Client) AcceptTransitGatewayVpcAttachment(ctx context.Context, params *AcceptTransitGatewayVpcAttachmentInput, optFns ...func(*Options)) (*AcceptTransitGatewayVpcAttachmentOutput, error) {
	stack := middleware.NewStack("AcceptTransitGatewayVpcAttachment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAcceptTransitGatewayVpcAttachmentMiddlewares(stack)
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
	addOpAcceptTransitGatewayVpcAttachmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptTransitGatewayVpcAttachment(options.Region), middleware.Before)
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
			OperationName: "AcceptTransitGatewayVpcAttachment",
			Err:           err,
		}
	}
	out := result.(*AcceptTransitGatewayVpcAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptTransitGatewayVpcAttachmentInput struct {
	// The ID of the attachment.
	TransitGatewayAttachmentId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AcceptTransitGatewayVpcAttachmentOutput struct {
	// The VPC attachment.
	TransitGatewayVpcAttachment *types.TransitGatewayVpcAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAcceptTransitGatewayVpcAttachmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAcceptTransitGatewayVpcAttachment{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAcceptTransitGatewayVpcAttachment{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptTransitGatewayVpcAttachment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AcceptTransitGatewayVpcAttachment",
	}
}
