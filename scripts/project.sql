

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_auth_node
-- ----------------------------
DROP TABLE IF EXISTS `t_auth_node`;
CREATE TABLE `t_auth_node`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父级ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由规则',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型:1=目录,2=菜单,3=按钮',
  `sort` int(11) NOT NULL DEFAULT 50 COMMENT '排序',
  `create_time` int(10) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` int(10) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单节点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_auth_node
-- ----------------------------
INSERT INTO `t_auth_node` VALUES (1, 0, '权限管理', 'el-icon-s-tools', '12312', 1, 1, 0, 1652450171);
INSERT INTO `t_auth_node` VALUES (18, 1, '管理员', '#', '/user/list', 2, 1, 0, 1652361618);
INSERT INTO `t_auth_node` VALUES (19, 1, '角色', '#', '/role/list', 2, 1, 0, 1652361631);
INSERT INTO `t_auth_node` VALUES (20, 1, '资源', '#', '/menu/list', 2, 1, 0, 1652449176);
INSERT INTO `t_auth_node` VALUES (21, 18, '新增管理员', '#', '/v1/user/create', 3, 1, 0, NULL);
INSERT INTO `t_auth_node` VALUES (22, 18, '编辑管理员', '#', '/v1/user/update', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (23, 18, '管理员列表', '#', '/user/list', 2, 1, NULL, 1652361595);
INSERT INTO `t_auth_node` VALUES (24, 18, '设置管理员状态', '#', '/v1/user/status', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (25, 18, '修改管理员密码', '#', '/v1/user/change/password', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (26, 18, '删除管理员', '#', '/v1/user/delete', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (33, 19, '新增角色', '#', '/v1/role/create', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (34, 19, '编辑角色', '#', '/v1/role/update', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (35, 19, '角色列表', '#', '/role/list', 2, 1, NULL, 1652362016);
INSERT INTO `t_auth_node` VALUES (36, 19, '删除角色', '#', '/v1/role/delete', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (37, 20, '新增资源', '#', '/v1/node/create', 3, 1, NULL, NULL);
INSERT INTO `t_auth_node` VALUES (38, 20, '资源列表', '#', '/menu/list', 2, 1, NULL, 1652362154);
INSERT INTO `t_auth_node` VALUES (52, 20, '编辑资源', '#', '/v1/node/update', 3, 3, 0, 0);
INSERT INTO `t_auth_node` VALUES (53, 20, '资源详情', '#', '/v1/node/detail', 3, 3, 0, 0);
INSERT INTO `t_auth_node` VALUES (65, 63, '定时任务', '1', '/logs/timing', 2, 2, 1657541204, 1657541204);

-- ----------------------------
-- Table structure for t_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `t_auth_role`;
CREATE TABLE `t_auth_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `sort` int(11) NOT NULL DEFAULT 50 COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '显隐:1=显示,2=隐藏',
  `create_time` int(10) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` int(10) NULL DEFAULT NULL COMMENT '更新时间',
  `is_delete` tinyint(1) NULL DEFAULT NULL COMMENT '是否删除:1=是,2=否',
  `delete_time` int(11) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_auth_role
-- ----------------------------
INSERT INTO `t_auth_role` VALUES (1, '角色名称', '描述', 2222, 0, 1651680978, 1652190989, 2, 1652177827);
INSERT INTO `t_auth_role` VALUES (2, '1111', '111111111', 1, 0, 1651680978, 1651805591, 2, NULL);
INSERT INTO `t_auth_role` VALUES (3, '6666', '8888', 111, 0, 1651834196, 1651834414, 2, 1652177825);
INSERT INTO `t_auth_role` VALUES (4, '1', '', 0, 0, 1652019835, 0, 1, 1652177779);
INSERT INTO `t_auth_role` VALUES (5, '1ggg', 'test', 100, 0, 1652019957, 0, 1, 1652177759);
INSERT INTO `t_auth_role` VALUES (6, '角色111', '角色11', 1, 0, 1652178004, 0, 2, 0);
INSERT INTO `t_auth_role` VALUES (7, '系统管理', '系统管理', 500000, 0, 1652186914, 1657272920, 2, 0);

-- ----------------------------
-- Table structure for t_auth_role_node
-- ----------------------------
DROP TABLE IF EXISTS `t_auth_role_node`;
CREATE TABLE `t_auth_role_node`  (
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `node_id` int(11) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`role_id`, `node_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色权限关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_auth_role_node
-- ----------------------------
INSERT INTO `t_auth_role_node` VALUES (1, 1);
INSERT INTO `t_auth_role_node` VALUES (1, 2);
INSERT INTO `t_auth_role_node` VALUES (1, 3);
INSERT INTO `t_auth_role_node` VALUES (1, 4);
INSERT INTO `t_auth_role_node` VALUES (5, 1);
INSERT INTO `t_auth_role_node` VALUES (5, 20);
INSERT INTO `t_auth_role_node` VALUES (5, 38);
INSERT INTO `t_auth_role_node` VALUES (6, 1);
INSERT INTO `t_auth_role_node` VALUES (6, 18);
INSERT INTO `t_auth_role_node` VALUES (6, 20);
INSERT INTO `t_auth_role_node` VALUES (6, 21);
INSERT INTO `t_auth_role_node` VALUES (6, 22);
INSERT INTO `t_auth_role_node` VALUES (6, 23);
INSERT INTO `t_auth_role_node` VALUES (6, 24);
INSERT INTO `t_auth_role_node` VALUES (6, 25);
INSERT INTO `t_auth_role_node` VALUES (6, 26);
INSERT INTO `t_auth_role_node` VALUES (6, 37);
INSERT INTO `t_auth_role_node` VALUES (6, 38);
INSERT INTO `t_auth_role_node` VALUES (7, 54);
INSERT INTO `t_auth_role_node` VALUES (7, 55);

-- ----------------------------
-- Table structure for t_auth_user
-- ----------------------------
DROP TABLE IF EXISTS `t_auth_user`;
CREATE TABLE `t_auth_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `is_root` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否超级管理员:1=是,2=否',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '名称',
  `account` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `create_time` int(10) NOT NULL COMMENT '创建时间',
  `update_time` int(10) NOT NULL COMMENT '修改时间',
  `login_time` int(10) NULL DEFAULT NULL COMMENT '最后登录时间',
  `login_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '最后登录ip',
  `disable` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用:1=是,2=否',
  `delete_time` int(11) NOT NULL DEFAULT 0 COMMENT '0为非删除状态,非0位删除时间',
  `login_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录详细地址',
  `is_delete` tinyint(1) NULL DEFAULT NULL COMMENT '是否删除:1=是,2=否',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_auth_user
-- ----------------------------
INSERT INTO `t_auth_user` VALUES (1, 1, 'daidai', 'daidai', 'dc2274734cf250cb7763e6ebb8ef3459', '4JqDI0m3IG9flrtiSpgBQcaUBJmB41sw', 0, 1652186925, 0, '0', 2, 0, '0', 2);
INSERT INTO `t_auth_user` VALUES (15, 2, 'ddd', 'Account', '457ac0468b45878ea9021d8b6cce24af', 'PM8HOlFla7lgfUVrLvemcMfypPxtH384', 0, 1658324612, 0, '', 1, 0, '', 2);
INSERT INTO `t_auth_user` VALUES (16, 2, '666', 'test', 'bdb199d25b2487a5bfaa3008efb15251', 'dHgQ5rTNVC2sgBmgpaEKFVfq5TEfGZHG', 0, 1652185014, 0, '', 1, 0, '', 2);
INSERT INTO `t_auth_user` VALUES (17, 2, 'ddddddddd', 'test11111', '948b56382d96b9d2d6d8637c1c6adf09', 'XGfXkjdtO2KPvVDco3ExPIvjNjvbYQXo', 0, 1658323695, 0, '', 1, 0, '', 2);
INSERT INTO `t_auth_user` VALUES (18, 2, 'hukai', 'hukai', '945ff70c3b4576c2e640285f60d4b739', '3uDaUVj753gtlye16ZApASWimf1eb5zA', 0, 1658324609, 0, '', 1, 1652083309, '', 2);
INSERT INTO `t_auth_user` VALUES (19, 2, '胡凯111', 'hukai11', '9cb54c9e7f626510443aa242b19a921a', 'RvCjEAo08viPhkP0mYpYKE8BubfhL048', 0, 1658324607, 0, '', 1, 0, '', 2);
INSERT INTO `t_auth_user` VALUES (20, 2, '漆超', 'qichao', '2d6a19211626172b6d002646e5cf9d7b', 'JcX0oV1FEpLwTfNBzDcvwF6rTWm3Vx2h', 0, 1652184975, 0, '', 2, 0, '', 2);
INSERT INTO `t_auth_user` VALUES (21, 2, 'jiajia', 'jiajia', '4403171446a02da351646cc98b00d7bf', 'KimlfTptTer4ee3fjaaC0924y98s5Oxa', 0, 0, 0, '', 2, 0, '', 2);

-- ----------------------------
-- Table structure for t_auth_user_role
-- ----------------------------
DROP TABLE IF EXISTS `t_auth_user_role`;
CREATE TABLE `t_auth_user_role`  (
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_auth_user_role
-- ----------------------------
INSERT INTO `t_auth_user_role` VALUES (1, 7);
INSERT INTO `t_auth_user_role` VALUES (15, 6);
INSERT INTO `t_auth_user_role` VALUES (16, 6);
INSERT INTO `t_auth_user_role` VALUES (17, 6);
INSERT INTO `t_auth_user_role` VALUES (18, 6);
INSERT INTO `t_auth_user_role` VALUES (19, 6);
INSERT INTO `t_auth_user_role` VALUES (20, 6);
INSERT INTO `t_auth_user_role` VALUES (22, 6);



-- ----------------------------
-- Records of t_mch
-- ----------------------------
INSERT INTO `t_mch` VALUES (6, '百度', 'edadaibcjgyvvjisngfoodycrxpvfvxq', 'pphiymrsfkhixvivpqaqfbfmtyqscqqj', 1, 1652865071, 1652865071, 2, 0, 1123, 'www.baidu.com', 'qyglMvRwH3oodB5BZyglMvRwH3oodB5B', '1nvPy5fzgem0lB6O');
INSERT INTO `t_mch` VALUES (7, '星力百货', 'jrzqwftgzjhywmjhcqpvkqhdkubnmyor', 'bafamhyrfljvixgcmwmjnbxbadplkjqc', 1, 1652865119, 1652865119, 2, 0, 12, 'test.xinglico.com', 'myglMvRwH3oodB5BZyglMvRwH3oodB5B', 'ZyglMvRwH3oodB5B');
INSERT INTO `t_mch` VALUES (8, '千百度', 'qnriswrfndvyegtmjdixjtqwqhoijjqf', 'qcfpussbvxdhxlfsbdflyklsqanmztaj', 1, 1652871388, 1652871388, 2, 0, 1123, 'www.baidu.com', 'ZyglMvRwH3oodB5BZyglMvRwH3oodB5B', 'ZyglMvRwH3oodB66');
INSERT INTO `t_mch` VALUES (9, '7千百度', 'jjjmywlqkakcanqdtbstjwtwcibunimc', 'dcbtdfhtztnwcnyfgwrjlcleottirqjc', 1, 1652875528, 1652875528, 2, 0, 1123, 'www.baidu.com', 'oyglMvRwH3oodB5BZyglMvRwH3oodB5B', '88glMvRwH3oodB5B');
INSERT INTO `t_mch` VALUES (10, '我是商户1144441', 'hybwjebbywgmvpxphnmwlbdyrwlieuko', 'cqxftlojeopobxbyopcdpwafaumtzfdh', 1, 1653892607, 1653892607, 2, 0, 1123, 'www.baidu.com', 'iyglMvRwH3oodB5BZyglMvRwH3oodB5B', 'ljiepvdxohhvukfg');
INSERT INTO `t_mch` VALUES (11, 'sdfs', 'nmylxoimtqtmurmhrtbpmcaxefpsxlzr', 'fgeynqtfyatlkemqyjhkkjxhnpqthhnw', 1, 1656927553, 1656927553, 2, 0, 23, 'https://element.eleme.io/#/zh-CN/component/input-number', 'pzcsfdtzjfkhmfmoqgpwaphbcuzzpesx', 'yhsiwvrxdqxtgfda');
INSERT INTO `t_mch` VALUES (12, '代总', 'eywrkwdrpxswzekcvhjhduzbnvujusol', 'bncnkyqpjqwnztdixteiijhzsmqnpmnc', 1, 1657178132, 1657178132, 2, 0, 1, 'https://uniapp.dcloud.io/plugin/', 'mnvcoazaboxkmcjnoyktimmuvrmrlsjk', 'pzbxnbkyhozxaznb');
