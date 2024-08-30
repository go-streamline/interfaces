package utils

import (
	"bytes"
	"text/template"
)

// ParseTemplate is a utility function that parses a given template string with the provided data.
var ParseTemplate = parseTemplate

// parseTemplate takes a template string and data to populate the template.
// It returns the parsed template as a string or an error if the template parsing or execution fails.
//
// Parameters:
// - tmpl: The template string to be parsed.
// - data: The data to be used in the template.
//
// Returns:
// - A string containing the parsed template.
// - An error if any issues occur during parsing or execution of the template.
func parseTemplate(tmpl string, data any) (string, error) {
	parsedTemplate, err := template.New("any").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = parsedTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
