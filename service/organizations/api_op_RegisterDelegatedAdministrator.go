// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the specified member account to administer the Organizations features of
// the specified AWS service. It grants read-only access to AWS Organizations
// service data. The account still requires IAM permissions to access and
// administer the AWS service. You can run this action only for AWS services that
// support this feature. For a current list of services that support it, see the
// column Supports Delegated Administrator in the table at AWS Services that you
// can use with AWS Organizations
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrated-services-list.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's master account.
func (c *Client) RegisterDelegatedAdministrator(ctx context.Context, params *RegisterDelegatedAdministratorInput, optFns ...func(*Options)) (*RegisterDelegatedAdministratorOutput, error) {
	stack := middleware.NewStack("RegisterDelegatedAdministrator", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterDelegatedAdministratorMiddlewares(stack)
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
	addOpRegisterDelegatedAdministratorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDelegatedAdministrator(options.Region), middleware.Before)
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
			OperationName: "RegisterDelegatedAdministrator",
			Err:           err,
		}
	}
	out := result.(*RegisterDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDelegatedAdministratorInput struct {
	// The service principal of the AWS service for which you want to make the member
	// account a delegated administrator.
	ServicePrincipal *string
	// The account ID number of the member account in the organization to register as a
	// delegated administrator.
	AccountId *string
}

type RegisterDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterDelegatedAdministratorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterDelegatedAdministrator{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterDelegatedAdministrator{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterDelegatedAdministrator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "RegisterDelegatedAdministrator",
	}
}
