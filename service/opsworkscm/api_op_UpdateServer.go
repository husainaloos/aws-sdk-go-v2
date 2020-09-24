// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates settings for a server. This operation is synchronous.
func (c *Client) UpdateServer(ctx context.Context, params *UpdateServerInput, optFns ...func(*Options)) (*UpdateServerOutput, error) {
	stack := middleware.NewStack("UpdateServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateServerMiddlewares(stack)
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
	addOpUpdateServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServer(options.Region), middleware.Before)
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
			OperationName: "UpdateServer",
			Err:           err,
		}
	}
	out := result.(*UpdateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServerInput struct {
	// DDD:HH:MM (weekly start time) or HH:MM (daily start time). Time windows always
	// use coordinated universal time (UTC). Valid strings for day of week (DDD) are:
	// Mon, Tue, Wed, Thr, Fri, Sat, or Sun.
	PreferredBackupWindow *string
	// The name of the server to update.
	ServerName *string
	// Sets the number of automated backups that you want to keep.
	BackupRetentionCount *int32
	// DDD:HH:MM (weekly start time) or HH:MM (daily start time). Time windows always
	// use coordinated universal time (UTC). Valid strings for day of week (DDD) are:
	// Mon, Tue, Wed, Thr, Fri, Sat, or Sun.
	PreferredMaintenanceWindow *string
	// Setting DisableAutomatedBackup to true disables automated or scheduled backups.
	// Automated backups are enabled by default.
	DisableAutomatedBackup *bool
}

type UpdateServerOutput struct {
	// Contains the response to a UpdateServer request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "UpdateServer",
	}
}
