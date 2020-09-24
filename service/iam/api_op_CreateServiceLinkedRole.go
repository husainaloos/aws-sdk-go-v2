// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an IAM role that is linked to a specific AWS service. The service
// controls the attached policies and when the role can be deleted. This helps
// ensure that the service is not broken by an unexpectedly changed or deleted
// role, which could put your AWS resources into an unknown state. Allowing the
// service to control the role helps improve service stability and proper cleanup
// when a service and its role are no longer needed. For more information, see
// Using Service-Linked Roles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in the IAM User Guide. To attach a policy to this service-linked role, you must
// make the request using the AWS service that depends on this role.
func (c *Client) CreateServiceLinkedRole(ctx context.Context, params *CreateServiceLinkedRoleInput, optFns ...func(*Options)) (*CreateServiceLinkedRoleOutput, error) {
	stack := middleware.NewStack("CreateServiceLinkedRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateServiceLinkedRoleMiddlewares(stack)
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
	addOpCreateServiceLinkedRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceLinkedRole(options.Region), middleware.Before)
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
			OperationName: "CreateServiceLinkedRole",
			Err:           err,
		}
	}
	out := result.(*CreateServiceLinkedRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceLinkedRoleInput struct {
	// The service principal for the AWS service to which this role is attached. You
	// use a string similar to a URL but without the http:// in front. For example:
	// elasticbeanstalk.amazonaws.com. Service principals are unique and
	// case-sensitive. To find the exact service principal for your service-linked
	// role, see AWS Services That Work with IAM
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html)
	// in the IAM User Guide. Look for the services that have Yes in the Service-Linked
	// Role column. Choose the Yes link to view the service-linked role documentation
	// for that service.
	AWSServiceName *string
	// The description of the role.
	Description *string
	// A string that you provide, which is combined with the service-provided prefix to
	// form the complete role name. If you make multiple requests for the same service,
	// then you must supply a different CustomSuffix for each request. Otherwise the
	// request fails with a duplicate role name error. For example, you could add -1 or
	// -debug to the suffix. Some services do not support the CustomSuffix parameter.
	// If you provide an optional suffix and the operation fails, try the operation
	// again without the suffix.
	CustomSuffix *string
}

type CreateServiceLinkedRoleOutput struct {
	// A Role () object that contains details about the newly created role.
	Role *types.Role

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateServiceLinkedRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateServiceLinkedRole{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateServiceLinkedRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateServiceLinkedRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateServiceLinkedRole",
	}
}
