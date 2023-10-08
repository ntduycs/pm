create table pa_pc
(
    id                  serial primary key,
    member_id           int          not null
        constraint pa_pc_member_id_fk references member (id) on delete restrict,
    technical_score     real         not null
        constraint pa_pc_technical_score_check check (technical_score >= 0),
    productivity_score  real         not null
        constraint pa_pc_productivity_score_check check (productivity_score >= 0),
    collaboration_score real         not null
        constraint pa_pc_collaboration_score_check check (collaboration_score >= 0),
    development_score   real         not null
        constraint pa_pc_development_score_check check (development_score >= 0),
    period              varchar(127) not null, -- in the format of yyyy-mm
    note                varchar(511) not null default ''
);

create unique index pa_pc_member_id_period_uindex
    on pa_pc (member_id, period);

create table pa_pc_technical_score
(
    id            serial primary key,
    pa_pc_id      int          not null
        constraint pa_pc_technical_score_pa_pc_id_fk references pa_pc (id) on delete cascade,
    type          varchar(31)  not null
        constraint pa_pc_technical_score_type_check check (type in ('soft-skills', 'hard-skills')),
    skill         varchar(127) not null,
    self_score    real         not null
        constraint pa_pc_technical_score_self_score_check check (self_score >= 0),
    peer_score    real
        constraint pa_pc_technical_score_peer_score_check check (peer_score >= 0),
    manager_score real         not null
        constraint pa_pc_technical_score_manager_score_check check (manager_score >= 0)
);

create unique index pa_pc_technical_score_pa_pc_id_type_skill_uindex
    on pa_pc_technical_score (pa_pc_id, type, skill);