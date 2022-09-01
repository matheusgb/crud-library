CREATE TABLE IF NOT EXISTS books (
  id SERIAL NOT NULL,
  title varchar(250) NOT NULL,
  category varchar(250) NOT NULL,
  author varchar(250) NOT NULL,
  synopsis varchar NOT NULL,
  PRIMARY KEY (id)
);