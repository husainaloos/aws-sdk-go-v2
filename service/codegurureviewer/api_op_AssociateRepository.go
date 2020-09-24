// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use to associate an AWS CodeCommit repository or a repostory managed by AWS
// CodeStar Connections with Amazon CodeGuru Reviewer. When you associate a
// repository, CodeGuru Reviewer reviews source code changes in the repository's
// pull requests and provides automatic recommendations. You can view
// recommendations using the CodeGuru Reviewer console. For more information, see
// Recommendations in Amazon CodeGuru Reviewer
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/recommendations.html)
// in the Amazon CodeGuru Reviewer User Guide.  <p>If you associate a CodeCommit
// repository, it must be in the same AWS Region and AWS account where its CodeGuru
// Reviewer code reviews are configured.</p> <p> Bitbucket and GitHub Enterprise
// Server repositories are managed by AWS CodeStar Connections to connect to
// CodeGuru Reviewer. For more information, see <a
// href="https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/reviewer-ug/step-one.html#select-repository-source-provider">Connect
// to a repository source provider</a> in the <i>Amazon CodeGuru Reviewer User
// Guide.</i> </p> <note> <p> You cannot use the CodeGuru Reviewer SDK or the AWS
// CLI to associate a GitHub repository with Amazon CodeGuru Reviewer. To associate
// a GitHub repository, use the console. For more information, see <a
// href="https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/getting-started-with-guru.html">Getting
// started with CodeGuru Reviewer</a> in the <i>CodeGuru Reviewer User Guide.</i>
// </p> </note>
func (c *Client) AssociateRepository(ctx context.Context, params *AssociateRepositoryInput, optFns ...func(*Options)) (*AssociateRepositoryOutput, error) {
	stack := middleware.NewStack("AssociateRepository", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociateRepositoryMiddlewares(stack)
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
	addIdempotencyToken_opAssociateRepositoryMiddleware(stack, options)
	addOpAssociateRepositoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateRepository(options.Region), middleware.Before)
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
			OperationName: "AssociateRepository",
			Err:           err,
		}
	}
	out := result.(*AssociateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateRepositoryInput struct {
	// The repository to associate.
	Repository *types.Repository
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. To add a new repository association, this parameter specifies a
	// unique identifier for the new repository association that helps ensure
	// idempotency.  <p>If you use the AWS CLI or one of the AWS SDKs to call this
	// operation, you can leave this parameter empty. The CLI or SDK generates a random
	// UUID for you and includes that in the request. If you don't use the SDK and
	// instead generate a raw HTTP request to the Secrets Manager service endpoint, you
	// must generate a ClientRequestToken yourself for new versions and include that
	// value in the request.</p> <p>You typically interact with this value if you
	// implement your own retry logic and want to ensure that a given repository
	// association is not created twice. We recommend that you generate a UUID-type
	// value to ensure uniqueness within the specified repository association.</p>
	// <p>Amazon CodeGuru Reviewer uses this value to prevent the accidental creation
	// of duplicate repository associations if there are failures and retries. </p>
	ClientRequestToken *string
}

type AssociateRepositoryOutput struct {
	// Information about the repository association.
	RepositoryAssociation *types.RepositoryAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociateRepositoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociateRepository{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateRepository{}, middleware.After)
}

type idempotencyToken_initializeOpAssociateRepository struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateRepository) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateRepository) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateRepositoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateRepositoryInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opAssociateRepositoryMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpAssociateRepository{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateRepository(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "AssociateRepository",
	}
}
