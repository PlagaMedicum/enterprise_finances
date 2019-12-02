create table if not exists employees
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

create table if not exists minimal_salaries
(
    date  text   not null,
    value serial not null
);

create table if not exists grades
(
    id    serial
        constraint grade_pk primary key,
    date  text   not null,
    coeff serial not null
);

-- Data migrations goes here:

insert into employees values (1, 'Johny', 'Osterbeiter', true);
insert into employees values (2, 'Pen', 'Pen', false);
insert into employees_grades values (1, 1);
insert into employees_grades values (2, 2);
insert into minimal_salaries values ('01.02.1999', 1023);
insert into minimal_salaries values ('4.05.2017', 7);
insert into grades values (1, '01.02.1999', 2);
insert into grades values (2, '27.01.2049', 5);
