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

// Returns domain configuration information about the specified Elasticsearch
// domains, including the domain ID, domain endpoint, and domain ARN.
func (c *Client) DescribeElasticsearchDomains(ctx context.Context, params *DescribeElasticsearchDomainsInput, optFns ...func(*Options)) (*DescribeElasticsearchDomainsOutput, error) {
	if params == nil {
		params = &DescribeElasticsearchDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeElasticsearchDomains", params, optFns, addOperationDescribeElasticsearchDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeElasticsearchDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeElasticsearchDomains () operation.
// By default, the API returns the status of all Elasticsearch domains.
type DescribeElasticsearchDomainsInput struct {

	// The Elasticsearch domains for which you want information.
	//
	// This member is required.
	DomainNames []*string
}

// The result of a DescribeElasticsearchDomains request. Contains the status of the
// specified domains or all domains owned by the account.
type DescribeElasticsearchDomainsOutput struct {

	// The status of the domains requested in the DescribeElasticsearchDomains request.
	//
	// This member is required.
	DomainStatusList []*types.ElasticsearchDomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeElasticsearchDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeElasticsearchDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeElasticsearchDomains{}, middleware.After)
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
	addOpDescribeElasticsearchDomainsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticsearchDomains(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeElasticsearchDomains(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeElasticsearchDomains",
	}
}
