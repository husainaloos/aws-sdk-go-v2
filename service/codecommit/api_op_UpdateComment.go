// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the contents of a comment.
func (c *Client) UpdateComment(ctx context.Context, params *UpdateCommentInput, optFns ...func(*Options)) (*UpdateCommentOutput, error) {
	stack := middleware.NewStack("UpdateComment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateCommentMiddlewares(stack)
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
	addOpUpdateCommentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComment(options.Region), middleware.Before)

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
			OperationName: "UpdateComment",
			Err:           err,
		}
	}
	out := result.(*UpdateCommentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCommentInput struct {
	// The updated content to replace the existing content of the comment.
	Content *string
	// The system-generated ID of the comment you want to update. To get this ID, use
	// GetCommentsForComparedCommit () or GetCommentsForPullRequest ().
	CommentId *string
}

type UpdateCommentOutput struct {
	// Information about the updated comment.
	Comment *types.Comment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateCommentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateComment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateComment{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateComment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdateComment",
	}
}