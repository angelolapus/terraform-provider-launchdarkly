package launchdarkly

// keys used in terraform files referencing keys in launchdarkly resource objects.
// The name of each constant is the same as its value.
const (
	//gofmts:sort
	ACTIONS                          = "actions"
	API_KEY                          = "api_key"
	APPROVAL_SETTINGS                = "approval_settings"
	ARCHIVED                         = "archived"
	ATTRIBUTE                        = "attribute"
	BUCKET_BY                        = "bucket_by"
	CAN_APPLY_DECLINED_CHANGES       = "can_apply_declined_changes"
	CAN_REVIEW_OWN_REQUEST           = "can_review_own_request"
	CLAUSES                          = "clauses"
	CLIENT_SIDE_AVAILABILITY         = "client_side_availability"
	CLIENT_SIDE_ID                   = "client_side_id"
	COLOR                            = "color"
	CONFIG                           = "config"
	CONFIRM_CHANGES                  = "confirm_changes"
	CREATION_DATE                    = "creation_date"
	CUSTOM_PROPERTIES                = "custom_properties"
	CUSTOM_ROLES                     = "custom_roles"
	DEFAULTS                         = "defaults"
	DEFAULT_API_VERSION              = "default_api_version"
	DEFAULT_CLIENT_SIDE_AVAILABILITY = "default_client_side_availability"
	DEFAULT_OFF_VARIATION            = "default_off_variation"
	DEFAULT_ON_VARIATION             = "default_on_variation"
	DEFAULT_TRACK_EVENTS             = "default_track_events"
	DEFAULT_TTL                      = "default_ttl"
	DESCRIPTION                      = "description"
	EFFECT                           = "effect"
	EMAIL                            = "email"
	ENABLED                          = "enabled"
	ENVIRONMENTS                     = "environments"
	ENV_KEY                          = "env_key"
	EXCLUDED                         = "excluded"
	EXPIRE                           = "expire"
	FALLTHROUGH                      = "fallthrough"
	FIRST_NAME                       = "first_name"
	FLAG_ID                          = "flag_id"
	FLAG_KEY                         = "flag_key"
	ID                               = "id"
	INCLUDED                         = "included"
	INCLUDE_IN_SNIPPET               = "include_in_snippet"
	INLINE_ROLES                     = "inline_roles"
	INTEGRATION_KEY                  = "integration_key"
	KEY                              = "key"
	KIND                             = "kind"
	LAST_NAME                        = "last_name"
	MAINTAINER_ID                    = "maintainer_id"
	MIN_NUM_APPROVALS                = "min_num_approvals"
	MOBILE_KEY                       = "mobile_key"
	NAME                             = "name"
	NEGATE                           = "negate"
	NOT_ACTIONS                      = "not_actions"
	NOT_RESOURCES                    = "not_resources"
	OFF_VARIATION                    = "off_variation"
	ON                               = "on"
	ON_VARIATION                     = "on_variation"
	OP                               = "op"
	POLICY                           = "policy"
	POLICY_STATEMENTS                = "policy_statements"
	PREREQUISITES                    = "prerequisites"
	PROJECT_KEY                      = "project_key"
	REQUIRED                         = "required"
	REQUIRED_APPROVAL_TAGS           = "required_approval_tags"
	REQUIRE_COMMENTS                 = "require_comments"
	RESOURCES                        = "resources"
	ROLE                             = "role"
	ROLLOUT_WEIGHTS                  = "rollout_weights"
	RULES                            = "rules"
	SECRET                           = "secret"
	SECURE_MODE                      = "secure_mode"
	SERVICE_TOKEN                    = "service_token"
	STATEMENTS                       = "statements"
	TAGS                             = "tags"
	TARGETS                          = "targets"
	TEMPORARY                        = "temporary"
	TOKEN                            = "token"
	TRACK_EVENTS                     = "track_events"
	URL                              = "url"
	USING_ENVIRONMENT_ID             = "using_environment_id"
	USING_MOBILE_KEY                 = "using_mobile_key"
	VALUE                            = "value"
	VALUES                           = "values"
	VALUE_TYPE                       = "value_type"
	VARIATION                        = "variation"
	VARIATIONS                       = "variations"
	VARIATION_TYPE                   = "variation_type"
	WEIGHT                           = "weight"
)
