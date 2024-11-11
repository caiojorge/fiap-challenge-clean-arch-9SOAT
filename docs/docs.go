// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/checkouts": {
            "post": {
                "description": "Efetiva o pagamento do cliente, via fake checkout nesse momento, e libera o pedido para preparação. A ordem muda de status nesse momento, para em preparação.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Checkouts"
                ],
                "summary": "Create Checkout",
                "parameters": [
                    {
                        "description": "cria novo Checkout",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.CheckoutInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.CheckoutOutputDTO"
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customers": {
            "get": {
                "description": "Get details of all customers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Get all customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usecase.CustomerFindAllOutputDTO"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create Customer in DB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Create Customer",
                "parameters": [
                    {
                        "description": "cria novo cliente",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.CustomerRegisterInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.CustomerRegisterOutputDTO"
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "409": {
                        "description": "customer already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/customers/{cpf}": {
            "get": {
                "description": "Get details of a customer by cpf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Get a customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer cpf",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.CustomerFindByCpfOutputDTO"
                        }
                    },
                    "404": {
                        "description": "Customer not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update details of a customer by cpf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Update a customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer cpf",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer data",
                        "name": "Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.CustomerUpdateInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.CustomerUpdateOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Customer not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/kitchens/orders": {
            "get": {
                "description": "Retorna todos os pedidos (orders) que estão na cozinha para inicio de preparação. Se não houver pedidos, retorna um erro (404).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kitchens"
                ],
                "summary": "Get all orders in the kitchen",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usecase.KitchenFindAllAOutputDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "Retorna todos os pedidos (orders) registrados no sistema. Se não houver pedidos, retorna um erro (404).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usecase.OrderFindAllOutputDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um peddo (order) no sistema. O cliente (customer) pode ou não de identificar. Se o cliente não se identificar, o pedido será registrado como anônimo. O produto, porém, deve ter sido previamente cadastrado.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": "cria nova Order",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.OrderCreateInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.OrderCreateOutputDTO"
                        }
                    },
                    "400": {
                        "description": "invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Order already exists",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/paid": {
            "get": {
                "description": "Retorna todos os pedidos (orders) registrados no sistema. Se não houver pedidos, retorna um erro (404).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get all paid orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usecase.OrderFindByParamOutputDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Get details of a Order and their items by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get a Order by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.OrderFindByIdOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get details of all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usecase.FindAllProductOutputDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalida data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "No products foundr",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new product in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "New Product Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.RegisterProductInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created",
                        "schema": {
                            "$ref": "#/definitions/usecase.RegisterProductOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid data format or missing fields",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Product already exists",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/category/{id}": {
            "get": {
                "description": "Get details of a Product by category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get a Product by category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product category",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usecase.FindProductByCategoryOutputDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Product not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get details of a Product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get a Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.FindProductByIDOutputDTO"
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update details of a Product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update a Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product data",
                        "name": "Product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.UpdateProductInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product updated",
                        "schema": {
                            "$ref": "#/definitions/usecase.UpdateProductOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete details of a Product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete a Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "shared.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "usecase.CheckoutInputDTO": {
            "type": "object",
            "properties": {
                "customer_cpf": {
                    "type": "string"
                },
                "gateway": {
                    "type": "string"
                },
                "gateway_id": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                }
            }
        },
        "usecase.CheckoutOutputDTO": {
            "type": "object",
            "properties": {
                "gateway_transaction_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "usecase.CustomerFindAllOutputDTO": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usecase.CustomerFindByCpfOutputDTO": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usecase.CustomerRegisterInputDTO": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usecase.CustomerRegisterOutputDTO": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usecase.CustomerUpdateInputDTO": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usecase.CustomerUpdateOutputDTO": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usecase.FindAllProductOutputDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "usecase.FindProductByCategoryOutputDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "usecase.FindProductByIDOutputDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "usecase.KitchenFindAllAOutputDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "item_order_id": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "product_name": {
                    "type": "string"
                },
                "responsible": {
                    "type": "string"
                }
            }
        },
        "usecase.OrderCreateInputDTO": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.OrderItemCreateInputDTO"
                    }
                }
            }
        },
        "usecase.OrderCreateOutputDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customercpf": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.OrderItemDTO"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "usecase.OrderFindAllOutputDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customercpf": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.OrderItemDTO"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "usecase.OrderFindByIdOutputDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customercpf": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.OrderItemDTO"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "usecase.OrderFindByParamOutputDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customercpf": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.OrderItemDTO"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "usecase.OrderItemCreateInputDTO": {
            "type": "object",
            "properties": {
                "productid": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "usecase.OrderItemDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "productid": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "usecase.RegisterProductInputDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "usecase.RegisterProductOutputDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "usecase.UpdateProductInputDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "usecase.UpdateProductOutputDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/kitchencontrol/api/v1",
	Schemes:          []string{},
	Title:            "Fiap Fase 2 Challenge Clean Arch API - 9SOAT",
	Description:      "This is fiap fase 2 challenge project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
