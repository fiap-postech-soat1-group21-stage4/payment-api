# Software Architecture - Tech Challenge

<details>

<summary>Entrega FASE 4 -Arquitetura de Microsserviços</summary>

# Software Architecture - FASE 4 - Tech Challenge

## Requisitos

|Recurso|Versão|Obrigatório|Nota|
|-|-|-|-|
|Docker Desktop| 4.21 ou mais atual|Sim|Necessário para rodar containers das APIs e banco de dados|
|Golang| 1.20|Não|Necessário apenas no caso de rodar localmente sem container|

## O que esse projeto faz e possui
### O que esse projeto faz
Através da API é possível simular o pagamento de um pedido.

#### O que esse projeto possui
 - [x] Dockerfile e DockerCompose

## O que esse projeto não faz e débitos técnicos
#### O que esse projeto não faz
- Não se comunica com outros microsserviços, além de não possuir banco para persistir o retorno do pagamento;

#### Débitos técnicos
- [ ] Comunicação com outras aplicações.
- [ ] Testes Unitários
- [ ] Comunicação com banco de dados

## Como executar o projeto
### Executar o projeto
É possivel executar o projeto através do Makefile, a partir da linha de comando. 
~~~bash
make run-project
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto

### Executar o Docker
Para executar o projeto, é necessário ter o `Docker Desktop` instalado. Com isso será possível criar as instancias usando o comando `docker compose` via IDE ou linha de comando conforme a seguir:
~~~bash
docker compose -f "docker-compose.yml" up -d --build
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto

### Utilizar Aplicação & Documentação API
1. Realiza pagamento de um pedido específico `[POST] localhost:8080/api/v1/payments/:id` 
