// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListPresets operation gets a list of the default presets included with
// Elastic Transcoder and the presets that you've added in an AWS region.
func (c *Client) ListPresets(ctx context.Context, params *ListPresetsInput, optFns ...func(*Options)) (*ListPresetsOutput, error) {
	stack := middleware.NewStack("ListPresets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPresetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPresets(options.Region), middleware.Before)
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
			OperationName: "ListPresets",
			Err:           err,
		}
	}
	out := result.(*ListPresetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListPresetsRequest structure.
type ListPresetsInput struct {
	// When Elastic Transcoder returns more than one page of results, use pageToken in
	// subsequent GET requests to get each successive page of results.
	PageToken *string
	// To list presets in chronological order by the date and time that they were
	// created, enter true. To list presets in reverse chronological order, enter
	// false.
	Ascending *string
}

// The ListPresetsResponse structure.
type ListPresetsOutput struct {
	// An array of Preset objects.
	Presets []*types.Preset
	// A value that you use to access the second and subsequent pages of results, if
	// any. When the presets fit on one page or when you've reached the last page of
	// results, the value of NextPageToken is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPresetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPresets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPresets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPresets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "ListPresets",
	}
}
