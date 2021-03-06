// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves metadata for a specified crawler.
func (c *Client) GetCrawler(ctx context.Context, params *GetCrawlerInput, optFns ...func(*Options)) (*GetCrawlerOutput, error) {
	if params == nil {
		params = &GetCrawlerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCrawler", params, optFns, addOperationGetCrawlerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCrawlerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCrawlerInput struct {

	// The name of the crawler to retrieve metadata for.
	//
	// This member is required.
	Name *string
}

type GetCrawlerOutput struct {

	// The metadata for the specified crawler.
	Crawler *types.Crawler

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCrawlerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCrawler{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCrawler{}, middleware.After)
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
	addOpGetCrawlerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCrawler(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCrawler(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetCrawler",
	}
}
