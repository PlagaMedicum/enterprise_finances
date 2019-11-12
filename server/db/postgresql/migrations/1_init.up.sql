create table if not exists employee
(
    id            serial
        constraint employee_pk primary key,
    name          text    not null,
    position      text    not null,
    tu_membership boolean not null
);

create table if not exists employees_grades
(
    e_id serial
        constraint eg_pk primary key,
    g_id serial not null
);

create table if not exists minimal_salary
(
    date  text   not null,
    value serial not null
);

create table if not exists grade
(
    id    serial
        constraint grade_pk primary key,
    date  text   not null,
    coeff serial not null
);
