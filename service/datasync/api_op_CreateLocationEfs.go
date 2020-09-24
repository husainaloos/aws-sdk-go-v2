// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an endpoint for an Amazon EFS file system.
func (c *Client) CreateLocationEfs(ctx context.Context, params *CreateLocationEfsInput, optFns ...func(*Options)) (*CreateLocationEfsOutput, error) {
	stack := middleware.NewStack("CreateLocationEfs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateLocationEfsMiddlewares(stack)
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
	addOpCreateLocationEfsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationEfs(options.Region), middleware.Before)
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
			OperationName: "CreateLocationEfs",
			Err:           err,
		}
	}
	out := result.(*CreateLocationEfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationEfsRequest
type CreateLocationEfsInput struct {
	// A subdirectory in the location’s path. This subdirectory in the EFS file system
	// is used to read data from the EFS source location or write data to the EFS
	// destination. By default, AWS DataSync uses the root directory. Subdirectory must
	// be specified with forward slashes. For example, /path/to/folder.
	Subdirectory *string
	// The key-value pair that represents a tag that you want to add to the resource.
	// The value can be an empty string. This value helps you manage, filter, and
	// search for your resources. We recommend that you create a name tag for your
	// location.
	Tags []*types.TagListEntry
	// The Amazon Resource Name (ARN) for the Amazon EFS file system.
	EfsFilesystemArn *string
	// The subnet and security group that the Amazon EFS file system uses. The security
	// group that you provide needs to be able to communicate with the security group
	// on the mount target in the subnet specified. The exact relationship between
	// security group M (of the mount target) and security group S (which you provide
	// for DataSync to use at this stage) is as follows:
	//
	//     * Security group M (which
	// you associate with the mount target) must allow inbound access for the
	// Transmission Control Protocol (TCP) on the NFS port (2049) from security group
	// S. You can enable inbound connections either by IP address (CIDR range) or
	// security group.
	//
	//     * Security group S (provided to DataSync to access EFS)
	// should have a rule that enables outbound connections to the NFS port on one of
	// the file system’s mount targets. You can enable outbound connections either by
	// IP address (CIDR range) or security group.  <p>For information about security
	// groups and mount targets, see Security Groups for Amazon EC2 Instances and Mount
	// Targets in the <i>Amazon EFS User Guide.</i> </p> </li> </ul>
	Ec2Config *types.Ec2Config
}

// CreateLocationEfs
type CreateLocationEfsOutput struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that is
	// created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateLocationEfsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationEfs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationEfs{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLocationEfs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationEfs",
	}
}
