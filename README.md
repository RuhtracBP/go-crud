
# Go Crud

Pequena Api feita em Go


## Instalação

Instale my-project com npm

```bash
    CompileDaemon -command="./go-crud" 
```
    
## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar adicionar as seguintes variáveis de ambiente no seu .env

`PORT`

`DB_URL` (Postgres URL )

## Documentação da API

#### Retorna todos os Posts

```http
  GET /posts/items
```



#### Retorna um Post

```http
  GET /posts/${id}
```

#### Cria um Post

```http
  POST /posts
```

| Req   | Tipo          |
| :---------- | :----   |
| `title`      | `string`  |
| `body`      | `string`  |

#### Update um Post

```http
  PUT /posts/${id}
```
| Req   | Tipo          |
| :---------- | :----   |
| `title`      | `string`  |
| `body`      | `string`  |

#### Deleta um Post

```http
  DELETE /posts/${id}
```








## Aprendizados

 - Foi utilizado o [Compile Daemon](https://github.com/githubnemo/CompileDaemon) para rodar e ter atualização automática
 - Também o [godotenv](https://github.com/joho/godotenv) para guardar e ter fácil acesso às variáveis de ambiente
 - para servir as rotas o [Gin](https://gin-gonic.com/)
 - e o [GORM](https://gorm.io/) como Blibioteca ORM
 - o Banco de Dados utilizado foi o [Postgres](https://www.postgresql.org/)
