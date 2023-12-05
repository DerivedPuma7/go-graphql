# Server http comunicando via graphql

## Para rodar, siga os passos:
1. é necessário ter o sqlite instalado, caso não tenha, instale usando:
    - sudo apt install sqlite3
2. crie o banco local e as tabelas  
    - na raiz, execute:
        - sqlite3 data.db  
    - na cli do sqlite, execute:
        - create table categories(id string, name string, description string);
        - create table courses(id string, name string, description string, category_id string);
3. suba o servidor go usando:
    - go run cmd/server/server.go
4. um playground será aberto na porta 8080, brinque a vontade!!!