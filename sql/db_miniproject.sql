/*
SQLyog Ultimate v12.5.1 (64 bit)
MySQL - 10.4.32-MariaDB : Database - miniproject
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`miniproject` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;

USE `miniproject`;

/*Table structure for table `jawaban_peserta` */

DROP TABLE IF EXISTS `jawaban_peserta`;

CREATE TABLE `jawaban_peserta` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `id_quiz` int(11) DEFAULT NULL,
  `id_pertanyaan` int(11) DEFAULT NULL,
  `jawaban_peserta` int(11) NOT NULL,
  `skor` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  KEY `id_quiz` (`id_quiz`),
  KEY `id_pertanyaan` (`id_pertanyaan`),
  CONSTRAINT `jawaban_peserta_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`),
  CONSTRAINT `jawaban_peserta_ibfk_2` FOREIGN KEY (`id_quiz`) REFERENCES `quiz` (`id`),
  CONSTRAINT `jawaban_peserta_ibfk_3` FOREIGN KEY (`id_pertanyaan`) REFERENCES `pertanyaan` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `jawaban_peserta` */

/*Table structure for table `pertanyaan` */

DROP TABLE IF EXISTS `pertanyaan`;

CREATE TABLE `pertanyaan` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pertanyaan` text NOT NULL,
  `opsi_jawaban` text NOT NULL,
  `jawaban_benar` int(11) NOT NULL,
  `id_quiz` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_quiz` (`id_quiz`),
  CONSTRAINT `pertanyaan_ibfk_1` FOREIGN KEY (`id_quiz`) REFERENCES `quiz` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `pertanyaan` */

/*Table structure for table `quiz` */

DROP TABLE IF EXISTS `quiz`;

CREATE TABLE `quiz` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `judul` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  `waktu_mulai` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT ' CURRENT_TIMESTAMP',
  `waktu_selesai` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT ' CURRENT_TIMESTAMP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `quiz` */

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `users` */

insert  into `users`(`id`,`nama`,`email`,`password`,`role`) values 
(1,'Roihan','roy@email.com','$2a$12$kShUgnkYB1Kna8OcVPD03O1QW8OxN6XIv0HrPKWe/6rxvkBAo3fFi','admin'),
(2,'Brian','bri@email.com','$2a$12$r6uzccjUoDy9Fj4e/I4Eb.7j0bjQSap67u2ga1uFjnft7MGksJL/i','user');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
