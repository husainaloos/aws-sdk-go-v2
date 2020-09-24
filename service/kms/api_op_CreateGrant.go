// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a grant to a customer master key (CMK). The grant allows the grantee
// principal to use the CMK when the conditions specified in the grant are met.
// When setting permissions, grants are an alternative to key policies. To create a
// grant that allows a cryptographic operation
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
// only when the request includes a particular encryption context
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context),
// use the Constraints parameter. For details, see GrantConstraints (). You can
// create grants on symmetric and asymmetric CMKs. However, if the grant allows an
// operation that the CMK does not support, CreateGrant fails with a
// ValidationException.  <ul> <li> <p>Grants for symmetric CMKs cannot allow
// operations that are not supported for symmetric CMKs, including <a>Sign</a>,
// <a>Verify</a>, and <a>GetPublicKey</a>. (There are limited exceptions to this
// rule for legacy operations, but you should not create a grant for an operation
// that AWS KMS does not support.)</p> </li> <li> <p>Grants for asymmetric CMKs
// cannot allow operations that are not supported for asymmetric CMKs, including
// operations that <a
// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey">generate
// data keys</a> or <a
// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKeyPair">data
// key pairs</a>, or operations related to <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html">automatic
// key rotation</a>, <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html">imported
// key material</a>, or CMKs in <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html">custom
// key stores</a>.</p> </li> <li> <p>Grants for asymmetric CMKs with a
// <code>KeyUsage</code> of <code>ENCRYPT_DECRYPT</code> cannot allow the
// <a>Sign</a> or <a>Verify</a> operations. Grants for asymmetric CMKs with a
// <code>KeyUsage</code> of <code>SIGN_VERIFY</code> cannot allow the
// <a>Encrypt</a> or <a>Decrypt</a> operations.</p> </li> <li> <p>Grants for
// asymmetric CMKs cannot include an encryption context grant constraint. An
// encryption context is not supported on asymmetric CMKs.</p> </li> </ul> <p>For
// information about symmetric and asymmetric CMKs, see <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html">Using
// Symmetric and Asymmetric CMKs</a> in the <i>AWS Key Management Service Developer
// Guide</i>.</p> <p>To perform this operation on a CMK in a different AWS account,
// specify the key  ARN in the value of the KeyId parameter. For more information
// about grants, see Grants
// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html) in the AWS
// Key Management Service Developer Guide . The CMK that you use for this operation
// must be in a compatible key state. For details, see How Key State Affects Use of
// a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) CreateGrant(ctx context.Context, params *CreateGrantInput, optFns ...func(*Options)) (*CreateGrantOutput, error) {
	stack := middleware.NewStack("CreateGrant", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateGrantMiddlewares(stack)
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
	addOpCreateGrantValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGrant(options.Region), middleware.Before)
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
			OperationName: "CreateGrant",
			Err:           err,
		}
	}
	out := result.(*CreateGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGrantInput struct {
	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string
	// A list of operations that the grant permits.
	Operations []types.GrantOperation
	// The principal that is given permission to perform the operations that the grant
	// permits. To specify the principal, use the Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// an AWS principal. Valid AWS principals include AWS accounts (root), IAM users,
	// IAM roles, federated users, and assumed role users. For examples of the ARN
	// syntax to use for specifying a principal, see AWS Identity and Access Management
	// (IAM)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the AWS General Reference.
	GranteePrincipal *string
	// The principal that is given permission to retire the grant by using RetireGrant
	// () operation. To specify the principal, use the Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// an AWS principal. Valid AWS principals include AWS accounts (root), IAM users,
	// federated users, and assumed role users. For examples of the ARN syntax to use
	// for specifying a principal, see AWS Identity and Access Management (IAM)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the AWS General Reference.
	RetiringPrincipal *string
	// A friendly name for identifying the grant. Use this value to prevent the
	// unintended creation of duplicate grants when retrying this request. When this
	// value is absent, all CreateGrant requests result in a new grant with a unique
	// GrantId even if all the supplied parameters are identical. This can result in
	// unintended duplicates when you retry the CreateGrant request. When this value is
	// present, you can retry a CreateGrant request with identical parameters; if the
	// grant already exists, the original GrantId is returned without creating a new
	// grant. Note that the returned grant token is unique with every CreateGrant
	// request, even when a duplicate GrantId is returned. All grant tokens obtained in
	// this way can be used interchangeably.
	Name *string
	// The unique identifier for the customer master key (CMK) that the grant applies
	// to.  <p>Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To
	// specify a CMK in a  different AWS account, you must use the key ARN. For
	// example:
	//
	//     * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys () or DescribeKey ().
	KeyId *string
	// Allows a cryptographic operation
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// only when the encryption context matches or includes the encryption context
	// specified in this structure. For more information about encryption context, see
	// Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide .
	Constraints *types.GrantConstraints
}

type CreateGrantOutput struct {
	// The unique identifier for the grant. You can use the GrantId in a subsequent
	// RetireGrant () or RevokeGrant () operation.
	GrantId *string
	// The grant token. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateGrantMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGrant{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGrant{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGrant(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "CreateGrant",
	}
}
