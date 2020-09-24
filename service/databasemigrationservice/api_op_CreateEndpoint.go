// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an endpoint using the provided settings.
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	stack := middleware.NewStack("CreateEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateEndpointMiddlewares(stack)
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
	addOpCreateEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before)
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
			OperationName: "CreateEndpoint",
			Err:           err,
		}
	}
	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateEndpointInput struct {
	// Provides information that defines an Amazon Redshift endpoint.
	RedshiftSettings *types.RedshiftSettings
	// One or more tags to be assigned to the endpoint.
	Tags []*types.Tag
	// The name of the server where the endpoint database resides.
	ServerName *string
	// Additional attributes associated with the connection. Each attribute is
	// specified as a name-value pair associated by an equal sign (=). Multiple
	// attributes are separated by a semicolon (;) with no additional white space. For
	// information on the attributes available for connecting your source or target
	// endpoint, see Working with AWS DMS Endpoints
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the
	// AWS Database Migration Service User Guide.
	ExtraConnectionAttributes *string
	// Settings in JSON format for the source MongoDB endpoint. For more information
	// about the available settings, see Using MongoDB as a Target for AWS Database
	// Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration)
	// in the AWS Database Migration Service User Guide.
	MongoDbSettings *types.MongoDbSettings
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn *string
	// The password to be used to log in to the endpoint database.
	Password *string
	// Settings in JSON format for the source and target SAP ASE endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using SAP ASE as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.ConnectionAttrib)
	// and Extra connection attributes when using SAP ASE as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	SybaseSettings *types.SybaseSettings
	// An AWS KMS key identifier that is used to encrypt the connection parameters for
	// the endpoint. If you don't specify a value for the KmsKeyId parameter, then AWS
	// DMS uses your default encryption key. AWS KMS creates the default encryption key
	// for your AWS account. Your AWS account has a different default encryption key
	// for each AWS Region.
	KmsKeyId *string
	// The name of the endpoint database.
	DatabaseName *string
	// The type of engine for the endpoint. Valid values, depending on the EndpointType
	// value, include "mysql", "oracle", "postgres", "mariadb", "aurora",
	// "aurora-postgresql", "redshift", "s3", "db2", "azuredb", "sybase", "dynamodb",
	// "mongodb", "kinesis", "kafka", "elasticsearch", "documentdb", "sqlserver", and
	// "neptune".
	EngineName *string
	// Settings in JSON format for the target Elasticsearch endpoint. For more
	// information about the available settings, see Extra Connection Attributes When
	// Using Elasticsearch as a Target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration)
	// in the AWS Database Migration Service User Guide.
	ElasticsearchSettings *types.ElasticsearchSettings
	// Settings in JSON format for the target Amazon DynamoDB endpoint. For information
	// about other available settings, see Using Object Mapping to Migrate Data to
	// DynamoDB
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html) in
	// the AWS Database Migration Service User Guide.
	DynamoDbSettings *types.DynamoDbSettings
	// The user name to be used to log in to the endpoint database.
	Username *string
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	// For more information about the available settings, see Using Amazon Kinesis Data
	// Streams as a Target for AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html) in
	// the AWS Database Migration Service User Guide.
	KinesisSettings *types.KinesisSettings
	// The external table definition.
	ExternalTableDefinition *string
	// Settings in JSON format for the source and target MySQL endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using MySQL as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.ConnectionAttrib)
	// and Extra connection attributes when using a MySQL-compatible database as a
	// target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	MySQLSettings *types.MySQLSettings
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default
	// is none
	SslMode types.DmsSslModeValue
	// The type of endpoint. Valid values are source and target.
	EndpointType types.ReplicationEndpointTypeValue
	// The Amazon Resource Name (ARN) for the service access role that you want to use
	// to create the endpoint.
	ServiceAccessRoleArn *string
	// Settings in JSON format for the source and target PostgreSQL endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using PostgreSQL as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.ConnectionAttrib)
	// and  Extra connection attributes when using PostgreSQL as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	PostgreSQLSettings *types.PostgreSQLSettings
	// The database endpoint identifier. Identifiers must begin with a letter and must
	// contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or
	// contain two consecutive hyphens.
	EndpointIdentifier *string
	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	// For information about other available settings, see Extra connection attributes
	// when using SQL Server as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.ConnectionAttrib)
	// and  Extra connection attributes when using SQL Server as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	MicrosoftSQLServerSettings *types.MicrosoftSQLServerSettings
	// Settings in JSON format for the source and target Oracle endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using Oracle as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.ConnectionAttrib)
	// and  Extra connection attributes when using Oracle as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	OracleSettings *types.OracleSettings
	// Settings in JSON format for the target Apache Kafka endpoint. For more
	// information about the available settings, see Using Apache Kafka as a Target for
	// AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html) in the
	// AWS Database Migration Service User Guide.
	KafkaSettings *types.KafkaSettings
	// The port used by the endpoint database.
	Port *int32
	// Settings in JSON format for the target Amazon Neptune endpoint. For more
	// information about the available settings, see Specifying Endpoint Settings for
	// Amazon Neptune as a Target
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings)
	// in the AWS Database Migration Service User Guide.
	NeptuneSettings *types.NeptuneSettings
	// The settings in JSON format for the DMS transfer type of source endpoint.
	// Possible settings include the following:
	//
	//     * ServiceAccessRoleArn - The IAM
	// role that has permission to access the Amazon S3 bucket.
	//
	//     * BucketName - The
	// name of the S3 bucket to use.
	//
	//     * CompressionType - An optional parameter to
	// use GZIP to compress the target files. To use GZIP, set this value to NONE (the
	// default). To keep the files uncompressed, don't use this value.
	//
	// Shorthand
	// syntax for these settings is as follows:
	// ServiceAccessRoleArn=string,BucketName=string,CompressionType=string JSON syntax
	// for these settings is as follows: { "ServiceAccessRoleArn": "string",
	// "BucketName": "string", "CompressionType": "none"|"gzip" }
	DmsTransferSettings *types.DmsTransferSettings
	// Settings in JSON format for the target Amazon S3 endpoint. For more information
	// about the available settings, see Extra Connection Attributes When Using Amazon
	// S3 as a Target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring)
	// in the AWS Database Migration Service User Guide.
	S3Settings *types.S3Settings
	// Settings in JSON format for the source IBM Db2 LUW endpoint. For information
	// about other available settings, see Extra connection attributes when using Db2
	// LUW as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.ConnectionAttrib)
	// in the AWS Database Migration Service User Guide.
	IBMDb2Settings *types.IBMDb2Settings
}

//
type CreateEndpointOutput struct {
	// The endpoint that was created.
	Endpoint *types.Endpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "CreateEndpoint",
	}
}
