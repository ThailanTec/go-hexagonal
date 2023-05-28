## Arquitetura Hexagonal

* A lógica da nossa aplicação vai ficar ao centro do código (pensando como um hexagono).
* Ao lado esquerdo vai ficar o "client", que é tudo que acessa a nossa aplicação. 
* Ao lado direto vai ficar o servidor (banco de dados, cloud e etc).

- Lado esquerdo{
    REST
    CLI
    GRPC
    GraphQL
    UI
}

- Lado direito{
    DB
    Redis(cache)
    Filesystem
    Lambda
    API(externa)
}

## Detalhes

* Modulos de alto nível não devem depender de módulos de baixo nível. Ambos devem depender de abstrações.

* Abstrações não devem depender de detalhes. Detalhes devem depender de abstrações.

## Por que utilizar hexagonal 

Isolar a logica da aplicação do que ela vai acessar ou quem vai acessar a mesma. 

## Service

é o cara que vai pegar a entidade e vai alterar ela no banco de dados. 

## Trabalhando com adapters

Eles serem como novas "tomadas" para entrada e saida de dados da nossa aplicação.


## Marshal, Unmarshal e NewEncoder

O Marshal e o Unmarshal eles pegam os dados e jogam dentro das variaveis. Enquando fazer com o NewEncoder, ele já pega a resposta que vem da aplicação e joga direto na variavel "acelerando" o processo.
```go

json.NewEncoder(w).Encode(product)

```

Então, já encodamos para Json e passamos para o usuario final a resposta.


## Trabalhando com DTO

Data transfer Object.

Objeto que se cria para transferir dados de uma camada para outra.