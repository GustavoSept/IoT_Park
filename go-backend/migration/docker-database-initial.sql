CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(20),
    last_name VARCHAR(100),
    office_level VARCHAR(25)
);


INSERT INTO users (first_name, last_name, office_level) VALUES
('Gustavo', 'Limeira', 'Operador'),
('Maria', 'Ferreira', 'Gerente'),
('José', 'Santos', 'Dono'),
('Ana', 'Silva', 'Operador'),
('Pedro', 'Costa', 'Operador'),
('Camila', 'Oliveira', 'Operador');
