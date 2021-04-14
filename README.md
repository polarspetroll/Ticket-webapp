# Installation :

```bash
go get -u github.com/go-sql-driver/mysql
go build -o main
export DBUSR='your database username'
export DBPWD='your database password'  
export DBADDR='your database address' 
./main

```

# Environment Variables :

- **DBUSR**   :  database username
- **DBPWD**   :  database password
- **DBADDR**  :  database address
---

#### Database :

```sql
CREATE DATABASE ticket;
USE ticket;
CREATE TABLE tickets (
  ticket VARCHAR(40) UNIQUE,
  name VARCHAR(30),
  username VARCHAR(30) UNIQUE,
  password VARCHAR(100),
  sold BOOLEAN
);
CREATE TABLE tokens (
  token_id VARCHAR(100) UNIQUE,
  used BOOLEAN DEFAULT false
);


```

- **DATABASE TYPE** : MYSQL
- <strong><a style="color:cyan;" href="https://github.com/go-sql-driver/mysql">SQL Driver</a></strong>
---
- **Default Port** : 8080
- **Dynamics** : templates/*.html
