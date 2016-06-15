package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Ldap struct {

	/*PasswordEncoder - Descr: Deprecated. Use uaa.ldap.passwordEncoder - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: org.cloudfoundry.identity.uaa.login.ldap.DynamicPasswordComparator
*/
	PasswordEncoder interface{} `yaml:"passwordEncoder,omitempty"`

	/*Enabled - Descr: Set to true to enable LDAP Default: false
*/
	Enabled interface{} `yaml:"enabled,omitempty"`

	/*Referral - Descr: Configures the UAA LDAP referral behavior. The following values are possible:
- follow -> Referrals are followed
- ignore -> Referrals are ignored and the partial result is returned
- throw  -> An error is thrown and the authentication is aborted
Reference: http://docs.oracle.com/javase/jndi/tutorial/ldap/referral/jndi.html
 Default: follow
*/
	Referral interface{} `yaml:"referral,omitempty"`

	/*Url - Descr: Deprecated. Use uaa.ldap.url - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

	/*ExternalGroupsWhitelist - Descr: Whitelist of external groups from LDAP that get added as roles in the ID Token Default: <nil>
*/
	ExternalGroupsWhitelist interface{} `yaml:"externalGroupsWhitelist,omitempty"`

	/*EmailDomain - Descr: Sets the whitelist of emails domains that the LDAP identity provider handles Default: <nil>
*/
	EmailDomain interface{} `yaml:"emailDomain,omitempty"`

	/*SearchFilter - Descr: Deprecated. Use uaa.ldap.searchFilter - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: cn={0}
*/
	SearchFilter interface{} `yaml:"searchFilter,omitempty"`

	/*UserPassword - Descr: Deprecated. Use uaa.ldap.userPassword - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	UserPassword interface{} `yaml:"userPassword,omitempty"`

	/*MailAttributeName - Descr: The name of the LDAP attribute that contains the users email address Default: mail
*/
	MailAttributeName interface{} `yaml:"mailAttributeName,omitempty"`

	/*UserDNPatternDelimiter - Descr: The delimiter character in between user DN patterns for simple-bind authentication Default: ;
*/
	UserDNPatternDelimiter interface{} `yaml:"userDNPatternDelimiter,omitempty"`

	/*LocalPasswordCompare - Descr: Deprecated. Use uaa.ldap.localPasswordCompare - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: true
*/
	LocalPasswordCompare interface{} `yaml:"localPasswordCompare,omitempty"`

	/*SslCertificateAlias - Descr: Deprecated. Use uaa.ldap.sslCertificateAlias - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	SslCertificateAlias interface{} `yaml:"sslCertificateAlias,omitempty"`

	/*ProfileType - Descr: The file to be used for configuring the LDAP authentication. Options are: 'simple-bind', 'search-and-bind', 'search-and-compare' Default: search-and-bind
*/
	ProfileType interface{} `yaml:"profile_type,omitempty"`

	/*PasswordAttributeName - Descr: Deprecated. Use uaa.ldap.passwordAttributeName - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: userPassword
*/
	PasswordAttributeName interface{} `yaml:"passwordAttributeName,omitempty"`

	/*MailSubstituteOverridesLdap - Descr: Set to true if you wish to override an LDAP user email address with a generated one Default: false
*/
	MailSubstituteOverridesLdap interface{} `yaml:"mailSubstituteOverridesLdap,omitempty"`

	/*SearchBase - Descr: Used with search-and-bind and search-and-compare. Define a base where the search starts at. Default: 
*/
	SearchBase interface{} `yaml:"searchBase,omitempty"`

	/*UserDNPattern - Descr: Deprecated. Use uaa.ldap.userDNPattern - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	UserDNPattern interface{} `yaml:"userDNPattern,omitempty"`

	/*SslCertificate - Descr: Used with ldaps:// URLs. The certificate, if self signed, to be trusted by this connection. Default: <nil>
*/
	SslCertificate interface{} `yaml:"sslCertificate,omitempty"`

	/*Groups - Descr: Boolean value, set to true to search below the search base Default: true
*/
	Groups *Groups `yaml:"groups,omitempty"`

	/*AttributeMappings - Descr: Specifies how UAA user attributes map to LDAP attributes. given_name, family_name, and phone_number are UAA user attributes, while other attributes should be included using the prefix `user.attribute` Default: <nil>
*/
	AttributeMappings interface{} `yaml:"attributeMappings,omitempty"`

	/*MailSubstitute - Descr: Defines an email pattern containing a {0} to generate an email address for an LDAP user during authentication Default: 
*/
	MailSubstitute interface{} `yaml:"mailSubstitute,omitempty"`

	/*UserDN - Descr: Deprecated. Use uaa.ldap.userDN - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	UserDN interface{} `yaml:"userDN,omitempty"`

}