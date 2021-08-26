CREATE DATABASE edagang;
USE edagang;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `product_id` int(11) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(100) NOT NULL,
  `product_category` varchar(20) NOT NULL,
  `quantity` int(11) NOT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

LOCK TABLES `products` WRITE;

INSERT INTO `products` VALUES 
	(1,'Lenovo A30','ELECTRONIC', 100),
	(2,'MSI MODERN 14','ELECTRONIC', 20),
	(3,'HONDA SUPRA X 125 2010','VEHICLE', 25),
	(4,'DETAILS CASUAL BLUE','CLOTHES', 220);

UNLOCK TABLES;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `role` varchar(20) NOT NULL,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `users` WRITE;

INSERT INTO `users` VALUES 
  (1,'admin','abc123','admin','2020-08-09 10:27:22'),
  (2,'2001','abc123','user','2020-08-09 10:27:22'),
  (3,'2000','abc123','user','2020-08-09 10:27:22');

UNLOCK TABLES;

