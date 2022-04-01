-- 创建数据库

CREATE
DATABASE ginson;

-- 创建数据表

CREATE TABLE `users`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `open_id`    varchar(50)             NOT NULL COMMENT '第三方平台Id',
    `source`     varchar(50)             NOT NULL COMMENT '注册来源',
    `nickname`   varchar(100)            NOT NULL comment '昵称',
    `avatar`     varchar(200) DEFAULT '' NOT NULL COMMENT '头像',
    `age`        int          DEFAULT 0  NOT NULL COMMENT '年龄',
    `gender`     int          DEFAULT 0  NOT NULL COMMENT '性别0-未知，1-男，2-女',
    `created_at` datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime     DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime     DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY          `users_open_id_source_uindex` (`open_id`, `source`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
COMMENT '用户表';

