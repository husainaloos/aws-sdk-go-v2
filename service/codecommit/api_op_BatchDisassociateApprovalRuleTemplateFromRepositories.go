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

// Removes the association between an approval rule template and one or more
// specified repositories.
func (c *Client) BatchDisassociateApprovalRuleTemplateFromRepositories(ctx context.Context, params *BatchDisassociateApprovalRuleTemplateFromRepositoriesInput, optFns ...func(*Options)) (*BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
	stack := middleware.NewStack("BatchDisassociateApprovalRuleTemplateFromRepositories", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchDisassociateApprovalRuleTemplateFromRepositoriesMiddlewares(stack)
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
	addOpBatchDisassociateApprovalRuleTemplateFromRepositoriesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateApprovalRuleTemplateFromRepositories(options.Region), middleware.Before)
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
			OperationName: "BatchDisassociateApprovalRuleTemplateFromRepositories",
			Err:           err,
		}
	}
	out := result.(*BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateApprovalRuleTemplateFromRepositoriesInput struct {
	// The repository names that you want to disassociate from the approval rule
	// template. The length constraint limit is for each string in the array. The array
	// itself can be empty.
	RepositoryNames []*string
	// The name of the template that you want to disassociate from one or more
	// repositories.
	ApprovalRuleTemplateName *string
}

type BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput struct {
	// A list of repository names that have had their association with the template
	// removed.
	DisassociatedRepositoryNames []*string
	// A list of any errors that might have occurred while attempting to remove the
	// association between the template and the repositories.
	Errors []*types.BatchDisassociateApprovalRuleTemplateFromRepositoriesError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchDisassociateApprovalRuleTemplateFromRepositoriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDisassociateApprovalRuleTemplateFromRepositories{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDisassociateApprovalRuleTemplateFromRepositories{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDisassociateApprovalRuleTemplateFromRepositories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "BatchDisassociateApprovalRuleTemplateFromRepositories",
	}
}
