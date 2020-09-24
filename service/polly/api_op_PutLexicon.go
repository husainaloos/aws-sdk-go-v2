// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stores a pronunciation lexicon in an AWS Region. If a lexicon with the same name
// already exists in the region, it is overwritten by the new lexicon. Lexicon
// operations have eventual consistency, therefore, it might take some time before
// the lexicon is available to the SynthesizeSpeech operation. For more
// information, see Managing Lexicons
// (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
func (c *Client) PutLexicon(ctx context.Context, params *PutLexiconInput, optFns ...func(*Options)) (*PutLexiconOutput, error) {
	stack := middleware.NewStack("PutLexicon", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutLexiconMiddlewares(stack)
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
	addOpPutLexiconValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLexicon(options.Region), middleware.Before)
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
			OperationName: "PutLexicon",
			Err:           err,
		}
	}
	out := result.(*PutLexiconOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLexiconInput struct {
	// Content of the PLS lexicon as string data.
	Content *string
	// Name of the lexicon. The name must follow the regular express format
	// [0-9A-Za-z]{1,20}. That is, the name is a case-sensitive alphanumeric string up
	// to 20 characters long.
	Name *string
}

type PutLexiconOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutLexiconMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutLexicon{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLexicon{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutLexicon(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "PutLexicon",
	}
}
