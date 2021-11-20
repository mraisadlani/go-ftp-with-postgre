-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               PostgreSQL 12.8 (Debian 12.8-1.pgdg110+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
-- Server OS:                    
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE "testing-db";

-- Dumping structure for table public.products
CREATE TABLE IF NOT EXISTS "products" (
	"product_id" INTEGER NOT NULL DEFAULT 'nextval(''products_product_id_seq''::regclass)',
	"product_code" VARCHAR(50) NULL DEFAULT NULL,
	"product_name" VARCHAR(200) NULL DEFAULT NULL,
	"product_slug" VARCHAR(200) NULL DEFAULT NULL,
	"product_description" TEXT NULL DEFAULT NULL,
	"qty" INTEGER NULL DEFAULT NULL,
	"min_qty" INTEGER NULL DEFAULT NULL,
	"max_qty" INTEGER NULL DEFAULT NULL,
	"weight" INTEGER NULL DEFAULT NULL,
	"volume" INTEGER NULL DEFAULT NULL,
	"create_at" TIMESTAMP NULL DEFAULT NULL,
	"update_at" TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY ("product_id")
);

-- Data exporting was unselected.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
