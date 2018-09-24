-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: restaurant
-- ------------------------------------------------------
-- Server version	5.5.5-10.1.33-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cuisines`
--

DROP TABLE IF EXISTS `cuisines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cuisines` (
  `cuisines_id` int(11) NOT NULL AUTO_INCREMENT,
  `cuisines_name` varchar(100) NOT NULL,
  PRIMARY KEY (`cuisines_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cuisines`
--

LOCK TABLES `cuisines` WRITE;
/*!40000 ALTER TABLE `cuisines` DISABLE KEYS */;
INSERT INTO `cuisines` VALUES (1,'chinese'),(2,'indian'),(3,'indonesian'),(4,'italian'),(5,'japanese'),(6,'korean'),(7,'pizza'),(8,'seafood'),(9,'sushi'),(10,'french');
/*!40000 ALTER TABLE `cuisines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_name` varchar(100) NOT NULL,
  `customer_phone` varchar(30) NOT NULL,
  `customer_email` varchar(100) NOT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES (1,'Boan','085210549044','boantuapasaribu@gmail.com');
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operational`
--

DROP TABLE IF EXISTS `operational`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `operational` (
  `operational_id` int(11) NOT NULL AUTO_INCREMENT,
  `operational_restaurant_id` int(11) NOT NULL,
  `operational_day` varchar(20) NOT NULL,
  `operational_open_hour` varchar(50) NOT NULL,
  `operational_closed_hour` varchar(50) NOT NULL,
  PRIMARY KEY (`operational_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operational`
--

LOCK TABLES `operational` WRITE;
/*!40000 ALTER TABLE `operational` DISABLE KEYS */;
INSERT INTO `operational` VALUES (1,1,'Saturday','10:30AM','10:30PM'),(2,1,'Sunday','10:30AM','10:30PM'),(3,1,'Monday','10:30AM','10:30PM'),(4,1,'Tuesday','10:30AM','10:30PM'),(5,1,'Wednesday','10:30AM','10:30PM'),(6,1,'Thursday','10:30AM','10:30PM'),(7,1,'Friday','10:30AM','10:30PM'),(8,2,'Saturday','10:30AM','10:30PM'),(9,2,'Sunday','10:30AM','10:30PM'),(10,2,'Monday','10:30AM','10:30PM'),(11,2,'Tuesday','10:30AM','10:30PM'),(12,2,'Wednesday','10:30AM','10:30PM'),(13,2,'Thursday','10:30AM','10:30PM'),(14,2,'Friday','10:30AM','10:30PM'),(15,3,'Saturday','10:30AM','10:30PM'),(16,3,'Sunday','10:30AM','10:30PM'),(17,3,'Monday','10:30AM','10:30PM'),(18,3,'Tuesday','10:30AM','10:30PM'),(19,3,'Wednesday','10:30AM','10:30PM'),(20,3,'Thursday','10:30AM','10:30PM'),(21,3,'Friday','10:30AM','10:30PM');
/*!40000 ALTER TABLE `operational` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reservation`
--

DROP TABLE IF EXISTS `reservation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reservation` (
  `reservation_id` int(11) NOT NULL AUTO_INCREMENT,
  `reservation_restaurant_id` int(11) NOT NULL,
  `reservation_code` varchar(50) NOT NULL,
  `reservation_total_guest` varchar(50) NOT NULL,
  `reservation_datetime` datetime NOT NULL,
  `reservation_customer_id` int(11) NOT NULL,
  `reservation_customer_name` varchar(50) NOT NULL,
  `reservation_customer_phone` varchar(30) NOT NULL,
  PRIMARY KEY (`reservation_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reservation`
--

LOCK TABLES `reservation` WRITE;
/*!40000 ALTER TABLE `reservation` DISABLE KEYS */;
INSERT INTO `reservation` VALUES (1,1,'1234567','2','2018-09-22 10:15:28',1,'Boan','085210549044'),(2,1,'7','','2018-09-25 12:00:00',0,'Boan',''),(3,2,'47382505','3','2018-09-26 01:00:00',1,'Boan','085210549044'),(4,2,'01322244','5','2018-09-27 03:00:00',1,'Boan','12343242524'),(5,3,'94464143','4','2018-09-30 02:00:00',1,'Boan','0865251421'),(6,1,'17885251','','2018-09-22 10:15:28',0,'Boan','08521054900'),(7,3,'28753318','4','2018-09-27 01:00:00',1,'Boan','12321312321'),(8,3,'86496495','3','2018-09-25 04:00:00',1,'Boan','12345'),(9,1,'85186993','','0000-00-00 00:00:00',0,'Boan','08521054900'),(10,1,'26442134','','0000-00-00 00:00:00',0,'Boan','08521054900'),(11,1,'51797753','5','2018-09-22 10:15:28',1,'Boan','08521054900'),(12,1,'36523429','5','2018-09-22 10:15:28',1,'Boan','08521054900'),(13,3,'88663523','3','2018-09-26 02:00:00',1,'Test','324234234'),(14,2,'44841937','3','2018-09-28 03:00:00',1,'test','325435656'),(15,3,'56118917','5','2018-09-26 12:00:00',1,'Boan Tua Pasaribu','085210549044'),(16,2,'56312493','4','2018-09-26 01:00:00',1,'test','23423423'),(17,1,'23525424','1','2018-09-30 05:00:00',1,'baoan','7766777777'),(18,3,'18667963','3','2018-09-27 12:00:00',1,'Boan','342423423'),(19,2,'47155241','4','2018-09-26 01:00:00',1,'Boan Test','12342323234'),(20,2,'39289779','4','2018-09-27 05:00:00',1,'Boan Final Test','12345678'),(21,3,'21367966','4','2018-09-26 10:00:00',1,'Boan Test Lagi','2342423423'),(22,3,'28348385','4','2018-09-25 11:00:00',1,'Test','1323423423');
/*!40000 ALTER TABLE `reservation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `restaurant`
--

DROP TABLE IF EXISTS `restaurant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `restaurant` (
  `restaurant_id` int(11) NOT NULL AUTO_INCREMENT,
  `restaurant_name` varchar(100) NOT NULL,
  `restaurant_url` varchar(255) NOT NULL,
  `restaurant_description` text NOT NULL,
  `restaurant_address` text NOT NULL,
  `restaurant_phone` varchar(30) NOT NULL,
  `restaurant_location` varchar(100) NOT NULL,
  `restaurant_cuisines_id` int(11) NOT NULL,
  `restaurant_latitude` double NOT NULL,
  `restaurant_longitude` double NOT NULL,
  `restaurant_image` varchar(255) NOT NULL,
  PRIMARY KEY (`restaurant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restaurant`
--

LOCK TABLES `restaurant` WRITE;
/*!40000 ALTER TABLE `restaurant` DISABLE KEYS */;
INSERT INTO `restaurant` VALUES (1,'Emilie French Restaurant','emillie','EMILIE OFFERS AN ARTFUL BLEND OF MODERN YET CLASSICALLY INSPIRED FRENCH CUISINE. SIMPLE, YET HIGH ATTENTION TO DETAIL, EACH DISH BENEFITS FROM THE BEST OF \r\nSEASONAL PRODUCE. ','Jalan Senopati 39, Jakarta Capital Region, RT.6/RW.3, Senayan, Kby. Baru, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12190','(021)5213626','Jakarta',10,-6.231031,106.809259,'french_jakarta.jpg'),(2,'Mikawa (Japanese Restaurant)','mikawa','Mikawa Japanese Sake Bar & Restaurant offers what Japan cuisine has to brag to their food about with its mouth-watering dishes and delectable foods that could satisfy your cravings with variety of selection of Sake, Sashimi, Sushi, Tempura, Yakitori and many more. This is a place to rest your mind and eat to your satisfaction. .','Jl. Sukaraja II No.32, Sukaraja, Cicendo, Kota Bandung, Jawa Barat 40175','(022)6073983','Bandung',5,-6.892334,107.574673,'japan_bandung.jpg'),(3,'Saung Kuring Sundanese Restaurant','saung_kuring','Sundanese traditional food from indonesia','Jalan KH Sholeh Iskandar No.9, Kedung Badak, Tanah Sereal, Kedung Badak, Tanah Sereal, Kota Bogor, Jawa Barat 16164','(0251)8331885','Bogor',3,-6.561851,106.799155,'indonesia_bogor.jpg');
/*!40000 ALTER TABLE `restaurant` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-09-24 21:46:29
