CREATE TABLE IF NOT EXISTS gotoken.userprofile (
    ID bigserial primary key not null,
    created timestamp without time zone default (now() at time zone 'utc') not null,
    lastChanges timestamp without time zone default (now() at time zone 'utc') not null,
    lastLogin timestamp without time zone default (now() at time zone 'utc') not null,
    identifier text not null,
    password text not null
);