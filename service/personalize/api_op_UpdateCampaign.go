// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a campaign by either deploying a new solution or changing the value of
// the campaign's minProvisionedTPS parameter. To update a campaign, the campaign
// status must be ACTIVE or CREATE FAILED. Check the campaign status using the
// DescribeCampaign () API. You must wait until the status of the updated campaign
// is ACTIVE before asking the campaign for recommendations. For more information
// on campaigns, see CreateCampaign ().
func (c *Client) UpdateCampaign(ctx context.Context, params *UpdateCampaignInput, optFns ...func(*Options)) (*UpdateCampaignOutput, error) {
	if params == nil {
		params = &UpdateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCampaign", params, optFns, addOperationUpdateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCampaignInput struct {

	// The Amazon Resource Name (ARN) of the campaign.
	//
	// This member is required.
	CampaignArn *string

	// Specifies the requested minimum provisioned transactions (recommendations) per
	// second that Amazon Personalize will support.
	MinProvisionedTPS *int32

	// The ARN of a new solution version to deploy.
	SolutionVersionArn *string
}

type UpdateCampaignOutput struct {

	// The same campaign ARN as given in the request.
	CampaignArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCampaign{}, middleware.After)
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
	addOpUpdateCampaignValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCampaign(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateCampaign(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "UpdateCampaign",
	}
}
