SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cake_table
-- ----------------------------
DROP TABLE IF EXISTS `cake_table`;
CREATE TABLE `cake_table`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `description` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `rating` int(4) NOT NULL,
  `image` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of cake_table
-- ----------------------------
INSERT INTO `cake_table` VALUES (1, 'Lemon cheesecake 3', 'A cheesecake made of lemon', 7, 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg', '2022-02-22 13:44:56', '2022-02-23 09:23:26');
INSERT INTO `cake_table` VALUES (2, 'Lemon cheesecake 4', 'A cheesecake made of lemon', 7, 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg', '2022-02-22 20:42:01', '2022-02-23 09:24:28');

SET FOREIGN_KEY_CHECKS = 1;
