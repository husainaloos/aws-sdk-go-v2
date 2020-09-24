// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the settings of a data set.
func (c *Client) UpdateDataset(ctx context.Context, params *UpdateDatasetInput, optFns ...func(*Options)) (*UpdateDatasetOutput, error) {
	stack := middleware.NewStack("UpdateDataset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDatasetMiddlewares(stack)
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
	addOpUpdateDatasetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataset(options.Region), middleware.Before)
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
			OperationName: "UpdateDataset",
			Err:           err,
		}
	}
	out := result.(*UpdateDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatasetInput struct {
	// The name of the data set to update.
	DatasetName *string
	// [Optional] How many versions of data set contents are kept. If not specified or
	// set to null, only the latest version plus the latest succeeded version (if they
	// are different) are kept for the time period specified by the "retentionPeriod"
	// parameter. (For more information, see
	// https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	VersioningConfiguration *types.VersioningConfiguration
	// When data set contents are created they are delivered to destinations specified
	// here.
	ContentDeliveryRules []*types.DatasetContentDeliveryRule
	// A list of "DatasetAction" objects.
	Actions []*types.DatasetAction
	// How long, in days, data set contents are kept for the data set.
	RetentionPeriod *types.RetentionPeriod
	// A list of "DatasetTrigger" objects. The list can be empty or can contain up to
	// five DataSetTrigger objects.
	Triggers []*types.DatasetTrigger
}

type UpdateDatasetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDatasetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataset{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDataset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "UpdateDataset",
	}
}
