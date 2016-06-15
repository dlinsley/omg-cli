package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Login struct {

	/*HomeRedirect - Descr: URL for configuring a custom home page Default: <nil>
*/
	HomeRedirect interface{} `yaml:"home_redirect,omitempty"`

	/*Prompt - Descr: The text used to prompt for a username during login Default: Email
*/
	Prompt *Prompt `yaml:"prompt,omitempty"`

	/*AssetBaseUrl - Descr: Base url for static assets, allows custom styling of the login server.  Use '/resources/pivotal' for Pivotal style. Default: /resources/oss
*/
	AssetBaseUrl interface{} `yaml:"asset_base_url,omitempty"`

	/*SpringProfiles - Descr: Deprecated. Use uaa.ldap.enabled - login.spring_profiles is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	SpringProfiles interface{} `yaml:"spring_profiles,omitempty"`

	/*Branding - Descr: This name is used on the UAA Pages and in account management related communication in UAA Default: <nil>
*/
	Branding *Branding `yaml:"branding,omitempty"`

	/*UaaBase - Descr: Deprecated. Use uaa.url for setting the location of UAA. Default: <nil>
*/
	UaaBase interface{} `yaml:"uaa_base,omitempty"`

	/*Analytics - Descr: Google analytics domain. If Google Analytics is desired set both login.analytics.code and login.analytics.domain Default: <nil>
*/
	Analytics *Analytics `yaml:"analytics,omitempty"`

	/*Logout - Descr: The Location of the redirect header following a logout of the the UAA (/logout.do). Default: /login
*/
	Logout *Logout `yaml:"logout,omitempty"`

	/*Links - Descr: A hash of home/passwd/signup URLS (see commented examples below) Default: <nil>
*/
	Links interface{} `yaml:"links,omitempty"`

	/*SelfServiceLinksEnabled - Descr: Enable self-service account creation and password resets links. Default: <nil>
*/
	SelfServiceLinksEnabled interface{} `yaml:"self_service_links_enabled,omitempty"`

	/*ClientSecret - Descr: Default login client secret, if no login client is defined Default: <nil>
*/
	ClientSecret interface{} `yaml:"client_secret,omitempty"`

	/*Saml - Descr: Read timeout in milliseconds for SAML metadata HTTP requests Default: 10000
*/
	Saml *Saml `yaml:"saml,omitempty"`

	/*Tiles - Descr: A list of links to other services to show on the landing page after log in. Default: <nil>
*/
	Tiles interface{} `yaml:"tiles,omitempty"`

	/*Messages - Descr: A nested or flat hash of messages that the login server uses to display UI message
This will be flattened into a java.util.Properties file. The example below will lead
to four properties, where the key is the concatenated value delimited by dot, for example scope.tokens.read=message
 Default: <nil>
*/
	Messages interface{} `yaml:"messages,omitempty"`

	/*Protocol - Descr: Scheme to use for HTTP communication (http/https) Default: https
*/
	Protocol interface{} `yaml:"protocol,omitempty"`

	/*SignupsEnabled - Descr: Deprecated. Use login.self_service_links_enabled. Instructs UAA to use 'enable account creation flow'. Enabled by default. Default: true
*/
	SignupsEnabled interface{} `yaml:"signups_enabled,omitempty"`

	/*EntityId - Descr: Deprecated. Use login.saml.entityid Default: <nil>
*/
	EntityId interface{} `yaml:"entity_id,omitempty"`

	/*Url - Descr: Set if you have an external login server.
The UAA uses this link on by its email service to create links
The UAA uses this as a base domain for internal hostnames so that subdomain can be detected
This defaults to the uaa.url property, and if not set, to login.<domain>
 Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

	/*Ldap - Descr: Deprecated. Use uaa.ldap.url - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	Ldap *Ldap `yaml:"ldap,omitempty"`

	/*Smtp - Descr: SMTP server password Default: <nil>
*/
	Smtp *Smtp `yaml:"smtp,omitempty"`

	/*InvitationsEnabled - Descr: Allows users to send invitations to email addresses outside the system and invite them to create an account. Disabled by default. Default: <nil>
*/
	InvitationsEnabled interface{} `yaml:"invitations_enabled,omitempty"`

	/*Notifications - Descr: The url for the notifications service (configure to use Notifications Service instead of SMTP server) Default: <nil>
*/
	Notifications *Notifications `yaml:"notifications,omitempty"`

}