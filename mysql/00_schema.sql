create table if not exists
user (
  id BINARY(16) PRIMARY KEY,
  id_text varchar(36) GENERATED ALWAYS as
  (
    insert(
      insert(
        insert(
          insert(
            hex(id), 9, 0, '-'
          ), 14, 0, '-'
        ), 19, 0, '-'
      ), 24, 0, '-'
    )
  ) VIRTUAL,
  created DATETIME NOT NULL
) CHARACTER SET 'utf8mb4';

create table if not exists
user_email (
  email varchar(128) not null primary key,
  id_user binary(16) not null,
  created datetime not null,
  index (id_user, created),
  constraint user_email_id_user foreign key (id_user) references user (id) on delete cascade on update cascade
) character set 'utf8mb4';

create table if not exists
user_handle (
  handle varchar(64) not null primary key,
  id_user binary(16) not null,
  created datetime not null,
  index (id_user, created),
  constraint user_handle_id_user foreign key (id_user) references user (id) on delete cascade on update cascade
) character set 'utf8mb4';

create table if not exists
user_temp_token (
  token varchar(32) not null,
  id_user binary(16) not null,
  created datetime not null,
  primary key (token),
  index (id_user, created),
  constraint user_temp_token_id_user foreign key (id_user) references user (id) on delete cascade on update cascade
) character set 'utf8mb4';

create table if not exists
user_setting (
  id_user binary(16) not null,
  created bigint not null,
  lang varchar(16) not null default 'en',
  primary key (id_user, created),
  constraint user_setting_id_user foreign key (id_user) references user (id) on delete cascade on update cascade
) character set 'utf8mb4';

create table if not exists
user_status (
  id_user binary(16) not null,
  created bigint not null,
  status varchar(16) not null,
  primary key (id_user, created),
  constraint user_status_id_user foreign key (id_user) references user (id) on delete cascade on update cascade
) character set 'utf8mb4';

