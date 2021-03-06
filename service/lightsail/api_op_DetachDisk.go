// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches a stopped block storage disk from a Lightsail instance. Make sure to
// unmount any file systems on the device within your operating system before
// stopping the instance and detaching the disk. The detach disk operation supports
// tag-based access control via resource tags applied to the resource identified by
// disk name. For more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) DetachDisk(ctx context.Context, params *DetachDiskInput, optFns ...func(*Options)) (*DetachDiskOutput, error) {
	if params == nil {
		params = &DetachDiskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachDisk", params, optFns, addOperationDetachDiskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachDiskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachDiskInput struct {

	// The unique name of the disk you want to detach from your instance (e.g.,
	// my-disk).
	//
	// This member is required.
	DiskName *string
}

type DetachDiskOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachDiskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachDisk{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachDisk{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDetachDiskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachDisk(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetachDisk(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DetachDisk",
	}
}
