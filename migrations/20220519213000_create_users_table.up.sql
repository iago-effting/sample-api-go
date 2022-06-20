CREATE TABLE IF NOT EXISTS users(
   id uuid DEFAULT uuid_generate_v4(),
   email VARCHAR (300) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,

   PRIMARY KEY (id)
);