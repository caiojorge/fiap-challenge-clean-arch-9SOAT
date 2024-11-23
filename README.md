# Kitchen Control | fase 2

## Descrição

A Kitchen Control API é uma aplicação para gerenciar clientes, produtos, pedidos e itens de pedidos. Esta API fornece endpoints para criar, buscar, atualizar e deletar registros.

## Tecnologias usadas no projeto

- Go
- GORM
- MySQL
- Gin Web Framework
- Docker
- Kind

## Instalação e acesso ao swagger

### Pré-requisitos

- wsl 2, macos ou linux
- Go 1.22.1 ou superior
- Docker
- Git
- make
- Kind (ou minikube / docker desktop)

### Como instalar o Kitchen Control

1. Com o projeto baixado / clonado, instale as dependências:
    ```bash
    go mod tidy
    ```
2. Execute os testes:
    ```bash
    make test (ou go test -v -cover ./...)
    ```
3. Se estiver usando kind como eu, crie o custer

    ```bash
    make setup-cluster
    ```
4. Crie os configmaps

    ```bash
    make setup-configmap
    ```
5. Crie os deployments, services, pvc e secrets 

    ```bash
    make setup-k8s
    ```
6. Se necessário, acessar os logs

    ```bash
    make logs-k8s
    ```
7. Para desligar tudo
    ```bash
    make shutdown
    ```

## Como usar o Kitchen Control

### Endpoints (acesso via swagger)

8. Swagger:
    ```bash
    http://localhost:8080/kitchencontrol/api/v1/docs/index.html
    ```
### Client para DB - Adminer
9. Adminer
    ```bash
    http://localhost:8282/
    ```
### Gerador de CPF para os testes
10. CPFs
    ```bash
    https://www.geradordecpf.org/

## Acesso ao projeto no github
- Será enviado aos professores via plataforma da fiap

## Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

