// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of the service quota increase request in your template.
func (c *Client) GetServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, params *GetServiceQuotaIncreaseRequestFromTemplateInput, optFns ...func(*Options)) (*GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	stack := middleware.NewStack("GetServiceQuotaIncreaseRequestFromTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetServiceQuotaIncreaseRequestFromTemplateMiddlewares(stack)
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
	addOpGetServiceQuotaIncreaseRequestFromTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(options.Region), middleware.Before)
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
			OperationName: "GetServiceQuotaIncreaseRequestFromTemplate",
			Err:           err,
		}
	}
	out := result.(*GetServiceQuotaIncreaseRequestFromTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceQuotaIncreaseRequestFromTemplateInput struct {
	// Specifies the quota you want.
	QuotaCode *string
	// Specifies the AWS Region for the quota that you want to use.
	AwsRegion *string
	// Specifies the service that you want to use.
	ServiceCode *string
}

type GetServiceQuotaIncreaseRequestFromTemplateOutput struct {
	// This object contains the details about the quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetServiceQuotaIncreaseRequestFromTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetServiceQuotaIncreaseRequestFromTemplate",
	}
}
