// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the timestamp and status of the last cache reset of a specific Amazon
// Lightsail content delivery network (CDN) distribution.
func (c *Client) GetDistributionLatestCacheReset(ctx context.Context, params *GetDistributionLatestCacheResetInput, optFns ...func(*Options)) (*GetDistributionLatestCacheResetOutput, error) {
	if params == nil {
		params = &GetDistributionLatestCacheResetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDistributionLatestCacheReset", params, optFns, addOperationGetDistributionLatestCacheResetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDistributionLatestCacheResetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDistributionLatestCacheResetInput struct {

	// The name of the distribution for which to return the timestamp of the last cache
	// reset.  <p>Use the <code>GetDistributions</code> action to get a list of
	// distribution names that you can specify.</p> <p>When omitted, the response
	// includes the latest cache reset timestamp of all your distributions.</p>
	DistributionName *string
}

type GetDistributionLatestCacheResetOutput struct {

	// The timestamp of the last cache reset (e.g., 1479734909.17) in Unix time format.
	CreateTime *time.Time

	// The status of the last cache reset.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDistributionLatestCacheResetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDistributionLatestCacheReset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDistributionLatestCacheReset{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDistributionLatestCacheReset(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDistributionLatestCacheReset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetDistributionLatestCacheReset",
	}
}
