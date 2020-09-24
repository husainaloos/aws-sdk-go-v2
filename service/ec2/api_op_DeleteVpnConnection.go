// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified VPN connection. If you're deleting the VPC and its
// associated components, we recommend that you detach the virtual private gateway
// from the VPC and delete the VPC before deleting the VPN connection. If you
// believe that the tunnel credentials for your VPN connection have been
// compromised, you can delete the VPN connection and create a new one that has new
// keys, without needing to delete the VPC or virtual private gateway. If you
// create a new VPN connection, you must reconfigure the customer gateway device
// using the new configuration information returned with the new VPN connection ID.
// For certificate-based authentication, delete all AWS Certificate Manager (ACM)
// private certificates used for the AWS-side tunnel endpoints for the VPN
// connection before deleting the VPN connection.
func (c *Client) DeleteVpnConnection(ctx context.Context, params *DeleteVpnConnectionInput, optFns ...func(*Options)) (*DeleteVpnConnectionOutput, error) {
	stack := middleware.NewStack("DeleteVpnConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteVpnConnectionMiddlewares(stack)
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
	addOpDeleteVpnConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpnConnection(options.Region), middleware.Before)
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
			OperationName: "DeleteVpnConnection",
			Err:           err,
		}
	}
	out := result.(*DeleteVpnConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteVpnConnection.
type DeleteVpnConnectionInput struct {
	// The ID of the VPN connection.
	VpnConnectionId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteVpnConnectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteVpnConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteVpnConnection{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteVpnConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVpnConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteVpnConnection",
	}
}
