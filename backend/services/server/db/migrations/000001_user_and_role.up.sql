CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  name text UNIQUE,
  description text
);

CREATE TABLE permissions (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  entity text,
  method text,
  description text
);

CREATE TABLE role_permission (
  role_id SERIAL,
  permission_id SERIAL,
  PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  username text UNIQUE,
  password text,
  email text UNIQUE,
  role_name text DEFAULT 'customer' NOT NULL,
  salt text
);

ALTER TABLE
  users
ADD
  CONSTRAINT fk_users_roles FOREIGN KEY (role_name) REFERENCES roles(name);

ALTER TABLE
  role_permission
ADD
  CONSTRAINT fk_role_permission_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE;

ALTER TABLE
  role_permission
ADD
  CONSTRAINT fk_role_permission_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE;

CREATE UNIQUE INDEX idx_roles_name ON public.roles USING btree (name);

CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);
CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);



INSERT INTO
  roles(created_at, name, description)
values
  (current_timestamp, 'customer', 'customer role');

INSERT INTO
  roles(created_at, name, description)
values
  (current_timestamp, 'superadmin', 'can do anything');

-- INSERT INTO
--   users(created_at, username, password, role_name, salt)
-- values
--   (current_timestamp, 'superadmin', 'customer role', 'superadmin', '');