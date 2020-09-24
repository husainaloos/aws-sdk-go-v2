// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a data source.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	stack := middleware.NewStack("CreateDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDataSourceMiddlewares(stack)
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
	addOpCreateDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before)
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
			OperationName: "CreateDataSource",
			Err:           err,
		}
	}
	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {
	// Secure Socket Layer (SSL) properties that apply when QuickSight connects to your
	// underlying source.
	SslProperties *types.SslProperties
	// A display name for the data source.
	Name *string
	// A list of resource permissions on the data source.
	Permissions []*types.ResourcePermission
	// The parameters that QuickSight uses to connect to your underlying source.
	DataSourceParameters *types.DataSourceParameters
	// The AWS account ID.
	AwsAccountId *string
	// The credentials QuickSight that uses to connect to your underlying source.
	// Currently, only credentials based on user name and password are supported.
	Credentials *types.DataSourceCredentials
	// An ID for the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string
	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the data source.
	Tags []*types.Tag
	// The type of the data source. Currently, the supported types for this operation
	// are: ATHENA, AURORA, AURORA_POSTGRESQL, MARIADB, MYSQL, POSTGRESQL, PRESTO,
	// REDSHIFT, S3, SNOWFLAKE, SPARK, SQLSERVER, TERADATA. Use ListDataSources to
	// return a list of all data sources.
	Type types.DataSourceType
	// Use this parameter only when you want QuickSight to use a VPC connection when
	// connecting to your underlying source.
	VpcConnectionProperties *types.VpcConnectionProperties
}

type CreateDataSourceOutput struct {
	// The AWS request ID for this operation.
	RequestId *string
	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string
	// The Amazon Resource Name (ARN) of the data source.
	Arn *string
	// The status of creating the data source.
	CreationStatus types.ResourceStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateDataSource",
	}
}
