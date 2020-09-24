// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about an organizational unit (OU). This operation can be
// called only from the organization's master account or by a member account that
// is a delegated administrator for an AWS service.
func (c *Client) DescribeOrganizationalUnit(ctx context.Context, params *DescribeOrganizationalUnitInput, optFns ...func(*Options)) (*DescribeOrganizationalUnitOutput, error) {
	stack := middleware.NewStack("DescribeOrganizationalUnit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeOrganizationalUnitMiddlewares(stack)
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
	addOpDescribeOrganizationalUnitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationalUnit(options.Region), middleware.Before)
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
			OperationName: "DescribeOrganizationalUnit",
			Err:           err,
		}
	}
	out := result.(*DescribeOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationalUnitInput struct {
	// The unique identifier (ID) of the organizational unit that you want details
	// about. You can get the ID from the ListOrganizationalUnitsForParent ()
	// operation. The regex pattern (http://wikipedia.org/wiki/regex) for an
	// organizational unit ID string requires "ou-" followed by from 4 to 32 lowercase
	// letters or digits (the ID of the root that contains the OU). This string is
	// followed by a second "-" dash and from 8 to 32 additional lowercase letters or
	// digits.
	OrganizationalUnitId *string
}

type DescribeOrganizationalUnitOutput struct {
	// A structure that contains details about the specified OU.
	OrganizationalUnit *types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeOrganizationalUnitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationalUnit{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationalUnit{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeOrganizationalUnit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeOrganizationalUnit",
	}
}
