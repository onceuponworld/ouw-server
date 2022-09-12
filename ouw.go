package main


const (
	DEFAULT_READ_BUFFER_SIZE			= 1024
	DEFAULT_WRITE_BUFFER_SIZE			= 1024
)


const (
	DEFAULT_HOST						= "127.0.0.1"
	DEFAULT_PORT            = "9012"
)


const (
	APP_CONF                = "./config.json"
	APP_NAME								= "ouw-server"
	APP_VERSION							= "1.0"
)


const (
	RESOURCE_PLAYER					= "player"
	RESOURCE_KINGDOM				= "kingdom"
	RESOURCE_MUNICIPAL			= "municipal"
	RESOURCE_PLOT						= "plot"
)


const (
	PLAYER_ACTION_STATS			= "stats"
)


type AppConfig struct {
	Host					string				`json:"host"`
	Port					string				`json:"port"`
}


type ClientAction struct {
	Resource			string				`json:"resource"`
	Command				string				`json:"command"`
	Action				string				`json:"action"`
	Options       map[string] string 	`json:"options"`
}
