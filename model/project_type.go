package model

import validation "github.com/go-ozzo/ozzo-validation"

// ProjectType represents a type of project Gimme can generate
type ProjectType struct {
	Language string             `json:"language" yaml:"language" mapstructure:"language"`
	SubTypes map[string]SubType `json:"subtypes" yaml:"subtypes" mapstructure:"subtypes"`
}

// Validate the ProjectType
func (pt ProjectType) Validate() error {
	return validation.ValidateStruct(&pt,
		validation.Field(&pt.Language, validation.Required),
		validation.Field(&pt.SubTypes, validation.Required),
	)
}

// SubType is a Project subtype
type SubType struct {
	Template string `json:"template" yaml:"template" mapstructure:"template"`
}

// Validate SubType...
func (st SubType) Validate() error {
	return validation.ValidateStruct(&st,
		validation.Field(&st.Template, validation.Required, validation.Length(2, 12)),
	)
}
