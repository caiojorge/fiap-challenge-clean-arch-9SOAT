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
- Kind v0.25.0 (ou minikube / docker desktop)

### Como instalar o Kitchen Control

- obs: A imagem do projeto esta no dockerhub; não é necessário gerar nova imagem.

1. Com o projeto baixado / clonado, instale as dependências:
    ```bash
    go mod tidy
    ```
2. Execute os testes:
    ```bash
    make test (ou go test -v -cover ./...)
    ```
3. Se estiver usando kind (v0.25.0) como eu, crie o custer

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
    - esse comando irá deletar o cluster e os dados no DB serão perdidos 
## Como usar o Kitchen Control

### Endpoints (acesso via swagger)

8. Swagger:
    ```bash
    http://localhost:30080/kitchencontrol/api/v1/docs/index.html
    ```
### Client para DB - Adminer
9. Adminer
    ```bash
    http://localhost:30000/
    ```

    - usuario: root
    - senha: root
    - servidor: mysql
    - db: dbcontrol

### Gerador de CPF para os testes
10. CPFs
    ```bash
    https://www.geradordecpf.org/

## Requisitos da Fase 2

### #1 Atualizar a aplicação desenvolvida na FASE 1 refatorando o código para seguir os padrões Clean Code e Clean Architecture

#### Todo código foi migrado para novos padrões de pastas que atendem melhor a ideia da clean arch (vejo que a arquitetura do projeto grita / deixa claro o que esta acontecendo)

```plaintext
internal/
├── domain/
│   └── (Entidades, objetos de valor, interafaces para repositório e gateway)
│
├── infraestructure/
│   └── (Implementações específicas de tecnologias, como repositórios, gateways e controllers)
│
├── shared/
│   └── (Pacotes compartilhados, helpers e utilitários usados em várias partes do projeto)
│
└── usecase/
    └── (Casos de uso que orquestram a lógica de negócios da aplicação e seus objetos de dados (dto's) de input e output)

```
 
 Também adaptei o código para o estilo clean arch; seguem alguns detalhes:

##### Domain
- DDD: Escolhi manter o DDD no centro do projeto. A clean arch não proíbe o uso da DDD, e define apenas a camada de entidade e que ali temos as regras de negócios do sistema. Entendi que a ideia é evoluir o projeto que construimos na FASE 1, sendo assim, decidi manter a camada de dominio.

- Gateway e Repositórios: Mesma coisa no contexto desse sistema. Decidi usar Gateway para o acesso à API de pagamentos de manter Repository para os acessos ao DB. Estou mantendo o conceito do DDD, então me agrada usar os repositórios com os agregadores. Nessa camada estou mantendo apenas as interfaces.

- Entities, agregates e value objects: Escolhi os agregadores depois de testar e repensar o funcionamento do sistema. As entidades / regras criadas representam as regras do sistema, bem como seu vocabulário comum (linguagem ubíqua) 

##### Usecases
- Usecases: É a camada de regras da aplicação; é o orquestrador e suporta a necessidade especifica do usuário, por isso, o nome é caso de uso. Escolhi usar um arquivo por caso de uso (acho mais fácil manter o caso de uso com apenas 1 responsabilidade) e também escolhi manter inputs e outputs espeficios por caso de uso. Não reuso os DTO's de entrada e saída. Eles são focados na regra de aplicação que irão suportar.

##### infraestructure

- Driven: É o output, e aqui deixamos a conexão com o DB e ou APIs externas
- Driver: É o input, e aqui temos os controlers e o server/api

### #2 refatorando o código para seguir os padrões Clean Code e Clean Architecture
### #3 Checkout Pedido que deverá receber os produtos solicitados e retornar à identificação do pedido.
### #4 Consultar status de pagamento pedido, que informa se o pagamento foi aprovado ou não.



## Acesso ao projeto no github
- Será enviado aos professores via plataforma da fiap
- https://github.com/caiojorge/fiap-challenge-clean-arch-9SOAT

## Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.




