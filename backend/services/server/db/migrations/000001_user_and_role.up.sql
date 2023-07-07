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

CREATE TABLE role_permissions (
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
  role_name text DEFAULT 'Free tier user' NOT NULL,
  salt text
);

ALTER TABLE
  users
ADD
  CONSTRAINT fk_users_roles FOREIGN KEY (role_name) REFERENCES roles(name);

ALTER TABLE
  role_permissions
ADD
  CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE;

ALTER TABLE
  role_permissions
ADD
  CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE;

CREATE UNIQUE INDEX idx_roles_name ON public.roles USING btree (name);

CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);

INSERT INTO
  permissions(created_at, entity, method, description)
values
  (
    current_timestamp,
    'USER',
    'DELETE',
    'Allow delete user'
  ),
  (
    current_timestamp,
    'ROLE',
    'CREATE',
    'Allow create new role'
  ),
  (
    current_timestamp,
    'ROLE',
    'DELETE',
    'Allow delete role'
  ),
  (
    current_timestamp,
    'ROLE',
    'UPDATE',
    'Allow update role'
  ),
  (
    current_timestamp,
    'CRAWLER',
    'CREATE',
    'Allow delete crawler'
  ),
  (
    current_timestamp,
    'CRAWLER',
    'UPDATE',
    'Allow update crawler'
  ),
  (
    current_timestamp,
    'ARTICLES_SOURCES',
    'UPDATE',
    'Allow update articles sources'
  ),
  (
    current_timestamp,
    'ARTICLES_SOURCES',
    'DELETE',
    'Allow delete articles sources'
  ),
  (
    current_timestamp,
    'CATEGORY',
    'CREATE',
    'Allow create new category'
  ),
  (
    current_timestamp,
    'CATEGORY',
    'DELETE',
    'Allow delete category'
  ),
  (
    current_timestamp,
    'CATEGORY',
    'UPDATE',
    'Allow update category'
  ),
  (
    current_timestamp,
    'TOPIC',
    'CREATE',
    'Allow create new topic'
  ),
  (
    current_timestamp,
    'TOPIC',
    'DELETE',
    'Allow delete topic'
  ),
  (
    current_timestamp,
    'TOPIC',
    'UPDATE',
    'Allow update topic'
  ),
  (
    current_timestamp,
    'ADMIN PAGE',
    'ACCESS',
    'Allow access to admin page'
  );

INSERT INTO
  roles(created_at, name, description)
values
  (
    current_timestamp,
    'Superadmin',
    'Can do anything'
  ),
  (
    current_timestamp,
    'Free tier user',
    'Can read article but not add new source'
  ),
  (
    current_timestamp,
    'Premium tier user',
    'Can read article and add new source'
  );

INSERT INTO
  role_permissions(role_id, permission_id)
values
  (1, 1),
  (1, 2),
  (1, 3),
  (1, 4),
  (1, 5),
  (1, 6),
  (1, 7),
  (1, 8),
  (1, 9),
  (1, 10),
  (1, 11),
  (1, 12),
  (1, 13),
  (1, 14),
  (1, 15),
  (2, 4),
  (3, 4),
  (3, 5);

INSERT INTO
  users(
    created_at,
    username,
    password,
    email,
    role_name,
    salt
  )
values
  (
    current_timestamp,
    'superadmin',
    '$argon2id$v=19$m=65536,t=3,p=2$olRPAZKGAI0$T49QphrVd7PJeF1ghQzL/Ba2aWofM8Sxp5QbN/MHU30',
    'superadmin@gmail.com',
    'Superadmin',
    'olRPAZKGAI0'
  );