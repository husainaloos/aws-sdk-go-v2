// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents the failure of a third party job as returned to the pipeline by a job
// worker. Used for partner actions only.
func (c *Client) PutThirdPartyJobFailureResult(ctx context.Context, params *PutThirdPartyJobFailureResultInput, optFns ...func(*Options)) (*PutThirdPartyJobFailureResultOutput, error) {
	stack := middleware.NewStack("PutThirdPartyJobFailureResult", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutThirdPartyJobFailureResultMiddlewares(stack)
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
	addOpPutThirdPartyJobFailureResultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutThirdPartyJobFailureResult(options.Region), middleware.Before)
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
			OperationName: "PutThirdPartyJobFailureResult",
			Err:           err,
		}
	}
	out := result.(*PutThirdPartyJobFailureResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutThirdPartyJobFailureResult action.
type PutThirdPartyJobFailureResultInput struct {
	// The clientToken portion of the clientId and clientToken pair used to verify that
	// the calling entity is allowed access to the job and its details.
	ClientToken *string
	// The ID of the job that failed. This is the same ID returned from
	// PollForThirdPartyJobs.
	JobId *string
	// Represents information about failure details.
	FailureDetails *types.FailureDetails
}

type PutThirdPartyJobFailureResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutThirdPartyJobFailureResultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutThirdPartyJobFailureResult{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutThirdPartyJobFailureResult{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutThirdPartyJobFailureResult(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutThirdPartyJobFailureResult",
	}
}
