// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an log pattern to a LogPatternSet.
func (c *Client) CreateLogPattern(ctx context.Context, params *CreateLogPatternInput, optFns ...func(*Options)) (*CreateLogPatternOutput, error) {
	stack := middleware.NewStack("CreateLogPattern", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateLogPatternMiddlewares(stack)
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
	addOpCreateLogPatternValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogPattern(options.Region), middleware.Before)
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
			OperationName: "CreateLogPattern",
			Err:           err,
		}
	}
	out := result.(*CreateLogPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogPatternInput struct {
	// Rank of the log pattern.
	Rank *int32
	// The name of the log pattern.
	PatternName *string
	// The name of the resource group.
	ResourceGroupName *string
	// The name of the log pattern set.
	PatternSetName *string
	// The log pattern.
	Pattern *string
}

type CreateLogPatternOutput struct {
	// The name of the resource group.
	ResourceGroupName *string
	// The successfully created log pattern.
	LogPattern *types.LogPattern

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateLogPatternMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLogPattern{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLogPattern{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLogPattern(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "CreateLogPattern",
	}
}
