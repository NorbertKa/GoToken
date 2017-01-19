CREATE TABLE IF NOT EXISTS gotoken.statistics (
    ID bigserial primary key not null,
    created timestamp without time zone default (now() at time zone 'utc') not null,
    registeredUsers integer default 0 not null,
    changedIdentities integer default 0 not null,
    changedPasswords integer default 0 not null,
    deletedUsers integer default 0 not null
);