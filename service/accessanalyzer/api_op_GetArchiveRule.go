// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about an archive rule.
func (c *Client) GetArchiveRule(ctx context.Context, params *GetArchiveRuleInput, optFns ...func(*Options)) (*GetArchiveRuleOutput, error) {
	stack := middleware.NewStack("GetArchiveRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetArchiveRuleMiddlewares(stack)
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
	addOpGetArchiveRuleValidationMiddleware(stack)
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
			OperationName: "GetArchiveRule",
			Err:           err,
		}
	}
	out := result.(*GetArchiveRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves an archive rule.
type GetArchiveRuleInput struct {
	// The name of the analyzer to retrieve rules from.
	AnalyzerName *string
	// The name of the rule to retrieve.
	RuleName *string
}

// The response to the request.
type GetArchiveRuleOutput struct {
	// Contains information about an archive rule.
	ArchiveRule *types.ArchiveRuleSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetArchiveRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetArchiveRule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetArchiveRule{}, middleware.After)
}
