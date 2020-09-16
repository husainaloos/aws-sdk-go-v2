// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about backups of AWS CloudHSM clusters. This is a paginated
// operation, which means that each response might contain only a subset of all the
// backups. When the response contains only a subset of backups, it includes a
// NextToken value. Use this value in a subsequent DescribeBackups request to get
// more backups. When you receive a response with no NextToken (or an empty or null
// value), that means there are no more backups to get.
func (c *Client) DescribeBackups(ctx context.Context, params *DescribeBackupsInput, optFns ...func(*Options)) (*DescribeBackupsOutput, error) {
	stack := middleware.NewStack("DescribeBackups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeBackupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackups(options.Region), middleware.Before)

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
			OperationName: "DescribeBackups",
			Err:           err,
		}
	}
	out := result.(*DescribeBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupsInput struct {
	// The NextToken value that you received in the previous response. Use this value
	// to get more backups.
	NextToken *string
	// The maximum number of backups to return in the response. When there are more
	// backups than the number you specify, the response contains a NextToken value.
	MaxResults *int32
	// Designates whether or not to sort the return backups by ascending chronological
	// order of generation.
	SortAscending *bool
	// One or more filters to limit the items returned in the response. Use the
	// backupIds filter to return only the specified backups. Specify backups by their
	// backup identifier (ID). Use the sourceBackupIds filter to return only the
	// backups created from a source backup. The sourceBackupID of a source backup is
	// returned by the CopyBackupToRegion () operation. Use the clusterIds filter to
	// return only the backups for the specified clusters. Specify clusters by their
	// cluster identifier (ID). Use the states filter to return only backups that match
	// the specified state.
	Filters map[string][]*string
}

type DescribeBackupsOutput struct {
	// An opaque string that indicates that the response contains only a subset of
	// backups. Use this value in a subsequent DescribeBackups request to get more
	// backups.
	NextToken *string
	// A list of backups.
	Backups []*types.Backup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeBackupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBackups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBackups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeBackups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DescribeBackups",
	}
}