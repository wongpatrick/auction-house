CREATE SCHEMA `auction` ;

-- ------------------------------------------------------------------------------------
-- The following tables are player/game data that are spread across different databases
-- ------------------------------------------------------------------------------------
-- Quick Note: 
--    - I created these tables more for my sake visually so I could see how they are linked
--    - the foriegn key is for testing purposes and not needed if it is across different databases

CREATE TABLE `auction`.`user` (
  `user_id` INT NOT NULL,
  `name` VARCHAR(50) NULL,
  `money` INT NULL,
  PRIMARY KEY (`user_id`));


CREATE TABLE `auction`.`rarity` (
  `rarity_id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  PRIMARY KEY (`rarity_id`));

CREATE TABLE `auction`.`item_type` (
  `item_type_id` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  PRIMARY KEY (`item_type_id`));

CREATE TABLE `auction`.`item` (
  `item_id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `rarity_id` INT NULL,
  `item_type_id` INT NULL,
  `description` VARCHAR(255) NULL,
  PRIMARY KEY (`item_id`),
  INDEX `rarity_id_idx` (`rarity_id` ASC) VISIBLE,
  INDEX `item_type_id_idx` (`item_type_id` ASC) VISIBLE,
  CONSTRAINT `rarity_id`
    FOREIGN KEY (`rarity_id`)
    REFERENCES `auction`.`rarity` (`rarity_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `item_type_id`
    FOREIGN KEY (`item_type_id`)
    REFERENCES `auction`.`item_type` (`item_type_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

-- ---------------------------------------------------------------------
-- The following tables are expected to be in the auction house database
-- ---------------------------------------------------------------------
-- Quick Note: the foriegn key is for testing purposes and not needed if it is across different databases

CREATE TABLE `auction`.`listing` (
  `listing_id` INT NOT NULL,
  `item_id` INT NULL,
  `seller_id` INT NULL,
  `starting_price` INT NULL,
  `buyout_price` INT NULL,
  `created_at` DATETIME NULL,
  `ended_at` DATETIME NULL,
  `status` CHAR(1) NULL,
  PRIMARY KEY (`listing_id`),
  INDEX `item_id_idx` (`item_id` ASC) VISIBLE,
  INDEX `seller_id_idx` (`seller_id` ASC) VISIBLE,
  CONSTRAINT `item_id`
    FOREIGN KEY (`item_id`)
    REFERENCES `auction`.`item` (`item_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `seller_id`
    FOREIGN KEY (`seller_id`)
    REFERENCES `auction`.`user` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

CREATE TABLE `auction`.`bid` (
  `bid_id` INT NOT NULL,
  `listing_id` INT NULL,
  `bidder_id` INT NULL,
  `bid_amount` INT NULL,
  `created_at` DATETIME NULL,
  `is_buyout` TINYINT NULL,
  PRIMARY KEY (`bid_id`));
