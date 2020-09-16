// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecated. Use DescribeLoa () instead. Gets the LOA-CFA for a connection. The
// Letter of Authorization - Connecting Facility Assignment (LOA-CFA) is a document
// that your APN partner or service provider uses when establishing your cross
// connect to AWS at the colocation facility. For more information, see Requesting
// Cross Connects at AWS Direct Connect Locations
// (https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html) in
// the AWS Direct Connect User Guide.
func (c *Client) DescribeConnectionLoa(ctx context.Context, params *DescribeConnectionLoaInput, optFns ...func(*Options)) (*DescribeConnectionLoaOutput, error) {
	stack := middleware.NewStack("DescribeConnectionLoa", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeConnectionLoaMiddlewares(stack)
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
	addOpDescribeConnectionLoaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectionLoa(options.Region), middleware.Before)

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
			OperationName: "DescribeConnectionLoa",
			Err:           err,
		}
	}
	out := result.(*DescribeConnectionLoaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectionLoaInput struct {
	// The name of the APN partner or service provider who establishes connectivity on
	// your behalf. If you specify this parameter, the LOA-CFA lists the provider name
	// alongside your company name as the requester of the cross connect.
	ProviderName *string
	// The standard media type for the LOA-CFA document. The only supported value is
	// application/pdf.
	LoaContentType types.LoaContentType
	// The ID of the connection.
	ConnectionId *string
}

type DescribeConnectionLoaOutput struct {
	// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA).
	Loa *types.Loa

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeConnectionLoaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConnectionLoa{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConnectionLoa{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeConnectionLoa(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeConnectionLoa",
	}
}