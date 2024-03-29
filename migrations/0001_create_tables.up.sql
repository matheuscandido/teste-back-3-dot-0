CREATE TABLE IF NOT EXISTS accounts (
    id serial primary key,
    document_number varchar(255) unique not null
);

CREATE TABLE IF NOT EXISTS operations_types (
    id serial primary key,
    description varchar(255) unique not null
);

INSERT INTO operations_types (description) values('COMPRA A VISTA');
INSERT INTO operations_types (description) values('COMPRA PARCELADA');
INSERT INTO operations_types (description) values('SAQUE');
INSERT INTO operations_types (description) values('PAGAMENTO');

CREATE TABLE IF NOT EXISTS transactions (
    id serial primary key,
    account_id int not null,
    operation_id int not null,
    amount numeric,
    event_date timestamp,
    CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts(id),
    CONSTRAINT fk_operation FOREIGN KEY (operation_id) REFERENCES operations_types(id)
);