// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables an account as a delegated administrator of Amazon Macie for an AWS
// organization.
func (c *Client) DisableOrganizationAdminAccount(ctx context.Context, params *DisableOrganizationAdminAccountInput, optFns ...func(*Options)) (*DisableOrganizationAdminAccountOutput, error) {
	stack := middleware.NewStack("DisableOrganizationAdminAccount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisableOrganizationAdminAccountMiddlewares(stack)
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
	addOpDisableOrganizationAdminAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableOrganizationAdminAccount(options.Region), middleware.Before)
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
			OperationName: "DisableOrganizationAdminAccount",
			Err:           err,
		}
	}
	out := result.(*DisableOrganizationAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableOrganizationAdminAccountInput struct {
	// The AWS account ID of the delegated administrator account.
	AdminAccountId *string
}

type DisableOrganizationAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisableOrganizationAdminAccountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisableOrganizationAdminAccount{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisableOrganizationAdminAccount{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableOrganizationAdminAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "DisableOrganizationAdminAccount",
	}
}
