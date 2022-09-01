# Library

CRUD desenvolvido para o desafio técnico da thagos. Nessa aplicação é possível cadastrar, editar, excluir e listar livros.

## 🚀 Começando

Essas instruções permitirão que você obtenha uma cópia do projeto em operação na sua máquina local para fins de desenvolvimento e teste.

## 🔧 Instalação

Uma série de exemplos passo-a-passo que informam o que você deve executar para ter um ambiente de desenvolvimento em execução.

Clone o repositório:

```
git clone git@github.com:matheusgb/crud-library.git && cd crud-library
```
Renomeie os arquivos `config.toml.example` e `.env.example` para `config.toml` e `.env` <br>
<br>
Rode o comando:
```
docker compose up -d --build
```

E inicialize a aplicação Golang:

```
go run main.go
```

## 📌  Rotas e retornos dos endpoints

`GET /` Lista todos os livros 

```
Retorno:

[
	{
		"id": 1,
		"title": "Título Genérico",
		"category": "Romance",
		"author": "Value Man",
		"synopsis": "Um livro com sinopse"
	},
	{
		"id": 2,
		"title": "Outro Título",
		"category": "Fantasia",
		"author": "Homem Valor",
		"synopsis": "Outro livro com sinopse"
	}
]
```
<br />

`GET /{id}` Lista o livro do id correspondente

```
Retorno:

{
	"id": 1,
	"title": "Título Genérico",
	"category": "Romance",
	"author": "Value Man",
	"synopsis": "Um livro com sinopse"
}
```
<br />

`POST /` Adiciona um livro

```
Corpo esperado da requisição:

{	
	"title": "Título Qualquer",
	"category": "Aventura",
	"author": "Autor desconhecido",
	"synopsis": "Um livro misterioso"
}


Retorno:

{
	"Error": false,
	"Message": "Book inserted! ID: 3"
}

```
<br />

`PUT /{id}` Edita um livro pelo id

```
Corpo esperado da requisição:

{	
	"title": "Título editado",
	"category": "Fantasia",
	"author": "Autor editado",
	"synopsis": "Um livro editado"
}


Retorno:

{
	"Error": false,
	"Message": "Successfully updated!"
}

```

<br />

`DELETE /{id}` Deleta um livro pelo id

```

Retorno:

{
	"Error": false,
	"Message": "Successfully deleted!"
}

```

## 🖇️ Considerações sobre o projeto

Primeiramente gostaria de agradecer ao Leonardo Magalhães e ao Christian Bittencort pela oportunidade de realizar esse desafio para a Thagos, gostei muito de desenvolver em Golang e do clima da conversa que tivemos, espero que eu atenda as expectativas para que ambos fatores se tornem rotineiros 😁.

#

### Sobre o desenvolvimento

Foram três dias de desenvolvimento e minha primeira vez lidando com Golang e Postgres.
No desafio havia a possibilidade de escolher um banco de dados SQL (Postgres) ou noSQL (MongoDB), optei pela primeira opção devido a afinidade com MYSQL que desenvolvi durante estudos na Trybe. MongoDB ainda estou aprendendo.

Durante a execução do projeto aprendi bastante, nos primeiros dias de desenvolvimento busquei entender fundamentos da linguagem Golang para conseguir compreender melhor conteúdos diversos que pesquisei enquanto me deparava com alguma dúvida.

Entendi o funcionamento do `go.mod` e consegui fazer uso de packages externos como `viper` (para setar variáveis de ambiente) e `chi` (para setar endpoints da API).

Tracei algumas semelhanças de Golang com Typescript, as tipagens, funções e interfaces funcionam de forma semelhante, isso trouxe certo conforto e melhor entendimento durante o projeto.

Ter utilizado o `docker-compose` para criar o banco de dados, tabelas e popular as tabelas facilitou muito desenvolvimento.

#

### Pontos de melhoria

Por se tratar do primeiro contato com a liguagem e poucos dias de estudo, ainda há conceitos que não entendi completamente como: ponteiros, arquitetura, clean code, tratamento de erros e implementação de http codes no retorno das endpoints.

---
⌨️ por [Matheus Gomes](https://www.linkedin.com/in/matheusgb/)
 