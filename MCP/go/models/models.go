package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tags interface{} `json:"Tags"`
}

// GetWebACLResponse represents the GetWebACLResponse schema from the OpenAPI specification
type GetWebACLResponse struct {
	Webacl interface{} `json:"WebACL,omitempty"`
}

// HTTPRequest represents the HTTPRequest schema from the OpenAPI specification
type HTTPRequest struct {
	Clientip interface{} `json:"ClientIP,omitempty"`
	Country interface{} `json:"Country,omitempty"`
	Httpversion interface{} `json:"HTTPVersion,omitempty"`
	Headers interface{} `json:"Headers,omitempty"`
	Method interface{} `json:"Method,omitempty"`
	Uri interface{} `json:"URI,omitempty"`
}

// ByteMatchSet represents the ByteMatchSet schema from the OpenAPI specification
type ByteMatchSet struct {
	Name interface{} `json:"Name,omitempty"`
	Bytematchsetid interface{} `json:"ByteMatchSetId"`
	Bytematchtuples interface{} `json:"ByteMatchTuples"`
}

// RateBasedRule represents the RateBasedRule schema from the OpenAPI specification
type RateBasedRule struct {
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Ratekey interface{} `json:"RateKey"`
	Ratelimit interface{} `json:"RateLimit"`
	Ruleid interface{} `json:"RuleId"`
	Matchpredicates interface{} `json:"MatchPredicates"`
}

// UpdateRegexPatternSetResponse represents the UpdateRegexPatternSetResponse schema from the OpenAPI specification
type UpdateRegexPatternSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// UpdateIPSetResponse represents the UpdateIPSetResponse schema from the OpenAPI specification
type UpdateIPSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// ListIPSetsRequest represents the ListIPSetsRequest schema from the OpenAPI specification
type ListIPSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// UpdateXssMatchSetRequest represents the UpdateXssMatchSetRequest schema from the OpenAPI specification
type UpdateXssMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Updates interface{} `json:"Updates"`
	Xssmatchsetid interface{} `json:"XssMatchSetId"`
}

// GetChangeTokenStatusResponse represents the GetChangeTokenStatusResponse schema from the OpenAPI specification
type GetChangeTokenStatusResponse struct {
	Changetokenstatus interface{} `json:"ChangeTokenStatus,omitempty"`
}

// ListWebACLsResponse represents the ListWebACLsResponse schema from the OpenAPI specification
type ListWebACLsResponse struct {
	Webacls interface{} `json:"WebACLs,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// DeleteWebACLResponse represents the DeleteWebACLResponse schema from the OpenAPI specification
type DeleteWebACLResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// CreateWebACLResponse represents the CreateWebACLResponse schema from the OpenAPI specification
type CreateWebACLResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Webacl interface{} `json:"WebACL,omitempty"`
}

// GetChangeTokenResponse represents the GetChangeTokenResponse schema from the OpenAPI specification
type GetChangeTokenResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// ListIPSetsResponse represents the ListIPSetsResponse schema from the OpenAPI specification
type ListIPSetsResponse struct {
	Ipsets interface{} `json:"IPSets,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// CreateGeoMatchSetRequest represents the CreateGeoMatchSetRequest schema from the OpenAPI specification
type CreateGeoMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// ListRateBasedRulesRequest represents the ListRateBasedRulesRequest schema from the OpenAPI specification
type ListRateBasedRulesRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// WebACLSummary represents the WebACLSummary schema from the OpenAPI specification
type WebACLSummary struct {
	Name interface{} `json:"Name"`
	Webaclid interface{} `json:"WebACLId"`
}

// DeleteSizeConstraintSetRequest represents the DeleteSizeConstraintSetRequest schema from the OpenAPI specification
type DeleteSizeConstraintSetRequest struct {
	Sizeconstraintsetid interface{} `json:"SizeConstraintSetId"`
	Changetoken interface{} `json:"ChangeToken"`
}

// CreateWebACLRequest represents the CreateWebACLRequest schema from the OpenAPI specification
type CreateWebACLRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Defaultaction interface{} `json:"DefaultAction"`
	Metricname interface{} `json:"MetricName"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// RuleSummary represents the RuleSummary schema from the OpenAPI specification
type RuleSummary struct {
	Name interface{} `json:"Name"`
	Ruleid interface{} `json:"RuleId"`
}

// UpdateSqlInjectionMatchSetRequest represents the UpdateSqlInjectionMatchSetRequest schema from the OpenAPI specification
type UpdateSqlInjectionMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Sqlinjectionmatchsetid interface{} `json:"SqlInjectionMatchSetId"`
	Updates interface{} `json:"Updates"`
}

// CreateGeoMatchSetResponse represents the CreateGeoMatchSetResponse schema from the OpenAPI specification
type CreateGeoMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Geomatchset interface{} `json:"GeoMatchSet,omitempty"`
}

// GeoMatchSetUpdate represents the GeoMatchSetUpdate schema from the OpenAPI specification
type GeoMatchSetUpdate struct {
	Action interface{} `json:"Action"`
	Geomatchconstraint interface{} `json:"GeoMatchConstraint"`
}

// IPSetUpdate represents the IPSetUpdate schema from the OpenAPI specification
type IPSetUpdate struct {
	Action interface{} `json:"Action"`
	Ipsetdescriptor interface{} `json:"IPSetDescriptor"`
}

// GetSqlInjectionMatchSetRequest represents the GetSqlInjectionMatchSetRequest schema from the OpenAPI specification
type GetSqlInjectionMatchSetRequest struct {
	Sqlinjectionmatchsetid interface{} `json:"SqlInjectionMatchSetId"`
}

// DeleteIPSetRequest represents the DeleteIPSetRequest schema from the OpenAPI specification
type DeleteIPSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Ipsetid interface{} `json:"IPSetId"`
}

// PutLoggingConfigurationRequest represents the PutLoggingConfigurationRequest schema from the OpenAPI specification
type PutLoggingConfigurationRequest struct {
	Loggingconfiguration interface{} `json:"LoggingConfiguration"`
}

// GetRegexPatternSetRequest represents the GetRegexPatternSetRequest schema from the OpenAPI specification
type GetRegexPatternSetRequest struct {
	Regexpatternsetid interface{} `json:"RegexPatternSetId"`
}

// ListRateBasedRulesResponse represents the ListRateBasedRulesResponse schema from the OpenAPI specification
type ListRateBasedRulesResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
}

// UpdateRegexMatchSetResponse represents the UpdateRegexMatchSetResponse schema from the OpenAPI specification
type UpdateRegexMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// RegexMatchSet represents the RegexMatchSet schema from the OpenAPI specification
type RegexMatchSet struct {
	Name interface{} `json:"Name,omitempty"`
	Regexmatchsetid interface{} `json:"RegexMatchSetId,omitempty"`
	Regexmatchtuples interface{} `json:"RegexMatchTuples,omitempty"`
}

// CreateSizeConstraintSetRequest represents the CreateSizeConstraintSetRequest schema from the OpenAPI specification
type CreateSizeConstraintSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// ListXssMatchSetsRequest represents the ListXssMatchSetsRequest schema from the OpenAPI specification
type ListXssMatchSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// ListXssMatchSetsResponse represents the ListXssMatchSetsResponse schema from the OpenAPI specification
type ListXssMatchSetsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Xssmatchsets interface{} `json:"XssMatchSets,omitempty"`
}

// GetSqlInjectionMatchSetResponse represents the GetSqlInjectionMatchSetResponse schema from the OpenAPI specification
type GetSqlInjectionMatchSetResponse struct {
	Sqlinjectionmatchset interface{} `json:"SqlInjectionMatchSet,omitempty"`
}

// CreateIPSetRequest represents the CreateIPSetRequest schema from the OpenAPI specification
type CreateIPSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// CreateByteMatchSetRequest represents the CreateByteMatchSetRequest schema from the OpenAPI specification
type CreateByteMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// DeleteSqlInjectionMatchSetRequest represents the DeleteSqlInjectionMatchSetRequest schema from the OpenAPI specification
type DeleteSqlInjectionMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Sqlinjectionmatchsetid interface{} `json:"SqlInjectionMatchSetId"`
}

// GetRateBasedRuleResponse represents the GetRateBasedRuleResponse schema from the OpenAPI specification
type GetRateBasedRuleResponse struct {
	Rule interface{} `json:"Rule,omitempty"`
}

// SizeConstraint represents the SizeConstraint schema from the OpenAPI specification
type SizeConstraint struct {
	Texttransformation interface{} `json:"TextTransformation"`
	Comparisonoperator interface{} `json:"ComparisonOperator"`
	Fieldtomatch interface{} `json:"FieldToMatch"`
	Size interface{} `json:"Size"`
}

// CreateRegexMatchSetRequest represents the CreateRegexMatchSetRequest schema from the OpenAPI specification
type CreateRegexMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// UpdateRateBasedRuleRequest represents the UpdateRateBasedRuleRequest schema from the OpenAPI specification
type UpdateRateBasedRuleRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Ratelimit interface{} `json:"RateLimit"`
	Ruleid interface{} `json:"RuleId"`
	Updates interface{} `json:"Updates"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// DeleteXssMatchSetRequest represents the DeleteXssMatchSetRequest schema from the OpenAPI specification
type DeleteXssMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Xssmatchsetid interface{} `json:"XssMatchSetId"`
}

// DeleteRuleGroupResponse represents the DeleteRuleGroupResponse schema from the OpenAPI specification
type DeleteRuleGroupResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GetChangeTokenStatusRequest represents the GetChangeTokenStatusRequest schema from the OpenAPI specification
type GetChangeTokenStatusRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
}

// CreateXssMatchSetRequest represents the CreateXssMatchSetRequest schema from the OpenAPI specification
type CreateXssMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// ListSqlInjectionMatchSetsRequest represents the ListSqlInjectionMatchSetsRequest schema from the OpenAPI specification
type ListSqlInjectionMatchSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// CreateSqlInjectionMatchSetResponse represents the CreateSqlInjectionMatchSetResponse schema from the OpenAPI specification
type CreateSqlInjectionMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Sqlinjectionmatchset interface{} `json:"SqlInjectionMatchSet,omitempty"`
}

// ListRulesRequest represents the ListRulesRequest schema from the OpenAPI specification
type ListRulesRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// GetSizeConstraintSetResponse represents the GetSizeConstraintSetResponse schema from the OpenAPI specification
type GetSizeConstraintSetResponse struct {
	Sizeconstraintset interface{} `json:"SizeConstraintSet,omitempty"`
}

// SizeConstraintSetSummary represents the SizeConstraintSetSummary schema from the OpenAPI specification
type SizeConstraintSetSummary struct {
	Sizeconstraintsetid interface{} `json:"SizeConstraintSetId"`
	Name interface{} `json:"Name"`
}

// UpdateWebACLResponse represents the UpdateWebACLResponse schema from the OpenAPI specification
type UpdateWebACLResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// CreateRuleRequest represents the CreateRuleRequest schema from the OpenAPI specification
type CreateRuleRequest struct {
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Changetoken interface{} `json:"ChangeToken"`
	Metricname interface{} `json:"MetricName"`
}

// DeleteByteMatchSetRequest represents the DeleteByteMatchSetRequest schema from the OpenAPI specification
type DeleteByteMatchSetRequest struct {
	Bytematchsetid interface{} `json:"ByteMatchSetId"`
	Changetoken interface{} `json:"ChangeToken"`
}

// ListRegexMatchSetsResponse represents the ListRegexMatchSetsResponse schema from the OpenAPI specification
type ListRegexMatchSetsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Regexmatchsets interface{} `json:"RegexMatchSets,omitempty"`
}

// WebACLUpdate represents the WebACLUpdate schema from the OpenAPI specification
type WebACLUpdate struct {
	Action interface{} `json:"Action"`
	Activatedrule interface{} `json:"ActivatedRule"`
}

// DeleteGeoMatchSetResponse represents the DeleteGeoMatchSetResponse schema from the OpenAPI specification
type DeleteGeoMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GeoMatchSet represents the GeoMatchSet schema from the OpenAPI specification
type GeoMatchSet struct {
	Geomatchconstraints interface{} `json:"GeoMatchConstraints"`
	Geomatchsetid interface{} `json:"GeoMatchSetId"`
	Name interface{} `json:"Name,omitempty"`
}

// RegexPatternSet represents the RegexPatternSet schema from the OpenAPI specification
type RegexPatternSet struct {
	Name interface{} `json:"Name,omitempty"`
	Regexpatternsetid interface{} `json:"RegexPatternSetId"`
	Regexpatternstrings interface{} `json:"RegexPatternStrings"`
}

// ListByteMatchSetsResponse represents the ListByteMatchSetsResponse schema from the OpenAPI specification
type ListByteMatchSetsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Bytematchsets interface{} `json:"ByteMatchSets,omitempty"`
}

// CreateRegexMatchSetResponse represents the CreateRegexMatchSetResponse schema from the OpenAPI specification
type CreateRegexMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Regexmatchset interface{} `json:"RegexMatchSet,omitempty"`
}

// XssMatchSet represents the XssMatchSet schema from the OpenAPI specification
type XssMatchSet struct {
	Name interface{} `json:"Name,omitempty"`
	Xssmatchsetid interface{} `json:"XssMatchSetId"`
	Xssmatchtuples interface{} `json:"XssMatchTuples"`
}

// GetRuleResponse represents the GetRuleResponse schema from the OpenAPI specification
type GetRuleResponse struct {
	Rule interface{} `json:"Rule,omitempty"`
}

// CreateRateBasedRuleResponse represents the CreateRateBasedRuleResponse schema from the OpenAPI specification
type CreateRateBasedRuleResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Rule interface{} `json:"Rule,omitempty"`
}

// CreateRegexPatternSetRequest represents the CreateRegexPatternSetRequest schema from the OpenAPI specification
type CreateRegexPatternSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// ExcludedRule represents the ExcludedRule schema from the OpenAPI specification
type ExcludedRule struct {
	Ruleid interface{} `json:"RuleId"`
}

// DeleteSizeConstraintSetResponse represents the DeleteSizeConstraintSetResponse schema from the OpenAPI specification
type DeleteSizeConstraintSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GetRateBasedRuleManagedKeysRequest represents the GetRateBasedRuleManagedKeysRequest schema from the OpenAPI specification
type GetRateBasedRuleManagedKeysRequest struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Ruleid interface{} `json:"RuleId"`
}

// ByteMatchSetSummary represents the ByteMatchSetSummary schema from the OpenAPI specification
type ByteMatchSetSummary struct {
	Bytematchsetid interface{} `json:"ByteMatchSetId"`
	Name interface{} `json:"Name"`
}

// UpdateIPSetRequest represents the UpdateIPSetRequest schema from the OpenAPI specification
type UpdateIPSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Ipsetid interface{} `json:"IPSetId"`
	Updates interface{} `json:"Updates"`
}

// GetIPSetRequest represents the GetIPSetRequest schema from the OpenAPI specification
type GetIPSetRequest struct {
	Ipsetid interface{} `json:"IPSetId"`
}

// SqlInjectionMatchSetSummary represents the SqlInjectionMatchSetSummary schema from the OpenAPI specification
type SqlInjectionMatchSetSummary struct {
	Name interface{} `json:"Name"`
	Sqlinjectionmatchsetid interface{} `json:"SqlInjectionMatchSetId"`
}

// DeleteXssMatchSetResponse represents the DeleteXssMatchSetResponse schema from the OpenAPI specification
type DeleteXssMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// WafAction represents the WafAction schema from the OpenAPI specification
type WafAction struct {
	TypeField interface{} `json:"Type"`
}

// GetSampledRequestsResponse represents the GetSampledRequestsResponse schema from the OpenAPI specification
type GetSampledRequestsResponse struct {
	Timewindow interface{} `json:"TimeWindow,omitempty"`
	Populationsize interface{} `json:"PopulationSize,omitempty"`
	Sampledrequests interface{} `json:"SampledRequests,omitempty"`
}

// DeleteRegexMatchSetRequest represents the DeleteRegexMatchSetRequest schema from the OpenAPI specification
type DeleteRegexMatchSetRequest struct {
	Regexmatchsetid interface{} `json:"RegexMatchSetId"`
	Changetoken interface{} `json:"ChangeToken"`
}

// ListWebACLsRequest represents the ListWebACLsRequest schema from the OpenAPI specification
type ListWebACLsRequest struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Limit interface{} `json:"Limit,omitempty"`
}

// CreateWebACLMigrationStackResponse represents the CreateWebACLMigrationStackResponse schema from the OpenAPI specification
type CreateWebACLMigrationStackResponse struct {
	S3objecturl interface{} `json:"S3ObjectUrl"`
}

// GetLoggingConfigurationRequest represents the GetLoggingConfigurationRequest schema from the OpenAPI specification
type GetLoggingConfigurationRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// GetPermissionPolicyResponse represents the GetPermissionPolicyResponse schema from the OpenAPI specification
type GetPermissionPolicyResponse struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// Predicate represents the Predicate schema from the OpenAPI specification
type Predicate struct {
	TypeField interface{} `json:"Type"`
	Dataid interface{} `json:"DataId"`
	Negated interface{} `json:"Negated"`
}

// ListSizeConstraintSetsResponse represents the ListSizeConstraintSetsResponse schema from the OpenAPI specification
type ListSizeConstraintSetsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Sizeconstraintsets interface{} `json:"SizeConstraintSets,omitempty"`
}

// DeleteRegexMatchSetResponse represents the DeleteRegexMatchSetResponse schema from the OpenAPI specification
type DeleteRegexMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Taginfoforresource interface{} `json:"TagInfoForResource,omitempty"`
}

// ActivatedRule represents the ActivatedRule schema from the OpenAPI specification
type ActivatedRule struct {
	Overrideaction interface{} `json:"OverrideAction,omitempty"`
	Priority interface{} `json:"Priority"`
	Ruleid interface{} `json:"RuleId"`
	TypeField interface{} `json:"Type,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Excludedrules interface{} `json:"ExcludedRules,omitempty"`
}

// UpdateByteMatchSetResponse represents the UpdateByteMatchSetResponse schema from the OpenAPI specification
type UpdateByteMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// DeleteRuleRequest represents the DeleteRuleRequest schema from the OpenAPI specification
type DeleteRuleRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Ruleid interface{} `json:"RuleId"`
}

// CreateWebACLMigrationStackRequest represents the CreateWebACLMigrationStackRequest schema from the OpenAPI specification
type CreateWebACLMigrationStackRequest struct {
	Ignoreunsupportedtype interface{} `json:"IgnoreUnsupportedType"`
	S3bucketname interface{} `json:"S3BucketName"`
	Webaclid interface{} `json:"WebACLId"`
}

// UpdateRuleRequest represents the UpdateRuleRequest schema from the OpenAPI specification
type UpdateRuleRequest struct {
	Ruleid interface{} `json:"RuleId"`
	Updates interface{} `json:"Updates"`
	Changetoken interface{} `json:"ChangeToken"`
}

// LoggingConfiguration represents the LoggingConfiguration schema from the OpenAPI specification
type LoggingConfiguration struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Logdestinationconfigs interface{} `json:"LogDestinationConfigs"`
	Redactedfields interface{} `json:"RedactedFields,omitempty"`
}

// GetXssMatchSetRequest represents the GetXssMatchSetRequest schema from the OpenAPI specification
type GetXssMatchSetRequest struct {
	Xssmatchsetid interface{} `json:"XssMatchSetId"`
}

// CreateRateBasedRuleRequest represents the CreateRateBasedRuleRequest schema from the OpenAPI specification
type CreateRateBasedRuleRequest struct {
	Name interface{} `json:"Name"`
	Ratekey interface{} `json:"RateKey"`
	Ratelimit interface{} `json:"RateLimit"`
	Tags interface{} `json:"Tags,omitempty"`
	Changetoken interface{} `json:"ChangeToken"`
	Metricname interface{} `json:"MetricName"`
}

// DeleteGeoMatchSetRequest represents the DeleteGeoMatchSetRequest schema from the OpenAPI specification
type DeleteGeoMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Geomatchsetid interface{} `json:"GeoMatchSetId"`
}

// GetByteMatchSetResponse represents the GetByteMatchSetResponse schema from the OpenAPI specification
type GetByteMatchSetResponse struct {
	Bytematchset interface{} `json:"ByteMatchSet,omitempty"`
}

// RuleGroupSummary represents the RuleGroupSummary schema from the OpenAPI specification
type RuleGroupSummary struct {
	Name interface{} `json:"Name"`
	Rulegroupid interface{} `json:"RuleGroupId"`
}

// IPSetDescriptor represents the IPSetDescriptor schema from the OpenAPI specification
type IPSetDescriptor struct {
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
}

// UpdateXssMatchSetResponse represents the UpdateXssMatchSetResponse schema from the OpenAPI specification
type UpdateXssMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GeoMatchSetSummary represents the GeoMatchSetSummary schema from the OpenAPI specification
type GeoMatchSetSummary struct {
	Name interface{} `json:"Name"`
	Geomatchsetid interface{} `json:"GeoMatchSetId"`
}

// DeleteLoggingConfigurationRequest represents the DeleteLoggingConfigurationRequest schema from the OpenAPI specification
type DeleteLoggingConfigurationRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// CreateSizeConstraintSetResponse represents the CreateSizeConstraintSetResponse schema from the OpenAPI specification
type CreateSizeConstraintSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Sizeconstraintset interface{} `json:"SizeConstraintSet,omitempty"`
}

// ListLoggingConfigurationsResponse represents the ListLoggingConfigurationsResponse schema from the OpenAPI specification
type ListLoggingConfigurationsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Loggingconfigurations interface{} `json:"LoggingConfigurations,omitempty"`
}

// WafOverrideAction represents the WafOverrideAction schema from the OpenAPI specification
type WafOverrideAction struct {
	TypeField interface{} `json:"Type"`
}

// DeleteRateBasedRuleRequest represents the DeleteRateBasedRuleRequest schema from the OpenAPI specification
type DeleteRateBasedRuleRequest struct {
	Ruleid interface{} `json:"RuleId"`
	Changetoken interface{} `json:"ChangeToken"`
}

// ByteMatchTuple represents the ByteMatchTuple schema from the OpenAPI specification
type ByteMatchTuple struct {
	Targetstring interface{} `json:"TargetString"`
	Texttransformation interface{} `json:"TextTransformation"`
	Fieldtomatch interface{} `json:"FieldToMatch"`
	Positionalconstraint interface{} `json:"PositionalConstraint"`
}

// UpdateRuleGroupRequest represents the UpdateRuleGroupRequest schema from the OpenAPI specification
type UpdateRuleGroupRequest struct {
	Rulegroupid interface{} `json:"RuleGroupId"`
	Updates interface{} `json:"Updates"`
	Changetoken interface{} `json:"ChangeToken"`
}

// GetRegexMatchSetRequest represents the GetRegexMatchSetRequest schema from the OpenAPI specification
type GetRegexMatchSetRequest struct {
	Regexmatchsetid interface{} `json:"RegexMatchSetId"`
}

// RegexMatchTuple represents the RegexMatchTuple schema from the OpenAPI specification
type RegexMatchTuple struct {
	Fieldtomatch interface{} `json:"FieldToMatch"`
	Regexpatternsetid interface{} `json:"RegexPatternSetId"`
	Texttransformation interface{} `json:"TextTransformation"`
}

// ListSubscribedRuleGroupsRequest represents the ListSubscribedRuleGroupsRequest schema from the OpenAPI specification
type ListSubscribedRuleGroupsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// GetWebACLRequest represents the GetWebACLRequest schema from the OpenAPI specification
type GetWebACLRequest struct {
	Webaclid interface{} `json:"WebACLId"`
}

// DeleteByteMatchSetResponse represents the DeleteByteMatchSetResponse schema from the OpenAPI specification
type DeleteByteMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// ByteMatchSetUpdate represents the ByteMatchSetUpdate schema from the OpenAPI specification
type ByteMatchSetUpdate struct {
	Action interface{} `json:"Action"`
	Bytematchtuple interface{} `json:"ByteMatchTuple"`
}

// GetRateBasedRuleManagedKeysResponse represents the GetRateBasedRuleManagedKeysResponse schema from the OpenAPI specification
type GetRateBasedRuleManagedKeysResponse struct {
	Managedkeys interface{} `json:"ManagedKeys,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// CreateByteMatchSetResponse represents the CreateByteMatchSetResponse schema from the OpenAPI specification
type CreateByteMatchSetResponse struct {
	Bytematchset interface{} `json:"ByteMatchSet,omitempty"`
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tagkeys interface{} `json:"TagKeys"`
}

// GetSampledRequestsRequest represents the GetSampledRequestsRequest schema from the OpenAPI specification
type GetSampledRequestsRequest struct {
	Ruleid interface{} `json:"RuleId"`
	Timewindow interface{} `json:"TimeWindow"`
	Webaclid interface{} `json:"WebAclId"`
	Maxitems interface{} `json:"MaxItems"`
}

// SqlInjectionMatchSetUpdate represents the SqlInjectionMatchSetUpdate schema from the OpenAPI specification
type SqlInjectionMatchSetUpdate struct {
	Action interface{} `json:"Action"`
	Sqlinjectionmatchtuple interface{} `json:"SqlInjectionMatchTuple"`
}

// UpdateWebACLRequest represents the UpdateWebACLRequest schema from the OpenAPI specification
type UpdateWebACLRequest struct {
	Defaultaction interface{} `json:"DefaultAction,omitempty"`
	Updates interface{} `json:"Updates,omitempty"`
	Webaclid interface{} `json:"WebACLId"`
	Changetoken interface{} `json:"ChangeToken"`
}

// DeleteLoggingConfigurationResponse represents the DeleteLoggingConfigurationResponse schema from the OpenAPI specification
type DeleteLoggingConfigurationResponse struct {
}

// GetRuleGroupResponse represents the GetRuleGroupResponse schema from the OpenAPI specification
type GetRuleGroupResponse struct {
	Rulegroup interface{} `json:"RuleGroup,omitempty"`
}

// ListRegexPatternSetsResponse represents the ListRegexPatternSetsResponse schema from the OpenAPI specification
type ListRegexPatternSetsResponse struct {
	Regexpatternsets interface{} `json:"RegexPatternSets,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// PutPermissionPolicyRequest represents the PutPermissionPolicyRequest schema from the OpenAPI specification
type PutPermissionPolicyRequest struct {
	Policy interface{} `json:"Policy"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// UpdateRateBasedRuleResponse represents the UpdateRateBasedRuleResponse schema from the OpenAPI specification
type UpdateRateBasedRuleResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// RuleUpdate represents the RuleUpdate schema from the OpenAPI specification
type RuleUpdate struct {
	Action interface{} `json:"Action"`
	Predicate interface{} `json:"Predicate"`
}

// CreateRuleGroupResponse represents the CreateRuleGroupResponse schema from the OpenAPI specification
type CreateRuleGroupResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Rulegroup interface{} `json:"RuleGroup,omitempty"`
}

// XssMatchSetUpdate represents the XssMatchSetUpdate schema from the OpenAPI specification
type XssMatchSetUpdate struct {
	Xssmatchtuple interface{} `json:"XssMatchTuple"`
	Action interface{} `json:"Action"`
}

// ListRuleGroupsResponse represents the ListRuleGroupsResponse schema from the OpenAPI specification
type ListRuleGroupsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Rulegroups interface{} `json:"RuleGroups,omitempty"`
}

// ListGeoMatchSetsResponse represents the ListGeoMatchSetsResponse schema from the OpenAPI specification
type ListGeoMatchSetsResponse struct {
	Geomatchsets interface{} `json:"GeoMatchSets,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// ListSqlInjectionMatchSetsResponse represents the ListSqlInjectionMatchSetsResponse schema from the OpenAPI specification
type ListSqlInjectionMatchSetsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Sqlinjectionmatchsets interface{} `json:"SqlInjectionMatchSets,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// CreateXssMatchSetResponse represents the CreateXssMatchSetResponse schema from the OpenAPI specification
type CreateXssMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Xssmatchset interface{} `json:"XssMatchSet,omitempty"`
}

// ListLoggingConfigurationsRequest represents the ListLoggingConfigurationsRequest schema from the OpenAPI specification
type ListLoggingConfigurationsRequest struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Limit interface{} `json:"Limit,omitempty"`
}

// CreateIPSetResponse represents the CreateIPSetResponse schema from the OpenAPI specification
type CreateIPSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Ipset interface{} `json:"IPSet,omitempty"`
}

// ListSizeConstraintSetsRequest represents the ListSizeConstraintSetsRequest schema from the OpenAPI specification
type ListSizeConstraintSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// UpdateByteMatchSetRequest represents the UpdateByteMatchSetRequest schema from the OpenAPI specification
type UpdateByteMatchSetRequest struct {
	Bytematchsetid interface{} `json:"ByteMatchSetId"`
	Changetoken interface{} `json:"ChangeToken"`
	Updates interface{} `json:"Updates"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Resourcearn interface{} `json:"ResourceARN"`
}

// WebACL represents the WebACL schema from the OpenAPI specification
type WebACL struct {
	Webaclarn interface{} `json:"WebACLArn,omitempty"`
	Webaclid interface{} `json:"WebACLId"`
	Defaultaction interface{} `json:"DefaultAction"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rules interface{} `json:"Rules"`
}

// RegexMatchSetUpdate represents the RegexMatchSetUpdate schema from the OpenAPI specification
type RegexMatchSetUpdate struct {
	Action interface{} `json:"Action"`
	Regexmatchtuple interface{} `json:"RegexMatchTuple"`
}

// ListSubscribedRuleGroupsResponse represents the ListSubscribedRuleGroupsResponse schema from the OpenAPI specification
type ListSubscribedRuleGroupsResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Rulegroups interface{} `json:"RuleGroups,omitempty"`
}

// ListRegexMatchSetsRequest represents the ListRegexMatchSetsRequest schema from the OpenAPI specification
type ListRegexMatchSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// SizeConstraintSetUpdate represents the SizeConstraintSetUpdate schema from the OpenAPI specification
type SizeConstraintSetUpdate struct {
	Action interface{} `json:"Action"`
	Sizeconstraint interface{} `json:"SizeConstraint"`
}

// FieldToMatch represents the FieldToMatch schema from the OpenAPI specification
type FieldToMatch struct {
	Data interface{} `json:"Data,omitempty"`
	TypeField interface{} `json:"Type"`
}

// GetByteMatchSetRequest represents the GetByteMatchSetRequest schema from the OpenAPI specification
type GetByteMatchSetRequest struct {
	Bytematchsetid interface{} `json:"ByteMatchSetId"`
}

// UpdateRuleResponse represents the UpdateRuleResponse schema from the OpenAPI specification
type UpdateRuleResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// XssMatchTuple represents the XssMatchTuple schema from the OpenAPI specification
type XssMatchTuple struct {
	Fieldtomatch interface{} `json:"FieldToMatch"`
	Texttransformation interface{} `json:"TextTransformation"`
}

// TimeWindow represents the TimeWindow schema from the OpenAPI specification
type TimeWindow struct {
	Endtime interface{} `json:"EndTime"`
	Starttime interface{} `json:"StartTime"`
}

// DeleteSqlInjectionMatchSetResponse represents the DeleteSqlInjectionMatchSetResponse schema from the OpenAPI specification
type DeleteSqlInjectionMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GetLoggingConfigurationResponse represents the GetLoggingConfigurationResponse schema from the OpenAPI specification
type GetLoggingConfigurationResponse struct {
	Loggingconfiguration interface{} `json:"LoggingConfiguration,omitempty"`
}

// ListRulesResponse represents the ListRulesResponse schema from the OpenAPI specification
type ListRulesResponse struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
}

// GetRegexMatchSetResponse represents the GetRegexMatchSetResponse schema from the OpenAPI specification
type GetRegexMatchSetResponse struct {
	Regexmatchset interface{} `json:"RegexMatchSet,omitempty"`
}

// DeleteRuleGroupRequest represents the DeleteRuleGroupRequest schema from the OpenAPI specification
type DeleteRuleGroupRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Rulegroupid interface{} `json:"RuleGroupId"`
}

// CreateRuleResponse represents the CreateRuleResponse schema from the OpenAPI specification
type CreateRuleResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Rule interface{} `json:"Rule,omitempty"`
}

// GetSizeConstraintSetRequest represents the GetSizeConstraintSetRequest schema from the OpenAPI specification
type GetSizeConstraintSetRequest struct {
	Sizeconstraintsetid interface{} `json:"SizeConstraintSetId"`
}

// ListRegexPatternSetsRequest represents the ListRegexPatternSetsRequest schema from the OpenAPI specification
type ListRegexPatternSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// GetGeoMatchSetRequest represents the GetGeoMatchSetRequest schema from the OpenAPI specification
type GetGeoMatchSetRequest struct {
	Geomatchsetid interface{} `json:"GeoMatchSetId"`
}

// CreateRuleGroupRequest represents the CreateRuleGroupRequest schema from the OpenAPI specification
type CreateRuleGroupRequest struct {
	Metricname interface{} `json:"MetricName"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Changetoken interface{} `json:"ChangeToken"`
}

// GetRuleRequest represents the GetRuleRequest schema from the OpenAPI specification
type GetRuleRequest struct {
	Ruleid interface{} `json:"RuleId"`
}

// HTTPHeader represents the HTTPHeader schema from the OpenAPI specification
type HTTPHeader struct {
	Value interface{} `json:"Value,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// PutLoggingConfigurationResponse represents the PutLoggingConfigurationResponse schema from the OpenAPI specification
type PutLoggingConfigurationResponse struct {
	Loggingconfiguration interface{} `json:"LoggingConfiguration,omitempty"`
}

// GetRuleGroupRequest represents the GetRuleGroupRequest schema from the OpenAPI specification
type GetRuleGroupRequest struct {
	Rulegroupid interface{} `json:"RuleGroupId"`
}

// SubscribedRuleGroupSummary represents the SubscribedRuleGroupSummary schema from the OpenAPI specification
type SubscribedRuleGroupSummary struct {
	Metricname interface{} `json:"MetricName"`
	Name interface{} `json:"Name"`
	Rulegroupid interface{} `json:"RuleGroupId"`
}

// UpdateSizeConstraintSetRequest represents the UpdateSizeConstraintSetRequest schema from the OpenAPI specification
type UpdateSizeConstraintSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Sizeconstraintsetid interface{} `json:"SizeConstraintSetId"`
	Updates interface{} `json:"Updates"`
}

// SizeConstraintSet represents the SizeConstraintSet schema from the OpenAPI specification
type SizeConstraintSet struct {
	Name interface{} `json:"Name,omitempty"`
	Sizeconstraintsetid interface{} `json:"SizeConstraintSetId"`
	Sizeconstraints interface{} `json:"SizeConstraints"`
}

// RuleGroupUpdate represents the RuleGroupUpdate schema from the OpenAPI specification
type RuleGroupUpdate struct {
	Action interface{} `json:"Action"`
	Activatedrule interface{} `json:"ActivatedRule"`
}

// RegexMatchSetSummary represents the RegexMatchSetSummary schema from the OpenAPI specification
type RegexMatchSetSummary struct {
	Name interface{} `json:"Name"`
	Regexmatchsetid interface{} `json:"RegexMatchSetId"`
}

// CreateRegexPatternSetResponse represents the CreateRegexPatternSetResponse schema from the OpenAPI specification
type CreateRegexPatternSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
	Regexpatternset interface{} `json:"RegexPatternSet,omitempty"`
}

// SqlInjectionMatchTuple represents the SqlInjectionMatchTuple schema from the OpenAPI specification
type SqlInjectionMatchTuple struct {
	Fieldtomatch interface{} `json:"FieldToMatch"`
	Texttransformation interface{} `json:"TextTransformation"`
}

// ListActivatedRulesInRuleGroupResponse represents the ListActivatedRulesInRuleGroupResponse schema from the OpenAPI specification
type ListActivatedRulesInRuleGroupResponse struct {
	Activatedrules interface{} `json:"ActivatedRules,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// SqlInjectionMatchSet represents the SqlInjectionMatchSet schema from the OpenAPI specification
type SqlInjectionMatchSet struct {
	Name interface{} `json:"Name,omitempty"`
	Sqlinjectionmatchsetid interface{} `json:"SqlInjectionMatchSetId"`
	Sqlinjectionmatchtuples interface{} `json:"SqlInjectionMatchTuples"`
}

// DeleteIPSetResponse represents the DeleteIPSetResponse schema from the OpenAPI specification
type DeleteIPSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// CreateSqlInjectionMatchSetRequest represents the CreateSqlInjectionMatchSetRequest schema from the OpenAPI specification
type CreateSqlInjectionMatchSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Name interface{} `json:"Name"`
}

// RegexPatternSetSummary represents the RegexPatternSetSummary schema from the OpenAPI specification
type RegexPatternSetSummary struct {
	Name interface{} `json:"Name"`
	Regexpatternsetid interface{} `json:"RegexPatternSetId"`
}

// UpdateGeoMatchSetResponse represents the UpdateGeoMatchSetResponse schema from the OpenAPI specification
type UpdateGeoMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// DeleteRegexPatternSetResponse represents the DeleteRegexPatternSetResponse schema from the OpenAPI specification
type DeleteRegexPatternSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// ListRuleGroupsRequest represents the ListRuleGroupsRequest schema from the OpenAPI specification
type ListRuleGroupsRequest struct {
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Limit interface{} `json:"Limit,omitempty"`
}

// TagInfoForResource represents the TagInfoForResource schema from the OpenAPI specification
type TagInfoForResource struct {
	Resourcearn interface{} `json:"ResourceARN,omitempty"`
	Taglist interface{} `json:"TagList,omitempty"`
}

// UpdateSizeConstraintSetResponse represents the UpdateSizeConstraintSetResponse schema from the OpenAPI specification
type UpdateSizeConstraintSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// PutPermissionPolicyResponse represents the PutPermissionPolicyResponse schema from the OpenAPI specification
type PutPermissionPolicyResponse struct {
}

// DeletePermissionPolicyResponse represents the DeletePermissionPolicyResponse schema from the OpenAPI specification
type DeletePermissionPolicyResponse struct {
}

// UpdateSqlInjectionMatchSetResponse represents the UpdateSqlInjectionMatchSetResponse schema from the OpenAPI specification
type UpdateSqlInjectionMatchSetResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GetChangeTokenRequest represents the GetChangeTokenRequest schema from the OpenAPI specification
type GetChangeTokenRequest struct {
}

// DeleteRuleResponse represents the DeleteRuleResponse schema from the OpenAPI specification
type DeleteRuleResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GetPermissionPolicyRequest represents the GetPermissionPolicyRequest schema from the OpenAPI specification
type GetPermissionPolicyRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// GetXssMatchSetResponse represents the GetXssMatchSetResponse schema from the OpenAPI specification
type GetXssMatchSetResponse struct {
	Xssmatchset interface{} `json:"XssMatchSet,omitempty"`
}

// GetRateBasedRuleRequest represents the GetRateBasedRuleRequest schema from the OpenAPI specification
type GetRateBasedRuleRequest struct {
	Ruleid interface{} `json:"RuleId"`
}

// ListActivatedRulesInRuleGroupRequest represents the ListActivatedRulesInRuleGroupRequest schema from the OpenAPI specification
type ListActivatedRulesInRuleGroupRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
	Rulegroupid interface{} `json:"RuleGroupId,omitempty"`
}

// IPSet represents the IPSet schema from the OpenAPI specification
type IPSet struct {
	Ipsetdescriptors interface{} `json:"IPSetDescriptors"`
	Ipsetid interface{} `json:"IPSetId"`
	Name interface{} `json:"Name,omitempty"`
}

// ListGeoMatchSetsRequest represents the ListGeoMatchSetsRequest schema from the OpenAPI specification
type ListGeoMatchSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// Rule represents the Rule schema from the OpenAPI specification
type Rule struct {
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Predicates interface{} `json:"Predicates"`
	Ruleid interface{} `json:"RuleId"`
}

// GeoMatchConstraint represents the GeoMatchConstraint schema from the OpenAPI specification
type GeoMatchConstraint struct {
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
}

// UpdateRegexPatternSetRequest represents the UpdateRegexPatternSetRequest schema from the OpenAPI specification
type UpdateRegexPatternSetRequest struct {
	Regexpatternsetid interface{} `json:"RegexPatternSetId"`
	Updates interface{} `json:"Updates"`
	Changetoken interface{} `json:"ChangeToken"`
}

// XssMatchSetSummary represents the XssMatchSetSummary schema from the OpenAPI specification
type XssMatchSetSummary struct {
	Name interface{} `json:"Name"`
	Xssmatchsetid interface{} `json:"XssMatchSetId"`
}

// DeleteWebACLRequest represents the DeleteWebACLRequest schema from the OpenAPI specification
type DeleteWebACLRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Webaclid interface{} `json:"WebACLId"`
}

// DeletePermissionPolicyRequest represents the DeletePermissionPolicyRequest schema from the OpenAPI specification
type DeletePermissionPolicyRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// RegexPatternSetUpdate represents the RegexPatternSetUpdate schema from the OpenAPI specification
type RegexPatternSetUpdate struct {
	Action interface{} `json:"Action"`
	Regexpatternstring interface{} `json:"RegexPatternString"`
}

// GetGeoMatchSetResponse represents the GetGeoMatchSetResponse schema from the OpenAPI specification
type GetGeoMatchSetResponse struct {
	Geomatchset interface{} `json:"GeoMatchSet,omitempty"`
}

// SampledHTTPRequest represents the SampledHTTPRequest schema from the OpenAPI specification
type SampledHTTPRequest struct {
	Rulewithinrulegroup interface{} `json:"RuleWithinRuleGroup,omitempty"`
	Timestamp interface{} `json:"Timestamp,omitempty"`
	Weight interface{} `json:"Weight"`
	Action interface{} `json:"Action,omitempty"`
	Request interface{} `json:"Request"`
}

// UpdateGeoMatchSetRequest represents the UpdateGeoMatchSetRequest schema from the OpenAPI specification
type UpdateGeoMatchSetRequest struct {
	Geomatchsetid interface{} `json:"GeoMatchSetId"`
	Updates interface{} `json:"Updates"`
	Changetoken interface{} `json:"ChangeToken"`
}

// UpdateRegexMatchSetRequest represents the UpdateRegexMatchSetRequest schema from the OpenAPI specification
type UpdateRegexMatchSetRequest struct {
	Updates interface{} `json:"Updates"`
	Changetoken interface{} `json:"ChangeToken"`
	Regexmatchsetid interface{} `json:"RegexMatchSetId"`
}

// UpdateRuleGroupResponse represents the UpdateRuleGroupResponse schema from the OpenAPI specification
type UpdateRuleGroupResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// DeleteRegexPatternSetRequest represents the DeleteRegexPatternSetRequest schema from the OpenAPI specification
type DeleteRegexPatternSetRequest struct {
	Changetoken interface{} `json:"ChangeToken"`
	Regexpatternsetid interface{} `json:"RegexPatternSetId"`
}

// IPSetSummary represents the IPSetSummary schema from the OpenAPI specification
type IPSetSummary struct {
	Ipsetid interface{} `json:"IPSetId"`
	Name interface{} `json:"Name"`
}

// GetRegexPatternSetResponse represents the GetRegexPatternSetResponse schema from the OpenAPI specification
type GetRegexPatternSetResponse struct {
	Regexpatternset interface{} `json:"RegexPatternSet,omitempty"`
}

// RuleGroup represents the RuleGroup schema from the OpenAPI specification
type RuleGroup struct {
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rulegroupid interface{} `json:"RuleGroupId"`
}

// ListByteMatchSetsRequest represents the ListByteMatchSetsRequest schema from the OpenAPI specification
type ListByteMatchSetsRequest struct {
	Limit interface{} `json:"Limit,omitempty"`
	Nextmarker interface{} `json:"NextMarker,omitempty"`
}

// DeleteRateBasedRuleResponse represents the DeleteRateBasedRuleResponse schema from the OpenAPI specification
type DeleteRateBasedRuleResponse struct {
	Changetoken interface{} `json:"ChangeToken,omitempty"`
}

// GetIPSetResponse represents the GetIPSetResponse schema from the OpenAPI specification
type GetIPSetResponse struct {
	Ipset interface{} `json:"IPSet,omitempty"`
}
