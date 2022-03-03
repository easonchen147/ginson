CREATE TABLE `users`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)      NOT NULL COMMENT '用户名',
    `email`      varchar(50)      NOT NULL COMMENT '邮箱',
    `password`   varchar(40)      NOT NULL COMMENT '密码',
    `salt`       varchar(100)     NOT NULL COMMENT '盐',
    `created_at` datetime default NULL COMMENT '创建时间',
    `updated_at` datetime default NULL COMMENT '更新时间',
    `deleted_at` datetime default NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

