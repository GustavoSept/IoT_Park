/* Usuários Dono */
INSERT INTO users (id, first_name, last_name, office_level) VALUES
('11111111-1111-1111-1111-111111111111','José', 'Santos', 'dono'),
('22222222-2222-2222-2222-222222222222','Patrícia', 'Souza', 'dono'),
('33333333-3333-3333-3333-333333333333','Rogério', 'da Silva Pereira Neves', 'dono');


/* Usuários Comuns */
INSERT INTO users (id, first_name, last_name, office_level) VALUES
('59721a63-9b2c-4754-8937-316f8a5942b3','Gustavo', 'Limeira', 'operador'),
('9b1153d4-02ba-494a-9327-89600e1a00b1','Maria', 'Ferreira', 'gerente'),
('9530df61-2f4c-4fae-8510-9520682ef2b6','Ana', 'Silva', 'operador'),
('af89d23c-c205-420f-8e03-8ecca3758d1f','Pedro', 'Costa', 'operador');

INSERT INTO users_authentication (user_id, email, password_hash, salt) VALUES -- all passwords are `testPassword`
('59721a63-9b2c-4754-8937-316f8a5942b3','gustLim@gmail.com', 'pVOFiaTlsSs+VwT/kwih1UpGYtPE3ihDQ0eJeyAR77Q', '0661798c29e62f313fcc20267a20cb6b'),
('9b1153d4-02ba-494a-9327-89600e1a00b1','abobrinha@gmail.com', 'pVOFiaTlsSs+VwT/kwih1UpGYtPE3ihDQ0eJeyAR77Q', '0661798c29e62f313fcc20267a20cb6b'),
('9530df61-2f4c-4fae-8510-9520682ef2b6','arroba2085_nox@gmail.com', 'pVOFiaTlsSs+VwT/kwih1UpGYtPE3ihDQ0eJeyAR77Q', '0661798c29e62f313fcc20267a20cb6b'),
('af89d23c-c205-420f-8e03-8ecca3758d1f','Pedro_costL@gmail.com', 'pVOFiaTlsSs+VwT/kwih1UpGYtPE3ihDQ0eJeyAR77Q', '0661798c29e62f313fcc20267a20cb6b');


INSERT INTO parking_lots (pl_name, addr_street, addr_number, cep, owner_id) VALUES
('Sample Name','Street Name 1', 123, '12345678', '11111111-1111-1111-1111-111111111111'),
('Estacionator v1','Street Name 2', 456, '23456789', '22222222-2222-2222-2222-222222222222'),
('Sample Name','Street Name 3', 789, '34567890', '33333333-3333-3333-3333-333333333333'),
('Sample Name','Street Name 4', 101, '45678901', '33333333-3333-3333-3333-333333333333'),
('Estacionator v2','Street Name 5', 202, '56789012', '22222222-2222-2222-2222-222222222222');
