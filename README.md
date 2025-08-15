# Lac Auction

## 🚀 Tecnologias Utilizadas

- **Go** (>= 1.23)
- **Docker** & **Docker Compose**

---

## ⚙️ Configuração do Ambiente

1. **Clonar o repositório**
   ```bash
   git clone https://github.com/aserafim/labs-auction-goexpert.git

2. **Subir os serviços com Docker**

   ```bash
   docker compose up --build
   ```

   Isso irá:

   * Compilar e iniciar a API na porta **8080**
   * Compilar e iniciar o mongodb na porta **27017**

---

## 📬 Como Utilizar

### 1. Enviar uma requisição para API

**Obs.: O tempo padrão para encerramento de uma auction é de 20s**

```bash
curl -X POST http://localhost:8080/auction \
-H "Content-Type: application/json" \
-d '{
  "product_name": "Produto Teste",
  "category": "Eletrônicos",
  "description": "Leilão de teste",
  "condition": 1
}'
```
Isso irá criar uma nova auction com os parâmetros utilizados.

---

### 2. Enviar uma requisição para listar as auctions ativas

```bash
curl http://localhost:8080/auction?status=0
```


**Resultado esperado**

```bash
[{"id":"11ad686f-e8c8-490c-9fd7-502cdd56cd87","product_name":"Produto Teste","category":"Eletrônicos","description":"Leilão de teste","condition":1,"status":0,"timestamp":"2025-08-15T00:58:25Z"}]
```

**Obs.: O tempo padrão para encerramento de uma auction é de 20s**

Após 20s, o mesmo comando retornará
```bash
null
```

### 2. Listar auctions concluídas

```bash
curl http://localhost:8080/auction?status=1
```

**Resultado esperado**

```bash
[{"id":"11ad686f-e8c8-490c-9fd7-502cdd56cd87","product_name":"Produto Teste","category":"Eletrônicos","description":"Leilão de teste","condition":1,"status":1,"timestamp":"2025-08-15T00:58:25Z"}]
```


## 📄 Licença

Este projeto é distribuído sob a licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.

```

