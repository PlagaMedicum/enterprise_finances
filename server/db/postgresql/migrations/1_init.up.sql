create table if not exists employee
(
    id     serial
        constraint employee_pk primary key,
    name   text   not null,
    salary serial not null
);

create table if not exists grade
(
    id    serial
        constraint grade_pk primary key,
    date  text   not null,
    coeff serial not null
);
