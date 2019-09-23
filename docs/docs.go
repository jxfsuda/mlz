// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-23 11:28:10.586285 +0800 CST m=+0.082494627

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
        "termsOfService": "http://git.vs9.cn",
        "contact": {
            "name": "jif",
            "url": "http://www.swagger.io/support",
            "email": "35802713@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/demo/index": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "演示",
                "parameters": [
                    {
                        "description": "参数对象,注意,此参数应该被包含在通用参数的data属性内",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/req.BaseSettingReqVO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0000,\"data\":{},\"message\":\"\",\"success\":true}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "req.BaseSettingReqVO": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "记录id",
                    "type": "string"
                },
                "keyer": {
                    "description": "key,代码定义",
                    "type": "string"
                },
                "name": {
                    "description": "设置名称",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "val": {
                    "description": "设置值",
                    "type": "string"
                }
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
	Version:     "1.0b125",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Gin API123",
	Description: "Golang API 演示, 主要是促成代码生成和数据库以及缓存使用,编写简单易于开发业务的框架 <br>技术栈:GOLang 1.12 ,  GIN , XORM , mysql , redis",
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
