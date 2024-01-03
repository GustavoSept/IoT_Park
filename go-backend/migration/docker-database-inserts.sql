/* Usuários Dono */
INSERT INTO users (id, first_name, last_name, office_level) VALUES
('11111111-1111-1111-1111-111111111111','José', 'Santos', 'dono'),
('22222222-2222-2222-2222-222222222222','Patrícia', 'Souza', 'dono'),
('33333333-3333-3333-3333-333333333333','Rogério', 'da Silva Pereira Neves', 'dono');


/* Usuários Comuns */
INSERT INTO users (first_name, last_name, office_level) VALUES
('Gustavo', 'Limeira', 'operador'),
('Maria', 'Ferreira', 'gerente'),
('Ana', 'Silva', 'operador'),
('Pedro', 'Costa', 'operador'),
('Camila', 'Oliveira', 'operador'),
('Beatriz', 'Santos', 'operador'),
('Eduardo', 'Ribeiro', 'operador'),
('Patrícia', 'Nunes', 'operador'),
('Bruno', 'Lopes', 'gerente'),
('Sofia', 'Martins', 'operador');

INSERT INTO parking_lots (addr_street, addr_number, cep, owner_id) VALUES
('Street Name 1', 123, '12345678', '11111111-1111-1111-1111-111111111111'),
('Street Name 2', 456, '23456789', '22222222-2222-2222-2222-222222222222'),
('Street Name 3', 789, '34567890', '33333333-3333-3333-3333-333333333333'),
('Street Name 4', 101, '45678901', '33333333-3333-3333-3333-333333333333'),
('Street Name 5', 202, '56789012', '22222222-2222-2222-2222-222222222222');
