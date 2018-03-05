package parserConfig

/*
import (
	"encoding/json"
)
*/
type Default struct {
	APIVersion    string `json:"api_version"`
	StartPosition string `json:"start_position"`
	StatInterval  int    `json:"stat_interval"`
}

type General struct {
	Plugins string  `json:"plugins"`
	Default Default `json:"default"`
}

type HTTP struct {
	Port string `json:"port"`
}

type LogHTTP struct {
	Type string `json:"type"`
}

type LogModulesInputSettings struct {
	Path string `json:"path"`
}

type LogModulesInput struct {
	Type     string                  `json:"type"`
	Settings LogModulesInputSettings `json:"settings"`
}

type LogModulesOutputSettings struct {
	Driver string `json:"driver"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Db     string `json:"db"`
}

type LogModulesOutput struct {
	Type     string                   `json:"type"`
	Settings LogModulesOutputSettings `json:"settings"`
}

type Modules struct { /////
	HTTP   LogHTTP          `json:"http"`
	Input  LogModulesInput  `json:"input"`
	Output LogModulesOutput `json:"output"`
}

type Log struct {
	Debug   bool    `json:"debug"`
	Modules Modules `json:"modules"`
}

type InputFile struct {
	Name          string   `json:"name"`
	APIVersion    string   `json:"api_version"`
	StartPosition string   `json:"start_position"`
	StatInterval  int      `json:"stat_interval"`
	Path          []string `json:"path"`
}

type InputTCP struct {
	Name string `json:"name"`
	Made string `json:"mode"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type InputDb struct {
	Name          string `json:"name"`
	APIVersion    string `json:"api_version"`
	StartPosition string `json:"start_position"`
	StatInterval  string `json:"stat_interval"`
	Type          string `json:"type"`
	Driver        string `json:"driver"`
	Host          string `json:"host"`
	port          int    `json:"port"`
}

type Input struct {
	Files []InputFile `json:"file"`
	TCP   []InputTCP  `json:"tcp"`
	DB    []InputDb   `json:"db"`
}

type Postgres struct {
	Input  []string `json:"input"`
	Driver string   `json:"driver"`
	Host   string   `json:"host"`
	Port   int      `json:"port"`
	User   string   `json:"user"`
	Pass   string   `json:"pass"`
	Db     string   `json:"db"`
}

type MongoDb struct {
	Input  []string `json:"input"`
	Driver string   `json:"driver"`
	Host   string   `json:"host"`
	Port   int      `json:"port"`
	User   string   `json:"user"`
	Pass   string   `json:"pass"`
	Db     string   `json:"db"`
}

type Output struct {
	Postgres Postgres `json:"postgres"`
	MongoDb  MongoDb  `json:"mongodb"`
}

type ParserConfig struct {
	General General `json:"general"`
	HTTP    HTTP    `json:"http"`
	Log     Log     `json:"log"`
	Input   Input   `json:"input"`
	Output  Output  `json:"output"`
}
