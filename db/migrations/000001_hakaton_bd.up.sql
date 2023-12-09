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
    total_points          int,
    final_exam_current    int,
    final_exam_max        int,            
    expired_at            timestamp null
);
-- INSERT INTO "user" (password_hash,registration_datetime) values('admin'); 
CREATE TABLE "course"
(
    id                    serial PRIMARY KEY not null unique,
    description           varchar(300),
    course_age            int,
    url                    varchar(255),
    name                  varchar(100)
    
);
INSERT INTO "course" (description, course_age, url, name) values('1',2,'3','4'); 
INSERT INTO "course" (description, course_age, url, name) values('1',2,'3','c2'); 
INSERT INTO "course" (description, course_age, url, name) values('1',2,'3','c3'); 
INSERT INTO "course" (description, course_age, url, name) values('1',2,'3','c4'); 

CREATE TABLE "lesson_type"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(100)
);
INSERT INTO "lesson_type" (name) values('material'); 
INSERT INTO "lesson_type" (name) values('test'); 

CREATE TABLE "lesson"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_type           int references "lesson_type"(id) on delete cascade,
    course_id             int references "course"(id) on delete cascade,
    value                 int,
    name                  varchar(100)
);

INSERT INTO "lesson" (course_id, lesson_type, value, name) values(1,1,4,'m1'); 
INSERT INTO "lesson" (course_id, lesson_type, value, name) values(1,2,0,'l1'); 
INSERT INTO "lesson" (course_id, lesson_type, value, name) values(1,1,12,'m2'); 
INSERT INTO "lesson" (course_id, lesson_type, value, name) values(1,2,13,'l2'); 
CREATE TABLE "user_lesson"
(
    id                  serial PRIMARY KEY not null unique,
    lesson_id           int references "lesson"(id) on delete cascade,
    course_id           int references "course"(id) on delete cascade,
    user_id             int references "user"(id) on delete cascade,
    finish              bool
);
-- INSERT INTO "user_lesson" (user_id, lesson_id, finish) values(1,1,true); 
-- INSERT INTO "user_lesson" (user_id, lesson_id, finish) values(1,2,true); 

-- lesson material 
CREATE TABLE "lesson_mat"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_id             int references "lesson"(id) on delete cascade,
    name                  varchar(255),
    lesson_text           text null
);
INSERT INTO "lesson_mat" (lesson_id, lesson_text) values(1,'mm1'); 
INSERT INTO "lesson_mat" (lesson_id, lesson_text) values(3,'mm3');

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
INSERT INTO "lesson_test" (lesson_id,name, lesson_text) values(2,'l2','ll2'); 
INSERT INTO "lesson_test" (lesson_id,name, lesson_text) values(4,'l4','ll4');
-- lesson_test_question
CREATE TABLE "lesson_test_question_type"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(255)
);
INSERT INTO "lesson_test_question_type" (name) values('test'); 
INSERT INTO "lesson_test_question_type" (name) values('images');
CREATE TABLE "lesson_test_question"
(
    id                    serial PRIMARY KEY not null unique,
    lesson_id             int references "lesson"(id) on delete cascade,
    lesson_test_question_type_id   int references "lesson_test_question_type"(id) on delete cascade,
    url                       varchar(255),
    question                  varchar(255)
);
INSERT INTO "lesson_test_question" (lesson_id,lesson_test_question_type_id,url,question) values(2,1,'1','?'); 
INSERT INTO "lesson_test_question" (lesson_id,lesson_test_question_type_id,url,question) values(2,1,'2','?');  

INSERT INTO "lesson_test_question" (lesson_id,lesson_test_question_type_id,url,question) values(4,1,'123','?'); 
CREATE TABLE "lesson_test_answer"
(
    id                      serial PRIMARY KEY not null unique,
    answer_text             text null,
    lesson_test_question_id int references "lesson_test_question"(id) on delete cascade,
    correct                 bool 
);
INSERT INTO "lesson_test_answer" (lesson_test_question_id,answer_text,correct) values(1,'--',false);
INSERT INTO "lesson_test_answer" (lesson_test_question_id,answer_text,correct) values(1,'+',true);
INSERT INTO "lesson_test_answer" (lesson_test_question_id,answer_text,correct) values(1,'=',false);
INSERT INTO "lesson_test_answer" (lesson_test_question_id,answer_text,correct) values(1,'/',false);

INSERT INTO "lesson_test_answer" (lesson_test_question_id,answer_text,correct) values(2,'[]',false);
INSERT INTO "lesson_test_answer" (lesson_test_question_id,answer_text,correct) values(2,'()',true);

INSERT INTO "lesson_test_answer" (lesson_test_question_id,answer_text,correct) values(3,'иволга',true);

CREATE TABLE "lesson_test_answer_src"
(
    id                     serial PRIMARY KEY not null unique,
    lesson_test_answer_id  int references "lesson_test_answer"(id) on delete cascade,
    url                    varchar(255)
);
INSERT INTO "lesson_test_answer_src" (lesson_test_answer_id,url) values(4,'photo ans');


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

-- achive 
CREATE TABLE "achive"
(
    id                     serial PRIMARY KEY not null unique,
    course                 int references "course"(id) on delete cascade, 
    name                   varchar(50),
    description            varchar(300),
    value                  int null
);
CREATE TABLE "user_achive"
(
    id                    serial PRIMARY KEY not null unique,
    achive_id             int references "achive"(id) on delete cascade,
    user_id               int references "user"(id) on delete cascade
);
CREATE TABLE "product"
(
    id                     serial PRIMARY KEY not null unique,
    name                   varchar(50),
    description            varchar(300),
    price                  int not null,
    url                    varchar(255)
);

-- news 
CREATE TABLE "news"
(
    id                     serial PRIMARY KEY not null unique,
    title                  varchar(50),
    news_text              text,
    creation_datetime      TIMESTAMP(0) NOT NULL,
    url_image              varchar(255),
    url_video              varchar(255)
);



-- final test
CREATE TABLE "final_test"
(
    id                    serial PRIMARY KEY not null unique,
    test_description      text
);
INSERT INTO "final_test" (test_description) values('main test');

CREATE TABLE "user_final"
(
    id                    serial PRIMARY KEY not null unique,
    final_test_id   int references "final_test"(id) on delete cascade,
    user_id         int references "user"(id) on delete cascade,
    max_result      int,
    last_result     int
);

CREATE TABLE "final_test_question_direction"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(255)
);
INSERT INTO "final_test_question_direction" (name) values('cyberJil');

CREATE TABLE "final_test_question"
(
    id                       serial PRIMARY KEY not null unique,
    final_test_question_direction_id   int references "final_test_question_direction"(id) on delete cascade,
    url                       varchar(255),
    question                  varchar(255)
);
INSERT INTO "final_test_question" (question,url,final_test_question_direction_id) values('meme','123',1);

CREATE TABLE "final_test_answer"
(
    id                      serial PRIMARY KEY not null unique,
    answer_text             text null,
    url                     varchar(255),
    final_test_question_id int references "final_test_question"(id) on delete cascade,
    correct                 bool 
);
INSERT INTO "final_test_answer" (final_test_question_id,answer_text,correct,url) values(1,'--',false,'123');
INSERT INTO "final_test_answer" (final_test_question_id,answer_text,correct) values(1,'+',true);
INSERT INTO "final_test_answer" (final_test_question_id,answer_text,correct) values(1,'=',false);
INSERT INTO "final_test_answer" (final_test_question_id,answer_text,correct) values(1,'/',false);