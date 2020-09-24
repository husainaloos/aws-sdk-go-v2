// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the evaluation results for the specified AWS Config rule. You can
// specify one AWS Config rule per request. After you delete the evaluation
// results, you can call the StartConfigRulesEvaluation () API to start evaluating
// your AWS resources against the rule.
func (c *Client) DeleteEvaluationResults(ctx context.Context, params *DeleteEvaluationResultsInput, optFns ...func(*Options)) (*DeleteEvaluationResultsOutput, error) {
	stack := middleware.NewStack("DeleteEvaluationResults", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteEvaluationResultsMiddlewares(stack)
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
	addOpDeleteEvaluationResultsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEvaluationResults(options.Region), middleware.Before)
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
			OperationName: "DeleteEvaluationResults",
			Err:           err,
		}
	}
	out := result.(*DeleteEvaluationResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteEvaluationResultsInput struct {
	// The name of the AWS Config rule for which you want to delete the evaluation
	// results.
	ConfigRuleName *string
}

// The output when you delete the evaluation results for the specified AWS Config
// rule.
type DeleteEvaluationResultsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteEvaluationResultsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEvaluationResults{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEvaluationResults{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteEvaluationResults(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteEvaluationResults",
	}
}
