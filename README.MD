# 1 脚本使用方法

需要在数据库里创建一个表：node_info ,数据库的连接地址在 database/core.go里面

sql如下  
```sql
/*
 Navicat Premium Data Transfer

 Source Server         : argo
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : 192.168.44.134:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 17/04/2024 14:41:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for node_info
-- ----------------------------
DROP TABLE IF EXISTS `node_info`;
CREATE TABLE `node_info`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服务器Ip',
  `datetime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服务器时间',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '服务器状态',
  `started_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '开机时间',
  `user_number` int(0) NULL DEFAULT NULL COMMENT '运行的用户数量 ',
  `loadaverage_1m` decimal(10, 2) NULL DEFAULT NULL COMMENT '1分钟负载',
  `loadaverage_5m` decimal(10, 2) NULL DEFAULT NULL COMMENT '5分钟负载',
  `loadaverage_15m` decimal(10, 2) NULL DEFAULT NULL COMMENT '15分钟负载',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of node_info
-- ----------------------------
INSERT INTO `node_info` VALUES (1, '192.168.44.134', '06:34:55', 'up', '5:25', 3, 0.00, 0.00, 0.00);
INSERT INTO `node_info` VALUES (2, '192.168.44.134', '06:37:42', 'up', '5:28', 3, 1.45, 1.41, 1.31);
INSERT INTO `node_info` VALUES (3, '192.168.44.134', '06:39:51', 'up', '5:30', 3, 1.45, 1.40, 1.32);

SET FOREIGN_KEY_CHECKS = 1;


```


```powershell
go run main.go sysinfo -i 192.168.44.134 -u root -P 22 -p 1 -c "uptime" --excel true #输入IP，用户，密码信息，以及要执行的 命令，这里只对 uptime进行了输出至数据库，这里的 "可以是任何命令，简单的实现跳板机功能,添加了 --excel true就会打印数据库的信息"
```