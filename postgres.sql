

CREATE TABLE auth (
  id SERIAL PRIMARY KEY,
  email varchar(100) NOT NULL,
  password varchar(100) not null,
  public_id varchar(100) not null,
  role varchar(20) NOT NULL DEFAULT 'user',
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW()
);

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  sku varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  stock INT NOT null DEFAULT 0,
  price INT NOT null DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  user_public_id varchar(100) NOT NULL,         
  product_id int NOT NULL,
  product_price int NOT NULL,
  amount int NOT NULL,
  sub_total int NOT NULL,
  platform_fee int DEFAULT 0,
  grand_total int NOT NULL,
  status int NOT NULL,
  product_snapshot jsonb,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);
