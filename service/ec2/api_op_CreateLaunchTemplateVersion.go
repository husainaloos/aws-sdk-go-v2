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

// Creates a new version for a launch template. You can specify an existing version
// of launch template from which to base the new version. Launch template versions
// are numbered in the order in which they are created. You cannot specify, change,
// or replace the numbering of launch template versions. For more information, see
// Managing launch template versions
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#manage-launch-template-versions)in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateLaunchTemplateVersion(ctx context.Context, params *CreateLaunchTemplateVersionInput, optFns ...func(*Options)) (*CreateLaunchTemplateVersionOutput, error) {
	stack := middleware.NewStack("CreateLaunchTemplateVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateLaunchTemplateVersionMiddlewares(stack)
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
	addOpCreateLaunchTemplateVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchTemplateVersion(options.Region), middleware.Before)
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
			OperationName: "CreateLaunchTemplateVersion",
			Err:           err,
		}
	}
	out := result.(*CreateLaunchTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchTemplateVersionInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The version number of the launch template version on which to base the new
	// version. The new version inherits the same launch parameters as the source
	// version, except for parameters that you specify in LaunchTemplateData. Snapshots
	// applied to the block device mapping are ignored when creating a new version
	// unless they are explicitly included.
	SourceVersion *string
	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	// Constraint: Maximum 128 ASCII characters.
	ClientToken *string
	// The name of the launch template. You must specify either the launch template ID
	// or launch template name in the request.
	LaunchTemplateName *string
	// The information for the launch template.
	LaunchTemplateData *types.RequestLaunchTemplateData
	// A description for the version of the launch template.
	VersionDescription *string
	// The ID of the launch template. You must specify either the launch template ID or
	// launch template name in the request.
	LaunchTemplateId *string
}

type CreateLaunchTemplateVersionOutput struct {
	// If the new version of the launch template contains parameters or parameter
	// combinations that are not valid, an error code and an error message are returned
	// for each issue that's found.
	Warning *types.ValidationWarning
	// Information about the launch template version.
	LaunchTemplateVersion *types.LaunchTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateLaunchTemplateVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateLaunchTemplateVersion{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLaunchTemplateVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLaunchTemplateVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLaunchTemplateVersion",
	}
}
