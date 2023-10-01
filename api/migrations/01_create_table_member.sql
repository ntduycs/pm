create table member
(
    id           serial primary key,
    email        varchar(127) not null
        constraint unique_member_email unique,
    name         varchar(127) not null,
    level        int          not null check ( level between 0 and 12),
    positions    varchar(256) not null,
    kpi          int          not null check ( kpi between 0 and 200 ),
    category     varchar(15)  not null check ( category in ('official', 'buffer') ),
    total_effort real         not null check ( total_effort between 0 and 100 ),
    start_date   date,
    end_date     date check ( end_date >= start_date ),
    status       varchar(15)  not null check ( status in ('active', 'inactive', 'pending') )
);