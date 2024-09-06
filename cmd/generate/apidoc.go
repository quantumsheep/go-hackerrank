package main

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

const APIDOC_URL = "https://www.hackerrank.com/apidoc"

type HackerRankApiDoc struct {
	Swagger             string                            `json:"swagger"`
	Info                *Info                             `json:"info"`
	SecurityDefinitions map[string]*SecurityDefinition    `json:"securityDefinitions"`
	Host                string                            `json:"host"`
	BasePath            string                            `json:"basePath"`
	Consumes            []string                          `json:"consumes"`
	Produces            []string                          `json:"produces"`
	Schemes             []string                          `json:"schemes"`
	Paths               map[string]map[string]*PathMethod `json:"paths"`
	Definitions         map[string]*Schema                `json:"definitions"`
}

func LoadHackerRankApiDoc() (*HackerRankApiDoc, error) {
	res, err := http.Get(APIDOC_URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var doc HackerRankApiDoc
	err = json.Unmarshal(bytes, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

var parameterRegex = regexp.MustCompile(`(s/)?{([^}]+)}`)

func (p *HackerRankApiDoc) GetPathFuncNameSnakeCase(path string, method string) string {
	path = strings.SplitN(path, "?", 2)[0]

	// Handle versioned paths
	var version string
	if strings.HasPrefix(path, "/x/api/v") {
		path = strings.TrimPrefix(path, "/x/api/v")
		version = strings.SplitN(path, "/", 2)[0]
		path = strings.TrimPrefix(path, version)[1:]
	}

	// Replace path parameters with snake case
	path = parameterRegex.ReplaceAllStringFunc(path, func(match string) string {
		match = strings.TrimPrefix(match, "s/")
		return "_by_" + strings.ToLower(match[1:len(match)-1])
	})

	// Replace method with create/update
	switch strings.ToUpper(method) {
	case http.MethodPost:
		method = "create"
	case http.MethodPatch:
		method = "update"
	}

	// Generate function name
	funcName := strings.ToLower(method) + "_" + strings.ReplaceAll(path, "/", "_")
	if version != "" {
		funcName = "v" + version + "_" + funcName
	}

	return strcase.ToSnake(funcName)
}

func (d *HackerRankApiDoc) GetPathSlugs(path string) []string {
	path = strings.SplitN(path, "?", 2)[0]

	submatches := parameterRegex.FindAllStringSubmatch(path, -1)
	slugs := make([]string, len(submatches))
	for i, submatch := range submatches {
		slugs[i] = submatch[2]
	}
	return slugs
}

type Info struct {
	Version        string       `json:"version"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	TermsOfService string       `json:"termsOfService"`
	Contact        *InfoContact `json:"contact"`
	License        *InfoLicense `json:"license"`
}

type InfoContact struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

type InfoLicense struct {
	Name string `json:"name"`
}

type SecurityDefinition struct {
	Type string `json:"type"`
	Name string `json:"name"`
	In   string `json:"in"`
}

type PathMethod struct {
	Summary     string                 `json:"summary"`
	Description string                 `json:"description"`
	Tags        []string               `json:"tags"`
	Parameters  []*PathMethodParameter `json:"parameters,omitempty"`
	Responses   *PathMethodResponses   `json:"responses,omitempty"`
}

func (m *PathMethod) DefinitionRequestBodyName() string {
	for _, param := range m.Parameters {
		if param.In == "body" {
			if param.Schema == nil {
				return ""
			}
			if param.Schema.Ref == "" {
				return ""
			}

			return strings.TrimPrefix(param.Schema.Ref, "#/definitions/")
		}
	}

	return ""
}

type PathMethodParameter struct {
	Name        string  `json:"name"`
	In          string  `json:"in"`
	Description string  `json:"description"`
	Required    bool    `json:"required,omitempty"`
	Type        string  `json:"type,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
	Ref         string  `json:"$ref,omitempty"`
}

func (p *PathMethodParameter) TrueType() string {
	return p.Type
}

type PathMethodResponses struct {
	Description *PathMethodResponsesDescription  `json:"description,omitempty"`
	Response200 *PathMethodResponsesResponseCode `json:"200,omitempty"`
}

type PathMethodResponsesResponseCode struct {
	Schema *Schema `json:"schema"`
}

type PathMethodResponsesDescription struct {
	Schema *Schema `json:"schema"`
}

type Schema struct {
	Ref        string                         `json:"$ref,omitempty"`
	Title      string                         `json:"title,omitempty"`
	Properties map[string]*DefinitionProperty `json:"properties,omitempty"`
	Required   []string                       `json:"required,omitempty"`
}

type DefinitionProperty struct {
	Type        string                  `json:"type"`
	Default     interface{}             `json:"default,omitempty"`
	Description string                  `json:"description,omitempty"`
	Enum        []interface{}           `json:"enum,omitempty"`
	Items       *DefinitionPropertyItem `json:"items,omitempty"`
	Required    bool                    `json:"required,omitempty"`
	Ref         string                  `json:"$ref,omitempty"`
}

func (p *DefinitionProperty) TrueType() string {
	if p.Default != nil {
		if _, ok := p.Default.(float64); ok {
			return "float"
		}
	}

	return p.Type
}

type DefinitionPropertyItem struct {
	Ref  string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}

func (i *DefinitionPropertyItem) TrueType() string {
	return i.Type
}
