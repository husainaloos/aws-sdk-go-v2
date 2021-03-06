// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a specific version of a solution. For more information on solutions,
// see CreateSolution ().
func (c *Client) DescribeSolutionVersion(ctx context.Context, params *DescribeSolutionVersionInput, optFns ...func(*Options)) (*DescribeSolutionVersionOutput, error) {
	if params == nil {
		params = &DescribeSolutionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSolutionVersion", params, optFns, addOperationDescribeSolutionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSolutionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSolutionVersionInput struct {

	// The Amazon Resource Name (ARN) of the solution version.
	//
	// This member is required.
	SolutionVersionArn *string
}

type DescribeSolutionVersionOutput struct {

	// The solution version.
	SolutionVersion *types.SolutionVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSolutionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSolutionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSolutionVersion{}, middleware.After)
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
	addOpDescribeSolutionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSolutionVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSolutionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeSolutionVersion",
	}
}
