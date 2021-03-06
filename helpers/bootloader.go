package helpers

import (
	"log"
	"path/filepath"
)

// DefaultAppSettings ...
func DefaultAppSettings() *AppSettings {
	return &AppSettings{
		ECSDK: ECSDK{
			URL:    "https://github.com/Enterprise-connect/ec-sdk/raw/dist/dist/",
			Folder: "dist",
			OS: OS{
				Linux:   "ecagent_linux_sys.tar.gz",
				ARM:     "ecagent_arm_sys.tar.gz",
				Darwin:  "ecagent_darwin_sys.tar.gz",
				Windows: "ecagent_windows_sys.exe.tar.gz",
			},
		},
		Output: Output{
			Root:        "output",
			FullDocFile: "ec-setup-full.md",
			Gateway: Gateway{
				Root:     "gateway",
				Script:   "ec-gateway.sh",
				Manifest: "manifest.yml",
				PushGatewayScript: PushScript{
					Unix:    "push-gateway.sh",
					Windows: "push-gateway.bat",
				},
			},
			Server: Server{
				Root: "server",
				Script: Script{
					Unix:    "ec-server.sh",
					Windows: "ec-server.bat",
				},
				Manifest: "manifest.yml",
				PushServerScript: PushScript{
					Unix:    "push-server.sh",
					Windows: "push-server.bat",
				},
			},
			Client: Client{
				Root: "client",
				Script: Script{
					Unix:    "ec-client.sh",
					Windows: "ec-client.bat",
				},
			},
		},
		Internal: Internal{
			Root: "ec-templates",
			Templates: Templates{
				GatewayTmpl: GatewayTmpl{
					Root:          "gateway",
					Manifest:      "gateway-manifest.tpl.yml",
					GatewayScript: "ec-gateway.tpl.sh",
					PushGatewayScript: PushScript{
						Unix:    "push-gateway.tpl.sh",
						Windows: "push-gateway.tpl.bat",
					},
				},
				ServerTmpl: ServerTmpl{
					Root:     "server",
					Manifest: "server-manifest.tpl.yml",
					ServerScriptScenarioOne: Script{
						Unix:    "ec-server.tpl.1.sh",
						Windows: "ec-server.tpl.1.bat",
					},
					ServerScriptScenarioTwo: Script{
						Unix: "ec-server.tpl.2.sh",
					},
					PushServerScript: PushScript{
						Unix:    "push-server.tpl.sh",
						Windows: "push-server.tpl.bat",
					},
				},
				ClientTmpl: ClientTmpl{
					Root: "client",
					ClientScriptScenarioOne: Script{
						Unix: "ec-client.tpl.1.sh",
					},
					ClientScriptScenarioTwo: Script{
						Unix:    "ec-client.tpl.2.sh",
						Windows: "ec-client.tpl.2.bat",
					},
				},
				FullDoc: FullDoc{
					ScenarioOne: "ec-scenario-1-full-doc.tpl.md",
					ScenarioTwo: "ec-scenario-2-full-doc.tpl.md",
				},
			},
		},
	}
}

type AppSettings struct {
	ECSDK    `json:"ec-sdk"`
	Output   `json:"output"`
	Internal `json:"internal"`
}

type ECSDK struct {
	URL    string `json:"url"`
	Folder string `json:"folder"`
	OS     `json:"os"`
}

type OS struct {
	Linux   string `json:"linux"`
	ARM     string `json:"arm"`
	Darwin  string `json:"darwin"`
	Windows string `json:"windows"`
}

type Output struct {
	Root        string `json:"root"`
	FullDocFile string `json:"full-doc-file"`
	Gateway     `json:"gateway"`
	Server      `json:"server"`
	Client      `json:"client"`
}

type Gateway struct {
	Root              string `json:"root"`
	Script            string `json:"script"`
	Manifest          string `json:"manifest"`
	PushGatewayScript PushScript
}

type Server struct {
	Root             string `json:"root"`
	Script           `json:"script"`
	Manifest         string `json:"manifest"`
	PushServerScript PushScript
}

type Client struct {
	Root   string `json:"root"`
	Script `json:"script"`
}

type Script struct {
	Unix    string `json:"unix"`
	Windows string `json:"windows"`
}

type Internal struct {
	Root      string `json:"root"`
	Templates `json:"templates"`
}

type ServerScriptScenarioOne = Script
type ServerScriptScenarioTwo = Script
type ClientScriptScenarioOne = Script
type ClientScriptScenarioTwo = Script
type PushScript = Script

type GatewayTmpl struct {
	Root              string `json:"root"`
	Manifest          string `json:"manifest"`
	GatewayScript     string `json:"gateway"`
	PushGatewayScript PushScript
}

type ServerTmpl struct {
	Root                    string `json:"root"`
	Manifest                string `json:"manifest"`
	ServerScriptScenarioOne `json:"serverOne"`
	ServerScriptScenarioTwo `json:"serverTwo"`
	PushServerScript        PushScript
}
type ClientTmpl struct {
	Root                    string `json:"root"`
	ClientScriptScenarioOne `json:"clientOne"`
	ClientScriptScenarioTwo `json:"clientTwo"`
}

type FullDoc struct {
	ScenarioOne string `json:"scenarioOne"`
	ScenarioTwo string `json:"scenarioTwo"`
}

type Templates struct {
	GatewayTmpl `json:"gateway"`
	ServerTmpl  `json:"server"`
	ClientTmpl  `json:"client"`
	FullDoc     `json:"full-doc"`
}

// InitAppStructure ...
func InitAppStructure(config *AppSettings) {
	log.Println("* Creating output folders...")
	// EC Gateway Folder
	gatewayFolder := filepath.Join(config.Output.Root, config.Output.Gateway.Root)
	CreateDirectory(gatewayFolder)
	log.Println("** ", gatewayFolder)
	// EC Server Folder
	serverFolder := filepath.Join(config.Output.Root, config.Output.Server.Root)
	CreateDirectory(serverFolder)
	log.Println("** ", serverFolder)
	// EC Client Folder
	clientFolder := filepath.Join(config.Output.Root, config.Output.Client.Root)
	CreateDirectory(clientFolder)
	log.Println("** ", clientFolder)
	// EC SDK Dist Folder
	ecSDKDistFolder := filepath.Join(config.Output.Root, config.ECSDK.Folder)
	CreateDirectory(ecSDKDistFolder)
	log.Println("** ", ecSDKDistFolder)
	// Done
	log.Println("** DONE")
}

// DownloadLatestECSDKVersion ...
func DownloadLatestECSDKVersion(config *AppSettings) {
	outputFolder := filepath.Join(config.Output.Root, config.ECSDK.Folder)
	urls := []string{
		config.ECSDK.URL + config.ECSDK.OS.Linux,
		config.ECSDK.URL + config.ECSDK.OS.ARM,
		config.ECSDK.URL + config.ECSDK.OS.Darwin,
		config.ECSDK.URL + config.ECSDK.OS.Windows,
	}
	downloadMultipleFiles(urls, outputFolder)
}
