/* Usuários Dono */
INSERT INTO users (id, email, first_name, last_name, office_level) VALUES
('11111111-1111-1111-1111-111111111111','jose@gmail.com','José', 'Santos', 'dono'),
('22222222-2222-2222-2222-222222222222','patricia@gmail.com','Patrícia', 'Souza', 'dono'),
('33333333-3333-3333-3333-333333333333','rogerio@gmail.com','Rogério', 'da Silva Pereira Neves', 'dono');


/* Usuários Comuns */
INSERT INTO users (email, first_name, last_name, office_level) VALUES
('normalUser01@gmail.com','Gustavo', 'Limeira', 'operador'),
('normalUser02@gmail.com','Maria', 'Ferreira', 'gerente'),
('normalUser03@gmail.com','Ana', 'Silva', 'operador'),
('normalUser04@gmail.com','Pedro', 'Costa', 'operador'),
('normalUser05@gmail.com','Camila', 'Oliveira', 'operador'),
('normalUser06@gmail.com','Beatriz', 'Santos', 'operador'),
('normalUser07@gmail.com','Eduardo', 'Ribeiro', 'operador'),
('normalUser08@gmail.com','Patrícia', 'Nunes', 'operador'),
('normalUser09@gmail.com','Bruno', 'Lopes', 'gerente'),
('normalUser10@gmail.com','Sofia', 'Martins', 'operador');

INSERT INTO parking_lots (addr_street, addr_number, cep, owner_id) VALUES
('Street Name 1', 123, '12345678', '11111111-1111-1111-1111-111111111111'),
('Street Name 2', 456, '23456789', '22222222-2222-2222-2222-222222222222'),
('Street Name 3', 789, '34567890', '33333333-3333-3333-3333-333333333333'),
('Street Name 4', 101, '45678901', '33333333-3333-3333-3333-333333333333'),
('Street Name 5', 202, '56789012', '22222222-2222-2222-2222-222222222222');
