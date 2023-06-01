# Clean Architecture Go Implementation With Gin

### Executando o projeto

A forma recomendada de executar o projeto é através do `docker compose`.

```bash
docker compose up -d
```

O projeto executa automaticamente as migrações presentes no diretório `/migrations` da raiz.

### Executando via Docker

Construir a imagem:

```bash
docker build . -t teste-back:latest
```

Executar um container

```bash
docker run -d --name test-back-container teste-back:latest
```