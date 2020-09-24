// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new DB instance.
func (c *Client) CreateDBInstance(ctx context.Context, params *CreateDBInstanceInput, optFns ...func(*Options)) (*CreateDBInstanceOutput, error) {
	stack := middleware.NewStack("CreateDBInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBInstanceMiddlewares(stack)
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
	addOpCreateDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBInstance(options.Region), middleware.Before)
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
			OperationName: "CreateDBInstance",
			Err:           err,
		}
	}
	out := result.(*CreateDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBInstanceInput struct {
	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string
	// A value that specifies the order in which an Read Replica is promoted to the
	// primary instance after a failure of the existing primary instance. Default: 1
	// Valid Values: 0 - 15
	PromotionTier *int32
	// The password for the master user. The password can include any printable ASCII
	// character except "/", """, or "@". Not used.
	MasterUserPassword *string
	// Specify the Active Directory Domain to create the instance in.
	Domain *string
	// The ARN for the IAM role that permits Neptune to send enhanced monitoring
	// metrics to Amazon CloudWatch Logs. For example,
	// arn:aws:iam:123456789012:role/emaccess. If MonitoringInterval is set to a value
	// other than 0, then you must supply a MonitoringRoleArn value.
	MonitoringRoleArn *string
	// The tags to assign to the new instance.
	Tags []*types.Tag
	// True to enable AWS Identity and Access Management (IAM) authentication for
	// Neptune. Default: false
	EnableIAMDatabaseAuthentication *bool
	// Specifies if the DB instance is a Multi-AZ deployment. You can't set the
	// AvailabilityZone parameter if the MultiAZ parameter is set to true.
	MultiAZ *bool
	// The ARN from the key store with which to associate the instance for TDE
	// encryption.
	TdeCredentialArn *string
	// (Not supported by Neptune)
	PerformanceInsightsKMSKeyId *string
	// This flag should no longer be used.
	PubliclyAccessible *bool
	// The identifier of the DB cluster that the instance will belong to. For
	// information on creating a DB cluster, see CreateDBCluster (). Type: String
	DBClusterIdentifier *string
	// The number of days for which automated backups are retained. Not applicable. The
	// retention period for automated backups is managed by the DB cluster. For more
	// information, see CreateDBCluster (). Default: 1 Constraints:
	//
	//     * Must be a
	// value from 0 to 35
	//
	//     * Cannot be set to 0 if the DB instance is a source to
	// Read Replicas
	BackupRetentionPeriod *int32
	// The port number on which the database accepts connections. Not applicable. The
	// port is managed by the DB cluster. For more information, see CreateDBCluster ().
	// Default: 8182 Type: Integer
	Port *int32
	// (Not supported by Neptune)
	CharacterSetName *string
	// A list of EC2 VPC security groups to associate with this DB instance. Not
	// applicable. The associated list of EC2 VPC security groups is managed by the DB
	// cluster. For more information, see CreateDBCluster (). Default: The default EC2
	// VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []*string
	// The version number of the database engine to use. Currently, setting this
	// parameter has no effect.
	EngineVersion *string
	// Specifies whether the DB instance is encrypted. Not applicable. The encryption
	// for DB instances is managed by the DB cluster. For more information, see
	// CreateDBCluster (). Default: false
	StorageEncrypted *bool
	// Specifies the storage type to be associated with the DB instance. Not
	// applicable. Storage is managed by the DB Cluster.
	StorageType *string
	// The DB instance identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	//     * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//
	// * First character must be a letter.
	//
	//     * Cannot end with a hyphen or contain
	// two consecutive hyphens.
	//
	// Example: mydbinstance
	DBInstanceIdentifier *string
	// (Not supported by Neptune)
	OptionGroupName *string
	// The amount of storage (in gibibytes) to allocate for the DB instance. Type:
	// Integer Not applicable. Neptune cluster volumes automatically grow as the amount
	// of data in your database increases, though you are only charged for the space
	// that you use in a Neptune cluster volume.
	AllocatedStorage *int32
	// (Not supported by Neptune)
	EnablePerformanceInsights *bool
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	EnableCloudwatchLogsExports []*string
	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance.
	Iops *int32
	// The compute and memory capacity of the DB instance, for example, db.m4.large.
	// Not all DB instance classes are available in all AWS Regions.
	DBInstanceClass *string
	// Indicates that minor engine upgrades are applied automatically to the DB
	// instance during the maintenance window. Default: true
	AutoMinorVersionUpgrade *bool
	// A list of DB security groups to associate with this DB instance. Default: The
	// default DB security group for the database engine.
	DBSecurityGroups []*string
	// The name of the database engine to be used for this instance. Valid Values:
	// neptune
	Engine *string
	// True to copy all tags from the DB instance to snapshots of the DB instance, and
	// otherwise false. The default is false.
	CopyTagsToSnapshot *bool
	// The daily time range during which automated backups are created. Not applicable.
	// The daily time range for creating automated backups is managed by the DB
	// cluster. For more information, see CreateDBCluster ().
	PreferredBackupWindow *string
	// The time range each week during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. Valid Days: Mon, Tue, Wed, Thu,
	// Fri, Sat, Sun. Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string
	// Not supported.
	DBName *string
	// The name for the master user. Not used.
	MasterUsername *string
	// The AWS KMS key identifier for an encrypted DB instance. The KMS key identifier
	// is the Amazon Resource Name (ARN) for the KMS encryption key. If you are
	// creating a DB instance with the same AWS account that owns the KMS encryption
	// key used to encrypt the new DB instance, then you can use the KMS key alias
	// instead of the ARN for the KM encryption key. Not applicable. The KMS key
	// identifier is managed by the DB cluster. For more information, see
	// CreateDBCluster (). If the StorageEncrypted parameter is true, and you do not
	// specify a value for the KmsKeyId parameter, then Amazon Neptune will use your
	// default encryption key. AWS KMS creates the default encryption key for your AWS
	// account. Your AWS account has a different default encryption key for each AWS
	// Region.
	KmsKeyId *string
	// The EC2 Availability Zone that the DB instance is created in Default: A random,
	// system-chosen Availability Zone in the endpoint's AWS Region. Example:
	// us-east-1d Constraint: The AvailabilityZone parameter can't be specified if the
	// MultiAZ parameter is set to true. The specified Availability Zone must be in the
	// same AWS Region as the current endpoint.
	AvailabilityZone *string
	// License model information for this DB instance. Valid values: license-included |
	// bring-your-own-license | general-public-license
	LicenseModel *string
	// The time zone of the DB instance.
	Timezone *string
	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0. If MonitoringRoleArn is specified, then
	// you must also set MonitoringInterval to a value other than 0. Valid Values: 0,
	// 1, 5, 10, 15, 30, 60
	MonitoringInterval *int32
	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. See Deleting a DB Instance
	// (https://docs.aws.amazon.com/neptune/latest/userguide/manage-console-instances-delete.html).
	// <p>DB instances in a DB cluster can be deleted even when deletion protection is
	// enabled in their parent DB cluster.</p>
	DeletionProtection *bool
	// A DB subnet group to associate with this DB instance. If there is no DB subnet
	// group, then it is a non-VPC DB instance.
	DBSubnetGroupName *string
	// The password for the given ARN from the key store in order to access the device.
	TdeCredentialPassword *string
	// The name of the DB parameter group to associate with this DB instance. If this
	// argument is omitted, the default DBParameterGroup for the specified engine is
	// used. Constraints:
	//
	//     * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//     *
	// First character must be a letter
	//
	//     * Cannot end with a hyphen or contain two
	// consecutive hyphens
	DBParameterGroupName *string
}

type CreateDBInstanceOutput struct {
	// Contains the details of an Amazon Neptune DB instance. This data type is used as
	// a response element in the DescribeDBInstances () action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBInstance",
	}
}
