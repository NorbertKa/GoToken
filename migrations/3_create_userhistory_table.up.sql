CREATE TABLE IF NOT EXISTS gotoken.userhistory (
    ID bigserial primary key not null,
    created timestamp without time zone default (now() at time zone 'utc') not null,
    identifier text not null,
    password text not null
);