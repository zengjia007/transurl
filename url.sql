/*
SQLyog Ultimate v8.32 
MySQL - 5.5.36 : Database - url
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`url` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `url`;

/*Table structure for table `tb_url` */

DROP TABLE IF EXISTS `tb_url`;

CREATE TABLE `tb_url` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `origin_url` varchar(500) DEFAULT NULL,
  `short_url` varchar(500) DEFAULT NULL,
  `url_code` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8;

/*Data for the table `tb_url` */

insert  into `tb_url`(`id`,`origin_url`,`short_url`,`url_code`) values (37,'http://fanyi.baidu.com/translate?aldtype=16047&query=&keyfrom=bault=dict&lang=auto2zh#auto/zh/','http://127.0.0.1:8088/B','469CNBb05896fd786f8e3a676e370c649d3d2b'),(38,'http://fanyi.baidu.com/translate?aldtype=16047&query=&keyfrom=bault=dict&lag=auto2zh#auto/zh/','http://127.0.0.1:8088/C','39W7Ol39982e39783ebfcecd038fc4ee8071a8'),(39,'http://fanyi.baidu.com/translate?aldtype=16047&query=&keyfrom=blt=dict&lag=auto2zh#auto/zh/','http://127.0.0.1:8088/D','e6Kjc70e8dcf6168f14d6af0ec51597b5f258'),(40,'http://fanyi.baidu.com/translate?aldtype=16047&query=&keyfom=blt=dict&lag=auto2zh#auto/zh/','http://127.0.0.1:8088/E','3ZFYDefc8c9335348dba7e6c5e4a98a26a553');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
