# 🔗 URL Shortener API (Go + Chi)

Este projeto é uma **API de encurtador de URLs** desenvolvida em **Go (Golang)** utilizando o framework **Chi**.  
Ela permite criar códigos curtos para URLs e redirecionar automaticamente quando o código é acessado.

---

## 🚀 Funcionalidades 

- Criar um código curto para uma URL (`POST /api/shorten`)  
- Redirecionar automaticamente para a URL original (`GET /{code}`)  
- Tratamento de erros com respostas em **JSON**  
- Middleware do **Chi** para:
  - Logs
  - ID de requisição
  - Recuperação de erros (panic recovery)

---

## 📦 Tecnologias utilizadas

- [Go](https://go.dev/)
- [Chi](https://github.com/go-chi/chi) (roteamento e middlewares)
- [slog](https://pkg.go.dev/log/slog) (logs estruturados)

---

## 📂 Estrutura do Código

- **NewHandler** → Cria o roteador e registra os endpoints.
- **handlePost** → Recebe a URL em JSON, gera um código aleatório e retorna no formato:
  ```json
  {
    "data": "abcXYZ12"
  }

###
1.handleGet → Recebe um código e redireciona para a URL original.
2.genCode → Gera um código curto aleatório com 8 caracteres.
3.sendJSON → Envia respostas JSON padronizadas.



### ▶️ Como rodar o projeto

1. Clone este repositório:
```
git clone https://github.com/Whofelisberto/API-shorten-url.git
cd API-shorten-url
```

2.Instale as dependências:
 ```
go mod tidy 
 ```

3.Rode a aplicação:
 ```
go run main.go 
 ```

4.A API estará disponível em:
 ```
http://localhost:8080 
 ```
