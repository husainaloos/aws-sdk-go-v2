// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation adds or updates tags for a specified domain. All tag operations
// are eventually consistent; subsequent operations might not immediately represent
// all issued operations.
func (c *Client) UpdateTagsForDomain(ctx context.Context, params *UpdateTagsForDomainInput, optFns ...func(*Options)) (*UpdateTagsForDomainOutput, error) {
	stack := middleware.NewStack("UpdateTagsForDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateTagsForDomainMiddlewares(stack)
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
	addOpUpdateTagsForDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTagsForDomain(options.Region), middleware.Before)
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
			OperationName: "UpdateTagsForDomain",
			Err:           err,
		}
	}
	out := result.(*UpdateTagsForDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdateTagsForDomainRequest includes the following elements.
type UpdateTagsForDomainInput struct {
	// The domain for which you want to add or update tags.
	DomainName *string
	// A list of the tag keys and values that you want to add or update. If you specify
	// a key that already exists, the corresponding value will be replaced.
	TagsToUpdate []*types.Tag
}

type UpdateTagsForDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateTagsForDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTagsForDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTagsForDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTagsForDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "UpdateTagsForDomain",
	}
}
