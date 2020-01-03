package snmpsim_restapi_client

/*
Labs is an array of Labs.
*/
type Labs []Lab

/*
Lab - Group of SNMP agents belonging to the same virtual laboratory. Some operations can be applied to them all at once.
*/
type Lab struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Agents Agents `json:"agents"`
	Power  string `json:"power"`
}

/*
Engines is an array of engines.
*/
type Engines []Engine

/*
Engine - Represents a unique, independent and fully operational SNMP engine, though not yet attached to any transport endpoints.
*/
type Engine struct {
	Id        int       `json:"id"`
	EngineId  string    `json:"engine_id"`
	Name      string    `json:"name"`
	Agents    Agents    `json:"agents"`
	Endpoints Endpoints `json:"endpoints"`
	Users     Users     `json:"users"`
}

/*
Agents is an array of agents.
*/
type Agents []Agent

/*
Agents - Represents SNMP agent. Consists of SNMP engine and transport endpoints it binds.
*/
type Agent struct {
	Id        int       `json:"id"`
	Engines   Engines   `json:"engines"`
	Name      string    `json:"name"`
	Endpoints Endpoints `json:"endpoints"`
	Labs      Labs      `json:"labs"`
	Selectors Selectors `json:"selectors"`
	DataDir   string    `json:"data_dir"`
}

/*
Endpoints is an array of endpoints.
*/
type Endpoints []Endpoint

/*
Endpoint - SNMP transport endpoint object. Each SNMP engine can bind one or more transport endpoints. Each transport endpoint can only be bound by one SNMP engine.
*/
type Endpoint struct {
	Id       int     `json:"id"`
	Engines  Engines `json:"engines"`
	Name     string  `json:"name"`
	Protocol string  `json:"protocol"`
	Address  string  `json:"address"`
}

/*
Recordings is an array of recordings.
*/
type Recordings []Recording

/*
Recording - Represents a single simulation data file residing by path under simulation data root.
*/
type Recording struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

/*
Users is an array of users.
*/
type Users []User

/*
User - SNMPv3 USM user object. Contains SNMPv3 credentials grouped by user name.
*/
type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	AuthKey   string  `json:"auth_key"`
	AuthProto string  `json:"auth_proto"`
	Engines   Engines `json:"engines"`
	PrivKey   string  `json:"priv_key"`
	PrivProto string  `json:"priv_proto"`
	User      string  `json:"user"`
}

//TODO: not implemented int the api yet, there is only one default selector for snmpv2c and one for snmpv3.
// We have to wait with the implementation until its implemented in the api.
// The strucuture is like this in the api doc but it might be different in the real api, this has to be checked first before it can be used.
/*
Selectors is an array of selectors.
*/
type Selectors []Selector

/*
Selector - Each selector should end up being a path to a simulation data file relative to the command responder's data directory.
The value of the selector can be static or, more likely, it contains templates that are expanded at run time. Each template can expand into some property of the inbound request.
Known templates include:
    ${context-engine-id}
    ${context-name}
    ${endpoint-id}
    ${source-address}
*/
type Selector struct {
	Id       int    `json:"id"`
	Comment  string `json:"comment"`
	Template string `json:"template"`
}

/*
ErrorResponse contains error information.
*/
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

/*
ProcessMetrics - SNMP simulator system is composed of many running processes. This object describes common properties of a process.
*/
type ProcessMetrics struct {
	Cmdline   *string           `json:"cmdline"`
	Uptime    *int              `json:"uptime"`
	Owner     *string           `json:"owner"`
	Memory    *int              `json:"memory"`
	Cpu       *int              `json:"cpu"`
	Files     *int              `json:"files"`
	StdOut    *string           `json:"stdout"`
	StdErr    *string           `json:"stderr"`
	LifeCycle *ProcessLifeCycle `json:"lifecycle"`
}

/*
ProcessLifeCycle - How this process has being doing.
*/
type ProcessLifeCycle struct {
	Exits    *int `json:"exits"`
	Restarts *int `json:"restarts"`
}

/*
ProcessesMetrics is an array of ProcessMetrics.
*/
type ProcessesMetrics []ProcessMetrics

/*
PacketMetrics - Transport endpoint related activity. Includes raw network packet counts as well as SNMP messages failed to get processed at the later stages.
*/
type PacketMetrics struct {
	FirstHit        *int   `json:"first_hit"`
	LastHit         *int   `json:"last_hit"`
	Total           *int64 `json:"total"`
	ParseFailures   *int64 `json:"parse_failures"`
	AuthFailures    *int64 `json:"auth_failures"`
	ContextFailures *int64 `json:"context_failures"`
}

/*
MessageMetrics - SNMP message level metrics.
*/
type MessageMetrics struct {
	FirstHit   *int          `json:"first_hit"`
	LastHit    *int          `json:"last_hit"`
	Pdus       *int64        `json:"pdus"`
	VarBinds   *int64        `json:"var_binds"`
	Failures   *int64        `json:"failures"`
	Variations []interface{} `json:"variations"`
}

/*
Variation - Variation module metrics.
*/
type Variation struct {
	FirstHit *int    `json:"first_hit"`
	LastHit  *int    `json:"last_hit"`
	Total    *int64  `json:"total"`
	Name     *string `json:"name"`
	Failures *int64  `json:"failures"`
}

/*
Variations is an array of variations.
*/
type Variations []Variation