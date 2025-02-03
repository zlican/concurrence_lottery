-- 创建inventory表的SQL语句
CREATE TABLE IF NOT EXISTS `inventory` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '商品名称',
  `description` text NOT NULL COMMENT '商品描述',
  `picture` varchar(255) COMMENT '商品图片路径',
  `price` decimal(10,2) NOT NULL COMMENT '商品价格',
  `count` int NOT NULL DEFAULT 0 COMMENT '商品库存数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 插入数据的SQL语句
INSERT INTO `inventory` (`id`, `name`, `description`, `picture`, `price`, `count`) VALUES
(1, '谢谢参与', '', 'img/face.png', 0, 0),
(10, '篮球', '', 'img/ball.jpeg', 100, 1000),
(11, '水杯', '', 'img/cup.jpeg', 80, 1000),
(12, '电脑', '', 'img/laptop.jpeg', 6000, 200),
(13, '平板', '', 'img/ipad.jpg', 4000, 300),
(15, '锅', '', 'img/pot.jpeg', 120, 1000),
(16, '茶叶', '', 'img/tea.jpeg', 90, 1000),
(17, '无人机', '', 'img/uav.jpeg', 400, 100),
(18, '酒', '', 'img/wine.jpeg', 160, 500),
(19, '手机', '', 'img/phone.jpeg', 5000, 400); 