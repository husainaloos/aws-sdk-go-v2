// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the permissions, if any, that have been assigned by a private CA.
// Permissions can be granted with the CreatePermission () action and revoked with
// the DeletePermission () action.
func (c *Client) ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
	stack := middleware.NewStack("ListPermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPermissionsMiddlewares(stack)
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
	addOpListPermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissions(options.Region), middleware.Before)
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
			OperationName: "ListPermissions",
			Err:           err,
		}
	}
	out := result.(*ListPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionsInput struct {
	// When paginating results, use this parameter in a subsequent request after you
	// receive a response with truncated results. Set it to the value of NextToken from
	// the response you just received.
	NextToken *string
	// When paginating results, use this parameter to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken value
	// in a subsequent request to retrieve additional items.
	MaxResults *int32
	// The Amazon Resource Number (ARN) of the private CA to inspect. You can find the
	// ARN by calling the ListCertificateAuthorities () action. This must be of the
	// form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// You can get a private CA's ARN by running the ListCertificateAuthorities ()
	// action.
	CertificateAuthorityArn *string
}

type ListPermissionsOutput struct {
	// Summary information about each permission assigned by the specified private CA,
	// including the action enabled, the policy provided, and the time of creation.
	Permissions []*types.Permission
	// When the list is truncated, this value is present and should be used for the
	// NextToken parameter in a subsequent pagination request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "ListPermissions",
	}
}
