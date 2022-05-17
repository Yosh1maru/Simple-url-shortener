CREATE TABLE url (
                     token VARCHAR NOT NULL,
                     full_url VARCHAR NOT NULL,
                     created_at DATE NOT NULL DEFAULT CURRENT_DATE,
                     expire_at TIMESTAMP NOT NULL DEFAULT CURRENT_DATE+1
);