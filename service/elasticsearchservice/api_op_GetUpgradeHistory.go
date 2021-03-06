// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the complete history of the last 10 upgrades that were performed on
// the domain.
func (c *Client) GetUpgradeHistory(ctx context.Context, params *GetUpgradeHistoryInput, optFns ...func(*Options)) (*GetUpgradeHistoryOutput, error) {
	if params == nil {
		params = &GetUpgradeHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUpgradeHistory", params, optFns, addOperationGetUpgradeHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUpgradeHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to GetUpgradeHistory () operation.
type GetUpgradeHistoryInput struct {

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter or
	// number and can contain the following characters: a-z (lowercase), 0-9, and -
	// (hyphen).
	//
	// This member is required.
	DomainName *string

	// Set this value to limit the number of results returned.
	MaxResults *int32

	// Paginated APIs accepts NextToken input to returns next page results and provides
	// a NextToken output in the response which can be used by the client to retrieve
	// more results.
	NextToken *string
}

// Container for response returned by GetUpgradeHistory () operation.
type GetUpgradeHistoryOutput struct {

	// Pagination token that needs to be supplied to the next call to get the next page
	// of results
	NextToken *string

	// A list of UpgradeHistory () objects corresponding to each Upgrade or Upgrade
	// Eligibility Check performed on a domain returned as part of
	// GetUpgradeHistoryResponse () object.
	UpgradeHistories []*types.UpgradeHistory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUpgradeHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUpgradeHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUpgradeHistory{}, middleware.After)
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
	addOpGetUpgradeHistoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUpgradeHistory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetUpgradeHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "GetUpgradeHistory",
	}
}
