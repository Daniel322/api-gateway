package config_manager

const CONFIG_FOLDER = "./config"

// static array with config keys
var SUPPORTED_KEYS = []string{"server.port", "wsconnection.useWs", "wsconnection.keepalive", "testInclude.asd"}

var DEFAULT_CONFIG = map[string]interface{}{
	"server": map[string]interface{}{
		"port": 5000,
	},
	"wsconnection": map[string]interface{}{
		"useWS":     true,
		"keepalive": 10,
	},
	"testInclude": map[string]interface{}{
		"asd": false,
	},
}
