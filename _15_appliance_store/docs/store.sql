CREATE DATABASE IF NOT EXISTS appliance_store DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE appliance_store;

-- ----------------------------
-- 产品分类表（多级分类）
-- ----------------------------
CREATE TABLE `product_category` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类ID',
	`parent_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父分类ID(0表示顶级)',
	`name` VARCHAR(50) NOT NULL COMMENT '分类名称',
	`level` TINYINT UNSIGNED NOT NULL COMMENT '分类层级(1/2/3)',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	INDEX `idx_parent` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品分类表（支持三级分类）';

-- ----------------------------
-- 产品基本信息表
-- ----------------------------
CREATE TABLE `product` (
   `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '产品ID',
   `name` VARCHAR(100) NOT NULL COMMENT '产品名称',
   `category_id` INT UNSIGNED NOT NULL COMMENT '分类ID',
   `brand` VARCHAR(50) NOT NULL COMMENT '品牌',
   `model` VARCHAR(50) NOT NULL COMMENT '型号',
   `spec` VARCHAR(100) COMMENT '规格',
   `description` TEXT COMMENT '详细描述',
   `price` DECIMAL(10,2) NOT NULL COMMENT '价格',
   `warranty_months` SMALLINT UNSIGNED COMMENT '保修月数',
   `service_terms` TEXT COMMENT '售后服务条款',
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`),
   FOREIGN KEY (`category_id`) REFERENCES `product_category`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品基本信息表';

-- ----------------------------
-- 产品多媒体表
-- ----------------------------
CREATE TABLE `product_media` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '媒体ID',
	`product_id` INT UNSIGNED NOT NULL COMMENT '产品ID',
	`type` TINYINT NOT NULL COMMENT '类型(1图片/2视频/33D模型)',
	`url` VARCHAR(255) NOT NULL COMMENT '资源地址',
	`sort_order` SMALLINT UNSIGNED DEFAULT 0 COMMENT '排序序号',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`product_id`) REFERENCES `product`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品多媒体资源表';

-- ----------------------------
-- SKU库存单元表
-- ----------------------------
CREATE TABLE `product_sku` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'SKU ID',
	`product_id` INT UNSIGNED NOT NULL COMMENT '产品ID',
	`sku_code` VARCHAR(50) NOT NULL COMMENT 'SKU编码',
	`attributes` JSON NOT NULL COMMENT '规格属性(JSON格式，如{"color":"白","capacity":"10L"})',
	`price` DECIMAL(10,2) NOT NULL COMMENT '实际售价',
	`stock` INT DEFAULT 0 COMMENT '库存数量',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`product_id`) REFERENCES `product`(`id`),
	UNIQUE `unq_sku` (`sku_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='SKU库存单元表';

-- ----------------------------
-- 门店信息表
-- ----------------------------
CREATE TABLE `store` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '门店ID',
	`name` VARCHAR(100) NOT NULL COMMENT '门店名称',
	`address` VARCHAR(255) NOT NULL COMMENT '详细地址',
	`contact_phone` VARCHAR(20) NOT NULL COMMENT '联系电话',
	`manager_id` INT UNSIGNED COMMENT '店长ID',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='门店信息表';


-- ----------------------------
-- 供应商表
-- ----------------------------
CREATE TABLE `supplier` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '供应商ID',
	`name` VARCHAR(100) NOT NULL COMMENT '供应商名称',
	`contact_person` VARCHAR(50) NOT NULL COMMENT '联系人',
	`phone` VARCHAR(20) NOT NULL COMMENT '联系电话',
	`payment_terms` TEXT COMMENT '结算条款',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='供应商信息表';

-- ------------------------------------------
-- 库存表（实时库存追踪及安全库存设置）
-- ------------------------------------------
CREATE TABLE `inventory` (
	`id` BIGINT AUTO_INCREMENT COMMENT '库存记录ID',
	`store_id` INT UNSIGNED NOT NULL COMMENT '门店ID',
	`sku_id` INT UNSIGNED NOT NULL COMMENT 'SKU ID',
	`quantity` INT NOT NULL DEFAULT 0 COMMENT '当前库存',
	`safety_stock` INT NOT NULL DEFAULT 0 COMMENT '安全库存',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`sku_id`) REFERENCES `product_sku`(`id`),
	FOREIGN KEY (`store_id`) REFERENCES `store`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='门店库存表（分店独立库存）';

-- -----------------------------------------
-- 库存操作记录表（入库、出库、调拨记录）
-- -----------------------------------------
CREATE TABLE `inventory_operation` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作ID',
	`inventory_id` BIGINT NOT NULL COMMENT '库存记录ID',
	`store_id` INT UNSIGNED NOT NULL COMMENT '门店ID',
	`sku_id` INT UNSIGNED NOT NULL COMMENT 'SKU ID',
	`type` TINYINT NOT NULL COMMENT '操作类型(1采购入库/2退货入库/3销售出库/4调拨出库/5调拨入库)',
	`quantity` INT NOT NULL COMMENT '操作数量',
	`related_order` VARCHAR(24) COMMENT '关联订单号',
	`remark` VARCHAR(500) COMMENT '备注',
	`operator` INT UNSIGNED NOT NULL COMMENT '操作人ID',
	`operate_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`inventory_id`) REFERENCES inventory(`id`),
	FOREIGN KEY (`store_id`) REFERENCES `store`(`id`),
	FOREIGN KEY (`sku_id`) REFERENCES `product_sku`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存操作记录表';

-- ----------------------------
-- 销售订单主表
-- ----------------------------
CREATE TABLE `sales_order` (
	`id` VARCHAR(24) NOT NULL COMMENT '订单号（规则：YYYYMMDDHHMMSSmmm+6位序列）',
	`store_id` INT UNSIGNED NOT NULL COMMENT '门店ID',
	`customer_id` INT UNSIGNED COMMENT '客户ID',
	`total_amount` DECIMAL(12,2) NOT NULL COMMENT '订单总额',
	`status` TINYINT NOT NULL COMMENT '状态(0待支付/1已支付/2配送中/3已完成/4已退货)',
	`payment_method` TINYINT COMMENT '支付方式(1现金/2微信/3支付宝/4刷卡)',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	INDEX `idx_store` (`store_id`),
	FOREIGN KEY (`store_id`) REFERENCES `store`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售订单主表';

use appliance_store;
-- ----------------------------
-- 销售订单明细表
-- ----------------------------
CREATE TABLE `sales_order_item` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '明细ID',
	`order_id` VARCHAR(24) NOT NULL COMMENT '订单号',
	`sku_id` INT UNSIGNED NOT NULL COMMENT 'SKU ID',
	`quantity` INT UNSIGNED NOT NULL COMMENT '购买数量',
	`price` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '产品单价',
	`unit_price` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '成交单价',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`order_id`) REFERENCES `sales_order`(`id`),
	FOREIGN KEY (`sku_id`) REFERENCES `product_sku`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细表';

-- ----------------------------
-- 采购订单主表
-- ----------------------------
CREATE TABLE `purchase_order` (
	`id` VARCHAR(24) NOT NULL COMMENT '采购单号（规则：YYYYMMDDHHMMSSmmm+6位序列）',
	`supplier_id` INT UNSIGNED NOT NULL COMMENT '供应商ID',
	`total_amount` DECIMAL(12,2) NOT NULL COMMENT '总金额',
	`status` TINYINT NOT NULL COMMENT '状态(0待审批/1已批准/2已到货)',
	`expected_date` DATE NOT NULL COMMENT '预计到货日期',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`supplier_id`) REFERENCES `supplier`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购订单主表';

-- ----------------------------
-- 采购订单明细表
-- ----------------------------
CREATE TABLE `purchase_order_item` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '明细ID',
	`order_id` VARCHAR(24) NOT NULL COMMENT '订单号',
	`sku_id` INT UNSIGNED NOT NULL COMMENT 'SKU ID',
	`quantity` INT UNSIGNED NOT NULL COMMENT '购买数量',
	`price` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '产品单价',
	`unit_price` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '成交单价',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`order_id`) REFERENCES `purchase_order`(`id`),
	FOREIGN KEY (`sku_id`) REFERENCES `product_sku`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购订单明细表';

-- -------------------------------------
-- 员工表（账号管理，密码加密存储）
-- -------------------------------------
CREATE TABLE `employee` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '员工ID',
	`name` VARCHAR(50) NOT NULL COMMENT '姓名',
	`phone` VARCHAR(20) NOT NULL COMMENT '手机号',
	`role` TINYINT NOT NULL COMMENT '角色(1管理员/2店长/3销售员/4财务)',
	`store_id` INT UNSIGNED COMMENT '所属门店ID',
	`password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希值',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`store_id`) REFERENCES `store`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工信息表';

-- ----------------------------
-- 员工绩效表
-- ----------------------------
CREATE TABLE `employee_performance` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
	`employee_id` INT UNSIGNED NOT NULL COMMENT '员工ID',
	`month` CHAR(7) NOT NULL COMMENT '统计月份(YYYY-MM)',
	`sales_amount` DECIMAL(12,2) NOT NULL DEFAULT 0 COMMENT '销售额',
	`commission` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '提成金额',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	UNIQUE `unq_emp_month` (`employee_id`, `month`),
	FOREIGN KEY (`employee_id`) REFERENCES `employee`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工月度绩效表';


-- ----------------------------
-- 客户信息表
-- ----------------------------
CREATE TABLE `customer` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '客户ID',
	`name` VARCHAR(50) NOT NULL COMMENT '客户姓名',
	`phone` VARCHAR(20) NOT NULL COMMENT '手机号',
	`email` VARCHAR(200) COMMENT '客户邮箱',
	`address` VARCHAR(200) COMMENT '地址',
	`birthday` DATE COMMENT '生日',
	`membership_level` TINYINT DEFAULT 0 COMMENT '会员等级(0普通客户)',
	`total_points` INT UNSIGNED DEFAULT 0 COMMENT '累计积分',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	UNIQUE `unq_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户信息表';

-- -----------------------------------------------
-- 促销活动表（满减、折扣、赠品、组合优惠等）
-- ------------------------------------------------
CREATE TABLE `promotion` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '活动ID',
	`name` VARCHAR(100) NOT NULL COMMENT '活动名称',
	`type` TINYINT NOT NULL COMMENT '类型(1满减/2折扣/3赠品/4组合优惠)',
	`rule` JSON NOT NULL COMMENT '活动规则（JSON格式）',
	`start_time` DATETIME NOT NULL COMMENT '开始时间',
	`end_time` DATETIME NOT NULL COMMENT '结束时间',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='促销活动表';

-- ---------------------------------------------
-- 操作日志表（记录增删改操作，审计日志）
-- ---------------------------------------------
CREATE TABLE `audit_log` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志ID',
	`operator_id` INT UNSIGNED NOT NULL COMMENT '操作人ID',
	`action_type` VARCHAR(50) NOT NULL COMMENT '操作类型',
	`target_table` VARCHAR(50) NOT NULL COMMENT '目标表名',
	`target_id` VARCHAR(50) NOT NULL COMMENT '目标记录ID',
	`old_value` JSON COMMENT '旧值',
	`new_value` JSON COMMENT '新值',
	`remark` VARCHAR(255) COMMENT '备注',
	`operate_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	PRIMARY KEY (`id`),
	INDEX `idx_operator` (`operator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志审计表';

-- -----------------------------------------------------
-- 财务流水表（销售收款、采购付款、其他费用等）
-- -----------------------------------------------------
CREATE TABLE `financial_transaction` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '流水ID',
	`type` TINYINT NOT NULL COMMENT '类型(1销售收款/2采购付款/3其他收入/4其他支出)',
	`amount` DECIMAL(12,2) NOT NULL COMMENT '金额',
	`related_id` VARCHAR(24) COMMENT '关联单据号(订单号/采购单号等)',
	`transaction_time` DATETIME NOT NULL COMMENT '交易时间',
	`operator` INT UNSIGNED NOT NULL COMMENT '操作人ID',
	`remark` VARCHAR(255) COMMENT '备注',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	INDEX `idx_related` (`related_id`),
	FOREIGN KEY (`operator`) REFERENCES `employee`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='财务收支流水表';

-- ------------------------
-- 售后服务表（记录退换货、维修等处理信息）
-- ------------------------
CREATE TABLE `after_sales` (
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '售后服务ID',
	`order_id` VARCHAR(24) NOT NULL COMMENT '订单ID',
	`type` VARCHAR(50) NOT NULL COMMENT '售后类型（退款、换货、维修）',
	`description` TEXT COMMENT '售后描述',
	`status` VARCHAR(50) COMMENT '售后状态',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	FOREIGN KEY (`order_id`) REFERENCES sales_order(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='售后服务表';

-- ------------------------
-- 维修工单表
-- ------------------------
CREATE TABLE `service_order` (
	`id` VARCHAR(24) NOT NULL COMMENT '工单号',
	`customer_id` INT UNSIGNED NOT NULL COMMENT '客户ID',
	`product_id` INT UNSIGNED NOT NULL COMMENT '产品ID',
	`fault_desc` TEXT NOT NULL COMMENT '故障描述',
	`status` TINYINT NOT NULL COMMENT '状态(0待处理/1维修中/2已完成/3已关闭)',
	`technician_id` INT UNSIGNED COMMENT '维修人员ID',
	`completion_time` DATETIME COMMENT '完成时间',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`customer_id`) REFERENCES `customer`(`id`),
	FOREIGN KEY (`product_id`) REFERENCES `product`(`id`),
	FOREIGN KEY (`technician_id`) REFERENCES `employee`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='售后服务维修工单表';

-- ----------------------------
-- 调拨单主表
-- ----------------------------
CREATE TABLE `transfer_order` (
	`id` VARCHAR(24) NOT NULL COMMENT '调拨单号',
	`from_store` INT UNSIGNED NOT NULL COMMENT '调出门店',
	`to_store` INT UNSIGNED NOT NULL COMMENT '调入门店',
	`total_quantity` INT NOT NULL COMMENT '总调拨数量',
	`status` TINYINT NOT NULL COMMENT '状态(0待处理/1已完成)',
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`complete_time` DATETIME COMMENT '完成时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`from_store`) REFERENCES `store`(`id`),
	FOREIGN KEY (`to_store`) REFERENCES `store`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存调拨单主表';

-- ----------------------------
-- 调拨明细表
-- ----------------------------
CREATE TABLE `transfer_detail` (
	`transfer_id` VARCHAR(24) NOT NULL COMMENT '调拨单号',
	`sku_id` INT UNSIGNED NOT NULL COMMENT 'SKU ID',
	`quantity` INT UNSIGNED NOT NULL COMMENT '调拨数量',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`transfer_id`, `sku_id`),
	FOREIGN KEY (`transfer_id`) REFERENCES `transfer_order`(`id`),
	FOREIGN KEY (`sku_id`) REFERENCES `product_sku`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存调拨明细表';

-- ----------------------------
-- 消息通知表
-- ----------------------------
CREATE TABLE `notification` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '通知ID',
	`type` TINYINT NOT NULL COMMENT '类型(1库存预警/2订单状态/3促销提醒)',
	`recipient` VARCHAR(100) NOT NULL COMMENT '接收人(手机号/邮箱/员工ID)',
	`content` TEXT NOT NULL COMMENT '通知内容',
	`send_status` TINYINT NOT NULL COMMENT '发送状态(0待发送/1已发送/2失败)',
	`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统消息通知表';

-- ----------------------------
-- 优惠券表
-- ----------------------------
CREATE TABLE `coupon` (
	`id` VARCHAR(20) NOT NULL COMMENT '优惠券ID',
	`type` TINYINT NOT NULL COMMENT '类型(1满减/2折扣/3代金券)',
	`value` DECIMAL(10,2) NOT NULL DEFAULT 0  COMMENT '面值/折扣率',
	`amount_condition` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '使用条件(满多少元可用)',
	`customer_id` INT UNSIGNED COMMENT '绑定客户ID',
	`status` TINYINT NOT NULL COMMENT '状态(0未发放/1未使用/2已使用/3已过期)',
	`expire_time` DATETIME NOT NULL COMMENT '过期时间',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`customer_id`) REFERENCES `customer`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券管理表';

-- ----------------------------
-- 应付账款表
-- ----------------------------
CREATE TABLE `account_payable` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
	`supplier_id` INT UNSIGNED NOT NULL COMMENT '供应商ID',
	`purchase_id` VARCHAR(24) NOT NULL COMMENT '采购单号',
	`due_amount` DECIMAL(12,2) NOT NULL COMMENT '应付款金额',
	`paid_amount` DECIMAL(12,2) NOT NULL DEFAULT 0 COMMENT '已付款金额',
	`due_date` DATE NOT NULL COMMENT '应付款日期',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`supplier_id`) REFERENCES `supplier`(`id`),
	FOREIGN KEY (`purchase_id`) REFERENCES `purchase_order`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='供应商应付账款表';

-- ----------------------------
-- 库存盘点主表
-- ----------------------------
CREATE TABLE `inventory_check` (
	`id` VARCHAR(20) NOT NULL COMMENT '盘点单号',
	`store_id` INT UNSIGNED NOT NULL COMMENT '门店ID',
	`operator_id` INT UNSIGNED NOT NULL COMMENT '操作人ID',
	`check_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '盘点时间',
	`total_diff` INT NOT NULL COMMENT '总差异数量',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`store_id`) REFERENCES `store`(`id`),
	FOREIGN KEY (`operator_id`) REFERENCES `employee`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存盘点主表';

-- ----------------------------
-- 盘点库存明细表
-- ----------------------------
CREATE TABLE `check_detail` (
	`check_id` VARCHAR(20) NOT NULL COMMENT '盘点单号',
	`sku_id` INT UNSIGNED NOT NULL COMMENT 'SKU ID',
	`system_qty` INT NOT NULL COMMENT '系统库存',
	`actual_qty` INT NOT NULL COMMENT '实际库存',
	PRIMARY KEY (`check_id`, `sku_id`),
	FOREIGN KEY (`check_id`) REFERENCES `inventory_check`(`id`),
	FOREIGN KEY (`sku_id`) REFERENCES `product_sku`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存盘点明细表';



-- ---------------------------- 插入数据 ----------------------------

use appliance_store;

-- 产品分类表（生成15条多级分类）
INSERT INTO product_category (parent_id, name, level) VALUES
	(0, '大家电', 1),
	(0, '厨卫电器', 1),
	(0, '生活电器', 1),
	(1, '冰箱', 2),
	(1, '洗衣机', 2),
	(1, '空调', 2),
	(2, '燃气灶', 2),
	(2, '抽油烟机', 2),
	(3, '电饭煲', 2),
	(3, '微波炉', 2),
	(4, '对开门冰箱', 3),
	(4, '三门冰箱', 3),
	(5, '滚筒洗衣机', 3),
	(5, '波轮洗衣机', 3),
	(6, '壁挂式空调', 3);

-- 产品表（10条）
INSERT INTO product (name, category_id, brand, model, spec, price) VALUES
	('海尔对开门冰箱', 11, '海尔', 'BCD-500W', '500L', 5999.00),
	('美的三门冰箱', 12, '美的', 'BCD-300M', '300L', 3999.00),
	('小天鹅滚筒洗衣机', 13, '小天鹅', 'TG100', '10kg', 2999.00),
	('西门子燃气灶', 7, '西门子', 'JZT-789', '天然气', 1999.00),
	('老板抽油烟机', 8, '老板', 'CXW-200', '侧吸式', 2499.00),
	('九阳电饭煲', 9, '九阳', 'F-40FZ', '4L', 299.00),
	('格兰仕微波炉', 10, '格兰仕', 'G80F23', '23L', 499.00),
	('格力壁挂空调', 15, '格力', 'KFR-35GW', '1.5匹', 3499.00),
	('海尔立柜空调', 6, '海尔', 'KFR-72LW', '3匹', 8999.00),
	('美的波轮洗衣机', 14, '美的', 'MB80-3000', '8kg', 1999.00);

-- 产品多媒体（10条）
INSERT INTO product_media (product_id, type, url) VALUES
	(1, 1, 'http://img.com/haier_fridge.jpg'),
	(2, 1, 'http://img.com/midea_fridge.jpg'),
	(3, 1, 'http://img.com/littleswan_washer.jpg'),
	(4, 1, 'http://img.com/siemens_stove.jpg'),
	(5, 1, 'http://img.com/robam_hood.jpg'),
	(6, 1, 'http://img.com/joyoung_ricecooker.jpg'),
	(7, 1, 'http://img.com/galanz_microwave.jpg'),
	(8, 1, 'http://img.com/gree_ac.jpg'),
	(9, 1, 'http://img.com/haier_ac.jpg'),
	(10, 1, 'http://img.com/midea_washer.jpg');

-- SKU表（10条）
INSERT INTO product_sku (product_id, sku_code, attributes, price, stock) VALUES
	(1, 'SKU001', '{"color": "白", "capacity": "500L"}', 5999.00, 50),
	(2, 'SKU002', '{"color": "银灰", "capacity": "300L"}', 3999.00, 30),
	(3, 'SKU003', '{"type": "滚筒", "capacity": "10kg"}', 2999.00, 40),
	(4, 'SKU004', '{"fuel": "天然气", "burners": 3}', 1999.00, 25),
	(5, 'SKU005', '{"type": "侧吸", "noise": "55dB"}', 2499.00, 20),
	(6, 'SKU006', '{"capacity": "4L", "material": "铝合金"}', 299.00, 100),
	(7, 'SKU007', '{"capacity": "23L", "control": "机械"}', 499.00, 80),
	(8, 'SKU008', '{"type": "变频", "power": "1.5匹"}', 3499.00, 15),
	(9, 'SKU009', '{"type": "立柜式", "power": "3匹"}', 8999.00, 10),
	(10, 'SKU010', '{"type": "波轮", "capacity": "8kg"}', 1999.00, 35);

-- 门店表（10条）
INSERT INTO store (name, address, contact_phone) VALUES
	('北京朝阳店', '朝阳区建国路100号', '010-12345678'),
	('上海浦东店', '浦东新区陆家嘴环路200号', '021-87654321'),
	('广州天河店', '天河区天河路300号', '020-11223344'),
	('深圳福田店', '福田区深南大道400号', '0755-55667788'),
	('成都锦江店', '锦江区春熙路500号', '028-33445566'),
	('杭州西湖店', '西湖区文三路600号', '0571-77889900'),
	('南京玄武店', '玄武区中山路700号', '025-44556677'),
	('武汉江汉店', '江汉区解放大道800号', '027-88990011'),
	('重庆渝中店', '渝中区解放碑900号', '023-66778899'),
	('西安雁塔店', '雁塔区科技路1000号', '029-55443322');

-- 员工表（10条，先插入后更新门店manager_id）
INSERT INTO employee (name, phone, role, store_id, password_hash) VALUES
	('张三店长', '13800001111', 2, 1, 'e10adc3949ba59abbe56e057f20f883e'),
	('李四销售', '13800002222', 3, 1, 'e10adc3949ba59abbe56e057f20f883e'),
	('王五财务', '13800003333', 4, 1, 'e10adc3949ba59abbe56e057f20f883e'),
	('陈六店长', '13800004444', 2, 2, 'e10adc3949ba59abbe56e057f20f883e'),
	('赵七销售', '13800005555', 3, 2, 'e10adc3949ba59abbe56e057f20f883e'),
	('周八店长', '13800006666', 2, 3, 'e10adc3949ba59abbe56e057f20f883e'),
	('吴九维修', '13800007777', 5, 3, 'e10adc3949ba59abbe56e057f20f883e'),
	('郑十财务', '13800008888', 4, 4, 'e10adc3949ba59abbe56e057f20f883e'),
	('孙十一店长', '13800009999', 2, 5, 'e10adc3949ba59abbe56e057f20f883e'),
	('刘十二销售', '13800001010', 3, 6, 'e10adc3949ba59abbe56e057f20f883e');

-- 更新门店店长
UPDATE store SET manager_id = 1 WHERE id = 1;
UPDATE store SET manager_id = 4 WHERE id = 2;
UPDATE store SET manager_id = 6 WHERE id = 3;
UPDATE store SET manager_id = 8 WHERE id = 4;
UPDATE store SET manager_id = 9 WHERE id = 5;


-- 员工绩效（10条）
INSERT INTO employee_performance (employee_id, month, sales_amount, commission) VALUES
	(2, '2024-02', 85000.00, 850.00),
	(3, '2024-02', 0.00, 0.00),
	(4, '2024-02', 120000.00, 1200.00),
	(5, '2024-02', 45000.00, 450.00),
	(6, '2024-02', 98000.00, 980.00),
	(7, '2024-02', 0.00, 0.00),
	(8, '2024-02', 0.00, 0.00),
	(9, '2024-02', 75000.00, 750.00),
	(10, '2024-02', 36000.00, 360.00),
	(1, '2024-02', 150000.00, 1500.00);

-- 客户表（10条）
INSERT INTO customer (name, phone, email) VALUES
	('王小明', '13912345678', 'wangxm@example.com'),
	('李晓红', '13923456789', 'lixh@example.com'),
	('张伟', '13934567890', 'zhangw@example.com'),
	('刘芳', '13945678901', 'liuf@example.com'),
	('陈刚', '13956789012', 'cheng@example.com'),
	('杨丽', '13967890123', 'yangli@example.com'),
	('黄强', '13978901234', 'huangq@example.com'),
	('徐敏', '13989012345', 'xum@example.com'),
	('周涛', '13990123456', 'zhout@example.com'),
	('吴霞', '13901234567', 'wux@example.com');

-- 促销活动（10条）
INSERT INTO promotion (name, type, rule, start_time, end_time) VALUES
	('春季满减', 1, '{"full": 1000, "minus": 200}', '2024-03-01 00:00:00', '2024-03-31 23:59:59'),
	('空调折扣', 2, '{"rate": 0.85}', '2024-03-05 00:00:00', '2024-03-15 23:59:59'),
	('赠品活动', 3, '{"gift": "电水壶"}', '2024-03-10 00:00:00', '2024-03-20 23:59:59'),
	('组合优惠', 4, '{"set": ["SKU001", "SKU008"], "price": 8999}', '2024-03-08 00:00:00', '2024-03-18 23:59:59'),
	('店庆特惠', 1, '{"full": 5000, "minus": 800}', '2024-03-15 00:00:00', '2024-03-25 23:59:59'),
	('冰箱直降', 2, '{"rate": 0.78}', '2024-03-12 00:00:00', '2024-03-22 23:59:59'),
	('洗衣机赠品', 3, '{"gift": "洗衣液"}', '2024-03-20 00:00:00', '2024-03-30 23:59:59'),
	('厨电组合', 4, '{"set": ["SKU004", "SKU005"], "price": 3999}', '2024-03-25 00:00:00', '2024-04-05 23:59:59'),
	('新人折扣', 2, '{"rate": 0.9}', '2024-03-01 00:00:00', '2024-03-10 23:59:59'),
	('全品类满减', 1, '{"full": 3000, "minus": 500}', '2024-03-18 00:00:00', '2024-03-28 23:59:59');

-- 操作日志（10条）
INSERT INTO audit_log (operator_id, action_type, target_table, target_id) VALUES
	(1, 'INSERT', 'product', 1),
	(2, 'UPDATE', 'sales_order', '20240301080000000001'),
	(3, 'DELETE', 'customer', 5),
	(4, 'INSERT', 'purchase_order', 'PO20240301002'),
	(5, 'UPDATE', 'inventory', 3),
	(6, 'INSERT', 'employee', 7),
	(7, 'DELETE', 'supplier', 9),
	(8, 'UPDATE', 'promotion', 2),
	(9, 'INSERT', 'transfer_order', 'TRANSFER005'),
	(10, 'DELETE', 'coupon', 'COUPON006');

-- 供应商表（10条）
INSERT INTO supplier (name, contact_person, phone) VALUES
	('海尔集团', '张经理', '400-100-1111'),
	('美的供应链', '李主管', '400-200-2222'),
	('格力电器', '王总', '400-300-3333'),
	('西门子中国', '赵经理', '400-400-4444'),
	('老板电器', '陈总监', '400-500-5555'),
	('九阳股份', '刘部长', '400-600-6666'),
	('格兰仕', '周经理', '400-700-7777'),
	('小天鹅', '吴主管', '400-800-8888'),
	('松下电器', '林总', '400-900-9999'),
	('飞利浦家电', '徐经理', '400-000-0000');


-- 库存表（10条）
INSERT INTO inventory (store_id, sku_id, quantity, safety_stock) VALUES
	(1, 1, 50, 10),
	(1, 2, 30, 5),
	(2, 3, 40, 8),
	(2, 4, 20, 5),
	(3, 5, 15, 3),
	(3, 6, 100, 20),
	(4, 7, 80, 15),
	(4, 8, 10, 2),
	(5, 9, 5, 1),
	(5, 10, 35, 5);

-- 销售订单（10条）
INSERT INTO sales_order (id, store_id, customer_id, total_amount, status, payment_method) VALUES
	('20240301080000000001', 1, 1, 5999.00, 1, 2),
	('20240301080000000002', 1, 2, 7998.00, 1, 3),
	('20240301080000000003', 2, 3, 2999.00, 0, NULL),
	('20240301080000000004', 2, 4, 4998.00, 2, 1),
	('20240301080000000005', 3, 5, 2499.00, 3, 4),
	('20240301080000000006', 3, 6, 299.00, 1, 2),
	('20240301080000000007', 4, 7, 499.00, 4, 3),
	('20240301080000000008', 4, 8, 3499.00, 1, 2),
	('20240301080000000009', 5, 9, 8999.00, 1, 3),
	('20240301080000000010', 5, 10, 1999.00, 0, NULL);

-- 销售订单明细（10条）
INSERT INTO sales_order_item (order_id, sku_id, quantity, unit_price) VALUES
	('20240301080000000001', 1, 1, 5999.00),
	('20240301080000000002', 2, 2, 3999.00),
	('20240301080000000003', 3, 1, 2999.00),
	('20240301080000000004', 4, 2, 1999.00),
	('20240301080000000005', 5, 1, 2499.00),
	('20240301080000000006', 6, 1, 299.00),
	('20240301080000000007', 7, 1, 499.00),
	('20240301080000000008', 8, 1, 3499.00),
	('20240301080000000009', 9, 1, 8999.00),
	('20240301080000000010', 10, 1, 1999.00);


-- 采购订单主表（10条）
INSERT INTO purchase_order (id, supplier_id, total_amount, status, expected_date) VALUES
	('PO202403010001', 1, 17997.00, 2, '2024-03-05'),  -- 海尔冰箱采购
	('PO202403010002', 2, 11997.00, 1, '2024-03-07'),  -- 美的洗衣机
	('PO202403010003', 3, 34990.00, 0, '2024-03-10'),  -- 格力空调
	('PO202403010004', 4, 5996.00, 2, '2024-03-12'),   -- 西门子燃气灶
	('PO202403010005', 5, 7497.00, 1, '2024-03-15'),   -- 老板油烟机
	('PO202403010006', 6, 2990.00, 2, '2024-03-18'),   -- 九阳电饭煲
	('PO202403010007', 7, 2495.00, 0, '2024-03-20'),   -- 格兰仕微波炉
	('PO202403010008', 8, 9990.00, 1, '2024-03-22'),   -- 小天鹅洗衣机
	('PO202403010009', 9, 17998.00, 2, '2024-03-25'),  -- 松下电器
	('PO202403010010', 10, 3499.00, 0, '2024-03-28');  -- 飞利浦家电

-- 采购订单明细（10条）
INSERT INTO purchase_order_item (order_id, sku_id, quantity, unit_price) VALUES
	('PO202403010001', 1, 3, 5999.00),
	('PO202403010002', 10, 3, 3999.00),
	('PO202403010003', 8, 10, 3499.00),
	('PO202403010004', 4, 3, 1999.00),
	('PO202403010005', 5, 3, 2499.00),
	('PO202403010006', 6, 10, 299.00),
	('PO202403010007', 7, 5, 499.00),
	('PO202403010008', 3, 5, 1999.00),
	('PO202403010009', 9, 2, 8999.00),
	('PO202403010010', 2, 1, 3499.00);

-- 库存操作记录（10条）
INSERT INTO inventory_operation (inventory_id, store_id, sku_id, type, quantity, operator) VALUES
	(1, 1, 1, 1, 50, 2),   -- 采购入库
	(2, 1, 2, 1, 30, 2),
	(3, 2, 3, 1, 40, 5),
	(4, 2, 4, 1, 25, 5),
	(5, 3, 5, 3, -2, 3),   -- 销售出库
	(6, 3, 6, 3, -1, 3),
	(7, 4, 7, 4, -5, 8),   -- 调拨出库
	(8, 4, 8, 5, 5, 8),    -- 调拨入库
	(9, 5, 9, 2, 3, 9),    -- 退货入库
	(10, 5, 10, 3, -2, 9);

-- 财务流水（10条）
INSERT INTO financial_transaction (type, amount, related_id, transaction_time, operator) VALUES
	(1, 5999.00, '20240301080000000001', '2024-03-01 08:05:00', 2),  -- 销售收款
	(2, 17997.00, 'PO202403010001', '2024-03-01 09:00:00', 3),       -- 采购付款
	(1, 7998.00, '20240301080000000002', '2024-03-01 10:30:00', 2),
	(3, 500.00, NULL, '2024-03-01 11:00:00', 3),                     -- 其他收入
	(4, 2000.00, NULL, '2024-03-01 14:00:00', 3),                    -- 其他支出
	(1, 2999.00, '20240301080000000003', '2024-03-01 15:20:00', 5),
	(2, 11997.00, 'PO202403010002', '2024-03-02 10:00:00', 3),
	(1, 3499.00, '20240301080000000008', '2024-03-02 11:30:00', 8),
	(2, 5996.00, 'PO202403010004', '2024-03-03 09:45:00', 3),
	(1, 8999.00, '20240301080000000009', '2024-03-03 14:15:00', 9);

-- 售后服务（10条）
INSERT INTO after_sales (order_id, type, description, status) VALUES
	('20240301080000000007', '退货', '微波炉门损坏', '已完成'),
	('20240301080000000004', '换货', '燃气灶点火故障', '处理中'),
	('20240301080000000002', '维修', '冰箱制冷异常', '待处理'),
	('20240301080000000005', '退款', '油烟机噪音过大', '已退款'),
	('20240301080000000001', '换货', '冰箱外观划痕', '已完成'),
	('20240301080000000010', '维修', '洗衣机脱水故障', '处理中'),
	('20240301080000000006', '退货', '电饭煲不加热', '待处理'),
	('20240301080000000003', '换货', '洗衣机漏水', '已完成'),
	('20240301080000000008', '维修', '空调制冷不足', '处理中'),
	('20240301080000000009', '退款', '空调安装延误', '已退款');

-- 维修工单（10条）
INSERT INTO service_order (id, customer_id, product_id, fault_desc, status, technician_id) VALUES
	('SV202403010001', 7, 7, '微波炉门损坏无法关闭', 2, 7),
	('SV202403010002', 4, 4, '燃气灶点火器故障', 1, 7),
	('SV202403010003', 2, 1, '冰箱制冷效果差', 0, NULL),
	('SV202403010004', 5, 5, '油烟机异响', 3, 7),
	('SV202403010005', 1, 1, '冰箱门密封条脱落', 2, 7),
	('SV202403010006', 10, 10, '洗衣机脱水时晃动', 1, 7),
	('SV202403010007', 6, 6, '电饭煲按键失灵', 0, NULL),
	('SV202403010008', 3, 3, '洗衣机排水管堵塞', 2, 7),
	('SV202403010009', 8, 8, '空调出风口漏水', 1, 7),
	('SV202403010010', 9, 9, '空调遥控器失灵', 3, 7);

-- 调拨单主表（10条）
INSERT INTO transfer_order (id, from_store, to_store, total_quantity, status) VALUES
	('TF20240301001', 1, 2, 10, 1),
	('TF20240301002', 2, 3, 5, 0),
	('TF20240301003', 3, 4, 8, 1),
	('TF20240301004', 4, 5, 15, 0),
	('TF20240301005', 5, 1, 3, 1),
	('TF20240301006', 1, 3, 20, 1),
	('TF20240301007', 2, 4, 12, 0),
	('TF20240301008', 3, 5, 6, 1),
	('TF20240301009', 4, 1, 9, 0),
	('TF20240301010', 5, 2, 4, 1);

-- 调拨明细（10条）
INSERT INTO transfer_detail (transfer_id, sku_id, quantity) VALUES
	('TF20240301001', 1, 5),
	('TF20240301001', 2, 5),
	('TF20240301002', 3, 5),
	('TF20240301003', 4, 8),
	('TF20240301004', 5, 15),
	('TF20240301005', 6, 3),
	('TF20240301006', 7, 20),
	('TF20240301007', 8, 12),
	('TF20240301008', 9, 6),
	('TF20240301009', 10, 9);

-- 消息通知（10条）
INSERT INTO notification (type, recipient, content, send_status) VALUES
	(1, '13800001111', 'SKU001库存低于安全库存', 1),          -- 店长手机
	(2, '20240301080000000001', '您的订单已支付完成', 1),     -- 订单号
	(3, 'all', '春季大促3月15日开启', 0),                    -- 全员促销
	(1, '13800004444', 'SKU005库存仅剩3件', 1),
	(2, '13912345678', '您的换货申请已受理', 1),
	(3, 'sales@store.com', '新员工销售培训通知', 2),         -- 发送失败
	(1, '13800009999', 'SKU009库存补货到仓', 1),
	(2, '20240301080000000005', '您的维修工单已完成', 1),
	(3, '13800002222', '本月销售目标提醒', 0),
	(1, '13800008888', 'SKU007库存盘点差异+2', 1);

-- 优惠券（10条）
INSERT INTO coupon (id, type, value, amount_condition, customer_id, status, expire_time) VALUES
	('COUPON20240301', 1, 200.00, 1000.00, 1, 1, '2024-03-31'),
	('COUPON20240302', 2, 0.90, 500.00, 2, 1, '2024-04-30'),
	('COUPON20240303', 3, 50.00, 300.00, 3, 2, '2024-03-15'),
	('COUPON20240304', 1, 100.00, 800.00, 4, 0, '2024-06-30'), -- 未发放
	('COUPON20240305', 2, 0.85, 2000.00, 5, 1, '2024-05-31'),
	('COUPON20240306', 3, 100.00, 300.00, 6, 3, '2024-03-10'),      -- 已过期
	('COUPON20240307', 1, 300.00, 1500.00, 7, 1, '2024-04-15'),
	('COUPON20240308', 2, 0.75, 1000.00, 8, 1, '2024-03-20'),
	('COUPON20240309', 3, 80.00, 300.00, 9, 2, '2024-03-25'),
	('COUPON20240310', 1, 150.00, 500.00, 10, 0, '2024-07-31');

-- 应付账款（10条）
INSERT INTO account_payable (supplier_id, purchase_id, due_amount, due_date) VALUES
	(1, 'PO202403010001', 17997.00, '2024-04-01'),
	(2, 'PO202403010002', 11997.00, '2024-04-05'),
	(4, 'PO202403010004', 5996.00, '2024-03-25'),
	(5, 'PO202403010005', 7497.00, '2024-04-10'),
	(6, 'PO202403010006', 2990.00, '2024-03-28'),
	(7, 'PO202403010007', 2495.00, '2024-04-15'),
	(8, 'PO202403010008', 9990.00, '2024-04-20'),
	(9, 'PO202403010009', 17998.00, '2024-04-25'),
	(10, 'PO202403010010', 3499.00, '2024-04-30'),
	(3, 'PO202403010003', 34990.00, '2024-05-01');

-- 库存盘点主表（10条）
INSERT INTO inventory_check (id, store_id, operator_id, total_diff) VALUES
	('CK2024030101', 1, 1, 2),
	('CK2024030102', 2, 4, -1),
	('CK2024030103', 3, 6, 0),
	('CK2024030104', 4, 8, 3),
	('CK2024030105', 5, 9, -2),
	('CK2024030106', 1, 1, 1),
	('CK2024030107', 2, 4, 0),
	('CK2024030108', 3, 6, -1),
	('CK2024030109', 4, 8, 2),
	('CK2024030110', 5, 9, 0);

-- 盘点明细（10条）
INSERT INTO check_detail (check_id, sku_id, system_qty, actual_qty) VALUES
	('CK2024030101', 1, 50, 52),   -- 差异+2
	('CK2024030102', 3, 40, 39),   -- 差异-1
	('CK2024030104', 7, 80, 83),   -- 差异+3
	('CK2024030105', 9, 5, 3),     -- 差异-2
	('CK2024030106', 2, 30, 31),   -- 差异+1
	('CK2024030108', 6, 100, 99),  -- 差异-1
	('CK2024030109', 8, 10, 12),   -- 差异+2
	('CK2024030103', 5, 15, 15),   -- 无差异
	('CK2024030107', 4, 25, 25),   -- 无差异
	('CK2024030110', 10, 35, 35);  -- 无差异
