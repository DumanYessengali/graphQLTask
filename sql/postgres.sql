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