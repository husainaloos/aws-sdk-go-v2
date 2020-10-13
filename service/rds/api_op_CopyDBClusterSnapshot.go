// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/query"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"net/http"
)

// Copies a snapshot of a DB cluster. To copy a DB cluster snapshot from a shared
// manual DB cluster snapshot, SourceDBClusterSnapshotIdentifier must be the Amazon
// Resource Name (ARN) of the shared DB cluster snapshot.  <p>You can copy an
// encrypted DB cluster snapshot from another AWS Region. In that case, the AWS
// Region where you call the <code>CopyDBClusterSnapshot</code> action is the
// destination AWS Region for the encrypted DB cluster snapshot to be copied to. To
// copy an encrypted DB cluster snapshot from another AWS Region, you must provide
// the following values:</p> <ul> <li> <p> <code>KmsKeyId</code> - The AWS Key
// Management System (AWS KMS) key identifier for the key to use to encrypt the
// copy of the DB cluster snapshot in the destination AWS Region.</p> </li> <li>
// <p> <code>PreSignedUrl</code> - A URL that contains a Signature Version 4 signed
// request for the <code>CopyDBClusterSnapshot</code> action to be called in the
// source AWS Region where the DB cluster snapshot is copied from. The pre-signed
// URL must be a valid request for the <code>CopyDBClusterSnapshot</code> API
// action that can be executed in the source AWS Region that contains the encrypted
// DB cluster snapshot to be copied.</p> <p>The pre-signed URL request must contain
// the following parameter values:</p> <ul> <li> <p> <code>KmsKeyId</code> - The
// KMS key identifier for the key to use to encrypt the copy of the DB cluster
// snapshot in the destination AWS Region. This is the same identifier for both the
// <code>CopyDBClusterSnapshot</code> action that is called in the destination AWS
// Region, and the action contained in the pre-signed URL.</p> </li> <li> <p>
// <code>DestinationRegion</code> - The name of the AWS Region that the DB cluster
// snapshot is to be created in.</p> </li> <li> <p>
// <code>SourceDBClusterSnapshotIdentifier</code> - The DB cluster snapshot
// identifier for the encrypted DB cluster snapshot to be copied. This identifier
// must be in the Amazon Resource Name (ARN) format for the source AWS Region. For
// example, if you are copying an encrypted DB cluster snapshot from the us-west-2
// AWS Region, then your <code>SourceDBClusterSnapshotIdentifier</code> looks like
// the following example:
// <code>arn:aws:rds:us-west-2:123456789012:cluster-snapshot:aurora-cluster1-snapshot-20161115</code>.</p>
// </li> </ul> <p>To learn how to generate a Signature Version 4 signed request,
// see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html">
// Authenticating Requests: Using Query Parameters (AWS Signature Version 4)</a>
// and <a
// href="https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html">
// Signature Version 4 Signing Process</a>.</p> <note> <p>If you are using an AWS
// SDK tool or the AWS CLI, you can specify <code>SourceRegion</code> (or
// <code>--source-region</code> for the AWS CLI) instead of specifying
// <code>PreSignedUrl</code> manually. Specifying <code>SourceRegion</code>
// autogenerates a pre-signed URL that is a valid request for the operation that
// can be executed in the source AWS Region.</p> </note> </li> <li> <p>
// <code>TargetDBClusterSnapshotIdentifier</code> - The identifier for the new copy
// of the DB cluster snapshot in the destination AWS Region.</p> </li> <li> <p>
// <code>SourceDBClusterSnapshotIdentifier</code> - The DB cluster snapshot
// identifier for the encrypted DB cluster snapshot to be copied. This identifier
// must be in the ARN format for the source AWS Region and is the same value as the
// <code>SourceDBClusterSnapshotIdentifier</code> in the pre-signed URL. </p> </li>
// </ul> <p>To cancel the copy operation once it is in progress, delete the target
// DB cluster snapshot identified by <code>TargetDBClusterSnapshotIdentifier</code>
// while that DB cluster snapshot is in "copying" status.</p> <p>For more
// information on copying encrypted DB cluster snapshots from one AWS Region to
// another, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CopySnapshot.html">
// Copying a Snapshot</a> in the <i>Amazon Aurora User Guide.</i> </p> <p>For more
// information on Amazon Aurora, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html">
// What Is Amazon Aurora?</a> in the <i>Amazon Aurora User Guide.</i> </p> <note>
// <p>This action only applies to Aurora DB clusters.</p> </note>
func (c *Client) CopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*Options)) (*CopyDBClusterSnapshotOutput, error) {
	if params == nil {
		params = &CopyDBClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBClusterSnapshot", params, optFns, addOperationCopyDBClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CopyDBClusterSnapshotInput struct {

	// The identifier of the DB cluster snapshot to copy. This parameter isn't
	// case-sensitive. You can't copy an encrypted, shared DB cluster snapshot from one
	// AWS Region to another. Constraints:
	//
	//     * Must specify a valid system snapshot
	// in the "available" state.
	//
	//     * If the source snapshot is in the same AWS
	// Region as the copy, specify a valid DB snapshot identifier.
	//
	//     * If the source
	// snapshot is in a different AWS Region than the copy, specify a valid DB cluster
	// snapshot ARN. For more information, go to  Copying Snapshots Across AWS Regions
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CopySnapshot.html#USER_CopySnapshot.AcrossRegions)
	// in the Amazon Aurora User Guide.
	//
	// Example: my-cluster-snapshot1
	//
	// This member is required.
	SourceDBClusterSnapshotIdentifier *string

	// The identifier of the new DB cluster snapshot to create from the source DB
	// cluster snapshot. This parameter isn't case-sensitive. Constraints:
	//
	//     * Must
	// contain from 1 to 63 letters, numbers, or hyphens.
	//
	//     * First character must
	// be a letter.
	//
	//     * Can't end with a hyphen or contain two consecutive
	// hyphens.
	//
	// Example: my-cluster-snapshot2
	//
	// This member is required.
	TargetDBClusterSnapshotIdentifier *string

	// A value that indicates whether to copy all tags from the source DB cluster
	// snapshot to the target DB cluster snapshot. By default, tags are not copied.
	CopyTags *bool

	// The AWS KMS key ID for an encrypted DB cluster snapshot. The KMS key ID is the
	// Amazon Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS
	// encryption key.  <p>If you copy an encrypted DB cluster snapshot from your AWS
	// account, you can specify a value for <code>KmsKeyId</code> to encrypt the copy
	// with a new KMS encryption key. If you don't specify a value for
	// <code>KmsKeyId</code>, then the copy of the DB cluster snapshot is encrypted
	// with the same KMS key as the source DB cluster snapshot. </p> <p>If you copy an
	// encrypted DB cluster snapshot that is shared from another AWS account, then you
	// must specify a value for <code>KmsKeyId</code>. </p> <p>To copy an encrypted DB
	// cluster snapshot to another AWS Region, you must set <code>KmsKeyId</code> to
	// the KMS key ID you want to use to encrypt the copy of the DB cluster snapshot in
	// the destination AWS Region. KMS encryption keys are specific to the AWS Region
	// that they are created in, and you can't use encryption keys from one AWS Region
	// in another AWS Region.</p> <p>If you copy an unencrypted DB cluster snapshot and
	// specify a value for the <code>KmsKeyId</code> parameter, an error is
	// returned.</p>
	KmsKeyId *string

	// The URL that contains a Signature Version 4 signed request for the
	// CopyDBClusterSnapshot API action in the AWS Region that contains the source DB
	// cluster snapshot to copy. The PreSignedUrl parameter must be used when copying
	// an encrypted DB cluster snapshot from another AWS Region. Don't specify
	// PreSignedUrl when you are copying an encrypted DB cluster snapshot in the same
	// AWS Region. The pre-signed URL must be a valid request for the
	// CopyDBClusterSnapshot API action that can be executed in the source AWS Region
	// that contains the encrypted DB cluster snapshot to be copied. The pre-signed URL
	// request must contain the following parameter values:  <ul> <li> <p>
	// <code>KmsKeyId</code> - The AWS KMS key identifier for the key to use to encrypt
	// the copy of the DB cluster snapshot in the destination AWS Region. This is the
	// same identifier for both the <code>CopyDBClusterSnapshot</code> action that is
	// called in the destination AWS Region, and the action contained in the pre-signed
	// URL.</p> </li> <li> <p> <code>DestinationRegion</code> - The name of the AWS
	// Region that the DB cluster snapshot is to be created in.</p> </li> <li> <p>
	// <code>SourceDBClusterSnapshotIdentifier</code> - The DB cluster snapshot
	// identifier for the encrypted DB cluster snapshot to be copied. This identifier
	// must be in the Amazon Resource Name (ARN) format for the source AWS Region. For
	// example, if you are copying an encrypted DB cluster snapshot from the us-west-2
	// AWS Region, then your <code>SourceDBClusterSnapshotIdentifier</code> looks like
	// the following example:
	// <code>arn:aws:rds:us-west-2:123456789012:cluster-snapshot:aurora-cluster1-snapshot-20161115</code>.</p>
	// </li> </ul> <p>To learn how to generate a Signature Version 4 signed request,
	// see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html">
	// Authenticating Requests: Using Query Parameters (AWS Signature Version 4)</a>
	// and <a
	// href="https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html">
	// Signature Version 4 Signing Process</a>.</p> <note> <p>If you are using an AWS
	// SDK tool or the AWS CLI, you can specify <code>SourceRegion</code> (or
	// <code>--source-region</code> for the AWS CLI) instead of specifying
	// <code>PreSignedUrl</code> manually. Specifying <code>SourceRegion</code>
	// autogenerates a pre-signed URL that is a valid request for the operation that
	// can be executed in the source AWS Region.</p> </note>
	PreSignedUrl *string

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag

	// Used by the SDK's PresignURL autofill customization to specify the region the of
	// the client's request.
	destinationRegion *string
}

type CopyDBClusterSnapshotOutput struct {

	// Contains the details for an Amazon RDS DB cluster snapshot This data type is
	// used as a response element in the DescribeDBClusterSnapshots action.
	DBClusterSnapshot *types.DBClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyDBClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addCopyDBClusterSnapshotPresignURLMiddleware(stack, options)
	addOpCopyDBClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBClusterSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func copyCopyDBClusterSnapshotInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCopyDBClusterSnapshotPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCopyDBClusterSnapshotSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCopyDBClusterSnapshotPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCopyDBClusterSnapshotdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}

type copyDBClusterSnapshotHTTPPresignURLClient struct {
	client    *Client
	presigner *v4.Signer
}

func newCopyDBClusterSnapshotHTTPPresignURLClient(options Options, optFns ...func(*Options)) *copyDBClusterSnapshotHTTPPresignURLClient {
	return &copyDBClusterSnapshotHTTPPresignURLClient{
		client:    New(options, optFns...),
		presigner: v4.NewSigner(),
	}
}
func (c *copyDBClusterSnapshotHTTPPresignURLClient) PresignCopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*Options)) (string, http.Header, error) {
	if params == nil {
		params = &CopyDBClusterSnapshotInput{}
	}

	optFns = append(optFns, func(o *Options) {
		o.HTTPClient = &smithyhttp.NopClient{}
	})

	ctx = presignedurlcust.WithIsPresigning(ctx)
	result, _, err := c.client.invokeOperation(ctx, "CopyDBClusterSnapshot", params, optFns,
		addOperationCopyDBClusterSnapshotMiddlewares,
		c.convertToPresignMiddleware,
	)
	if err != nil {
		return ``, nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out.URL, out.SignedHeader, nil
}
func (c *copyDBClusterSnapshotHTTPPresignURLClient) convertToPresignMiddleware(stack *middleware.Stack, options Options) (err error) {
	stack.Finalize.Clear()
	stack.Deserialize.Clear()
	stack.Build.Remove(awsmiddleware.RequestInvocationIDMiddleware{}.ID())
	err = stack.Finalize.Add(v4.NewPresignHTTPRequestMiddleware(options.Credentials, c.presigner), middleware.After)
	if err != nil {
		return err
	}
	err = query.AddAsGetRequestMiddleware(stack)
	if err != nil {
		return err
	}
	return nil
}
func addCopyDBClusterSnapshotPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL:      getCopyDBClusterSnapshotPreSignedUrl,
			GetSourceRegion:      getCopyDBClusterSnapshotSourceRegion,
			CopyInput:            copyCopyDBClusterSnapshotInputForPresign,
			SetDestinationRegion: setCopyDBClusterSnapshotdestinationRegion,
			SetPresignedURL:      setCopyDBClusterSnapshotPreSignedUrl,
		},
		Presigner: &presignAutoFillCopyDBClusterSnapshotClient{client: newCopyDBClusterSnapshotHTTPPresignURLClient(options)},
	})
}

type presignAutoFillCopyDBClusterSnapshotClient struct {
	client *copyDBClusterSnapshotHTTPPresignURLClient
}

func (c *presignAutoFillCopyDBClusterSnapshotClient) PresignURL(ctx context.Context, region string, params interface{}) (string, http.Header, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return ``, nil, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = region
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	return c.client.PresignCopyDBClusterSnapshot(ctx, input, optFn)
}

func newServiceMetadataMiddleware_opCopyDBClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBClusterSnapshot",
	}
}
