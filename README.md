# Desafio GOLang Clean Architecture FullCycle 

Aplicação em Go sendo: 
  - Servidor HTTP Rest
  - Servidor GraphQL 
  - Servidor gRPC 
&nbsp;
- **rodando em Multithreads com Clean Archictecture**

## Funcionalidades

- **Cadastro de Orders:**
  - O servidor permite cadastrar Ordens utilizando HTTP, GraphQL e gRPC.

- **Listagem de Orders:**
  - O servidor permite listar todas as Ordens utilizando HTTP, GraphQL e gRPC.


## Como Utilizar localmente:

1. **Requisitos:** 
   - Certifique-se de ter o Go instalado em sua máquina.
   - Certifique-se de ter Evans e gRPC instalado em sua máquina.
   - Certifique-se de ter GraphQL (gqlgen) instalado em sua máquina.
   - Certifique-se de ter o Docker instalado em sua máquina.
&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/tiago-g-sales/clean-arch.git
```
&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd clean-arch
```
&nbsp;
4. **Rode o docker para subir o serviço RabbitMQ e MySQL:**
&nbsp;

```bash 
docker-compose up -d
```
5. **Acesse a pasta cmd/orderssystem e rode o main.go:**
&nbsp;

```bash 
cd cmd/ordersystem
```

```bash 
go run main.go
```

Observação: Ao iniciar a aplicação, o processo de migração é executado automaticamente, não sendo necessário realizá-lo manualmente.

## Como testar localmente:

### Portas
HTTP server on port :8000 <br />
gRPC server on port :50051 <br />
GraphQL server on port :8080

### HTTP
 - Acesse a pasta api/ e dispare os dois arquivos:
 - create_order.http
 - list_order.http
** Necessário instalar o plugin REST Client no VSCode. **

** Ou utilizar o arquivo Makefile para execução dos comandos. **
 - Executar o comando make execute_POST 
** Para inserir 5 order. **

 - Executar o comando make execute_GET
** Para listar todas as order. ** 

### GraphQL
 - Abra a página do GraphQL na porta 8080 e execute a mutation ou query abaixo:
 <a href="http://localhost:8080/" target="_blank">http://localhost:8080/</a>

 ```graphql
    mutation CreateOrder {
      createOrder(input: {id: "order 6",Price: 1.0, Tax: 3.0}){
        id
        Price
        Tax
        FinalPrice
      }
    }


     mutation CreateOrder {
      createOrder(input: {id: "order 7",Price: 3.0, Tax: 2.0}){
        id
        Price
        Tax
        FinalPrice
      }
    }

    mutation CreateOrder {
      createOrder(input: {id: "order 8",Price: 1.0, Tax: 3.0}){
        id
        Price
        Tax
        FinalPrice
      }
    }

    mutation CreateOrder {
      createOrder(input: {id: "order 9",Price: 13.0, Tax: 2.3}){
        id
        Price
        Tax
        FinalPrice
      }
    }

    mutation CreateOrder {
      createOrder(input: {id: "order 10",Price: 15.0, Tax: 3.1}){
        id
        Price
        Tax
        FinalPrice
      }
    }


    query ListOrders {
      ListOrders{
        id
        Price
        Taxevans 
        FinalPrice
      }
    }

 ```

### gRPC
 - Rode o evans:

```bash
evans -r repl
```
```bash
package pb
```
```bash
service OrderService
```

Para criar orders utilize
```bash
call CreateOrder 
```

Para listar as orders utilize
```bash
call ListOrders
```

