// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-04 16:44:34.472196615 +0800 CST m=+0.036664922

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/info": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查看账户信息"
            }
        },
        "/api/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "登录"
            }
        },
        "/api/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "注册"
            }
        },
        "/v1/todo": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取待办事项列表"
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增一个待办事项"
            }
        },
        "/v1/todo/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "更改指定待办事项状态"
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除一个待办事项"
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "myList API",
	Description: "一个todo清单后端，支持多用户注册，登录，每个用户有独立的清单",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}