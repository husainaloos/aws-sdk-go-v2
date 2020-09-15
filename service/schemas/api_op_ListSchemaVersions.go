// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides a list of the schema versions and related information.
func (c *Client) ListSchemaVersions(ctx context.Context, params *ListSchemaVersionsInput, optFns ...func(*Options)) (*ListSchemaVersionsOutput, error) {
	stack := middleware.NewStack("ListSchemaVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSchemaVersionsMiddlewares(stack)
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
	addOpListSchemaVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemaVersions(options.Region), middleware.Before)

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
			OperationName: "ListSchemaVersions",
			Err:           err,
		}
	}
	out := result.(*ListSchemaVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemaVersionsInput struct {
	// The name of the registry.
	RegistryName *string
	// The name of the schema.
	SchemaName *string
	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string
	Limit     *int32
}

type ListSchemaVersionsOutput struct {
	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string
	// An array of schema version summaries.
	SchemaVersions []*types.SchemaVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSchemaVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSchemaVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSchemaVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSchemaVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "ListSchemaVersions",
	}
}