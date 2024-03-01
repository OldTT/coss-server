// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplategroup = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/group/create": {
            "post": {
                "description": "创建群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "创建群聊",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/group/delete": {
            "post": {
                "description": "删除群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除群聊",
                "parameters": [
                    {
                        "description": "群聊id",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DeleteGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/group/getBatch": {
            "get": {
                "description": "批量获取群聊信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "批量获取群聊信息",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "群聊ID列表",
                        "name": "group_ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/group/info": {
            "get": {
                "description": "获取群聊信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取群聊信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "群聊ID",
                        "name": "group_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/group/update/": {
            "post": {
                "description": "更新群聊信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新群聊信息",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateGroupRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "member": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "model.DeleteGroupRequest": {
            "type": "object",
            "required": [
                "group_id"
            ],
            "properties": {
                "group_id": {
                    "type": "integer"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.UpdateGroupRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfogroup holds exported Swagger Info so clients can modify it
var SwaggerInfogroup = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "coss-user服务",
	Description:      "",
	InfoInstanceName: "group",
	SwaggerTemplate:  docTemplategroup,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfogroup.InstanceName(), SwaggerInfogroup)
}
