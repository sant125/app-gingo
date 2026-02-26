# gin-tattoo

API de estilos e curiosidades de tatuagem, escrita em Go + Gin. Serve o frontend estático junto, expõe métricas Prometheus em `/metrics` e documenta os endpoints via Swagger.

## Rodando local

```bash
docker compose up
```

API em `http://localhost:8080`, Swagger em `http://localhost:8080/swagger/index.html`.

## Endpoints

| Método | Path | |
|--------|------|---|
| GET | `/health` | health check |
| GET | `/metrics` | métricas Prometheus |
| GET | `/api/v1/styles` | estilos de tatuagem |
| GET | `/api/v1/styles/:id` | detalhe de um estilo |
| GET | `/api/v1/curiosities` | curiosidades |
| GET | `/api/v1/curiosities/:id` | detalhe de uma curiosidade |
| GET | `/swagger/*` | Swagger UI |

## Env

| Variável | Exemplo |
|----------|---------|
| `DATABASE_URL` | `postgres://user:pass@host:5432/db` |
| `DB_SCHEMA` | `homolog` / `prod` |
| `GIN_MODE` | `release` |

## Testes

```bash
go test ./...
```

## CI/CD

GitHub Actions — `feature/*` roda lint, test e análise estática; `developer` builda, sobe pro ECR e atualiza o manifest de homolog; `main` vai pra prod.

Manifests e infra em [aws-devops](https://github.com/sant125/aws-devops).
