// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the resource policy set on a domain.
func (c *Client) DeleteDomainPermissionsPolicy(ctx context.Context, params *DeleteDomainPermissionsPolicyInput, optFns ...func(*Options)) (*DeleteDomainPermissionsPolicyOutput, error) {
	stack := middleware.NewStack("DeleteDomainPermissionsPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteDomainPermissionsPolicyMiddlewares(stack)
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
	addOpDeleteDomainPermissionsPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDomainPermissionsPolicy(options.Region), middleware.Before)
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
			OperationName: "DeleteDomainPermissionsPolicy",
			Err:           err,
		}
	}
	out := result.(*DeleteDomainPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDomainPermissionsPolicyInput struct {
	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
	// The name of the domain associated with the resource policy to be deleted.
	Domain *string
	// The current revision of the resource policy to be deleted. This revision is used
	// for optimistic locking, which prevents others from overwriting your changes to
	// the domain's resource policy.
	PolicyRevision *string
}

type DeleteDomainPermissionsPolicyOutput struct {
	// Information about the deleted resource policy after processing the request.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteDomainPermissionsPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDomainPermissionsPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDomainPermissionsPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDomainPermissionsPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DeleteDomainPermissionsPolicy",
	}
}
