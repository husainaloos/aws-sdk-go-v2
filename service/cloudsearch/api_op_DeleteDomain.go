// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Permanently deletes a search domain and all of its data. Once a domain has been
// deleted, it cannot be recovered. For more information, see Deleting a Search
// Domain
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/deleting-domains.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DeleteDomain(ctx context.Context, params *DeleteDomainInput, optFns ...func(*Options)) (*DeleteDomainOutput, error) {
	if params == nil {
		params = &DeleteDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDomain", params, optFns, addOperationDeleteDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteDomain () operation. Specifies the
// name of the domain you want to delete.
type DeleteDomainInput struct {

	// The name of the domain you want to permanently delete.
	//
	// This member is required.
	DomainName *string
}

// The result of a DeleteDomain request. Contains the status of a newly deleted
// domain, or no status if the domain has already been completely deleted.
type DeleteDomainOutput struct {

	// The current status of the search domain.
	DomainStatus *types.DomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDomain{}, middleware.After)
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
	addOpDeleteDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDomain(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DeleteDomain",
	}
}
