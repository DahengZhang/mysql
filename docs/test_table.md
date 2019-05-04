CREATE TABLE test_table (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(40) NOT NULL,
    `content` TEXT NOT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO test_table (title, content) VALUE ("流浪地球", "无论结果将人类的历史导向何处，我们决定，选择希望！");

UPDATE test_table SET title="流浪地球法案", content="刹车时代，逃逸时代，流浪时代" WHERE id=1;

SELECT * FROM test_table;

SELECT * FROM test_table WHERE id=1;

DELETE FROM test_table WHERE id=1;
