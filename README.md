[![CI pipeline](https://github.com/fiap-postech-soat1-group21-stage4/payment-api/actions/workflows/github-ci.yml/badge.svg)](https://github.com/fiap-postech-soat1-group21-stage4/payment-api/actions/workflows/github-ci.yml) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=fiap-postech-soat1-group21-stage4_payment-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=fiap-postech-soat1-group21-stage4_payment-api) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=fiap-postech-soat1-group21-stage4_payment-api&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=fiap-postech-soat1-group21-stage4_payment-api) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=fiap-postech-soat1-group21-stage4_payment-api&metric=bugs)](https://sonarcloud.io/summary/new_code?id=fiap-postech-soat1-group21-stage4_payment-api)

[![CD pipeline](https://github.com/fiap-postech-soat1-group21-stage4/payment-api/actions/workflows/github-cd.yml/badge.svg)](https://github.com/fiap-postech-soat1-group21-stage4/payment-api/actions/workflows/github-cd.yml)

# Software Architecture - Tech Challenge

<details>

<summary>Entrega FASE 4 -Arquitetura de Microsserviços</summary>

## Requisitos

|Recurso|Versão|Obrigatório|Nota|
|-|-|-|-|
|Docker Desktop| >=4.21 |Sim|Necessário para rodar containers das APIs e banco de dados|
|SonarQube Cli| >=5.0 |Não|Apenas testes ShiftLeft locais|
|Golang| 1.21 |Não|Necessário apenas executar o projeto local e rodar testes|

## O que esse projeto faz e possui
### O que esse projeto faz
Através da API é possível gerar pagamentos de um pedido.

#### O que esse projeto possui
 - [x] Workflow/ Pipeline para Validação e Deploy da Imagem
 - [x] Dockerfile e DockerCompose
 - [x] Documentação para Consumo das API
 - [x] Testes Unitários e BDD

## O que esse projeto não faz e débitos técnicos
#### Débitos técnicos
- [ ] Comunicação integral com outras aplicações.
- [ ] Cobertura completa dos testes

## Como executar o projeto
### Criar Variáveis de Ambiente
Criar um arquivo nomedo como `.env` na raiz do projeto contendo os seguintes valores.
~~~bash
API_PORT=8080
API_HOST_PORT=8080
API_CONTAINER_PORT=8080
~~~

### Executar o projeto
É possivel executar o projeto através do Makefile, a partir da linha de comando.
~~~bash
make run-project
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto e deve possuir o .env com as variáveis preenchidas

### Executar o Docker
Para executar o projeto, é necessário ter o `Docker Desktop` instalado. Com isso será possível criar as instancias usando o comando `docker compose` via IDE ou linha de comando conforme a seguir:
~~~bash
docker compose -f "docker-compose.yml" up -d --build
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto e deve possuir o .env com as variáveis preenchida

### Executar testes
Para executar os testes contidos nos projetos, poderá ser aplicado o seguinte comando:
~~~bash
go test -coverprofile=coverage.out ./... ;    go tool cover -func=coverage.ou
~~~
Notas: o comando executará todos os testes e gerará um relatório de cobertura dos testes

### Utilizar Aplicação & Documentação API
1. Realiza pagamento de um pedido específico `[POST] localhost:8080/api/v1/payments/:id` 
