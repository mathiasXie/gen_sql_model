create table users
(
    id                  bigint unsigned auto_increment comment '自增id',
    user_name           varchar(32)  not null default '' comment '用户名',
    gmt_created         timestamp    not null default CURRENT_TIMESTAMP comment '创建时间',
    gmt_modified        timestamp    not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_name` (`user_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 comment '用户表';