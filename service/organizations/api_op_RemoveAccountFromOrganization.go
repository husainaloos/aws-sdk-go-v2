// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified account from the organization. The removed account becomes
// a standalone account that isn't a member of any organization. It's no longer
// subject to any policies and is responsible for its own bill payments. The
// organization's master account is no longer charged for any expenses accrued by
// the member account after it's removed from the organization. This operation can
// be called only from the organization's master account. Member accounts can
// remove themselves with LeaveOrganization () instead. You can remove an account
// from your organization only if the account is configured with the information
// required to operate as a standalone account. When you create an account in an
// organization using the AWS Organizations console, API, or CLI commands, the
// information required of standalone accounts is not automatically collected. For
// an account that you want to make standalone, you must accept the end user
// license agreement (EULA), choose a support plan, provide and verify the required
// contact information, and provide a current payment method. AWS uses the payment
// method to charge for any billable (not free tier) AWS activity that occurs while
// the account isn't attached to an organization. To remove an account that doesn't
// yet have this information, you must sign in as the member account and follow the
// steps at  To leave an organization when all required account information has not
// yet been provided
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html#leave-without-all-info)
// in the AWS Organizations User Guide.
func (c *Client) RemoveAccountFromOrganization(ctx context.Context, params *RemoveAccountFromOrganizationInput, optFns ...func(*Options)) (*RemoveAccountFromOrganizationOutput, error) {
	if params == nil {
		params = &RemoveAccountFromOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveAccountFromOrganization", params, optFns, addOperationRemoveAccountFromOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveAccountFromOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAccountFromOrganizationInput struct {

	// The unique identifier (ID) of the member account that you want to remove from
	// the organization. The regex pattern (http://wikipedia.org/wiki/regex) for an
	// account ID string requires exactly 12 digits.
	//
	// This member is required.
	AccountId *string
}

type RemoveAccountFromOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveAccountFromOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveAccountFromOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveAccountFromOrganization{}, middleware.After)
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
	addOpRemoveAccountFromOrganizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAccountFromOrganization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveAccountFromOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "RemoveAccountFromOrganization",
	}
}
