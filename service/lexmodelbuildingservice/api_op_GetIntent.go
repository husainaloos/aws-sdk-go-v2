// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about an intent. In addition to the intent name, you must
// specify the intent version. This operation requires permissions to perform the
// lex:GetIntent action.
func (c *Client) GetIntent(ctx context.Context, params *GetIntentInput, optFns ...func(*Options)) (*GetIntentOutput, error) {
	stack := middleware.NewStack("GetIntent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetIntentMiddlewares(stack)
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
	addOpGetIntentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntent(options.Region), middleware.Before)
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
			OperationName: "GetIntent",
			Err:           err,
		}
	}
	out := result.(*GetIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntentInput struct {
	// The name of the intent. The name is case sensitive.
	Name *string
	// The version of the intent.
	Version *string
}

type GetIntentOutput struct {
	// A description of the intent.
	Description *string
	// Configuration information, if any, to connect to an Amazon Kendra index with the
	// AMAZON.KendraSearchIntent intent.
	KendraConfiguration *types.KendraConfiguration
	// A unique identifier for a built-in intent.
	ParentIntentSignature *string
	// After the Lambda function specified in the fulfillmentActivity element fulfills
	// the intent, Amazon Lex conveys this statement to the user.
	ConclusionStatement *types.Statement
	// An array of sample utterances configured for the intent.
	SampleUtterances []*string
	// The date that the intent was created.
	CreatedDate *time.Time
	// If the user answers "no" to the question defined in confirmationPrompt, Amazon
	// Lex responds with this statement to acknowledge that the intent was canceled.
	RejectionStatement *types.Statement
	// Checksum of the intent.
	Checksum *string
	// If defined in the bot, Amazon Amazon Lex invokes this Lambda function for each
	// user input. For more information, see PutIntent ().
	DialogCodeHook *types.CodeHook
	// An array of intent slots configured for the intent.
	Slots []*types.Slot
	// The name of the intent.
	Name *string
	// The date that the intent was updated. When you create a resource, the creation
	// date and the last updated date are the same.
	LastUpdatedDate *time.Time
	// If defined in the bot, Amazon Lex uses prompt to confirm the intent before
	// fulfilling the user's request. For more information, see PutIntent ().
	ConfirmationPrompt *types.Prompt
	// Describes how the intent is fulfilled. For more information, see PutIntent ().
	FulfillmentActivity *types.FulfillmentActivity
	// If defined in the bot, Amazon Lex uses this prompt to solicit additional user
	// activity after the intent is fulfilled. For more information, see PutIntent ().
	FollowUpPrompt *types.FollowUpPrompt
	// The version of the intent.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetIntentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetIntent{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntent{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIntent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetIntent",
	}
}
