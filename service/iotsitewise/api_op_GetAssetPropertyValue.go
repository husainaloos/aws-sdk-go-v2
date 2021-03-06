// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets an asset property's current value. For more information, see Querying
// Current Property Values
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#current-values)
// in the AWS IoT SiteWise User Guide. To identify an asset property, you must
// specify one of the following:
//
//     * The assetId and propertyId of an asset
// property.
//
//     * A propertyAlias, which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature). To define an asset property's alias,
// see UpdateAssetProperty
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
func (c *Client) GetAssetPropertyValue(ctx context.Context, params *GetAssetPropertyValueInput, optFns ...func(*Options)) (*GetAssetPropertyValueOutput, error) {
	if params == nil {
		params = &GetAssetPropertyValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssetPropertyValue", params, optFns, addOperationGetAssetPropertyValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssetPropertyValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetPropertyValueInput struct {

	// The ID of the asset.
	AssetId *string

	// The property alias that identifies the property, such as an OPC-UA server data
	// stream path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping Industrial Data Streams to Asset Properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide.
	PropertyAlias *string

	// The ID of the asset property.
	PropertyId *string
}

type GetAssetPropertyValueOutput struct {

	// The current asset property value.
	PropertyValue *types.AssetPropertyValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAssetPropertyValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAssetPropertyValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssetPropertyValue{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssetPropertyValue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAssetPropertyValue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "GetAssetPropertyValue",
	}
}
