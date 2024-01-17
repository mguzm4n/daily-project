CREATE TABLE IF NOT EXISTS permissions (
  id bigserial PRIMARY KEY,
  code text NOT NULL
);

CREATE TABLE IF NOT EXISTS roles (
  id bigserial PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS roles_users (
  role_id bigint NOT NULL REFERENCES roles ON DELETE CASCADE,
  user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
  PRIMARY KEY (user_id, role_id)
);


CREATE TABLE IF NOT EXISTS roles_permissions (
  role_id bigint NOT NULL REFERENCES roles ON DELETE CASCADE,
  permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
  PRIMARY KEY (role_id, permission_id)
);

-- Add the two permissions to the table.
INSERT 
INTO permissions (code)
VALUES
('notes:read'),
('notes:write');

INSERT 
INTO roles (name)
VALUES
('USER'),
('ADMIN');
