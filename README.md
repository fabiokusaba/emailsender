# 📧 Email Sender - API

Aplicação backend para **envio e gerenciamento de campanhas por e-mail**, desenvolvida seguindo boas práticas de REST API, arquitetura limpa e foco em escalabilidade, segurança e testes.

O sistema permite criar campanhas, processar envios de forma assíncrona via filas e garantir autenticação e autorização através do **Keycloak**.

## 🚀 Tecnologias Utilizadas

- **Golang**

- **Chi** – Router HTTP leve e idiomático

- **PostgreSQL** – Banco de dados relacional

- **Docker & Docker Compose** – Conteinerização e ambiente de desenvolvimento

- **Keycloak** – Autenticação e autorização (OAuth2 / OpenID Connect)

- **Filas (Message Queue)** – Processamento assíncrono de envios

- **Testes Unitários** – Garantia de qualidade e confiabilidade

## 🔐 Autenticação e Autorização

A segurança da API é garantida através do Keycloak, utilizando OAuth2 / OpenID Connect.

- Autenticação via **JWT Bearer Token**

- Validação de token através de middleware

- Controle de acesso por roles e scopes

## 📬 Processamento de Campanhas

O envio de e-mails é feito de forma assíncrona, garantindo performance e escalabilidade:

1. API recebe a solicitação de criação da campanha

2. Campanha é persistida no banco de dados

3. Mensagem é publicada na fila

4. Worker consome a fila e realiza os envios

5. Status da campanha é atualizado

Benefícios:

- Evita bloqueio da API

- Permite alto volume de envios

- Facilita retry e observabilidade

## 🧪 Testes

A aplicação possui testes unitários cobrindo:

- Casos de uso

- Regras de negócio

- Serviços de domínio

Ferramentas:

- testing (padrão Go)

- Mocks para repositórios e serviços externos

- assert

## 🐳 Docker

O projeto utiliza Docker Compose para subir todo o ambiente:

- PostgreSQL

- Keycloak

## 📈 Boas Práticas Aplicadas

- REST API

- Separação de responsabilidades

- Processamento assíncrono

- Segurança com OAuth2

- Testes automatizados