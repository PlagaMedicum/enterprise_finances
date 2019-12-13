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
    e_id int
        constraint eg_pk primary key,
    g_id int not null
);

create table if not exists minimal_salaries
(
    date  date
        constraint ms_pk primary key,
    value int not null
);

create table if not exists grades
(
    id    serial
        constraint grade_pk primary key,
    num   int  not null, -- order number of a grade.
    date  date not null,
    coeff int  not null
);

-- Data migrations goes here:

insert into employees (name, position, tu_membership)
values ('Daniil Dankovskij', 'Bachelor', true),
       ('Haruspex', 'Medic', false),
       ('Gordon Freeman', 'Physician', false);
insert into employees_grades
values (1, 1),
       (2, 2),
       (3, 3);
insert into minimal_salaries
values ('02/01/1999', 1023),
       ('05/04/2017', 7);
insert into grades (num, date, coeff)
values (5, '02/01/1999', 2),
       (7, '01/27/2047', 5),
       (10, '01/05/1980', 9);
