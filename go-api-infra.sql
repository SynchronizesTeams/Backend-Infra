/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-12.0.2-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: go-api-infra
-- ------------------------------------------------------
-- Server version	12.0.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `achievements`
--

DROP TABLE IF EXISTS `achievements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `achievements` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `achievement_date` longtext DEFAULT NULL,
  `news_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_achievements_deleted_at` (`deleted_at`),
  KEY `fk_achievements_news` (`news_id`),
  CONSTRAINT `fk_achievements_news` FOREIGN KEY (`news_id`) REFERENCES `news` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `achievements`
--

LOCK TABLES `achievements` WRITE;
/*!40000 ALTER TABLE `achievements` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `achievements` VALUES
(1,'2025-10-17 08:45:57.352','2025-10-17 08:53:31.557','2025-10-17 09:01:36.749','juara 1 JHIC','sdsadks','','2025-10-03',12),
(2,'2025-10-17 08:47:09.871','2025-10-17 08:47:09.871','2025-10-17 09:01:39.998','juara 1 JHIC','dsajdsjpoadjsaodjposja','','2025-10-03',12),
(3,'2025-10-17 08:49:05.913','2025-10-17 08:53:50.317','2025-10-17 09:01:47.137','juara 1 JHIC','sdsadks','','2025-10-03',12),
(4,'2025-10-17 08:49:21.790','2025-10-17 17:22:13.292',NULL,'juara 1 JHIC','<p>sdsassqdk<strong>ssdwdqwsw</strong></p>','uploads/achievement/1760665761_Screenshot_06-Oct_14-26-05_kitty.png','2025-10-03',12),
(5,'2025-10-17 14:36:02.204','2025-10-17 14:36:02.204','2025-10-17 17:22:49.091','juara 1 JHIC','dsajdsjpoadjsaodjposja','','2025-10-03',12),
(6,'2025-10-17 14:36:14.204','2025-10-17 14:36:14.204',NULL,'juara 1 JHIC','dsajdsjpoadjsaodjposja','uploads/achievement/1760686574_swappy-20251015-220044.png','2025-10-03',12),
(7,'2025-10-17 17:06:40.054','2025-10-17 17:06:40.054',NULL,'Menang osn','<p>MAHHH AKU MENANG</p>','uploads/achievement/1760695600_AI.png','2025-10-16',11),
(8,'2025-10-17 17:07:46.789','2025-10-17 17:07:46.789',NULL,'juara renang si goit','<p>gelokkk menang jier</p>','uploads/achievement/1760695666_download.jpeg','2025-10-22',11),
(9,'2025-10-17 17:26:02.881','2025-10-17 17:26:02.881',NULL,'juara 1 JHIC','dsajdsjpoadjsaodjposja','uploads/achievement/1760696762_komporr.jpg','2025-10-03',3),
(10,'2025-10-17 17:26:06.788','2025-10-17 17:26:06.788',NULL,'juara 1 JHIC','dsajdsjpoadjsaodjposja','uploads/achievement/1760696766_komporr.jpg','2025-10-03',3),
(11,'2025-10-17 17:26:10.773','2025-10-17 17:26:10.773',NULL,'juara 1 JHIC','dsajdsjpoadjsaodjposja','uploads/achievement/1760696770_komporr.jpg','2025-10-03',3);
/*!40000 ALTER TABLE `achievements` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `certifications`
--

DROP TABLE IF EXISTS `certifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `certifications` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `issuer` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `certification_number` longtext DEFAULT NULL,
  `issue_date` longtext DEFAULT NULL,
  `expiry_date` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_certifications_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `certifications`
--

LOCK TABLES `certifications` WRITE;
/*!40000 ALTER TABLE `certifications` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `certifications` VALUES
(1,'2025-10-17 13:31:52.419','2025-10-17 13:32:44.023','2025-10-17 13:36:49.405','sertifikat bodong','dikodong','dsadlksdlksmlkajdjsadsa','uploads/certifications/1760682712_Screenshot_06-Oct_14-26-17_kitty.png','cert_djskadklsjadsadsdsa','2025-12-03','2026-12-03'),
(2,'2025-10-17 13:36:57.109','2025-10-17 13:36:57.109',NULL,'sertifikat bodong','dikodong','dsadlksdlksmlkajdjsadsa','uploads/certifications/1760683017_Screenshot_06-Oct_14-26-17_kitty.png','cert_djskadklsja','2025-12-03','2026-12-03'),
(3,'2025-10-17 13:36:58.928','2025-10-17 13:36:58.928','2025-10-19 23:37:49.984','sertifikat bodong','dikodong','dsadlksdlksmlkajdjsadsa','uploads/certifications/1760683018_Screenshot_06-Oct_14-26-17_kitty.png','cert_djskadklsja','2025-12-03','2026-12-03');
/*!40000 ALTER TABLE `certifications` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `eskuls`
--

DROP TABLE IF EXISTS `eskuls`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `eskuls` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `description` text DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `pembina` longtext DEFAULT NULL,
  `pembina_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_eskuls_deleted_at` (`deleted_at`),
  KEY `fk_teachers_eskuls` (`pembina_id`),
  CONSTRAINT `fk_teachers_eskuls` FOREIGN KEY (`pembina_id`) REFERENCES `teachers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eskuls`
--

LOCK TABLES `eskuls` WRITE;
/*!40000 ALTER TABLE `eskuls` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `eskuls` VALUES
(1,'2025-10-16 21:18:42.807','2025-10-16 21:55:08.690',NULL,'progayming','ini eskul programming','',NULL,3),
(2,'2025-10-16 21:19:33.096','2025-10-16 21:19:33.096',NULL,'programming','ini eskul programming','',NULL,1),
(3,'2025-10-16 21:19:57.291','2025-10-16 21:19:57.291',NULL,'programming','ini eskul programming','',NULL,3),
(4,'2025-10-16 21:21:32.999','2025-10-16 21:21:32.999',NULL,'programming','ini eskul programming','',NULL,3),
(5,'2025-10-16 22:01:12.413','2025-10-19 22:54:20.919',NULL,'progayming','<p>ini eskul <strong>programming</strong></p>','uploads/eskul/1760626872_wallhaven-5dww19.png',NULL,3);
/*!40000 ALTER TABLE `eskuls` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `description` text DEFAULT NULL,
  `category` enum('kunjungan_industri','tamu','acara_sekolah','ujian','libur','lainnya') DEFAULT 'tamu',
  `visibility` enum('public','private') DEFAULT 'private',
  `start_date` longtext DEFAULT NULL,
  `end_date` longtext DEFAULT NULL,
  `location` longtext DEFAULT NULL,
  `organizer` longtext DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_events_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `events` VALUES
(1,'2025-10-14 12:56:17.133','2025-10-15 08:07:16.134','2025-10-15 08:25:56.617','Pas dsuahduish','lore ipusm','tamu','private','0000-00-00','0000-00-00','nirwana','devaccto','uploads/events/1760490436_Screenshot_06-Oct_14-22-46_kitty.png'),
(2,'2025-10-14 12:56:33.440','2025-10-15 11:25:07.767',NULL,'Pas djsajdsoiajdoisajoia','doskpokdposaopdskadsa','tamu','public','0000-00-00','0000-00-00','nirwana','devaccto',''),
(3,'2025-10-14 13:02:32.259','2025-10-19 19:09:56.099',NULL,'Pas djsajdsoiajdoisajoia','doskpokdposaopdskadsa','tamu','public','0000-00-00','0000-00-00','nirwana','devaccto','uploads/events/1760421752_ai ohto.jpg'),
(4,'2025-10-17 01:10:09.631','2025-10-19 19:09:51.057',NULL,'Pas djsajdsoiajdoisajoia','doskpokdposaopdskadsa','tamu','public','2025-10-13','2025-10-14','nirwana','devaccto',''),
(5,'2025-10-17 01:11:23.017','2025-10-19 19:09:48.863',NULL,'Pas djsajdsoiajdoisajoia','doskpokdposaopdskadsa','tamu','public','2025-10-13','2025-10-14','nirwana','devaccto','uploads/events/1760638283_Screenshot_06-Oct_14-25-20_kitty.png'),
(6,'2025-10-19 19:07:21.082','2025-10-19 19:09:45.195',NULL,'Pas djsajdsoiajdoisajoia','doskpokdposaopdskadsa','tamu','public','2025-10-13','2025-10-19','nirwana','devaccto','uploads/events/1760875641_ascii-art.webp'),
(7,'2025-10-19 19:09:18.698','2025-10-19 19:09:42.841',NULL,'Maulid','Perayaan maulid nabi muhammad','tamu','public','2025-10-28','2025-10-28','Bogor','SMK Plus Pelita Nusantara','uploads/events/1760875758_Video.png'),
(8,'2025-10-19 19:09:43.390','2025-10-19 19:10:49.927',NULL,'lorem ipsum dolor sit amet','lorem ipsum dolor sit amet','tamu','public','2025-10-13','2025-10-19','nirwana','devaccto','uploads/events/1760875783_atamerica24.webp');
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `forum_post_votes`
--

DROP TABLE IF EXISTS `forum_post_votes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `forum_post_votes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `post_id` bigint(20) unsigned DEFAULT NULL,
  `vote_type` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_forum_post_votes_deleted_at` (`deleted_at`),
  KEY `fk_forum_post_votes_user` (`user_id`),
  KEY `fk_forum_post_votes_post` (`post_id`),
  CONSTRAINT `fk_forum_post_votes_post` FOREIGN KEY (`post_id`) REFERENCES `forum_posts` (`id`),
  CONSTRAINT `fk_forum_post_votes_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_post_votes`
--

LOCK TABLES `forum_post_votes` WRITE;
/*!40000 ALTER TABLE `forum_post_votes` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `forum_post_votes` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `forum_posts`
--

DROP TABLE IF EXISTS `forum_posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `forum_posts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext NOT NULL,
  `content` text DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `file` longtext DEFAULT NULL,
  `status` enum('pending','approved','rejected','hidden') DEFAULT 'pending',
  `upvote` bigint(20) DEFAULT NULL,
  `downvote` bigint(20) DEFAULT NULL,
  `reply_count` bigint(20) DEFAULT NULL,
  `is_hidden` tinyint(1) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `section_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_forum_posts_deleted_at` (`deleted_at`),
  KEY `fk_forum_posts_section` (`section_id`),
  KEY `fk_forum_posts_user` (`user_id`),
  CONSTRAINT `fk_forum_posts_section` FOREIGN KEY (`section_id`) REFERENCES `forum_sections` (`id`),
  CONSTRAINT `fk_forum_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_users_post` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_posts`
--

LOCK TABLES `forum_posts` WRITE;
/*!40000 ALTER TABLE `forum_posts` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `forum_posts` VALUES
(1,'2025-10-13 09:36:21.113','2025-10-13 10:32:59.191','2025-10-13 10:54:19.608','Gua muak sama penusssss','kalian pernah gasih muak banget sama penus','uploads/forum-posts/1760322981_ai ohto.jpg','','pending',0,0,0,0,1,2),
(2,'2025-10-13 10:32:49.028','2025-10-13 14:38:10.821',NULL,'Gua muak sama penusssss','kalian pernah gasih muak banget sama penus','','','approved',4,6,0,0,1,2),
(3,'2025-10-13 11:09:55.015','2025-10-13 11:09:55.015',NULL,'Gua muak sama penus','kalian pernah gasih muak banget sama penus','uploads/forum-posts/1760328595_ai ohto.jpg','','pending',0,0,0,0,1,2);
/*!40000 ALTER TABLE `forum_posts` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `forum_replies`
--

DROP TABLE IF EXISTS `forum_replies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `forum_replies` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `content` text NOT NULL,
  `upvote` bigint(20) DEFAULT NULL,
  `downvote` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `post_id` bigint(20) unsigned DEFAULT NULL,
  `parent_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_forum_replies_deleted_at` (`deleted_at`),
  KEY `fk_forum_replies_post` (`post_id`),
  KEY `fk_forum_replies_children` (`parent_id`),
  KEY `fk_users_replies` (`user_id`),
  CONSTRAINT `fk_forum_replies_children` FOREIGN KEY (`parent_id`) REFERENCES `forum_replies` (`id`),
  CONSTRAINT `fk_forum_replies_post` FOREIGN KEY (`post_id`) REFERENCES `forum_posts` (`id`),
  CONSTRAINT `fk_users_replies` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_replies`
--

LOCK TABLES `forum_replies` WRITE;
/*!40000 ALTER TABLE `forum_replies` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `forum_replies` VALUES
(2,'2025-10-13 15:37:21.803','2025-10-13 15:37:21.803',NULL,'lorem ipsum',0,0,1,2,NULL),
(3,'2025-10-14 08:46:07.141','2025-10-14 08:46:07.141',NULL,'lorem ipsum',NULL,NULL,1,2,NULL),
(4,'2025-10-14 10:07:12.236','2025-10-14 10:07:12.236',NULL,'lorem ipsum',NULL,NULL,1,2,NULL),
(5,'2025-10-14 10:07:30.756','2025-10-14 10:07:30.756',NULL,'lorem ipsum',NULL,NULL,1,2,2),
(6,'2025-10-14 10:11:30.765','2025-10-14 10:11:30.765',NULL,'lorem ipsum',NULL,NULL,1,2,NULL),
(7,'2025-10-14 10:11:51.009','2025-10-14 10:11:51.009',NULL,'lorem ipsum',NULL,NULL,1,2,2);
/*!40000 ALTER TABLE `forum_replies` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `forum_sections`
--

DROP TABLE IF EXISTS `forum_sections`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `forum_sections` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `slug` longtext DEFAULT NULL,
  `description` text DEFAULT NULL,
  `icon` longtext DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `idx_forum_sections_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_sections`
--

LOCK TABLES `forum_sections` WRITE;
/*!40000 ALTER TABLE `forum_sections` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `forum_sections` VALUES
(1,'2025-10-12 20:10:32.849','2025-10-13 08:29:56.424','2025-10-13 08:32:24.450','Forum Alumnisasasa','forum-alumnisasasa','forum untukl alumni','uploads/forum-section/1760318858_Screenshot_06-Oct_14-25-20_kitty.png',1),
(2,'2025-10-12 20:13:10.008','2025-10-12 20:13:10.008',NULL,'Forum Alumni','forum-alumni','forum untukl alumni','uploads/forum-section/1760274790_1400309.jpg',1),
(3,'2025-10-12 20:23:22.765','2025-10-12 20:23:22.765',NULL,'Forum Alumni','forum-alumni','forum untukl alumni','uploads/forum-section/1760275402_1400309.jpg',1),
(4,'2025-10-12 20:23:48.553','2025-10-12 20:23:48.553',NULL,'Forum Alumni','forum-alumni','forum untukl alumni','uploads/forum-section/1760275428_1400309.jpg',1),
(5,'2025-10-17 18:50:41.804','2025-10-17 18:50:41.804',NULL,'Forum Pembalap','forum-pembalap','forum untukl Pembalap','uploads/forum-section/1760701841_Motor.jpg',1),
(6,'2025-10-18 23:01:24.752','2025-10-18 23:01:24.752',NULL,'Forum Publik','forum-publik','Forum terbuka untuk semua orang','uploads/forum-section/1760803284_ascii-art.webp',1),
(7,'2025-10-18 23:02:15.393','2025-10-18 23:02:15.393',NULL,'Forum Alumni','forum-alumni','Forum khusus untuk alumni, siswa, dan guru','uploads/forum-section/1760803335_ascii-art.webp',1),
(8,'2025-10-18 23:02:31.176','2025-10-18 23:02:31.176',NULL,'Forum Siswa','forum-siswa','Forum khusus untuk siswa aktif dan guru','uploads/forum-section/1760803351_ascii-art.webp',1);
/*!40000 ALTER TABLE `forum_sections` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `guest_books`
--

DROP TABLE IF EXISTS `guest_books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `guest_books` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `instance_name` longtext DEFAULT NULL,
  `contact_person` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `proposed_date` date DEFAULT NULL,
  `status` enum('pending','approved','rejected','hidden') DEFAULT 'pending',
  `rejection_reason` text DEFAULT NULL,
  `approved_at` datetime(3) DEFAULT NULL,
  `request_date` date DEFAULT NULL,
  `show_in_calendar` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_guest_books_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `guest_books`
--

LOCK TABLES `guest_books` WRITE;
/*!40000 ALTER TABLE `guest_books` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `guest_books` VALUES
(1,'2025-10-15 13:35:03.779','2025-10-19 18:20:50.017',NULL,'Kedatangan Imam Mahdi','','','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet','0000-00-00','approved','','2025-10-19 18:20:50.016','0000-00-00',0),
(2,'2025-10-15 13:41:02.743','2025-10-17 18:27:01.234',NULL,'Kedatangan Imam Mahdi','','','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet','0000-00-00','approved','','2025-10-17 18:27:01.234','0000-00-00',0),
(3,'2025-10-15 13:45:37.374','2025-10-19 18:21:05.169',NULL,'Kedatangan Imam Mahdi','','','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet','0000-00-00','approved','','2025-10-19 18:21:05.169','0000-00-00',0),
(4,'2025-10-15 13:46:00.463','2025-10-19 18:20:54.419',NULL,'Kedatangan Imam Mahdi','','','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet','0000-00-00','approved','','2025-10-19 18:20:54.418','0000-00-00',0),
(5,'2025-10-15 13:46:49.044','2025-10-15 13:46:49.044',NULL,'Kedatangan Imam Mahdi','','','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet','0000-00-00','pending','','0000-00-00 00:00:00.000','0000-00-00',0),
(6,'2025-10-15 14:03:47.529','2025-10-15 14:03:47.529',NULL,'Kedatangan Imam Mahdi','penbus','doskaodka@.omsasa','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet','0000-00-00','pending','','0000-00-00 00:00:00.000','0000-00-00',0),
(7,'2025-10-15 14:06:16.427','2025-10-16 10:22:20.757',NULL,'Kedatangan Imam Mahdi','penbus','doskaodka@.omsasa','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet','0000-00-00','approved','lo ngger bangsat','2025-10-16 10:22:20.757','2025-10-13',0),
(8,'2025-10-15 15:00:37.351','2025-10-16 10:19:07.398',NULL,'Kedatangan Imam Mahdidsa','penbusdasd','doskaodka@.omsasadsads','kdskskidkiskia@ds.comdsadas','08161287823','lorem ipsum dolor sit ametdsadsa','0000-00-00','approved','','0000-00-00 00:00:00.000','2025-10-12',0),
(9,'2025-10-16 08:54:24.947','2025-10-16 08:54:24.947',NULL,'Kedatangan Imam Mahdi','penbus','doskaodka@.omsasa','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet',NULL,'pending','','0000-00-00 00:00:00.000','2025-10-13',0),
(10,'2025-10-17 17:33:59.442','2025-10-17 17:33:59.442',NULL,'Kedatangan Imam Mahdi','penbus','doskaodka@.omsasa','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet',NULL,'pending','','0000-00-00 00:00:00.000','2025-10-13',0),
(11,'2025-10-19 19:06:13.379','2025-10-19 19:06:13.379',NULL,'Kedatangan Imam Mahdidsa','penbusdasd','doskaodka@.omsasadsads','kdskskidkiskia@ds.comdsadas','08161287823','lorem ipsum dolor sit ametdsadsa',NULL,'pending','','0000-00-00 00:00:00.000','2025-10-19',0),
(12,'2025-10-19 19:07:37.245','2025-10-19 19:07:37.245',NULL,'Kedatangan Imam Mahdi','penbus','doskaodka@.omsasa','kdskskidkiskia@ds.com','08161287821','lorem ipsum dolor sit amet',NULL,'pending','','0000-00-00 00:00:00.000','2025-10-19',0);
/*!40000 ALTER TABLE `guest_books` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `images`
--

DROP TABLE IF EXISTS `images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `images` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_path` longtext DEFAULT NULL,
  `category` enum('berita','mading','galeri','teacher','eskul','achievement','testimonial','saprol','certification','portal','mitra','profile') DEFAULT 'galeri',
  `uploaded_by` longtext DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `event_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_images_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `images`
--

LOCK TABLES `images` WRITE;
/*!40000 ALTER TABLE `images` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `images` VALUES
(1,'2025-10-10 09:33:27.630','2025-10-10 09:33:27.630','2025-10-10 09:48:50.760','uploads/images/1760063607_1400309.jpg','galeri',NULL,NULL,NULL),
(2,'2025-10-10 09:33:51.483','2025-10-10 09:33:51.483',NULL,'uploads/images/1760063631_1400309.jpg','galeri',NULL,NULL,NULL),
(3,'2025-10-10 09:52:20.162','2025-10-10 09:52:20.162','2025-10-10 09:52:42.438','uploads/images/1760064740_1400309.jpg','galeri',NULL,'dsadsadsa',NULL),
(4,'2025-10-12 08:47:19.002','2025-10-12 08:47:19.006',NULL,'uploads/images/1760233639_1400309.jpg','galeri',NULL,'dsadsadsa',NULL),
(5,'2025-10-12 08:48:02.651','2025-10-12 08:48:02.651',NULL,'uploads/images/1760233682_1400309.jpg','galeri',NULL,'dsadsadsa',NULL),
(6,'2025-10-12 08:48:15.954','2025-10-12 08:48:15.954',NULL,'uploads/images/1760233695_1400309.jpg','galeri',NULL,'dsadsadsa',NULL),
(7,'2025-10-15 20:24:08.433','2025-10-15 20:24:08.433',NULL,'uploads/images/1760534648_♡ — Ai Ohto Icons —.jpg','galeri',NULL,'dsadsadsa',0),
(8,'2025-10-15 20:33:38.233','2025-10-15 20:33:38.233',NULL,'uploads/images/1760535218_♡ — Ai Ohto Icons —.jpg','galeri',NULL,'dsadsadsa',0),
(9,'2025-10-18 07:23:41.580','2025-10-18 07:23:41.580',NULL,'uploads/images/1760747021_IMG_9961.webp','galeri',NULL,'Duquba Pagi',0);
/*!40000 ALTER TABLE `images` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `industries`
--

DROP TABLE IF EXISTS `industries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `industries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `logo` longtext DEFAULT NULL,
  `website` longtext DEFAULT NULL,
  `description` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_industries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `industries`
--

LOCK TABLES `industries` WRITE;
/*!40000 ALTER TABLE `industries` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `industries` VALUES
(1,'2025-10-17 12:26:53.655','2025-10-19 19:20:42.101',NULL,'Indomaret','uploads/industry/1760876442_mpk.png','www.indomaret.co.id','<p>jujur gua kecewa</p>'),
(2,'2025-10-17 12:36:50.694','2025-10-17 12:36:50.694','2025-10-19 19:32:36.389','Sertifikasi guru','uploads/industry/1760679410_swappy-20251016-214909.png','https://dssd.com','disjaodjsoiajdoisajoidjsjia'),
(3,'2025-10-19 18:47:07.273','2025-10-19 18:47:07.273',NULL,'PT Telkom','uploads/industry/1760874427_1594112895830_compress_PNG Icon Telkom.png','https://www.telkom.co.id/sites','<p>Sejak awal berdiri, Telkom Indonesia berkomitmen kuat untuk membangun infrastruktur telekomunikasi dan informasi serta dunia digital Indonesia. Berikut adalah tonggak sejarah penting bagaimana Telkom memenuhi komitmennya untuk Indonesia dan Dunia.</p>');
/*!40000 ALTER TABLE `industries` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `madings`
--

DROP TABLE IF EXISTS `madings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `madings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `content` text DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `status` enum('draft','published','archieve') DEFAULT 'draft',
  `is_active` tinyint(1) DEFAULT NULL,
  `type` enum('infographic','announcement') DEFAULT 'infographic',
  PRIMARY KEY (`id`),
  KEY `idx_madings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `madings`
--

LOCK TABLES `madings` WRITE;
/*!40000 ALTER TABLE `madings` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `madings` VALUES
(1,'2025-10-16 12:36:55.175','2025-10-16 13:02:43.416','2025-10-19 20:58:46.068','pikri lutungss','adsoapkopsakjinaijhoiasj','uploads/madings/1760593015_Screenshot_06-Oct_14-22-46_kitty.png','draft',0,'infographic'),
(2,'2025-10-16 13:17:44.779','2025-10-19 21:50:08.006',NULL,'pikri lutung','<p>adsoapkopsakjinaijhoiasj</p>','uploads/madings/1760595464_Screenshot_06-Oct_14-22-46_kitty.png','published',0,'infographic'),
(3,'2025-10-16 13:17:46.013','2025-10-16 13:17:46.013','2025-10-19 20:58:35.106','pikri lutung','adsoapkopsakjinaijhoiasj','uploads/madings/1760595466_Screenshot_06-Oct_14-22-46_kitty.png','draft',0,'infographic'),
(4,'2025-10-19 20:46:53.360','2025-10-19 20:46:53.360',NULL,'pikri lutung','adsoapkopsakjinaijhoiasj','','draft',0,'infographic'),
(5,'2025-10-19 20:47:50.298','2025-10-19 20:47:50.298',NULL,'pikri lutung','adsoapkopsakjinaijhoiasj','','draft',0,'infographic'),
(6,'2025-10-19 20:57:31.736','2025-10-19 20:57:31.736',NULL,'MAULIDDD','<p>bayar bayar</p>','','draft',0,'infographic');
/*!40000 ALTER TABLE `madings` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `news`
--

DROP TABLE IF EXISTS `news`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `news` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `subtitle` longtext DEFAULT NULL,
  `thumbnail` longtext DEFAULT NULL,
  `author` longtext DEFAULT NULL,
  `description` text DEFAULT NULL,
  `slug` longtext DEFAULT NULL,
  `content` text DEFAULT NULL,
  `excerpt` longtext DEFAULT NULL,
  `status` enum('draft','published','archived') DEFAULT 'draft',
  `view_count` bigint(20) DEFAULT NULL,
  `author_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_news_deleted_at` (`deleted_at`),
  KEY `fk_users_news` (`author_id`),
  CONSTRAINT `fk_users_news` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `news`
--

LOCK TABLES `news` WRITE;
/*!40000 ALTER TABLE `news` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `news` VALUES
(1,'2025-10-08 11:28:12.921','2025-10-08 13:51:31.402','2025-10-08 19:42:04.642','Prabogos','sadsadsadsads','uploads/news/wallhaven-5dww19.png','himro','',NULL,NULL,NULL,NULL,NULL,NULL),
(2,'2025-10-08 11:38:55.162','2025-10-10 13:35:54.182','2025-10-10 13:57:56.975','Makanan MBG Bergizi Gratisssss','dsajoidjsaoidjsaoijsoijdsa','uploads/news/Screenshot_06-Oct_14-45-16_zen.png','adeli','dsadsdsadsadsdsdsadsadadsa','makanan-mbg-bergizi-gratisssss','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','archived',12,1),
(3,'2025-10-10 08:24:53.901','2025-10-10 08:24:53.901','2025-10-10 13:33:42.549','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760059493_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'dsajoidjsaoidjsaoijsoijdsa','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(4,'2025-10-10 09:11:50.591','2025-10-10 09:11:50.591','2025-10-10 13:56:21.732','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760062310_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(5,'2025-10-10 11:01:26.292','2025-10-10 11:01:26.292','2025-10-10 13:58:03.396','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760068886_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(6,'2025-10-10 11:08:20.393','2025-10-10 11:08:20.393','2025-10-10 13:58:05.641','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760069300_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(7,'2025-10-10 11:17:27.573','2025-10-10 11:17:27.573','2025-10-10 13:58:07.481','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760069847_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(8,'2025-10-10 11:19:33.213','2025-10-10 11:19:33.213','2025-10-10 13:58:10.888','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760069973_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(9,'2025-10-10 11:27:20.416','2025-10-10 11:27:20.416','2025-10-10 13:58:13.580','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760070440_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(10,'2025-10-10 11:27:38.177','2025-10-10 11:27:38.177','2025-10-10 13:58:15.874','Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760070458_Screenshot_06-Oct_14-22-46_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,1),
(11,'2025-10-16 14:20:01.246','2025-10-16 14:20:01.246',NULL,'Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760599201_Screenshot_06-Oct_14-26-05_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,7),
(12,'2025-10-16 14:20:30.925','2025-10-16 14:20:30.925',NULL,'Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760599230_Screenshot_06-Oct_14-26-05_kitty.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,7),
(13,'2025-10-17 22:27:14.606','2025-10-17 22:27:14.606',NULL,'Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760714834_Screenshot_06-Oct_14-45-16_zen.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,7),
(14,'2025-10-17 22:27:22.396','2025-10-17 22:27:22.396',NULL,'Makanan MBG Bergizi Gratis',NULL,'uploads/news/1760714842_Screenshot_06-Oct_14-45-16_zen.png',NULL,NULL,'makanan-mbg-bergizi-gratis','dsadsdsadsadsdsdsadsadadsa','dsadsdsadsadsadsadsad','draft',0,7);
/*!40000 ALTER TABLE `news` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `news_news_tags`
--

DROP TABLE IF EXISTS `news_news_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `news_news_tags` (
  `news_tag_id` bigint(20) unsigned NOT NULL,
  `news_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`news_tag_id`,`news_id`),
  KEY `fk_news_news_tags_news` (`news_id`),
  CONSTRAINT `fk_news_news_tags_news` FOREIGN KEY (`news_id`) REFERENCES `news` (`id`),
  CONSTRAINT `fk_news_news_tags_news_tag` FOREIGN KEY (`news_tag_id`) REFERENCES `news_tags` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `news_news_tags`
--

LOCK TABLES `news_news_tags` WRITE;
/*!40000 ALTER TABLE `news_news_tags` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `news_news_tags` VALUES
(2,11),
(3,11),
(5,11),
(6,11),
(9,12),
(10,12),
(11,12),
(12,12),
(9,13),
(10,13),
(11,13),
(12,13),
(9,14),
(10,14),
(11,14),
(12,14);
/*!40000 ALTER TABLE `news_news_tags` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `news_tags`
--

DROP TABLE IF EXISTS `news_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `news_tags` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_news_tags_name` (`name`),
  KEY `idx_news_tags_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `news_tags`
--

LOCK TABLES `news_tags` WRITE;
/*!40000 ALTER TABLE `news_tags` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `news_tags` VALUES
(1,'2025-10-10 11:01:26.235','2025-10-10 11:01:26.235',NULL,'[\"edan\"'),
(2,'2025-10-10 11:01:26.289','2025-10-10 11:01:26.289',NULL,'\"icikiwir\"'),
(3,'2025-10-10 11:01:26.290','2025-10-10 11:01:26.290',NULL,'\"buston\"'),
(4,'2025-10-10 11:01:26.291','2025-10-10 11:01:26.291',NULL,'\"ghultom\"]'),
(5,'2025-10-10 11:08:20.377','2025-10-10 11:08:20.377',NULL,'\"edan\"'),
(6,'2025-10-10 11:08:20.392','2025-10-10 11:08:20.392',NULL,'\"ghultom\"'),
(7,'2025-10-10 11:37:27.274','2025-10-10 11:37:27.274',NULL,'go'),
(8,'2025-10-10 11:37:27.284','2025-10-10 11:37:27.284',NULL,'bun'),
(9,'2025-10-16 14:20:30.911','2025-10-16 14:20:30.911',NULL,'edan'),
(10,'2025-10-16 14:20:30.922','2025-10-16 14:20:30.922',NULL,'icikiwir'),
(11,'2025-10-16 14:20:30.923','2025-10-16 14:20:30.923',NULL,'lmao'),
(12,'2025-10-16 14:20:30.924','2025-10-16 14:20:30.924',NULL,'ikiw');
/*!40000 ALTER TABLE `news_tags` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `portals`
--

DROP TABLE IF EXISTS `portals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `portals` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `url` longtext DEFAULT NULL,
  `logo` longtext DEFAULT NULL,
  `category` longtext DEFAULT NULL,
  `is_sso_enabled` tinyint(1) DEFAULT 0,
  `is_active` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `idx_portals_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `portals`
--

LOCK TABLES `portals` WRITE;
/*!40000 ALTER TABLE `portals` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `portals` VALUES
(1,'2025-10-17 09:35:22.737','2025-10-17 09:50:36.552','2025-10-17 10:04:26.392','Instagram','intagram penus kali','dsoaodsoai','uploads/portal/1760669436_1400309.jpg','social media',0,1),
(2,'2025-10-17 09:36:14.038','2025-10-17 09:36:14.038','2025-10-20 00:00:38.983','','intagram penus','dsoaodsoai','uploads/portal/1760668574_wallhaven-5dww19.png','social media',0,1),
(3,'2025-10-17 09:42:25.341','2025-10-20 00:00:25.359',NULL,'Instagram','intagram penus','akun oersona 3','uploads/portal/1760668945_wallhaven-5dww19.png','social media',0,1),
(4,'2025-10-18 23:05:26.507','2025-10-18 23:05:26.507',NULL,'Instagram','Akun instagram resmi SMK Plus Pelita Nusantara','https://www.instagram.com/smkpluspelitanusantara/','uploads/portal/1760803526_instagram.png','social media',0,1),
(5,'2025-10-19 23:59:57.134','2025-10-19 23:59:57.134',NULL,'Youtube','channel youtube','https://www.youtube.com/@smkpluspelitanusantara9719','uploads/portal/1760893197_Youtube_logo.png','social media',0,1);
/*!40000 ALTER TABLE `portals` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `teachers`
--

DROP TABLE IF EXISTS `teachers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `teachers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nig` longtext DEFAULT NULL,
  `full_name` longtext DEFAULT NULL,
  `position` longtext DEFAULT NULL,
  `subject` longtext DEFAULT NULL,
  `photo` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_teachers_deleted_at` (`deleted_at`),
  KEY `fk_users_teacher` (`user_id`),
  CONSTRAINT `fk_users_teacher` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teachers`
--

LOCK TABLES `teachers` WRITE;
/*!40000 ALTER TABLE `teachers` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `teachers` VALUES
(1,'2025-10-16 17:33:35.952','2025-10-16 17:33:35.952','2025-10-16 17:39:06.194','09832101','Adly Fahreza','guru','B. Inggris','','saya guru bahasa inggris',7),
(2,'2025-10-16 17:33:37.370','2025-10-16 17:33:37.370','2025-10-16 17:39:10.092','09832101','Adly Fahreza','guru','B. Inggris','','saya guru bahasa inggris',7),
(3,'2025-10-16 21:19:45.196','2025-10-16 21:19:45.196',NULL,'09832101','Adly Fahreza','guru','B. Inggris','','saya guru bahasa inggris',7),
(4,'2025-10-19 19:51:51.841','2025-10-19 20:04:10.513',NULL,'09832101','Adly Fahrezaz','guru','B. Inggris','uploads/teachers/1760879050_atamerica24.webp','saya guru bahasa inggris',8);
/*!40000 ALTER TABLE `teachers` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `testimonials`
--

DROP TABLE IF EXISTS `testimonials`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `testimonials` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `position` longtext DEFAULT NULL,
  `photo` longtext DEFAULT NULL,
  `testimonial` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_testimonials_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `testimonials`
--

LOCK TABLES `testimonials` WRITE;
/*!40000 ALTER TABLE `testimonials` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `testimonials` VALUES
(1,'2025-10-17 14:11:47.634','2025-10-17 14:17:03.005','2025-10-17 14:18:28.397','Adelidsads','ceo','','great'),
(2,'2025-10-17 14:12:36.983','2025-10-17 15:30:33.825',NULL,'Adelidsads','ceo','','great'),
(3,'2025-10-17 14:12:46.424','2025-10-17 14:12:46.424','2025-10-19 23:38:47.953','Adeli','ceo','uploads/testimonial/1760685166_swappy-20251016-214909.png','great'),
(4,'2025-10-17 14:16:56.981','2025-10-19 23:36:14.373',NULL,'fahri','software enginer','','kurikulum yg sangat membantu saya untk berkembang'),
(5,'2025-10-19 23:37:40.402','2025-10-19 23:37:40.402',NULL,'John Obama','S1 Teknik informatika','','Saya merasa terbantu dengan materi yg ada, membantu saya dalam hard skill maupun soft skill');
/*!40000 ALTER TABLE `testimonials` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `user_links`
--

DROP TABLE IF EXISTS `user_links`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_links` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `url` longtext DEFAULT NULL,
  `icon` longtext DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_links_deleted_at` (`deleted_at`),
  KEY `fk_users_user_links` (`user_id`),
  CONSTRAINT `fk_users_user_links` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_links`
--

LOCK TABLES `user_links` WRITE;
/*!40000 ALTER TABLE `user_links` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `user_links` VALUES
(1,'2025-10-11 19:35:53.013','2025-10-11 20:29:21.156','2025-10-11 21:01:39.771','bukan ig','https://instagram.comsss','uploads/user_links/1760186153_Screenshot_06-Oct_14-45-16_zen.png',1),
(2,'2025-10-11 19:43:18.489','2025-10-11 19:43:18.489',NULL,'Github','https://instagram.coms','uploads/user_links/1760186598_Screenshot_06-Oct_14-45-16_zen.png',1),
(3,'2025-10-11 19:45:22.651','2025-10-11 19:45:22.651',NULL,'Linkedin','https://instagram.comsss','uploads/user_links/1760186722_Screenshot_06-Oct_14-45-16_zen.png',1),
(4,'2025-10-11 19:48:52.276','2025-10-11 19:48:52.276',NULL,'Linkedindsadas','https://instagram.comsss','uploads/user_links/1760186932_Screenshot_06-Oct_14-45-16_zen.png',1);
/*!40000 ALTER TABLE `user_links` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `no_induk` longtext DEFAULT NULL,
  `role` enum('admin','guru','siswa','orang_tua') DEFAULT 'siswa',
  `photo_url` longtext DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `jabatan` longtext DEFAULT NULL,
  `tahun_ajaran_mulai` year(4) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `users` VALUES
(1,'2025-10-07 12:37:41.406','2025-10-07 12:37:41.406',NULL,'adeley','mitsuhaadly@gmail.com','$2a$10$r2RSwn4q.6wWOKr3BS1ikO0xCYCxBs6tHdOyYICfIrjEL/tjFxoEm',NULL,'siswa',NULL,NULL,NULL,NULL,NULL),
(3,'2025-10-07 12:41:28.571','2025-10-07 12:41:28.571',NULL,'adeleys','mitsuhaadlys@gmail.com','$2a$10$2h8UIx.75vV7OEDGm17h1.ZxLpXzdDhRECCLa./BlzznSmeaS.4uC',NULL,'siswa',NULL,NULL,NULL,NULL,NULL),
(5,'2025-10-08 13:08:17.999','2025-10-08 13:08:17.999',NULL,'adeleys','mitsuhaadlysss@gmail.com','$2a$10$/Y.42sX01LxYGo9D4KT3c.M9A1TbfLCeV6k88CWyXvD4h30XsI9Z6',NULL,'siswa',NULL,NULL,NULL,NULL,NULL),
(7,'2025-10-15 20:33:07.309','2025-10-17 14:53:51.520',NULL,'adeleyss','mitsuhaadlyss@gmail.com','$2a$10$vPcIAOuKzwkKtd8UOfB87.MNtLzhLYOGhc9sZVWGBbcl3kPUS38py','232410003','siswa','uploads/user/1760687631_Screenshot_06-Oct_14-22-46_kitty.png','018208208102','sdajoidjsaijdsa',' siswa',2025),
(8,'2025-10-17 16:41:55.347','2025-10-17 16:42:18.481',NULL,'Gheraldy','gheraldy@synchonizeteams.com','$2a$10$3gZVKiY9M1Dojq0AGbS97.0QzKCOTJCbBQyqt9ZanKhzp06PhHcmO','232410028','siswa','uploads/user/1760694138_mockupNutric.png','082163387780','sdajoidjsaijdsa','siswa',2025),
(9,'2025-10-18 21:27:52.064','2025-10-18 21:27:52.064',NULL,'john','johnobama@gmail.com','$2a$10$lZUbUMiNzx6QR/kNvwTZkOVDdKL7VRElvf5S1AWzuzVaqiO1YbilS','232410028','siswa','','018208208102','cilangkap','siswa',2025);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
commit;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2025-10-20  0:10:46
