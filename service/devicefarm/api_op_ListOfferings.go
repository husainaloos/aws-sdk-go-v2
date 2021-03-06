// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of products or offerings that the user can manage through the
// API. Each offering record indicates the recurring price per unit and the
// frequency for that offering. The API returns a NotEligible error if the user is
// not permitted to invoke the operation. If you must be able to invoke this
// operation, contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com).
func (c *Client) ListOfferings(ctx context.Context, params *ListOfferingsInput, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	if params == nil {
		params = &ListOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOfferings", params, optFns, addOperationListOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list all offerings.
type ListOfferingsInput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the return values of the list of offerings.
type ListOfferingsOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// A value that represents the list offering results.
	Offerings []*types.Offering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOfferings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListOfferings",
	}
}
