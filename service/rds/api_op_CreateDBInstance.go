// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new DB instance.
func (c *Client) CreateDBInstance(ctx context.Context, params *CreateDBInstanceInput, optFns ...func(*Options)) (*CreateDBInstanceOutput, error) {
	if params == nil {
		params = &CreateDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBInstance", params, optFns, addOperationCreateDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBInstanceInput struct {

	// The compute and memory capacity of the DB instance, for example, db.m4.large.
	// Not all DB instance classes are available in all AWS Regions, or for all
	// database engines. For the full list of DB instance classes, and availability for
	// your engine, see DB Instance Class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide.
	//
	// This member is required.
	DBInstanceClass *string

	// The DB instance identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	//     * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//
	// * First character must be a letter.
	//
	//     * Can't end with a hyphen or contain
	// two consecutive hyphens.
	//
	// Example: mydbinstance
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name of the database engine to be used for this instance.  <p>Not every
	// database engine is available for every AWS Region. </p> <p>Valid Values: </p>
	// <ul> <li> <p> <code>aurora</code> (for MySQL 5.6-compatible Aurora)</p> </li>
	// <li> <p> <code>aurora-mysql</code> (for MySQL 5.7-compatible Aurora)</p> </li>
	// <li> <p> <code>aurora-postgresql</code> </p> </li> <li> <p> <code>mariadb</code>
	// </p> </li> <li> <p> <code>mysql</code> </p> </li> <li> <p>
	// <code>oracle-ee</code> </p> </li> <li> <p> <code>oracle-se2</code> </p> </li>
	// <li> <p> <code>oracle-se1</code> </p> </li> <li> <p> <code>oracle-se</code> </p>
	// </li> <li> <p> <code>postgres</code> </p> </li> <li> <p>
	// <code>sqlserver-ee</code> </p> </li> <li> <p> <code>sqlserver-se</code> </p>
	// </li> <li> <p> <code>sqlserver-ex</code> </p> </li> <li> <p>
	// <code>sqlserver-web</code> </p> </li> </ul>
	//
	// This member is required.
	Engine *string

	// The amount of storage (in gibibytes) to allocate for the DB instance. Type:
	// Integer Amazon Aurora Not applicable. Aurora cluster volumes automatically grow
	// as the amount of data in your database increases, though you are only charged
	// for the space that you use in an Aurora cluster volume.  <p> <b>MySQL</b> </p>
	// <p>Constraints to the amount of storage for each storage type are the following:
	// </p> <ul> <li> <p>General Purpose (SSD) storage (gp2): Must be an integer from
	// 20 to 65536.</p> </li> <li> <p>Provisioned IOPS storage (io1): Must be an
	// integer from 100 to 65536.</p> </li> <li> <p>Magnetic storage (standard): Must
	// be an integer from 5 to 3072.</p> </li> </ul> <p> <b>MariaDB</b> </p>
	// <p>Constraints to the amount of storage for each storage type are the following:
	// </p> <ul> <li> <p>General Purpose (SSD) storage (gp2): Must be an integer from
	// 20 to 65536.</p> </li> <li> <p>Provisioned IOPS storage (io1): Must be an
	// integer from 100 to 65536.</p> </li> <li> <p>Magnetic storage (standard): Must
	// be an integer from 5 to 3072.</p> </li> </ul> <p> <b>PostgreSQL</b> </p>
	// <p>Constraints to the amount of storage for each storage type are the following:
	// </p> <ul> <li> <p>General Purpose (SSD) storage (gp2): Must be an integer from
	// 20 to 65536.</p> </li> <li> <p>Provisioned IOPS storage (io1): Must be an
	// integer from 100 to 65536.</p> </li> <li> <p>Magnetic storage (standard): Must
	// be an integer from 5 to 3072.</p> </li> </ul> <p> <b>Oracle</b> </p>
	// <p>Constraints to the amount of storage for each storage type are the following:
	// </p> <ul> <li> <p>General Purpose (SSD) storage (gp2): Must be an integer from
	// 20 to 65536.</p> </li> <li> <p>Provisioned IOPS storage (io1): Must be an
	// integer from 100 to 65536.</p> </li> <li> <p>Magnetic storage (standard): Must
	// be an integer from 10 to 3072.</p> </li> </ul> <p> <b>SQL Server</b> </p>
	// <p>Constraints to the amount of storage for each storage type are the following:
	// </p> <ul> <li> <p>General Purpose (SSD) storage (gp2):</p> <ul> <li>
	// <p>Enterprise and Standard editions: Must be an integer from 200 to 16384.</p>
	// </li> <li> <p>Web and Express editions: Must be an integer from 20 to 16384.</p>
	// </li> </ul> </li> <li> <p>Provisioned IOPS storage (io1):</p> <ul> <li>
	// <p>Enterprise and Standard editions: Must be an integer from 200 to 16384.</p>
	// </li> <li> <p>Web and Express editions: Must be an integer from 100 to
	// 16384.</p> </li> </ul> </li> <li> <p>Magnetic storage (standard):</p> <ul> <li>
	// <p>Enterprise and Standard editions: Must be an integer from 200 to 1024.</p>
	// </li> <li> <p>Web and Express editions: Must be an integer from 20 to 1024.</p>
	// </li> </ul> </li> </ul>
	AllocatedStorage *int32

	// A value that indicates whether minor engine upgrades are applied automatically
	// to the DB instance during the maintenance window. By default, minor engine
	// upgrades are applied automatically.
	AutoMinorVersionUpgrade *bool

	// The Availability Zone (AZ) where the database will be created. For information
	// on AWS Regions and Availability Zones, see Regions and Availability Zones
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS Region.
	// Example: us-east-1d Constraint: The AvailabilityZone parameter can't be
	// specified if the DB instance is a Multi-AZ deployment. The specified
	// Availability Zone must be in the same AWS Region as the current endpoint. If
	// you're creating a DB instance in an RDS on VMware environment, specify the
	// identifier of the custom Availability Zone to create the DB instance in. For
	// more information about RDS on VMware, see the  RDS on VMware User Guide.
	// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
	AvailabilityZone *string

	// The number of days for which automated backups are retained. Setting this
	// parameter to a positive number enables backups. Setting this parameter to 0
	// disables automated backups. Amazon Aurora Not applicable. The retention period
	// for automated backups is managed by the DB cluster. Default: 1 Constraints:
	//
	//
	// * Must be a value from 0 to 35
	//
	//     * Can't be set to 0 if the DB instance is a
	// source to read replicas
	BackupRetentionPeriod *int32

	// For supported engines, indicates that the DB instance should be associated with
	// the specified CharacterSet.  <p> <b>Amazon Aurora</b> </p> <p>Not applicable.
	// The character set is managed by the DB cluster. For more information, see
	// <code>CreateDBCluster</code>.</p>
	CharacterSetName *string

	// A value that indicates whether to copy tags from the DB instance to snapshots of
	// the DB instance. By default, tags are not copied. Amazon Aurora Not applicable.
	// Copying tags to snapshots is managed by the DB cluster. Setting this value for
	// an Aurora DB instance has no effect on the DB cluster setting.
	CopyTagsToSnapshot *bool

	// The identifier of the DB cluster that the instance will belong to.
	DBClusterIdentifier *string

	// The meaning of this parameter differs according to the database engine you use.
	// MySQL The name of the database to create when the DB instance is created. If
	// this parameter isn't specified, no database is created in the DB instance.
	// Constraints:
	//
	//     * Must contain 1 to 64 letters or numbers.
	//
	//     * Can't be a
	// word reserved by the specified database engine
	//
	// MariaDB The name of the database
	// to create when the DB instance is created. If this parameter isn't specified, no
	// database is created in the DB instance. Constraints:
	//
	//     * Must contain 1 to 64
	// letters or numbers.
	//
	//     * Can't be a word reserved by the specified database
	// engine
	//
	// PostgreSQL The name of the database to create when the DB instance is
	// created. If this parameter isn't specified, the default "postgres" database is
	// created in the DB instance. Constraints:
	//
	//     * Must contain 1 to 63 letters,
	// numbers, or underscores.
	//
	//     * Must begin with a letter or an underscore.
	// Subsequent characters can be letters, underscores, or digits (0-9).
	//
	//     * Can't
	// be a word reserved by the specified database engine
	//
	// Oracle The Oracle System ID
	// (SID) of the created DB instance. If you specify null, the default value ORCL is
	// used. You can't specify the string NULL, or any other reserved word, for DBName.
	// Default: ORCL Constraints:
	//
	//     * Can't be longer than 8 characters
	//
	// SQL Server
	// Not applicable. Must be null. Amazon Aurora The name of the database to create
	// when the primary instance of the DB cluster is created. If this parameter isn't
	// specified, no database is created in the DB instance. Constraints:
	//
	//     * Must
	// contain 1 to 64 letters or numbers.
	//
	//     * Can't be a word reserved by the
	// specified database engine
	DBName *string

	// The name of the DB parameter group to associate with this DB instance. If you do
	// not specify a value, then the default DB parameter group for the specified DB
	// engine and version is used. Constraints:
	//
	//     * Must be 1 to 255 letters,
	// numbers, or hyphens.
	//
	//     * First character must be a letter
	//
	//     * Can't end
	// with a hyphen or contain two consecutive hyphens
	DBParameterGroupName *string

	// A list of DB security groups to associate with this DB instance. Default: The
	// default DB security group for the database engine.
	DBSecurityGroups []*string

	// A DB subnet group to associate with this DB instance. If there is no DB subnet
	// group, then it is a non-VPC DB instance.
	DBSubnetGroupName *string

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see  Deleting a DB
	// Instance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	// Amazon Aurora Not applicable. You can enable or disable deletion protection for
	// the DB cluster. For more information, see CreateDBCluster. DB instances in a DB
	// cluster can be deleted even when deletion protection is enabled for the DB
	// cluster.
	DeletionProtection *bool

	// The Active Directory directory ID to create the DB instance in. Currently, only
	// Microsoft SQL Server and Oracle DB instances can be created in an Active
	// Directory Domain. For Microsoft SQL Server DB instances, Amazon RDS can use
	// Windows Authentication to authenticate users that connect to the DB instance.
	// For more information, see  Using Windows Authentication with an Amazon RDS DB
	// Instance Running Microsoft SQL Server
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_SQLServerWinAuth.html)
	// in the Amazon RDS User Guide. For Oracle DB instances, Amazon RDS can use
	// Kerberos Authentication to authenticate users that connect to the DB instance.
	// For more information, see  Using Kerberos Authentication with Amazon RDS for
	// Oracle
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos.html) in
	// the Amazon RDS User Guide.
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string

	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Relational Database Service User Guide.
	EnableCloudwatchLogsExports []*string

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// You can enable IAM database authentication for the following database engines:
	// <p> <b>Amazon Aurora</b> </p> <p>Not applicable. Mapping AWS IAM accounts to
	// database accounts is managed by the DB cluster.</p> <p> <b>MySQL</b> </p> <ul>
	// <li> <p>For MySQL 5.6, minor version 5.6.34 or higher</p> </li> <li> <p>For
	// MySQL 5.7, minor version 5.7.16 or higher</p> </li> <li> <p>For MySQL 8.0, minor
	// version 8.0.16 or higher</p> </li> </ul> <p> <b>PostgreSQL</b> </p> <ul> <li>
	// <p>For PostgreSQL 9.5, minor version 9.5.15 or higher</p> </li> <li> <p>For
	// PostgreSQL 9.6, minor version 9.6.11 or higher</p> </li> <li> <p>PostgreSQL
	// 10.6, 10.7, and 10.9</p> </li> </ul> <p>For more information, see <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html">
	// IAM Database Authentication for MySQL and PostgreSQL</a> in the <i>Amazon RDS
	// User Guide.</i> </p>
	EnableIAMDatabaseAuthentication *bool

	// A value that indicates whether to enable Performance Insights for the DB
	// instance. For more information, see Using Amazon Performance Insights
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon Relational Database Service User Guide.
	EnablePerformanceInsights *bool

	// The version number of the database engine to use. For a list of valid engine
	// versions, use the DescribeDBEngineVersions action. The following are the
	// database engines and links to information about the major and minor versions
	// that are available with Amazon RDS. Not every database engine is available for
	// every AWS Region.  <p> <b>Amazon Aurora</b> </p> <p>Not applicable. The version
	// number of the database engine to be used by the DB instance is managed by the DB
	// cluster.</p> <p> <b>MariaDB</b> </p> <p>See <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt">MariaDB
	// on Amazon RDS Versions</a> in the <i>Amazon RDS User Guide.</i> </p> <p>
	// <b>Microsoft SQL Server</b> </p> <p>See <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.FeatureSupport">Version
	// and Feature Support on Amazon RDS</a> in the <i>Amazon RDS User Guide.</i> </p>
	// <p> <b>MySQL</b> </p> <p>See <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt">MySQL
	// on Amazon RDS Versions</a> in the <i>Amazon RDS User Guide.</i> </p> <p>
	// <b>Oracle</b> </p> <p>See <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html">Oracle
	// Database Engine Release Notes</a> in the <i>Amazon RDS User Guide.</i> </p> <p>
	// <b>PostgreSQL</b> </p> <p>See <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts.General.DBVersions">Supported
	// PostgreSQL Database Versions</a> in the <i>Amazon RDS User Guide.</i> </p>
	EngineVersion *string

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance. For information about valid Iops
	// values, see Amazon RDS Provisioned IOPS Storage to Improve Performance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide. Constraints: For MariaDB, MySQL, Oracle, and
	// PostgreSQL DB instances, must be a multiple between .5 and 50 of the storage
	// amount for the DB instance. For SQL Server DB instances, must be a multiple
	// between 1 and 50 of the storage amount for the DB instance.
	Iops *int32

	// The AWS KMS key identifier for an encrypted DB instance. The KMS key identifier
	// is the Amazon Resource Name (ARN) for the KMS encryption key. If you are
	// creating a DB instance with the same AWS account that owns the KMS encryption
	// key used to encrypt the new DB instance, then you can use the KMS key alias
	// instead of the ARN for the KM encryption key. Amazon Aurora Not applicable. The
	// KMS key identifier is managed by the DB cluster. For more information, see
	// CreateDBCluster.  <p>If <code>StorageEncrypted</code> is enabled, and you do not
	// specify a value for the <code>KmsKeyId</code> parameter, then Amazon RDS will
	// use your default encryption key. AWS KMS creates the default encryption key for
	// your AWS account. Your AWS account has a different default encryption key for
	// each AWS Region.</p>
	KmsKeyId *string

	// License model information for this DB instance. Valid values: license-included |
	// bring-your-own-license | general-public-license
	LicenseModel *string

	// The password for the master user. The password can include any printable ASCII
	// character except "/", """, or "@".  <p> <b>Amazon Aurora</b> </p> <p>Not
	// applicable. The password for the master user is managed by the DB cluster.</p>
	// <p> <b>MariaDB</b> </p> <p>Constraints: Must contain from 8 to 41
	// characters.</p> <p> <b>Microsoft SQL Server</b> </p> <p>Constraints: Must
	// contain from 8 to 128 characters.</p> <p> <b>MySQL</b> </p> <p>Constraints: Must
	// contain from 8 to 41 characters.</p> <p> <b>Oracle</b> </p> <p>Constraints: Must
	// contain from 8 to 30 characters.</p> <p> <b>PostgreSQL</b> </p> <p>Constraints:
	// Must contain from 8 to 128 characters.</p>
	MasterUserPassword *string

	// The name for the master user.  <p> <b>Amazon Aurora</b> </p> <p>Not applicable.
	// The name for the master user is managed by the DB cluster. </p> <p>
	// <b>MariaDB</b> </p> <p>Constraints:</p> <ul> <li> <p>Required for MariaDB.</p>
	// </li> <li> <p>Must be 1 to 16 letters or numbers.</p> </li> <li> <p>Can't be a
	// reserved word for the chosen database engine.</p> </li> </ul> <p> <b>Microsoft
	// SQL Server</b> </p> <p>Constraints:</p> <ul> <li> <p>Required for SQL
	// Server.</p> </li> <li> <p>Must be 1 to 128 letters or numbers.</p> </li> <li>
	// <p>The first character must be a letter.</p> </li> <li> <p>Can't be a reserved
	// word for the chosen database engine.</p> </li> </ul> <p> <b>MySQL</b> </p>
	// <p>Constraints:</p> <ul> <li> <p>Required for MySQL.</p> </li> <li> <p>Must be 1
	// to 16 letters or numbers.</p> </li> <li> <p>First character must be a
	// letter.</p> </li> <li> <p>Can't be a reserved word for the chosen database
	// engine.</p> </li> </ul> <p> <b>Oracle</b> </p> <p>Constraints:</p> <ul> <li>
	// <p>Required for Oracle.</p> </li> <li> <p>Must be 1 to 30 letters or
	// numbers.</p> </li> <li> <p>First character must be a letter.</p> </li> <li>
	// <p>Can't be a reserved word for the chosen database engine.</p> </li> </ul> <p>
	// <b>PostgreSQL</b> </p> <p>Constraints:</p> <ul> <li> <p>Required for
	// PostgreSQL.</p> </li> <li> <p>Must be 1 to 63 letters or numbers.</p> </li> <li>
	// <p>First character must be a letter.</p> </li> <li> <p>Can't be a reserved word
	// for the chosen database engine.</p> </li> </ul>
	MasterUsername *string

	// The upper limit to which Amazon RDS can automatically scale the storage of the
	// DB instance.
	MaxAllocatedStorage *int32

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0. If MonitoringRoleArn is specified, then
	// you must also set MonitoringInterval to a value other than 0. Valid Values: 0,
	// 1, 5, 10, 15, 30, 60
	MonitoringInterval *int32

	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to
	// Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For
	// information on creating a monitoring role, go to Setting Up and Enabling
	// Enhanced Monitoring
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling)
	// in the Amazon RDS User Guide. If MonitoringInterval is set to a value other than
	// 0, then you must supply a MonitoringRoleArn value.
	MonitoringRoleArn *string

	// A value that indicates whether the DB instance is a Multi-AZ deployment. You
	// can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ
	// deployment.
	MultiAZ *bool

	// Indicates that the DB instance should be associated with the specified option
	// group. Permanent options, such as the TDE option for Oracle Advanced Security
	// TDE, can't be removed from an option group. Also, that option group can't be
	// removed from a DB instance once it is associated with a DB instance
	OptionGroupName *string

	// The AWS KMS key identifier for encryption of Performance Insights data. The KMS
	// key ID is the Amazon Resource Name (ARN), KMS key identifier, or the KMS key
	// alias for the KMS encryption key. If you do not specify a value for
	// PerformanceInsightsKMSKeyId, then Amazon RDS uses your default encryption key.
	// AWS KMS creates the default encryption key for your AWS account. Your AWS
	// account has a different default encryption key for each AWS Region.
	PerformanceInsightsKMSKeyId *string

	// The amount of time, in days, to retain Performance Insights data. Valid values
	// are 7 or 731 (2 years).
	PerformanceInsightsRetentionPeriod *int32

	// The port number on which the database accepts connections. MySQL Default: 3306
	// Valid values: 1150-65535 Type: Integer MariaDB Default: 3306 Valid values:
	// 1150-65535 Type: Integer PostgreSQL Default: 5432 Valid values: 1150-65535 Type:
	// Integer Oracle Default: 1521 Valid values: 1150-65535 SQL Server Default: 1433
	// Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and
	// 49152-49156. Amazon Aurora Default: 3306 Valid values: 1150-65535 Type: Integer
	Port *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. For more
	// information, see The Backup Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow)
	// in the Amazon RDS User Guide.  <p> <b>Amazon Aurora</b> </p> <p>Not applicable.
	// The daily time range for creating automated backups is managed by the DB
	// cluster.</p> <p> The default is a 30-minute window selected at random from an
	// 8-hour block of time for each AWS Region. To see the time blocks available, see
	// <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow">
	// Adjusting the Preferred DB Instance Maintenance Window</a> in the <i>Amazon RDS
	// User Guide</i>. </p> <p>Constraints:</p> <ul> <li> <p>Must be in the format
	// <code>hh24:mi-hh24:mi</code>.</p> </li> <li> <p>Must be in Universal Coordinated
	// Time (UTC).</p> </li> <li> <p>Must not conflict with the preferred maintenance
	// window.</p> </li> <li> <p>Must be at least 30 minutes.</p> </li> </ul>
	PreferredBackupWindow *string

	// The time range each week during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). For more information, see Amazon RDS Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).
	// Format: ddd:hh24:mi-ddd:hh24:mi The default is a 30-minute window selected at
	// random from an 8-hour block of time for each AWS Region, occurring on a random
	// day of the week. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun. Constraints:
	// Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []*types.ProcessorFeature

	// A value that specifies the order in which an Aurora Replica is promoted to the
	// primary instance after a failure of the existing primary instance. For more
	// information, see  Fault Tolerance for an Aurora DB Cluster
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance)
	// in the Amazon Aurora User Guide. Default: 1 Valid Values: 0 - 15
	PromotionTier *int32

	// A value that indicates whether the DB instance is publicly accessible. When the
	// DB instance is publicly accessible, its DNS endpoint resolves to the private IP
	// address from within the DB instance's VPC, and to the public IP address from
	// outside of the DB instance's VPC. Access to the DB instance is ultimately
	// controlled by the security group it uses, and that public access is not
	// permitted if the security group assigned to the DB instance doesn't permit it.
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address. Default: The default
	// behavior varies depending on whether DBSubnetGroupName is specified. If
	// DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the
	// following applies:
	//
	//     * If the default VPC in the target region doesn’t have
	// an Internet gateway attached to it, the DB instance is private.
	//
	//     * If the
	// default VPC in the target region has an Internet gateway attached to it, the DB
	// instance is public.
	//
	// If DBSubnetGroupName is specified, and PubliclyAccessible
	// isn't specified, the following applies:
	//
	//     * If the subnets are part of a VPC
	// that doesn’t have an Internet gateway attached to it, the DB instance is
	// private.
	//
	//     * If the subnets are part of a VPC that has an Internet gateway
	// attached to it, the DB instance is public.
	PubliclyAccessible *bool

	// A value that indicates whether the DB instance is encrypted. By default, it
	// isn't encrypted.  <p> <b>Amazon Aurora</b> </p> <p>Not applicable. The
	// encryption for DB instances is managed by the DB cluster.</p>
	StorageEncrypted *bool

	// Specifies the storage type to be associated with the DB instance. Valid values:
	// standard | gp2 | io1 If you specify io1, you must also include a value for the
	// Iops parameter. Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string

	// Tags to assign to the DB instance.
	Tags []*types.Tag

	// The ARN from the key store with which to associate the instance for TDE
	// encryption.
	TdeCredentialArn *string

	// The password for the given ARN from the key store in order to access the device.
	TdeCredentialPassword *string

	// The time zone of the DB instance. The time zone parameter is currently supported
	// only by Microsoft SQL Server
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).
	Timezone *string

	// A list of Amazon EC2 VPC security groups to associate with this DB instance.
	// <p> <b>Amazon Aurora</b> </p> <p>Not applicable. The associated list of EC2 VPC
	// security groups is managed by the DB cluster.</p> <p>Default: The default EC2
	// VPC security group for the DB subnet group's VPC.</p>
	VpcSecurityGroupIds []*string
}

type CreateDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBInstance{}, middleware.After)
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
	addOpCreateDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBInstance",
	}
}
