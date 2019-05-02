// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

const (

	// ErrCodeConcurrentModification for service response error code
	// "ConcurrentModification".
	//
	// Another user submitted a request to create, update, or delete the object
	// at the same time that you did. Retry the request.
	ErrCodeConcurrentModification = "ConcurrentModification"

	// ErrCodeConflictingDomainExists for service response error code
	// "ConflictingDomainExists".
	//
	// The cause of this error depends on whether you're trying to create a public
	// or a private hosted zone:
	//
	//    * Public hosted zone: Two hosted zones that have the same name or that
	//    have a parent/child relationship (example.com and test.example.com) can't
	//    have any common name servers. You tried to create a hosted zone that has
	//    the same name as an existing hosted zone or that's the parent or child
	//    of an existing hosted zone, and you specified a delegation set that shares
	//    one or more name servers with the existing hosted zone. For more information,
	//    see CreateReusableDelegationSet (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html).
	//
	//    * Private hosted zone: You specified an Amazon VPC that you're already
	//    using for another hosted zone, and the domain that you specified for one
	//    of the hosted zones is a subdomain of the domain that you specified for
	//    the other hosted zone. For example, you can't use the same Amazon VPC
	//    for the hosted zones for example.com and test.example.com.
	ErrCodeConflictingDomainExists = "ConflictingDomainExists"

	// ErrCodeConflictingTypes for service response error code
	// "ConflictingTypes".
	//
	// You tried to update a traffic policy instance by using a traffic policy version
	// that has a different DNS type than the current type for the instance. You
	// specified the type in the JSON document in the CreateTrafficPolicy or CreateTrafficPolicyVersionrequest.
	ErrCodeConflictingTypes = "ConflictingTypes"

	// ErrCodeDelegationSetAlreadyCreated for service response error code
	// "DelegationSetAlreadyCreated".
	//
	// A delegation set with the same owner and caller reference combination has
	// already been created.
	ErrCodeDelegationSetAlreadyCreated = "DelegationSetAlreadyCreated"

	// ErrCodeDelegationSetAlreadyReusable for service response error code
	// "DelegationSetAlreadyReusable".
	//
	// The specified delegation set has already been marked as reusable.
	ErrCodeDelegationSetAlreadyReusable = "DelegationSetAlreadyReusable"

	// ErrCodeDelegationSetInUse for service response error code
	// "DelegationSetInUse".
	//
	// The specified delegation contains associated hosted zones which must be deleted
	// before the reusable delegation set can be deleted.
	ErrCodeDelegationSetInUse = "DelegationSetInUse"

	// ErrCodeDelegationSetNotAvailable for service response error code
	// "DelegationSetNotAvailable".
	//
	// You can create a hosted zone that has the same name as an existing hosted
	// zone (example.com is common), but there is a limit to the number of hosted
	// zones that have the same name. If you get this error, Amazon Route 53 has
	// reached that limit. If you own the domain name and Route 53 generates this
	// error, contact Customer Support.
	ErrCodeDelegationSetNotAvailable = "DelegationSetNotAvailable"

	// ErrCodeDelegationSetNotReusable for service response error code
	// "DelegationSetNotReusable".
	//
	// A reusable delegation set with the specified ID does not exist.
	ErrCodeDelegationSetNotReusable = "DelegationSetNotReusable"

	// ErrCodeHealthCheckAlreadyExists for service response error code
	// "HealthCheckAlreadyExists".
	//
	// The health check you're attempting to create already exists. Amazon Route
	// 53 returns this error when you submit a request that has the following values:
	//
	//    * The same value for CallerReference as an existing health check, and
	//    one or more values that differ from the existing health check that has
	//    the same caller reference.
	//
	//    * The same value for CallerReference as a health check that you created
	//    and later deleted, regardless of the other settings in the request.
	ErrCodeHealthCheckAlreadyExists = "HealthCheckAlreadyExists"

	// ErrCodeHealthCheckInUse for service response error code
	// "HealthCheckInUse".
	//
	// This error code is not in use.
	ErrCodeHealthCheckInUse = "HealthCheckInUse"

	// ErrCodeHealthCheckVersionMismatch for service response error code
	// "HealthCheckVersionMismatch".
	//
	// The value of HealthCheckVersion in the request doesn't match the value of
	// HealthCheckVersion in the health check.
	ErrCodeHealthCheckVersionMismatch = "HealthCheckVersionMismatch"

	// ErrCodeHostedZoneAlreadyExists for service response error code
	// "HostedZoneAlreadyExists".
	//
	// The hosted zone you're trying to create already exists. Amazon Route 53 returns
	// this error when a hosted zone has already been created with the specified
	// CallerReference.
	ErrCodeHostedZoneAlreadyExists = "HostedZoneAlreadyExists"

	// ErrCodeHostedZoneNotEmpty for service response error code
	// "HostedZoneNotEmpty".
	//
	// The hosted zone contains resource records that are not SOA or NS records.
	ErrCodeHostedZoneNotEmpty = "HostedZoneNotEmpty"

	// ErrCodeHostedZoneNotFound for service response error code
	// "HostedZoneNotFound".
	//
	// The specified HostedZone can't be found.
	ErrCodeHostedZoneNotFound = "HostedZoneNotFound"

	// ErrCodeHostedZoneNotPrivate for service response error code
	// "HostedZoneNotPrivate".
	//
	// The specified hosted zone is a public hosted zone, not a private hosted zone.
	ErrCodeHostedZoneNotPrivate = "HostedZoneNotPrivate"

	// ErrCodeIncompatibleVersion for service response error code
	// "IncompatibleVersion".
	//
	// The resource you're trying to access is unsupported on this Amazon Route
	// 53 endpoint.
	ErrCodeIncompatibleVersion = "IncompatibleVersion"

	// ErrCodeInsufficientCloudWatchLogsResourcePolicy for service response error code
	// "InsufficientCloudWatchLogsResourcePolicy".
	//
	// Amazon Route 53 doesn't have the permissions required to create log streams
	// and send query logs to log streams. Possible causes include the following:
	//
	//    * There is no resource policy that specifies the log group ARN in the
	//    value for Resource.
	//
	//    * The resource policy that includes the log group ARN in the value for
	//    Resource doesn't have the necessary permissions.
	//
	//    * The resource policy hasn't finished propagating yet.
	ErrCodeInsufficientCloudWatchLogsResourcePolicy = "InsufficientCloudWatchLogsResourcePolicy"

	// ErrCodeInvalidArgument for service response error code
	// "InvalidArgument".
	//
	// Parameter name is invalid.
	ErrCodeInvalidArgument = "InvalidArgument"

	// ErrCodeInvalidChangeBatch for service response error code
	// "InvalidChangeBatch".
	//
	// This exception contains a list of messages that might contain one or more
	// error messages. Each error message indicates one error in the change batch.
	ErrCodeInvalidChangeBatch = "InvalidChangeBatch"

	// ErrCodeInvalidDomainName for service response error code
	// "InvalidDomainName".
	//
	// The specified domain name is not valid.
	ErrCodeInvalidDomainName = "InvalidDomainName"

	// ErrCodeInvalidInput for service response error code
	// "InvalidInput".
	//
	// The input is not valid.
	ErrCodeInvalidInput = "InvalidInput"

	// ErrCodeInvalidPaginationToken for service response error code
	// "InvalidPaginationToken".
	//
	// The value that you specified to get the second or subsequent page of results
	// is invalid.
	ErrCodeInvalidPaginationToken = "InvalidPaginationToken"

	// ErrCodeInvalidTrafficPolicyDocument for service response error code
	// "InvalidTrafficPolicyDocument".
	//
	// The format of the traffic policy document that you specified in the Document
	// element is invalid.
	ErrCodeInvalidTrafficPolicyDocument = "InvalidTrafficPolicyDocument"

	// ErrCodeInvalidVPCId for service response error code
	// "InvalidVPCId".
	//
	// The VPC ID that you specified either isn't a valid ID or the current account
	// is not authorized to access this VPC.
	ErrCodeInvalidVPCId = "InvalidVPCId"

	// ErrCodeLastVPCAssociation for service response error code
	// "LastVPCAssociation".
	//
	// The VPC that you're trying to disassociate from the private hosted zone is
	// the last VPC that is associated with the hosted zone. Amazon Route 53 doesn't
	// support disassociating the last VPC from a hosted zone.
	ErrCodeLastVPCAssociation = "LastVPCAssociation"

	// ErrCodeLimitsExceeded for service response error code
	// "LimitsExceeded".
	//
	// This operation can't be completed either because the current account has
	// reached the limit on reusable delegation sets that it can create or because
	// you've reached the limit on the number of Amazon VPCs that you can associate
	// with a private hosted zone. To get the current limit on the number of reusable
	// delegation sets, see GetAccountLimit (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
	// To get the current limit on the number of Amazon VPCs that you can associate
	// with a private hosted zone, see GetHostedZoneLimit (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetHostedZoneLimit.html).
	// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
	// with the AWS Support Center.
	ErrCodeLimitsExceeded = "LimitsExceeded"

	// ErrCodeNoSuchChange for service response error code
	// "NoSuchChange".
	//
	// A change with the specified change ID does not exist.
	ErrCodeNoSuchChange = "NoSuchChange"

	// ErrCodeNoSuchCloudWatchLogsLogGroup for service response error code
	// "NoSuchCloudWatchLogsLogGroup".
	//
	// There is no CloudWatch Logs log group with the specified ARN.
	ErrCodeNoSuchCloudWatchLogsLogGroup = "NoSuchCloudWatchLogsLogGroup"

	// ErrCodeNoSuchDelegationSet for service response error code
	// "NoSuchDelegationSet".
	//
	// A reusable delegation set with the specified ID does not exist.
	ErrCodeNoSuchDelegationSet = "NoSuchDelegationSet"

	// ErrCodeNoSuchGeoLocation for service response error code
	// "NoSuchGeoLocation".
	//
	// Amazon Route 53 doesn't support the specified geographic location.
	ErrCodeNoSuchGeoLocation = "NoSuchGeoLocation"

	// ErrCodeNoSuchHealthCheck for service response error code
	// "NoSuchHealthCheck".
	//
	// No health check exists with the specified ID.
	ErrCodeNoSuchHealthCheck = "NoSuchHealthCheck"

	// ErrCodeNoSuchHostedZone for service response error code
	// "NoSuchHostedZone".
	//
	// No hosted zone exists with the ID that you specified.
	ErrCodeNoSuchHostedZone = "NoSuchHostedZone"

	// ErrCodeNoSuchQueryLoggingConfig for service response error code
	// "NoSuchQueryLoggingConfig".
	//
	// There is no DNS query logging configuration with the specified ID.
	ErrCodeNoSuchQueryLoggingConfig = "NoSuchQueryLoggingConfig"

	// ErrCodeNoSuchTrafficPolicy for service response error code
	// "NoSuchTrafficPolicy".
	//
	// No traffic policy exists with the specified ID.
	ErrCodeNoSuchTrafficPolicy = "NoSuchTrafficPolicy"

	// ErrCodeNoSuchTrafficPolicyInstance for service response error code
	// "NoSuchTrafficPolicyInstance".
	//
	// No traffic policy instance exists with the specified ID.
	ErrCodeNoSuchTrafficPolicyInstance = "NoSuchTrafficPolicyInstance"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// Associating the specified VPC with the specified hosted zone has not been
	// authorized.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodePriorRequestNotComplete for service response error code
	// "PriorRequestNotComplete".
	//
	// If Amazon Route 53 can't process a request before the next request arrives,
	// it will reject subsequent requests for the same hosted zone and return an
	// HTTP 400 error (Bad request). If Route 53 returns this error repeatedly for
	// the same request, we recommend that you wait, in intervals of increasing
	// duration, before you try the request again.
	ErrCodePriorRequestNotComplete = "PriorRequestNotComplete"

	// ErrCodePublicZoneVPCAssociation for service response error code
	// "PublicZoneVPCAssociation".
	//
	// You're trying to associate a VPC with a public hosted zone. Amazon Route
	// 53 doesn't support associating a VPC with a public hosted zone.
	ErrCodePublicZoneVPCAssociation = "PublicZoneVPCAssociation"

	// ErrCodeQueryLoggingConfigAlreadyExists for service response error code
	// "QueryLoggingConfigAlreadyExists".
	//
	// You can create only one query logging configuration for a hosted zone, and
	// a query logging configuration already exists for this hosted zone.
	ErrCodeQueryLoggingConfigAlreadyExists = "QueryLoggingConfigAlreadyExists"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The limit on the number of requests per second was exceeded.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTooManyHealthChecks for service response error code
	// "TooManyHealthChecks".
	//
	// This health check can't be created because the current account has reached
	// the limit on the number of active health checks.
	//
	// For information about default limits, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// For information about how to get the current limit for an account, see GetAccountLimit
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
	// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
	// with the AWS Support Center.
	//
	// You have reached the maximum number of active health checks for an AWS account.
	// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
	// with the AWS Support Center.
	ErrCodeTooManyHealthChecks = "TooManyHealthChecks"

	// ErrCodeTooManyHostedZones for service response error code
	// "TooManyHostedZones".
	//
	// This operation can't be completed either because the current account has
	// reached the limit on the number of hosted zones or because you've reached
	// the limit on the number of hosted zones that can be associated with a reusable
	// delegation set.
	//
	// For information about default limits, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// To get the current limit on hosted zones that can be created by an account,
	// see GetAccountLimit (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
	//
	// To get the current limit on hosted zones that can be associated with a reusable
	// delegation set, see GetReusableDelegationSetLimit (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetReusableDelegationSetLimit.html).
	//
	// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
	// with the AWS Support Center.
	ErrCodeTooManyHostedZones = "TooManyHostedZones"

	// ErrCodeTooManyTrafficPolicies for service response error code
	// "TooManyTrafficPolicies".
	//
	// This traffic policy can't be created because the current account has reached
	// the limit on the number of traffic policies.
	//
	// For information about default limits, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// To get the current limit for an account, see GetAccountLimit (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
	//
	// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
	// with the AWS Support Center.
	ErrCodeTooManyTrafficPolicies = "TooManyTrafficPolicies"

	// ErrCodeTooManyTrafficPolicyInstances for service response error code
	// "TooManyTrafficPolicyInstances".
	//
	// This traffic policy instance can't be created because the current account
	// has reached the limit on the number of traffic policy instances.
	//
	// For information about default limits, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// For information about how to get the current limit for an account, see GetAccountLimit
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
	//
	// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
	// with the AWS Support Center.
	ErrCodeTooManyTrafficPolicyInstances = "TooManyTrafficPolicyInstances"

	// ErrCodeTooManyTrafficPolicyVersionsForCurrentPolicy for service response error code
	// "TooManyTrafficPolicyVersionsForCurrentPolicy".
	//
	// This traffic policy version can't be created because you've reached the limit
	// of 1000 on the number of versions that you can create for the current traffic
	// policy.
	//
	// To create more traffic policy versions, you can use GetTrafficPolicy (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetTrafficPolicy.html)
	// to get the traffic policy document for a specified traffic policy version,
	// and then use CreateTrafficPolicy (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateTrafficPolicy.html)
	// to create a new traffic policy using the traffic policy document.
	ErrCodeTooManyTrafficPolicyVersionsForCurrentPolicy = "TooManyTrafficPolicyVersionsForCurrentPolicy"

	// ErrCodeTooManyVPCAssociationAuthorizations for service response error code
	// "TooManyVPCAssociationAuthorizations".
	//
	// You've created the maximum number of authorizations that can be created for
	// the specified hosted zone. To authorize another VPC to be associated with
	// the hosted zone, submit a DeleteVPCAssociationAuthorization request to remove
	// an existing authorization. To get a list of existing authorizations, submit
	// a ListVPCAssociationAuthorizations request.
	ErrCodeTooManyVPCAssociationAuthorizations = "TooManyVPCAssociationAuthorizations"

	// ErrCodeTrafficPolicyAlreadyExists for service response error code
	// "TrafficPolicyAlreadyExists".
	//
	// A traffic policy that has the same value for Name already exists.
	ErrCodeTrafficPolicyAlreadyExists = "TrafficPolicyAlreadyExists"

	// ErrCodeTrafficPolicyInUse for service response error code
	// "TrafficPolicyInUse".
	//
	// One or more traffic policy instances were created by using the specified
	// traffic policy.
	ErrCodeTrafficPolicyInUse = "TrafficPolicyInUse"

	// ErrCodeTrafficPolicyInstanceAlreadyExists for service response error code
	// "TrafficPolicyInstanceAlreadyExists".
	//
	// There is already a traffic policy instance with the specified ID.
	ErrCodeTrafficPolicyInstanceAlreadyExists = "TrafficPolicyInstanceAlreadyExists"

	// ErrCodeVPCAssociationAuthorizationNotFound for service response error code
	// "VPCAssociationAuthorizationNotFound".
	//
	// The VPC that you specified is not authorized to be associated with the hosted
	// zone.
	ErrCodeVPCAssociationAuthorizationNotFound = "VPCAssociationAuthorizationNotFound"

	// ErrCodeVPCAssociationNotFound for service response error code
	// "VPCAssociationNotFound".
	//
	// The specified VPC and hosted zone are not currently associated.
	ErrCodeVPCAssociationNotFound = "VPCAssociationNotFound"
)
