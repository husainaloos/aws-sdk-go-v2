// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the information for an existing site. To remove information for any of
// the parameters, specify an empty string.
func (c *Client) UpdateSite(ctx context.Context, params *UpdateSiteInput, optFns ...func(*Options)) (*UpdateSiteOutput, error) {
	stack := middleware.NewStack("UpdateSite", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateSiteMiddlewares(stack)
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
	addOpUpdateSiteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSite(options.Region), middleware.Before)
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
			OperationName: "UpdateSite",
			Err:           err,
		}
	}
	out := result.(*UpdateSiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSiteInput struct {
	// The ID of the global network.
	GlobalNetworkId *string
	// The ID of your site.
	SiteId *string
	// The site location:
	//
	//     * Address: The physical address of the site.
	//
	//     *
	// Latitude: The latitude of the site.
	//
	//     * Longitude: The longitude of the site.
	Location *types.Location
	// A description of your site. Length Constraints: Maximum length of 256
	// characters.
	Description *string
}

type UpdateSiteOutput struct {
	// Information about the site.
	Site *types.Site

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateSiteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSite{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSite{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSite(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "UpdateSite",
	}
}
