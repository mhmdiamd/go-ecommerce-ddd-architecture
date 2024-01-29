

CREATE TABLE auth (
  id SERIAL PRIMARY KEY,
  email varchar(100) NOT NULL,
  password varchar(100) not null,
  public_id varchar(100) not null,
  role varchar(20) NOT NULL DEFAULT 'user',
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW()
);
