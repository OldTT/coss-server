// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplaterelation = `{
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
        "/relation/delete_blacklist": {
            "post": {
                "description": "删除黑名单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除黑名单",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.deleteBlacklistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/group/join": {
            "post": {
                "description": "加入群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "加入群聊",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.joinGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/group/manage_join_group": {
            "post": {
                "description": "管理加入群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "管理加入群聊",
                "parameters": [
                    {
                        "description": "Status (0: rejected, 1: joined)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.approveJoinGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/group/member": {
            "get": {
                "description": "群聊成员列表",
                "produces": [
                    "application/json"
                ],
                "summary": "群聊成员列表",
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
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/group/quit": {
            "post": {
                "description": "退出群聊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "退出群聊",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.quitGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/group/remove": {
            "post": {
                "description": "将用户从群聊移除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "将用户从群聊移除",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.removeUserFromGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/group/request_list": {
            "get": {
                "description": "群聊申请列表",
                "produces": [
                    "application/json"
                ],
                "summary": "群聊申请列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/http.requestListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/relation/user/add_blacklist": {
            "post": {
                "description": "添加黑名单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "添加黑名单",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.addBlacklistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/user/add_friend": {
            "post": {
                "description": "添加好友",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "添加好友",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.addFriendRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/user/blacklist": {
            "get": {
                "description": "黑名单",
                "produces": [
                    "application/json"
                ],
                "summary": "黑名单",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/user/delete_friend": {
            "post": {
                "description": "删除好友",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除好友",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.deleteFriendRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/user/friend_list": {
            "get": {
                "description": "好友列表",
                "produces": [
                    "application/json"
                ],
                "summary": "好友列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/user/manage_friend": {
            "post": {
                "description": "管理好友请求",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "管理好友请求",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.ManageFriendRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/user/request_list": {
            "get": {
                "description": "好友申请列表",
                "produces": [
                    "application/json"
                ],
                "summary": "好友申请列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/relation/user/switch/e2e/key": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "交换用户端到端公钥",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "交换用户端到端公钥",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.switchUserE2EPublicKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.ManageFriendRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "e2e_public_key": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.addBlacklistRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.addFriendRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "e2e_public_key": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.approveJoinGroupRequest": {
            "type": "object",
            "required": [
                "group_id",
                "user_id"
            ],
            "properties": {
                "group_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.deleteBlacklistRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.deleteFriendRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.joinGroupRequest": {
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
        "http.quitGroupRequest": {
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
        "http.removeUserFromGroupRequest": {
            "type": "object",
            "required": [
                "group_id",
                "user_id"
            ],
            "properties": {
                "group_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.requestListResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "request_at": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.switchUserE2EPublicKeyRequest": {
            "type": "object",
            "required": [
                "public_key",
                "user_id"
            ],
            "properties": {
                "public_key": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "utils.Response": {
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
        }
    }
}`

// SwaggerInforelation holds exported Swagger Info so clients can modify it
var SwaggerInforelation = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "coss-relation-bff服务",
	Description:      "",
	InfoInstanceName: "relation",
	SwaggerTemplate:  docTemplaterelation,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInforelation.InstanceName(), SwaggerInforelation)
}
