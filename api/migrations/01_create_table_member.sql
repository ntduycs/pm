create table member
(
    id           serial primary key,
    email        varchar(127)     not null constraint unique_member_email unique,
    name         varchar(127)     not null,
    level        varchar(15)      not null,
    positions    varchar(256)     not null,
    kpi          double precision not null check ( kpi between 0 and 1 ),
    category     varchar(15)      not null,
    total_effort double precision not null check ( total_effort between 0 and 100 ),
    start_date   date,
    end_date     date,
    status       varchar(15)      not null check ( status in ('active', 'inactive') )
);