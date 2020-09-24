// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds tags to an Amazon EMR resource. Tags make it easier to associate clusters
// in various ways, such as grouping clusters to track your Amazon EMR resource
// allocation costs. For more information, see Tag Clusters
// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-tags.html).
func (c *Client) AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) {
	stack := middleware.NewStack("AddTags", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddTagsMiddlewares(stack)
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
	addOpAddTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTags(options.Region), middleware.Before)
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
			OperationName: "AddTags",
			Err:           err,
		}
	}
	out := result.(*AddTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input identifies a cluster and a list of tags to attach.
type AddTagsInput struct {
	// A list of tags to associate with a cluster and propagate to EC2 instances. Tags
	// are user-defined key/value pairs that consist of a required key string with a
	// maximum of 128 characters, and an optional value string with a maximum of 256
	// characters.
	Tags []*types.Tag
	// The Amazon EMR resource identifier to which tags will be added. This value must
	// be a cluster identifier.
	ResourceId *string
}

// This output indicates the result of adding tags to a resource.
type AddTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddTagsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddTags{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTags{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "AddTags",
	}
}
