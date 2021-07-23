create table users
(
    id       serial primary key,
    username varchar(255),
    password varchar(255),
    role     integer
);

create table auth
(
    id            serial primary key,
    refresh_token varchar(255),
    user_id       integer
);

create table vulnerability
(
    id          serial primary key,
    name        varchar(255),
    description varchar(255),
    created     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

create table project
(
    id                serial primary key,
    name              varchar(255),
    short_description varchar(255),
    description       varchar(255),
    private           bool,
    closed            bool,
    org_id            integer,
    created           TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated           TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

create table report
(
    id                serial primary key,
    name              varchar(255),
    description       varchar(255),
    status            varchar(255),
    seriousness       varchar(255),
    archive           bool,
    delete            bool,
    reward            integer,
    point             integer,
    project_id        integer,
    vulnerability_id  integer,
    user_id           integer,
    assignee          integer,
    unread_comments   bool,
    comments          varchar(255),
    sent_report_date  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    last_comment_time TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created           TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated           TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

create table level_achievements
(
    id                      serial primary key,
    achievement_name        varchar(255),
    achievement_description varchar(255),
    points                  integer,
    ordered_number          integer
);

insert into level_achievements(achievement_name, achievement_description, points, ordered_number)
values ('Слава Kali', 'Поклоняющийся человек кали', 50),
       ('Оруженосец', 'Тут все понятно', 100),
       ('Дизенфектор', 'Что-то с жуком и газом', 150),
       ('Гроза муравьев', 'Гроза муравьев', 200),
       ('Лучше, чем ничего', 'Обыграть какую-то идею с 0', 250),
       ('Мальчик с пальчик', 'Тут все понятно', 300),
       ('Мамкин хацкер', 'Пальцы нажимающие на клаве на “х” и “у”', 350),
       ('Вступил в комьюнити', 'Отсылка к лого КХС', 500),
       ('Можешь пробовать себя в CTF', 'Идея с флагом на башне', 550),
       ('Охотник', 'Человек со шкурой льва и копьем', 600),
       ('Ученик превосходит учителя', 'Молодой пацанчик стал выше старого', 850),
       ('Знатный воин', 'Какой-то из атрибутов воина', 900),
       ('Защитник', 'Барьер и летящий на него огненный шар', 950),
       ('Военачальник', 'Армия сзади и человек спереди', 1300),
       ('Тебе пора вести мастерклассы', 'Микрофон', 1500),
       ('Легенда', 'Воткнутый джедайский меч в землю', 1600),
       ('Правильный путь', 'чистый путь, а соседний с обрывом', 2000);

create table user_achievements
(
    id                    serial primary key,
    user_id               integer,
    level_achievements_id integer,
    created     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);