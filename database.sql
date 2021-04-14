CREATE DATABASE ticket;
USE ticket;
CREATE TABLE tickets (
  ticket VARCHAR(40) UNIQUE,
  name VARCHAR(30),
  username VARCHAR(30) UNIQUE,
  password VARCHAR(50),
  sold BOOLEAN,
  checkin BOOLEAN DEFAULT false
);

CREATE TABLE tokens (
  token_id VARCHAR(100) UNIQUE,
  used BOOLEAN DEFAULT false
);
