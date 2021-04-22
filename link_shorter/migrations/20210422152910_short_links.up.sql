CREATE TABLE short_links
(
    id         BIGINT UNSIGNED AUTO_INCREMENT UNIQUE
        COMMENT 'PK',
    short_link_key VARCHAR(255) DEFAULT NULL UNIQUE
        COMMENT 'ключ для сокращенной ссылки',
    link       VARCHAR(255) NOT NULL UNIQUE COMMENT 'полная ссылка',
    PRIMARY KEY (id),
    UNIQUE INDEX k_short_link (short_link)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci