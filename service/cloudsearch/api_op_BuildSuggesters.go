// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Indexes the search suggestions. For more information, see Configuring Suggesters
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html#configuring-suggesters)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) BuildSuggesters(ctx context.Context, params *BuildSuggestersInput, optFns ...func(*Options)) (*BuildSuggestersOutput, error) {
	stack := middleware.NewStack("BuildSuggesters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpBuildSuggestersMiddlewares(stack)
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
	addOpBuildSuggestersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBuildSuggesters(options.Region), middleware.Before)
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
			OperationName: "BuildSuggesters",
			Err:           err,
		}
	}
	out := result.(*BuildSuggestersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the BuildSuggester () operation. Specifies the
// name of the domain you want to update.
type BuildSuggestersInput struct {
	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	DomainName *string
}

// The result of a BuildSuggester request. Contains a list of the fields used for
// suggestions.
type BuildSuggestersOutput struct {
	// A list of field names.
	FieldNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpBuildSuggestersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpBuildSuggesters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpBuildSuggesters{}, middleware.After)
}

func newServiceMetadataMiddleware_opBuildSuggesters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "BuildSuggesters",
	}
}
