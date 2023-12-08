CREATE TABLE "role"
(
    id                      serial  PRIMARY KEY not null unique,
    name                    varchar(50) not null
);
INSERT INTO "role" (name) values('user'); 
INSERT INTO "role" (name) values('admin'); 

CREATE TABLE "user"
(
    id                    serial PRIMARY KEY not null unique,
    password_hash         varchar(255) not null,
    role_id               int references "role"(id) on delete cascade,
    email                 VARCHAR(40),
    registration_datetime TIMESTAMP(0) NOT NULL,
    age                   int,
    refresh               varchar(100) null,
    expired_at            timestamp null
);
CREATE TABLE "course"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(100)
);
CREATE TABLE "lesson_type"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(100)
);
CREATE TABLE "lesson"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_type           int references "lesson_type"(id) on delete cascade,
    course_id             int references "course"(id) on delete cascade,
    name                  varchar(100)
);
CREATE TABLE "user_lesson"
(
    id                  serial PRIMARY KEY not null unique,
    lesson_id           int references "lesson"(id) on delete cascade,
    user_id             int references "user"(id) on delete cascade,
    finish              bool
);

-- lesson material 
CREATE TABLE "lesson_mat"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_id             int references "lesson"(id) on delete cascade,
    name                  varchar(255),
    lesson_text           text null
);
CREATE TABLE "lesson_mat_src"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_mat_id         int references "lesson_mat"(id) on delete cascade,
    url                   varchar(255)
);


-- lesson_test
CREATE TABLE "lesson_test"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_id             int references "lesson"(id) on delete cascade,
    name                  varchar(255),
    lesson_text           text
);


-- lesson_test_question
CREATE TABLE "lesson_test_question_type"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(255)
);
CREATE TABLE "lesson_test_answer"
(
    id                    serial PRIMARY KEY not null unique,
    answer_text           text null
);
CREATE TABLE "lesson_test_answer_src"
(
    id                     serial PRIMARY KEY not null unique,
    lesson_test_answer_id  int references "lesson_test_answer"(id) on delete cascade,
    url                    varchar(255)
);
CREATE TABLE "lesson_test_question"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_id             int references "lesson"(id) on delete cascade,
    lesson_test_question_type_id   int references "lesson_test_question_type"(id) on delete cascade,
    name                  varchar(255),
    lesson_right_answer_id      int references "lesson_test_answer"(id) on delete cascade
);


-- forward 
CREATE TABLE "forward"
(
    id                     serial PRIMARY KEY not null unique,
    name                    varchar(100),
    forward_text            text,
    creation_datetime       TIMESTAMP(0) NOT NULL
);
CREATE TABLE "forward_src"
(
    id                     serial PRIMARY KEY not null unique,
    forward_id             int references "forward"(id) on delete cascade,
    url                    varchar(255)
);