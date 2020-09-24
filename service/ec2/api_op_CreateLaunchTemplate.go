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

// Creates a launch template. A launch template contains the parameters to launch
// an instance. When you launch an instance using RunInstances (), you can specify
// a launch template instead of providing the launch parameters in the request. For
// more information, see Launching an instance from a launch template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateLaunchTemplate(ctx context.Context, params *CreateLaunchTemplateInput, optFns ...func(*Options)) (*CreateLaunchTemplateOutput, error) {
	stack := middleware.NewStack("CreateLaunchTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateLaunchTemplateMiddlewares(stack)
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
	addOpCreateLaunchTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchTemplate(options.Region), middleware.Before)
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
			OperationName: "CreateLaunchTemplate",
			Err:           err,
		}
	}
	out := result.(*CreateLaunchTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchTemplateInput struct {
	// The tags to apply to the launch template during creation.
	TagSpecifications []*types.TagSpecification
	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	// Constraint: Maximum 128 ASCII characters.
	ClientToken *string
	// A name for the launch template.
	LaunchTemplateName *string
	// The information for the launch template.
	LaunchTemplateData *types.RequestLaunchTemplateData
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// A description for the first version of the launch template.
	VersionDescription *string
}

type CreateLaunchTemplateOutput struct {
	// Information about the launch template.
	LaunchTemplate *types.LaunchTemplate
	// If the launch template contains parameters or parameter combinations that are
	// not valid, an error code and an error message are returned for each issue that's
	// found.
	Warning *types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateLaunchTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateLaunchTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLaunchTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLaunchTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLaunchTemplate",
	}
}
