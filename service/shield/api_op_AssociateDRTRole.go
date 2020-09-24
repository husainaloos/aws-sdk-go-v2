// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Authorizes the DDoS Response Team (DRT), using the specified role, to access
// your AWS account to assist with DDoS attack mitigation during potential attacks.
// This enables the DRT to inspect your AWS WAF configuration and create or update
// AWS WAF rules and web ACLs. You can associate only one RoleArn with your
// subscription. If you submit an AssociateDRTRole request for an account that
// already has an associated role, the new RoleArn will replace the existing
// RoleArn. Prior to making the AssociateDRTRole request, you must attach the
// AWSShieldDRTAccessPolicy
// (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy)
// managed policy to the role you will specify in the request. For more information
// see Attaching and Detaching IAM Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html).
// The role must also trust the service principal  drt.shield.amazonaws.com. For
// more information, see IAM JSON Policy Elements: Principal
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html).
// <p>The DRT will have access only to your AWS WAF and Shield resources. By
// submitting this request, you authorize the DRT to inspect your AWS WAF and
// Shield configuration and create and update AWS WAF rules and web ACLs on your
// behalf. The DRT takes these actions only if explicitly authorized by you.</p>
// <p>You must have the <code>iam:PassRole</code> permission to make an
// <code>AssociateDRTRole</code> request. For more information, see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html">Granting
// a User Permissions to Pass a Role to an AWS Service</a>. </p> <p>To use the
// services of the DRT and make an <code>AssociateDRTRole</code> request, you must
// be subscribed to the <a
// href="https://aws.amazon.com/premiumsupport/business-support/">Business Support
// plan</a> or the <a
// href="https://aws.amazon.com/premiumsupport/enterprise-support/">Enterprise
// Support plan</a>.</p>
func (c *Client) AssociateDRTRole(ctx context.Context, params *AssociateDRTRoleInput, optFns ...func(*Options)) (*AssociateDRTRoleOutput, error) {
	stack := middleware.NewStack("AssociateDRTRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateDRTRoleMiddlewares(stack)
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
	addOpAssociateDRTRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDRTRole(options.Region), middleware.Before)
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
			OperationName: "AssociateDRTRole",
			Err:           err,
		}
	}
	out := result.(*AssociateDRTRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDRTRoleInput struct {
	// The Amazon Resource Name (ARN) of the role the DRT will use to access your AWS
	// account. Prior to making the AssociateDRTRole request, you must attach the
	// AWSShieldDRTAccessPolicy
	// (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy)
	// managed policy to this role. For more information see Attaching and Detaching
	// IAM Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html).
	RoleArn *string
}

type AssociateDRTRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateDRTRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateDRTRole{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateDRTRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateDRTRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "AssociateDRTRole",
	}
}
