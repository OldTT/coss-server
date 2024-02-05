// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplatelive = `{
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
        "/live/user/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "创建用户通话",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "创建用户通话",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserCallRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserCallResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/live/user/join": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "加入通话",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "加入通话",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserJoinRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/live/user/leave": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "结束通话",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "结束通话",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLeaveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/live/user/reject": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "拒绝通话",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "拒绝通话",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRejectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/live/user/show": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取通话房间信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "liveUser"
                ],
                "summary": "获取通话房间信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "房间名",
                        "name": "room",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "participant=通话参与者",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserShowResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CallOption": {
            "type": "object",
            "properties": {
                "audio_enabled": {
                    "description": "是否启用音频",
                    "type": "boolean"
                },
                "codec": {
                    "description": "编解码器",
                    "type": "string"
                },
                "frame_rate": {
                    "description": "帧率",
                    "type": "integer"
                },
                "resolution": {
                    "description": "分辨率",
                    "type": "string"
                },
                "video_enabled": {
                    "description": "是否启用视频",
                    "type": "boolean"
                }
            }
        },
        "dto.ParticipantInfo": {
            "type": "object",
            "properties": {
                "identity": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "is_publisher": {
                    "description": "创建者",
                    "type": "boolean"
                },
                "joined_at": {
                    "description": "加入时间",
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sid": {
                    "description": "房间id",
                    "type": "string"
                },
                "state": {
                    "description": "用户状态",
                    "allOf": [
                        {
                            "$ref": "#/definitions/dto.ParticipantState"
                        }
                    ]
                }
            }
        },
        "dto.ParticipantState": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-comments": {
                "ParticipantInfo_ACTIVE": "双方都已加入通话",
                "ParticipantInfo_DISCONNECTED": "断开连接",
                "ParticipantInfo_JOINED": "已加入通话，对方未响应",
                "ParticipantInfo_JOINING": "websocket已连接，未加入通话"
            },
            "x-enum-varnames": [
                "ParticipantInfo_JOINING",
                "ParticipantInfo_JOINED",
                "ParticipantInfo_ACTIVE",
                "ParticipantInfo_DISCONNECTED"
            ]
        },
        "dto.Response": {
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
        "dto.UserCallRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "option": {
                    "$ref": "#/definitions/dto.CallOption"
                },
                "user_id": {
                    "description": "接收视频通话的用户ID",
                    "type": "string"
                }
            }
        },
        "dto.UserCallResponse": {
            "type": "object",
            "properties": {
                "room": {
                    "description": "房间名称",
                    "type": "string"
                },
                "room_id": {
                    "description": "房间id",
                    "type": "string"
                },
                "token": {
                    "description": "加入通话的token",
                    "type": "string"
                },
                "url": {
                    "description": "webRtc服务器地址",
                    "type": "string"
                }
            }
        },
        "dto.UserJoinRequest": {
            "type": "object",
            "required": [
                "room"
            ],
            "properties": {
                "option": {
                    "$ref": "#/definitions/dto.CallOption"
                },
                "room": {
                    "description": "房间名称",
                    "type": "string"
                }
            }
        },
        "dto.UserLeaveRequest": {
            "type": "object",
            "required": [
                "room"
            ],
            "properties": {
                "room": {
                    "description": "房间名称",
                    "type": "string"
                }
            }
        },
        "dto.UserRejectRequest": {
            "type": "object",
            "required": [
                "room"
            ],
            "properties": {
                "room": {
                    "description": "房间名称",
                    "type": "string"
                }
            }
        },
        "dto.UserShowResponse": {
            "type": "object",
            "properties": {
                "duration": {
                    "description": "房间持续时间",
                    "type": "integer"
                },
                "participant": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ParticipantInfo"
                    }
                },
                "room": {
                    "description": "房间名称",
                    "type": "string"
                },
                "start_at": {
                    "description": "创建房间时间",
                    "type": "integer"
                },
                "video_call_record_url": {
                    "type": "string"
                },
                "video_call_status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfolive holds exported Swagger Info so clients can modify it
var SwaggerInfolive = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "",
	InfoInstanceName: "live",
	SwaggerTemplate:  docTemplatelive,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfolive.InstanceName(), SwaggerInfolive)
}