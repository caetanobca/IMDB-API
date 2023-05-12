# IMDB-API

### 1. Dependencias
* [Redis Server](https://redis.io/)
* [Golang](https://go.dev/)

### 2. Importando os dados  

Primeiro, será necessário baixar o [dump.rdb](https://drive.google.com/file/d/1drUkHa-SQ6z0UQWWa3j_aOtTLz9R5GVf/view?usp=sharing).

Uma vez com o dump, para importar os dados para o Redis, tudo o que você precisa fazer é parar o Redis e copiar o arquivo `dump.rdb` para o diretório do seu Redis. Para encontrar o caminho desse diretório você pode usar o seguinte comando no 
`redis-cli`:

```
CONFIG get dir
```

### 3. Rodando a aplicação

Uma vez estando no diretório `api`, primeiro será necessário instalar as dependências:

```
go mod download
```

Agora, basta rodar com o seguinte comando:

```
go run src/main.go
```
