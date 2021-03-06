// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the versions of a device definition.
func (c *Client) ListDeviceDefinitionVersions(ctx context.Context, params *ListDeviceDefinitionVersionsInput, optFns ...func(*Options)) (*ListDeviceDefinitionVersionsOutput, error) {
	if params == nil {
		params = &ListDeviceDefinitionVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeviceDefinitionVersions", params, optFns, addOperationListDeviceDefinitionVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeviceDefinitionVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeviceDefinitionVersionsInput struct {

	// The ID of the device definition.
	//
	// This member is required.
	DeviceDefinitionId *string

	// The maximum number of results to be returned per request.
	MaxResults *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type ListDeviceDefinitionVersionsOutput struct {

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// Information about a version.
	Versions []*types.VersionInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeviceDefinitionVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeviceDefinitionVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeviceDefinitionVersions{}, middleware.After)
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
	addOpListDeviceDefinitionVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeviceDefinitionVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDeviceDefinitionVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListDeviceDefinitionVersions",
	}
}
