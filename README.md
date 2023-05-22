# Pismo Back Teste 3.0

Implementação do teste de backend da Pismo.

### Executando o projeto

A forma recomendada de executar o projeto é através do `docker compose`.

### Executando via Docker

Construir a imagem:

```bash
docker build . -t teste-back:latest
```

Executar um container

```bash
docker run -d --name test-back-container teste-back:latest
```