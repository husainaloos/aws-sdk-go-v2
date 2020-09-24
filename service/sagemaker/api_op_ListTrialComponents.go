// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the trial components in your account. You can sort the list by trial
// component name or creation time. You can filter the list to show only components
// that were created in a specific time range. You can also filter on one of the
// following:
//
//     * ExperimentName
//
//     * SourceArn
//
//     * TrialName
func (c *Client) ListTrialComponents(ctx context.Context, params *ListTrialComponentsInput, optFns ...func(*Options)) (*ListTrialComponentsOutput, error) {
	stack := middleware.NewStack("ListTrialComponents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTrialComponentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTrialComponents(options.Region), middleware.Before)
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
			OperationName: "ListTrialComponents",
			Err:           err,
		}
	}
	out := result.(*ListTrialComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrialComponentsInput struct {
	// If the previous call to ListTrialComponents didn't return the full set of
	// components, the call returns a token for getting the next set of components.
	NextToken *string
	// A filter that returns only components created before the specified time.
	CreatedBefore *time.Time
	// A filter that returns only components that have the specified source Amazon
	// Resource Name (ARN). If you specify SourceArn, you can't filter by
	// ExperimentName or TrialName.
	SourceArn *string
	// The maximum number of components to return in the response. The default value is
	// 10.
	MaxResults *int32
	// A filter that returns only components that are part of the specified experiment.
	// If you specify ExperimentName, you can't filter by SourceArn or TrialName.
	ExperimentName *string
	// A filter that returns only components that are part of the specified trial. If
	// you specify TrialName, you can't filter by ExperimentName or SourceArn.
	TrialName *string
	// A filter that returns only components created after the specified time.
	CreatedAfter *time.Time
	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortTrialComponentsBy
	// The sort order. The default value is Descending.
	SortOrder types.SortOrder
}

type ListTrialComponentsOutput struct {
	// A token for getting the next set of components, if there are any.
	NextToken *string
	// A list of the summaries of your trial components.
	TrialComponentSummaries []*types.TrialComponentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTrialComponentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTrialComponents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrialComponents{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTrialComponents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrialComponents",
	}
}
