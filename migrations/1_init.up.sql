CREATE TABLE IF NOT EXISTS `orders` (
  PRIMARY KEY (`id`)
  `id` varchar(255) NOT NULL,
  `price` float NOT NULL,
  `tax` float NOT NULL,
  `final_price` float NOT NULL,
)
