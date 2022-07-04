-- MySQL Script generated by MySQL Workbench
-- Mon Jul  4 11:02:29 2022
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mercadofresco
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mercadofresco
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mercadofresco` DEFAULT CHARACTER SET utf8 ;
USE `mercadofresco` ;

-- -----------------------------------------------------
-- Table `mercadofresco`.`countries`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`countries` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `country_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`provinces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`provinces` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `province_name` VARCHAR(255) NOT NULL,
  `country_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `country_id_idx` (`country_id` ASC) VISIBLE,
  CONSTRAINT `fk_country_provinces`
    FOREIGN KEY (`country_id`)
    REFERENCES `mercadofresco`.`countries` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`localities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`localities` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `locality_name` VARCHAR(255) NOT NULL,
  `province_name` VARCHAR(255) NOT NULL,
  `country_name` VARCHAR(255) NOT NULL,
  `province_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `province_id_idx` (`province_id` ASC) VISIBLE,
  CONSTRAINT `fk_province_localities`
    FOREIGN KEY (`province_id`)
    REFERENCES `mercadofresco`.`provinces` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`sellers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`sellers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cid` VARCHAR(255) NOT NULL,
  `company_name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `locality_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `locality_id_idx` (`locality_id` ASC) VISIBLE,
  UNIQUE INDEX `cid_UNIQUE` (`cid` ASC) VISIBLE,
  CONSTRAINT `fk_locality_sellers`
    FOREIGN KEY (`locality_id`)
    REFERENCES `mercadofresco`.`localities` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`product_types`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`product_types` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`products` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `product_code` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `width` VARCHAR(45) NOT NULL,
  `height` DECIMAL(19,2) NOT NULL,
  `length` DECIMAL(19,2) NOT NULL,
  `net_weight` DECIMAL(19,2) NOT NULL,
  `expiration_rate` DECIMAL(19,2) NOT NULL,
  `recommended_freezing_temperature` DECIMAL(19,2) NOT NULL,
  `freezing_rate` DECIMAL(19,2) NOT NULL,
  `product_type_id` INT NOT NULL,
  `seller_id` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `seller_id_idx` (`seller_id` ASC) VISIBLE,
  INDEX `product_type_id_idx` (`product_type_id` ASC) VISIBLE,
  UNIQUE INDEX `product_code_UNIQUE` (`product_code` ASC) VISIBLE,
  CONSTRAINT `fk_seller_products`
    FOREIGN KEY (`seller_id`)
    REFERENCES `mercadofresco`.`sellers` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_product_type_products`
    FOREIGN KEY (`product_type_id`)
    REFERENCES `mercadofresco`.`product_types` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`warehouses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`warehouses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `warehouse_code` VARCHAR(255) NOT NULL,
  `minimun_capacity` INT NOT NULL,
  `minimun_temperature` DECIMAL(19,2) NOT NULL,
  `locality_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `locality_id_idx` (`locality_id` ASC) VISIBLE,
  UNIQUE INDEX `warehouse_code_UNIQUE` (`warehouse_code` ASC) VISIBLE,
  CONSTRAINT `fk_locality_warehouse`
    FOREIGN KEY (`locality_id`)
    REFERENCES `mercadofresco`.`localities` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`sections`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`sections` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `section_number` INT NOT NULL,
  `current_temperature` DECIMAL(19,2) NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
  `current_capacity` INT NOT NULL,
  `minimum_capacity` INT NOT NULL,
  `maximum_capacity` INT NOT NULL,
  `warehouse_id` INT NOT NULL,
  `product_type_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_type_id_idx` (`product_type_id` ASC) VISIBLE,
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  UNIQUE INDEX `section_number_UNIQUE` (`section_number` ASC) VISIBLE,
  CONSTRAINT `fk_product_type_sections`
    FOREIGN KEY (`product_type_id`)
    REFERENCES `mercadofresco`.`product_types` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_warehouse_sections`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`product_baches`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`product_baches` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `batch_number` VARCHAR(255) NOT NULL,
  `current_quantity` INT NOT NULL,
  `current_temperature` DECIMAL(19,2) NOT NULL,
  `due_date` DATETIME(6) NOT NULL,
  `initial_quantity` INT NOT NULL,
  `manufacturing_date` DATETIME(6) NOT NULL,
  `manufacturing_hour` DATETIME(6) NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
  `product_id` INT NOT NULL,
  `section_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_id_idx` (`product_id` ASC) VISIBLE,
  INDEX `section_id_idx` (`section_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_product_baches`
    FOREIGN KEY (`product_id`)
    REFERENCES `mercadofresco`.`products` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_section_product_baches`
    FOREIGN KEY (`section_id`)
    REFERENCES `mercadofresco`.`sections` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`product_records`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`product_records` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `last_update_date` DATETIME(6) NOT NULL,
  `purchase_price` DECIMAL(19,2) NOT NULL,
  `sale_price` DECIMAL(19,2) NOT NULL,
  `product_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_id_idx` (`product_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_product_records`
    FOREIGN KEY (`product_id`)
    REFERENCES `mercadofresco`.`products` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`buyers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`buyers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `card_number_id` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `card_number_id_UNIQUE` (`card_number_id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`carrier`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`carrier` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cid` VARCHAR(255) NOT NULL,
  `company_name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `locality_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `locality_id_idx` (`locality_id` ASC) VISIBLE,
  CONSTRAINT `fk_locality_carrier`
    FOREIGN KEY (`locality_id`)
    REFERENCES `mercadofresco`.`localities` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`order_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`order_status` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`purchase_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`purchase_orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `order_number` VARCHAR(255) NOT NULL,
  `order_date` DATETIME(6) NOT NULL,
  `tracking_code` VARCHAR(255) NOT NULL,
  `buyer_id` INT NOT NULL,
  `carrier_id` INT NOT NULL,
  `order_status_id` INT NOT NULL,
  `warehouse_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `buyer_id_idx` (`buyer_id` ASC) VISIBLE,
  INDEX `carrier_id_idx` (`carrier_id` ASC) VISIBLE,
  INDEX `order_status_id_idx` (`order_status_id` ASC) VISIBLE,
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  CONSTRAINT `fk_buyer_purchase_orders`
    FOREIGN KEY (`buyer_id`)
    REFERENCES `mercadofresco`.`buyers` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_carrier_purchase_orders`
    FOREIGN KEY (`carrier_id`)
    REFERENCES `mercadofresco`.`carrier` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_order_status_purchase_orders`
    FOREIGN KEY (`order_status_id`)
    REFERENCES `mercadofresco`.`order_status` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_warehouse_purchase_orders`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`order_details`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`order_details` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `clean_liness_status` VARCHAR(255) NOT NULL,
  `quantity` INT NOT NULL,
  `temperature` DECIMAL(19,2) NOT NULL,
  `product_record_id` INT NOT NULL,
  `purchase_order_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `product_record_id_idx` (`product_record_id` ASC) VISIBLE,
  INDEX `purchase_order_id_idx` (`purchase_order_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_record_order_details`
    FOREIGN KEY (`product_record_id`)
    REFERENCES `mercadofresco`.`product_records` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_purchase_order_order_details`
    FOREIGN KEY (`purchase_order_id`)
    REFERENCES `mercadofresco`.`purchase_orders` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`employees`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`employees` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `card_number_id` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `warehouse_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  UNIQUE INDEX `card_number_id_UNIQUE` (`card_number_id` ASC) VISIBLE,
  CONSTRAINT `fk_warehouse_employees`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`inbound_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`inbound_orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `order_date` DATETIME(6) NOT NULL,
  `order_number` VARCHAR(255) NOT NULL,
  `employee_id` INT NOT NULL,
  `product_batch_id` INT NOT NULL,
  `warehouse_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `employee_id_idx` (`employee_id` ASC) VISIBLE,
  INDEX `product_batch_id_idx` (`product_batch_id` ASC) VISIBLE,
  INDEX `warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  CONSTRAINT `fk_employee_inbound_orders`
    FOREIGN KEY (`employee_id`)
    REFERENCES `mercadofresco`.`employees` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_product_batch_inbound_orders`
    FOREIGN KEY (`product_batch_id`)
    REFERENCES `mercadofresco`.`product_baches` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_warehouse_inbound_orders`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`roles`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`roles` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  `rol_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `passoword` VARCHAR(255) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercadofresco`.`user_rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`user_rol` (
  `usuario_id` INT NOT NULL AUTO_INCREMENT,
  `rol_id` INT NOT NULL,
  INDEX `usuario_id_idx` (`usuario_id` ASC) VISIBLE,
  INDEX `rol_id_idx` (`rol_id` ASC) VISIBLE,
  CONSTRAINT `fk_usuario_user_rol`
    FOREIGN KEY (`usuario_id`)
    REFERENCES `mercadofresco`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_rol_user_rol`
    FOREIGN KEY (`rol_id`)
    REFERENCES `mercadofresco`.`roles` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
