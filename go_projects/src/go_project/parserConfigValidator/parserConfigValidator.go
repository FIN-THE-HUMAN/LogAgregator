package parserConfigValidator

import (
	"go_project/errors"
	"go_project/parserConfig"
)

func IsGeneralDefaultStartPositionValid(parserConfig *parserConfig.ParserConfig) (bool, errors.ValidateError) {
	var err errors.ValidateError
	if parserConfig.General.Default.StartPosition == "end" || parserConfig.General.Default.StartPosition == "begin" {
		err.AddMessage("general gefault start_position is Invalid")
		return false, err
	}
	return true, err
}

func IsGeneralValid(general *parserConfig.General) (bool, errors.ValidateError) {
	var IsValid = true
	var err errors.ValidateError
	//if general.Plugins - проверка на правильность пути к плагинам /path/to/plugins/folder/
	//if general.Default.APIVersion - проверка на правильность версии general.Default.APIVersion[0]==v general.Default.APIVersion[1:lenght] - is int
	if general.Default.StartPosition != "end" && general.Default.StartPosition != "begin" {
		err.AddMessage("general default start_position must be end of begin")
		IsValid = false
	}
	if general.Default.StatInterval <= 0 {
		err.AddMessage("general default stat_interval must be > 0")
		IsValid = false
	}
	return IsValid, err
}

//подключить файл со структурой, на данный момент вместо нее ParserConfig
func IsParserConfigValid(parserConfig *parserConfig.ParserConfig) (bool, errors.ValidateError) {
	var IsConfigValid = true
	var err errors.ValidateError
	if parserConfig != nil {
		err.AddMessage("Config is NULL")
	}
	return IsConfigValid, err
}
