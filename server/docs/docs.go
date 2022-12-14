// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/sysTable/getDBTableInfos": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "代码生成器"
                ],
                "summary": "获取当前数据库所有表信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.JsonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.TableInfoResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/sysTable/importTables": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "代码生成器"
                ],
                "summary": "导入表",
                "parameters": [
                    {
                        "description": "表名",
                        "name": "tableNames",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/sysTable/previewCode": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "代码生成器"
                ],
                "summary": "预览代码",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "表名",
                        "name": "tableId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/sysUser/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SysUserCreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/sysUser/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.IdInfoDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/sysUser/detail": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id编号",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.JsonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.SysUserVo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/sysUser/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.JsonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.SysUserVo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/sysUser/page": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户分页查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户昵称",
                        "name": "nickName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.JsonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/response.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/dto.SysUserVo"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/sysUser/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SysUserUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/test/demo": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "测试"
                ],
                "summary": "测试demo",
                "parameters": [
                    {
                        "description": "表名",
                        "name": "tableNames",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.SysUserCreateDto": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string"
                },
                "avatar": {
                    "description": "头像地址",
                    "type": "string"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号码",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "sex": {
                    "description": "用户性别（0未知，1男，2女）",
                    "type": "integer"
                },
                "status": {
                    "description": "帐号状态（0正常 1停用）",
                    "type": "integer"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string"
                },
                "userType": {
                    "description": "用户类型（0普通账号，1超级管理员）",
                    "type": "integer"
                }
            }
        },
        "dto.SysUserUpdateDto": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string"
                },
                "avatar": {
                    "description": "头像地址",
                    "type": "string"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "编号",
                    "type": "integer"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号码",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "sex": {
                    "description": "用户性别（0未知，1男，2女）",
                    "type": "integer"
                },
                "status": {
                    "description": "帐号状态（0正常 1停用）",
                    "type": "integer"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string"
                },
                "userType": {
                    "description": "用户类型（0普通账号，1超级管理员）",
                    "type": "integer"
                }
            }
        },
        "dto.SysUserVo": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "用户账号",
                    "type": "string"
                },
                "avatar": {
                    "description": "头像地址",
                    "type": "string"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号码",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "sex": {
                    "description": "用户性别（0未知，1男，2女）",
                    "type": "integer"
                },
                "status": {
                    "description": "帐号状态（0正常 1停用）",
                    "type": "integer"
                },
                "userType": {
                    "description": "用户类型（0普通账号，1超级管理员）",
                    "type": "integer"
                }
            }
        },
        "dto.TableInfoResponse": {
            "type": "object",
            "properties": {
                "columnList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysGenTableColumn"
                    }
                },
                "createTime": {
                    "type": "string"
                },
                "functionName": {
                    "type": "string"
                },
                "modelName": {
                    "type": "string"
                },
                "paramName": {
                    "type": "string"
                },
                "tableComment": {
                    "type": "string"
                },
                "tableName": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                }
            }
        },
        "model.SysGenTableColumn": {
            "type": "object",
            "properties": {
                "columnComment": {
                    "type": "string"
                },
                "columnName": {
                    "type": "string"
                },
                "columnType": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "goField": {
                    "type": "string"
                },
                "goType": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isEdit": {
                    "type": "boolean"
                },
                "isList": {
                    "type": "boolean"
                },
                "isPk": {
                    "type": "boolean"
                },
                "isQuery": {
                    "type": "boolean"
                },
                "jsonField": {
                    "type": "string"
                },
                "queryMethod": {
                    "type": "string"
                },
                "tableId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "request.IdInfoDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "id编号",
                    "type": "integer"
                }
            }
        },
        "response.JsonResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "message": {
                    "description": "消息",
                    "type": "string"
                },
                "success": {
                    "description": "是否成功",
                    "type": "boolean"
                }
            }
        },
        "response.PageResult": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "数据"
                },
                "totalCount": {
                    "description": "总条数",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "系统模块",
	Description:      "gin-fast-admin 接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
