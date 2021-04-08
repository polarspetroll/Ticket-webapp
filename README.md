# Installation :

```bash
go build -o main
export DBUSR='your database username'  #//
export DBPWD='your database password'  #//for running as a service add this lines to .bashrc file
export DBADDR='your database address'  #//
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

```

- **DATABASE TYPE** : <strong style="color:cyan;;">MYSQL</strong>
- <strong><a style="color:cyan;" href="https://github.com/go-sql-driver/mysql">SQL Driver</a></strong>
---
- **Default Port** : 8080
- **Admin Pannel** : /admin/
- **Dynamics** : templates/*.html

##### Report Bugs :

- [github](https://github.com/polarspetroll)
- [email](mailto:polarspetroll@protonmail.com)
- [website](https://polarspetroll.github.io)
