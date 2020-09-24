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

// Runs an on-demand evaluation for the specified AWS Config rules against the last
// known configuration state of the resources. Use StartConfigRulesEvaluation when
// you want to test that a rule you updated is working as expected.
// StartConfigRulesEvaluation does not re-record the latest configuration state for
// your resources. It re-runs an evaluation against the last known state of your
// resources. You can specify up to 25 AWS Config rules per request.  <p>An
// existing <code>StartConfigRulesEvaluation</code> call for the specified rules
// must complete before you can call the API again. If you chose to have AWS Config
// stream to an Amazon SNS topic, you will receive a
// <code>ConfigRuleEvaluationStarted</code> notification when the evaluation
// starts.</p> <note> <p>You don't need to call the
// <code>StartConfigRulesEvaluation</code> API to run an evaluation for a new rule.
// When you create a rule, AWS Config evaluates your resources against the rule
// automatically. </p> </note> <p>The <code>StartConfigRulesEvaluation</code> API
// is useful if you want to run on-demand evaluations, such as the following
// example:</p> <ol> <li> <p>You have a custom rule that evaluates your IAM
// resources every 24 hours.</p> </li> <li> <p>You update your Lambda function to
// add additional conditions to your rule.</p> </li> <li> <p>Instead of waiting for
// the next periodic evaluation, you call the
// <code>StartConfigRulesEvaluation</code> API.</p> </li> <li> <p>AWS Config
// invokes your Lambda function and evaluates your IAM resources.</p> </li> <li>
// <p>Your custom rule will still run periodic evaluations every 24 hours.</p>
// </li> </ol>
func (c *Client) StartConfigRulesEvaluation(ctx context.Context, params *StartConfigRulesEvaluationInput, optFns ...func(*Options)) (*StartConfigRulesEvaluationOutput, error) {
	stack := middleware.NewStack("StartConfigRulesEvaluation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartConfigRulesEvaluationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartConfigRulesEvaluation(options.Region), middleware.Before)
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
			OperationName: "StartConfigRulesEvaluation",
			Err:           err,
		}
	}
	out := result.(*StartConfigRulesEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type StartConfigRulesEvaluationInput struct {
	// The list of names of AWS Config rules that you want to run evaluations for.
	ConfigRuleNames []*string
}

// The output when you start the evaluation for the specified AWS Config rule.
type StartConfigRulesEvaluationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartConfigRulesEvaluationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartConfigRulesEvaluation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartConfigRulesEvaluation{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartConfigRulesEvaluation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "StartConfigRulesEvaluation",
	}
}
