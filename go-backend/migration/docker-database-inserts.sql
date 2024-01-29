/* Usuários Dono */
INSERT INTO users (id, first_name, last_name, office_level) VALUES
('11111111-1111-1111-1111-111111111111','José', 'Santos', 'dono'),
('22222222-2222-2222-2222-222222222222','Patrícia', 'Souza', 'dono'),
('33333333-3333-3333-3333-333333333333','Rogério', 'da Silva Pereira Neves', 'dono');


/* Usuários Comuns */
INSERT INTO users (id, first_name, last_name, office_level) VALUES
('59721a63-9b2c-4754-8937-316f8a5942b3','Gustavo', 'Limeira', 'operador'),
('9b1153d4-02ba-494a-9327-89600e1a00b1','Maria', 'Ferreira', 'gerente');

INSERT INTO users_authentication (user_id, email, password_hash, salt) VALUES -- all passwords are `testPassword`
('59721a63-9b2c-4754-8937-316f8a5942b3','gustLim@gmail.com', '2vAqglwoBay7pEHne5WFW6Kio5/uVJA0CqJtN7lbcew', '0661798c29e62f313fcc20267a20cb6b'),
('9b1153d4-02ba-494a-9327-89600e1a00b1','abobrinha@gmail.com', 'hA1+B91+vhIawdiydhqr/bf+XyHG/ph72r5h0XF6UZw', '27103a4a063a5ad67f9652ec3d725f2c'),
('11111111-1111-1111-1111-111111111111','adm@admin.com', 'PMdqYn1t++gXg50MXlgwEjm8bxbFbVQdXaB2TBDXVLU', '15653de26d13ea6861651915ba4ae2c6'), -- pass is: Adm1098!
('22222222-2222-2222-2222-222222222222','Pathy+app@gmail.com', 'aZ/S6dt1tzBHhPEADbcRvUI6pqOk4p10KrXO94q5Y0E', 'a8fc781214146125518847fdb3720fc1'),
('33333333-3333-3333-3333-333333333333','Roger9000@hotmail.com', '434+TvM1iZW6o7hWRd1aPR3uxW8liSobh1Dzz33HDQk', '49406bc22617e929a37463995d0fa7de');


INSERT INTO parking_lots (pl_name, addr_street, addr_number, cep, owner_id) VALUES
('Sample Name','Street Name 1', 123, '12345678', '11111111-1111-1111-1111-111111111111'),
('Estacionator v1','Street Name 2', 456, '23456789', '22222222-2222-2222-2222-222222222222'),
('Sample Name','Street Name 3', 789, '34567890', '33333333-3333-3333-3333-333333333333'),
('Sample Name','Street Name 4', 101, '45678901', '33333333-3333-3333-3333-333333333333'),
('Estacionator v2','Street Name 5', 202, '56789012', '22222222-2222-2222-2222-222222222222');
