// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the account that's designated as the delegated
// administrator of Amazon Macie for an AWS organization.
func (c *Client) ListOrganizationAdminAccounts(ctx context.Context, params *ListOrganizationAdminAccountsInput, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) {
	stack := middleware.NewStack("ListOrganizationAdminAccounts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListOrganizationAdminAccountsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationAdminAccounts(options.Region), middleware.Before)
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
			OperationName: "ListOrganizationAdminAccounts",
			Err:           err,
		}
	}
	out := result.(*ListOrganizationAdminAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationAdminAccountsInput struct {
	// The maximum number of items to include in each page of a paginated response.
	MaxResults *int32
	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string
}

type ListOrganizationAdminAccountsOutput struct {
	// An array of objects, one for each account that's designated as a delegated
	// administrator of Amazon Macie for the AWS organization. Of those accounts, only
	// one can have a status of ENABLED.
	AdminAccounts []*types.AdminAccount
	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListOrganizationAdminAccountsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListOrganizationAdminAccounts{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrganizationAdminAccounts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListOrganizationAdminAccounts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "ListOrganizationAdminAccounts",
	}
}
