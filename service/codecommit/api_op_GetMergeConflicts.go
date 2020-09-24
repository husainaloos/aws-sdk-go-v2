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

// Returns information about merge conflicts between the before and after commit
// IDs for a pull request in a repository.
func (c *Client) GetMergeConflicts(ctx context.Context, params *GetMergeConflictsInput, optFns ...func(*Options)) (*GetMergeConflictsOutput, error) {
	stack := middleware.NewStack("GetMergeConflicts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetMergeConflictsMiddlewares(stack)
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
	addOpGetMergeConflictsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMergeConflicts(options.Region), middleware.Before)
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
			OperationName: "GetMergeConflicts",
			Err:           err,
		}
	}
	out := result.(*GetMergeConflictsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMergeConflictsInput struct {
	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum
	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	DestinationCommitSpecifier *string
	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	SourceCommitSpecifier *string
	// The maximum number of files to include in the output.
	MaxConflictFiles *int32
	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum
	// The merge option or strategy you want to use to merge the code.
	MergeOption types.MergeOptionTypeEnum
	// The name of the repository where the pull request was created.
	RepositoryName *string
}

type GetMergeConflictsOutput struct {
	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	DestinationCommitId *string
	// The commit ID of the source commit specifier that was used in the merge
	// evaluation.
	SourceCommitId *string
	// A Boolean value that indicates whether the code is mergeable by the specified
	// merge option.
	Mergeable *bool
	// The commit ID of the merge base.
	BaseCommitId *string
	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string
	// A list of metadata for any conflicting files. If the specified merge strategy is
	// FAST_FORWARD_MERGE, this list is always empty.
	ConflictMetadataList []*types.ConflictMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetMergeConflictsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetMergeConflicts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMergeConflicts{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMergeConflicts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetMergeConflicts",
	}
}
