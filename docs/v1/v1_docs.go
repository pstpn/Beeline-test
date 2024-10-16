// Package v1 Code generated by swaggo/swag. DO NOT EDIT
package v1

import "github.com/swaggo/swag"

const docTemplatev1 = `{
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
        "/healthcheck": {
            "get": {
                "description": "Проверка на жизнеспособность",
                "tags": [
                    "system"
                ],
                "summary": "Проверка здоровья",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/profiles": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Метод для создания профиля пользователя",
                "tags": [
                    "admin"
                ],
                "summary": "Создание профиля пользователя",
                "parameters": [
                    {
                        "description": "Создание профиля пользователя",
                        "name": "createProfileRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.createProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Профиль успешно создан",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректное тело запроса",
                        "schema": {
                            "$ref": "#/definitions/http.StatusBadRequest"
                        }
                    },
                    "401": {
                        "description": "Авторизация неуспешна",
                        "schema": {
                            "$ref": "#/definitions/http.StatusUnauthorized"
                        }
                    },
                    "403": {
                        "description": "Отсутствуют права доступа",
                        "schema": {
                            "$ref": "#/definitions/http.StatusForbidden"
                        }
                    }
                }
            }
        },
        "/profiles/{email}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Метод для просмотра профиля пользователя",
                "tags": [
                    "user"
                ],
                "summary": "Просмотр профиля пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email пользователя",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Профиль успешно получен",
                        "schema": {
                            "$ref": "#/definitions/user.getProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректное тело запроса",
                        "schema": {
                            "$ref": "#/definitions/http.StatusBadRequest"
                        }
                    },
                    "401": {
                        "description": "Авторизация неуспешна",
                        "schema": {
                            "$ref": "#/definitions/http.StatusUnauthorized"
                        }
                    },
                    "403": {
                        "description": "Отсутствуют права доступа",
                        "schema": {
                            "$ref": "#/definitions/http.StatusForbidden"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Метод для изменения профиля пользователя",
                "tags": [
                    "admin"
                ],
                "summary": "Изменение профиля пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email пользователя",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Изменение профиля пользователя",
                        "name": "updateProfileRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.updateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Профиль успешно изменен",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректное тело запроса",
                        "schema": {
                            "$ref": "#/definitions/http.StatusBadRequest"
                        }
                    },
                    "401": {
                        "description": "Авторизация неуспешна",
                        "schema": {
                            "$ref": "#/definitions/http.StatusUnauthorized"
                        }
                    },
                    "403": {
                        "description": "Отсутствуют права доступа",
                        "schema": {
                            "$ref": "#/definitions/http.StatusForbidden"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Метод для удаления профиля пользователя",
                "tags": [
                    "admin"
                ],
                "summary": "Удаление профиля пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email пользователя",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Профиль успешно удален",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректное тело запроса",
                        "schema": {
                            "$ref": "#/definitions/http.StatusBadRequest"
                        }
                    },
                    "401": {
                        "description": "Авторизация неуспешна",
                        "schema": {
                            "$ref": "#/definitions/http.StatusUnauthorized"
                        }
                    },
                    "403": {
                        "description": "Отсутствуют права доступа",
                        "schema": {
                            "$ref": "#/definitions/http.StatusForbidden"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Метод для регистрации пользователя",
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "Регистрация пользователя",
                        "name": "registerRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.registerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Пользователь успешно зарегистрирован",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректное тело запроса",
                        "schema": {
                            "$ref": "#/definitions/http.StatusBadRequest"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.createProfileRequest": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string",
                    "example": "Beeland"
                },
                "email": {
                    "type": "string",
                    "example": "beebee@bee.bee"
                },
                "name": {
                    "type": "string",
                    "example": "Bee"
                },
                "surname": {
                    "type": "string",
                    "example": "Beevich"
                }
            }
        },
        "admin.updateProfileRequest": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string",
                    "example": "NewBeeland"
                },
                "name": {
                    "type": "string",
                    "example": "Bee"
                },
                "surname": {
                    "type": "string",
                    "example": "Beevich"
                }
            }
        },
        "auth.registerRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "beebee@bee.bee"
                },
                "password": {
                    "type": "string",
                    "example": "123"
                }
            }
        },
        "http.StatusBadRequest": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Incorrect request body"
                }
            }
        },
        "http.StatusForbidden": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Incorrect user role"
                }
            }
        },
        "http.StatusUnauthorized": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Cant login user"
                }
            }
        },
        "user.getProfileResponse": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfov1 holds exported Swagger Info so clients can modify it
var SwaggerInfov1 = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Bee app",
	Description:      "",
	InfoInstanceName: "v1",
	SwaggerTemplate:  docTemplatev1,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfov1.InstanceName(), SwaggerInfov1)
}
