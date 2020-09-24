// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// You can call DescribeValidDBInstanceModifications () to learn what modifications
// you can make to your DB instance. You can use this information when you call
// ModifyDBInstance ().
func (c *Client) DescribeValidDBInstanceModifications(ctx context.Context, params *DescribeValidDBInstanceModificationsInput, optFns ...func(*Options)) (*DescribeValidDBInstanceModificationsOutput, error) {
	stack := middleware.NewStack("DescribeValidDBInstanceModifications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeValidDBInstanceModificationsMiddlewares(stack)
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
	addOpDescribeValidDBInstanceModificationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeValidDBInstanceModifications(options.Region), middleware.Before)
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
			OperationName: "DescribeValidDBInstanceModifications",
			Err:           err,
		}
	}
	out := result.(*DescribeValidDBInstanceModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeValidDBInstanceModificationsInput struct {
	// The customer identifier or the ARN of your DB instance.
	DBInstanceIdentifier *string
}

type DescribeValidDBInstanceModificationsOutput struct {
	// Information about valid modifications that you can make to your DB instance.
	// Contains the result of a successful call to the
	// DescribeValidDBInstanceModifications () action. You can use this information
	// when you call ModifyDBInstance ().
	ValidDBInstanceModificationsMessage *types.ValidDBInstanceModificationsMessage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeValidDBInstanceModificationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeValidDBInstanceModifications{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeValidDBInstanceModifications{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeValidDBInstanceModifications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeValidDBInstanceModifications",
	}
}
