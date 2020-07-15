package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Configs is the main configuration object
type Configs struct {
	RootDir      string                 `json:"root_dir"      yaml:"root_dir"      mapstructure:"root_dir"`
	ProjectTypes map[string]ProjectType `json:"project_types" yaml:"project_types" mapstructure:"project_types"`
}

// Validate the Configurations
func (confs Configs) Validate() error {
	return validation.ValidateStruct(&confs,
		validation.Field(&confs.RootDir, validation.Required, validation.Length(5, 50)),
		validation.Field(&confs.ProjectTypes),
	)
}

// DefaultConfigs returns the sensible default configs
func DefaultConfigs() *Configs {
	// pth, err := utils.AbsPath("~/gimme")
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to expand ~/gimme to absolute path: %d", err))
	// }
	// return &Configs{
	// 	RootDir: pth,
	// }
	return &Configs{}
}
