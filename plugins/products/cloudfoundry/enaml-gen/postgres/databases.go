package postgres 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Databases struct {

	/*Address - Descr: The database address Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*LogLinePrefix - Descr: The postgres `printf` style string that is output at the beginning of each log line Default: %m: 
*/
	LogLinePrefix interface{} `yaml:"log_line_prefix,omitempty"`

	/*DbScheme - Descr: The database scheme Default: <nil>
*/
	DbScheme interface{} `yaml:"db_scheme,omitempty"`

	/*Databases - Descr: A list of databases and associated properties to create Default: <nil>
*/
	Databases interface{} `yaml:"databases,omitempty"`

	/*Roles - Descr: A list of database roles and associated properties to create Default: <nil>
*/
	Roles interface{} `yaml:"roles,omitempty"`

	/*Port - Descr: The database port Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*CollectStatementStatistics - Descr: Enable the `pg_stat_statements` extension and collect statement execution statistics Default: false
*/
	CollectStatementStatistics interface{} `yaml:"collect_statement_statistics,omitempty"`

	/*MaxConnections - Descr: Maximum number of database connections Default: <nil>
*/
	MaxConnections interface{} `yaml:"max_connections,omitempty"`

	/*AdditionalConfig - Descr: A map of additional key/value pairs to include as extra configuration properties Default: <nil>
*/
	AdditionalConfig interface{} `yaml:"additional_config,omitempty"`

}