-- Create Online Platform table
DROP Table IF EXISTS t_platform;
CREATE TABLE t_platform (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    name VARCHAR(100) UNIQUE COMMENT '平台名称',
    logo_url VARCHAR(100) DEFAULT '' COMMENT '平台logo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '平台创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '平台最后更新时间',
    deleted_at TIMESTAMP COMMENT '平台软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

-- Insert example platforms
INSERT INTO t_platform (name) VALUES 
('Wish'),
('eBay'),
('Lazada'),
('Amazon');

-- Create Online Platform Shop table
DROP Table IF EXISTS t_shop;
CREATE TABLE t_shop (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    platform_name VARCHAR(255) COMMENT '平台名称',
    shop_id VARCHAR(255) COMMENT '平台店铺唯一标识符',
    shop_name VARCHAR(255) COMMENT '店铺名称',
    client_id VARCHAR(255) COMMENT 'API Client ID',
    client_cert VARCHAR(255) COMMENT 'API Client Cert',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '账号创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '账号最后更新时间',
    deleted_at TIMESTAMP COMMENT '账号软删除时间（如果适用）'
) CHARACTER SET utf8mb4;
ALTER TABLE t_shop ADD CONSTRAINT t_shop UNIQUE (platform_name, shop_id); -- 在同一个平台下店铺唯一
INSERT INTO t_shop (platform_name, shop_id, shop_name) VALUES 
('eBay', 'ShopID001', 'Shop 1');

-- Create Category table
DROP Table IF EXISTS t_category;
CREATE TABLE t_category (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    domain ENUM('Product', 'Goods') COMMENT '类型场景',
    name VARCHAR(100) COMMENT '类别名称',
    parent_id BIGINT UNSIGNED COMMENT '父类别的标识符',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '类别创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '类别最后更新时间',
    deleted_at TIMESTAMP COMMENT '类别软删除时间（如果适用）'
) CHARACTER SET utf8mb4;
ADD CONSTRAINT t_category UNIQUE (domain, name); -- 在同一个场景下类别唯一
INSERT INTO t_category (domain, name, parent_id) VALUES 
('Product', '所有分类', 0),
('Goods', '所有分类', 0);

-- Create Product table
DROP Table IF EXISTS t_product;
CREATE TABLE t_product (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    sku VARCHAR(50) UNIQUE COMMENT '产品的唯一标识符',
    summary TEXT COMMENT '产品摘要',
    product_name VARCHAR(255) COMMENT '产品名称',
    description TEXT COMMENT '产品描述',
    category_id BIGINT UNSIGNED COMMENT '产品所属的类别标识符',
    unit_price DECIMAL(10, 2) COMMENT '产品的单价',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '产品创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '产品最后更新时间',
    deleted_at TIMESTAMP COMMENT '产品软删除时间（如果适用）'
) CHARACTER SET utf8mb4;
-- Insert sample data into t_product table
INSERT INTO t_product (sku, summary, product_name, description, category_id, unit_price)
VALUES
    ('SKU001', 'Product summary 1', 'Product 1', 'Description for Product 1', 1, 19.99),
    ('SKU002', 'Product summary 2', 'Product 2', 'Description for Product 2', 2, 29.99),
    ('SKU003', 'Product summary 3', 'Product 3', 'Description for Product 3', 1, 39.99),
    ('SKU004', 'Product summary 4', 'Product 4', 'Description for Product 4', 3, 49.99),
    ('SKU005', 'Product summary 5', 'Product 5', 'Description for Product 5', 2, 59.99);

-- Create image table
DROP Table IF EXISTS t_image;
CREATE TABLE t_image (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    domain ENUM('Product', 'Goods') COMMENT '类型场景',
    domain_id BIGINT UNSIGNED COMMENT '关联产品的标识符',
    image_url VARCHAR(255) COMMENT '图片的URL',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '图片创建时间'
) CHARACTER SET utf8mb4;

-- Create ProductAttribute table
DROP Table IF EXISTS t_product_attribute;
CREATE TABLE t_product_attribute (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    product_sku VARCHAR(50)COMMENT '关联产品的标识符',
    attribute_key VARCHAR(255) COMMENT '属性名称',
    attribute_value VARCHAR(255) COMMENT '属性值',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '属性创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '属性最后更新时间',
    deleted_at TIMESTAMP COMMENT '属性软删除时间（如果适用）'
) CHARACTER SET utf8mb4;
INSERT INTO t_product_attribute (product_sku, attribute_key, attribute_value)
VALUES
    ('SKU001', 'Color', 'Red'),
    ('SKU001', 'Size', 'Large'),
    ('SKU002', 'Color', 'Blue'),
    ('SKU002', 'Size', 'Medium'),
    ('SKU003', 'Color', 'Green'),
    ('SKU003', 'Size', 'Small'),
    ('SKU004', 'Color', 'Black'),
    ('SKU004', 'Size', 'XL'),
    ('SKU005', 'Color', 'White'),
    ('SKU005', 'Size', 'Small');

-- Create Order table
DROP Table IF EXISTS t_order;
CREATE TABLE t_order (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    platform_name VARCHAR(50) COMMENT '电商平台',
    shop_id VARCHAR(255) COMMENT '店铺唯一标识符',
    shop_name VARCHAR(50)  COMMENT '店铺名称',
    platform_order_id VARCHAR(50) COMMENT '店铺订单号',

    total_amount DECIMAL(12, 2) COMMENT '订单总金额',
    discount DECIMAL(5, 2) COMMENT '订单折扣金额',
    tax DECIMAL(8, 2) COMMENT '订单税额',
    
    status_id BIGINT UNSIGNED COMMENT '订单当前状态ID',
    status_name VARCHAR(50) COMMENT '订单当前状态',

    customer_id VARCHAR(50) COMMENT '客户的标识符',
    customer_name VARCHAR(50) COMMENT '客户名称',
    customer_tel VARCHAR(50) COMMENT '客户电话',
    customer_email VARCHAR(50) COMMENT '客户邮箱',
    customer_remarks VARCHAR(2048) COMMENT '客户备注',

    payment_method VARCHAR(50) COMMENT '支付方式',
    payment_at TIMESTAMP COMMENT '付款时间',
    
    -- 配送信息
    shipping_order_id VARCHAR(50) COMMENT '运单号',
    shipping_cost DECIMAL(8, 2) COMMENT '订单运费',
    shipping_street_address VARCHAR(255) COMMENT '街道地址',
    shipping_city VARCHAR(100) COMMENT '城市',
    shipping_state VARCHAR(100) COMMENT '省/州',
    shipping_postal_code VARCHAR(20) COMMENT '邮政编码',
    shipping_country VARCHAR(100) COMMENT '国家',

    ship_at TIMESTAMP COMMENT '发货时间',
    refund_at TIMESTAMP COMMENT '退款时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单最后更新时间',
    deleted_at TIMESTAMP COMMENT '订单软删除时间（如果适用）'
) CHARACTER SET utf8mb4;
-- Insert sample data into t_order table
INSERT INTO t_order (
    platform_name, shop_id, shop_name, platform_order_id,
    total_amount, discount, tax,
    status_id,
    customer_id, customer_name, customer_tel, customer_email, customer_remarks,
    payment_method, payment_at,
    shipping_order_id, shipping_cost,
    shipping_street_address, shipping_city, shipping_state, shipping_postal_code, shipping_country,
    ship_at, refund_at
) VALUES
    ('eBay', 'ShopID001', 'Shop 1', 'OrderID001', 100.00, 10.00, 5.00, 1, 'CustomerID001', 'John Doe', '123456789', 'john@example.com', 'Special instructions for the order', 'Credit Card', '2023-01-01 12:00:00', 'Shipping123', 8.00, '123 Main St', 'City1', 'State1', '12345', 'Country1', '2023-01-02 10:00:00', NULL),
    ('eBay', 'ShopID002', 'Shop 2', 'OrderID002', 150.00, 15.00, 7.50, 2, 'CustomerID002', 'Jane Smith', '987654321', 'jane@example.com', 'Additional notes for the order', 'PayPal', '2023-01-02 15:30:00', 'Shipping456', 12.00, '456 Oak St', 'City2', 'State2', '54321', 'Country2', '2023-01-03 09:45:00', NULL),
    ('eBay', 'ShopID003', 'Shop 3', 'OrderID003', 80.00, 8.00, 4.00, 3, 'CustomerID003', 'Bob Johnson', '555111222', 'bob@example.com', 'No special instructions', 'Bank Transfer', '2023-01-03 18:20:00', 'Shipping789', 6.50, '789 Pine St', 'City3', 'State3', '67890', 'Country3', '2023-01-04 11:20:00', NULL);


-- Create Order Item table
DROP Table IF EXISTS t_order_item;
CREATE TABLE t_order_item (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    order_id VARCHAR(50) COMMENT '关联订单的标识符',
    product_sku VARCHAR(50) COMMENT '关联产品的标识符',
    product_url VARCHAR(512) COMMENT '产品图片', -- 图片一般是订单同步时，平台接口直接返回的图片信息。
    attributes TEXT COMMENT '产品属性', -- 平台返回的产品的颜色、尺寸等属性信息
    quantity INT COMMENT '订单中产品的数量',
    unit_price DECIMAL(10, 2) COMMENT '产品的单价',
    total_amount DECIMAL(12, 2) COMMENT '订单项总金额',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '订单项创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单项最后更新时间',
    deleted_at TIMESTAMP COMMENT '订单项软删除时间（如果适用）'
) CHARACTER SET utf8mb4;
-- Insert sample data into t_order_item table
INSERT INTO t_order_item (
    order_id, platform_order_id, product_sku, product_url, attributes, quantity, unit_price, total_amount
) VALUES
    (1, 'OrderID001', 'SKU001', 'https://example.com/product1.jpg', '{"color": "Red", "size": "Large"}', 2, 25.00, 50.00),
    (1, 'OrderID001', 'SKU002', 'https://example.com/product2.jpg', '{"color": "Blue", "size": "Medium"}', 1, 30.00, 30.00),
    (2, 'OrderID002', 'SKU003', 'https://example.com/product3.jpg', '{"color": "Green", "size": "Small"}', 3, 15.00, 45.00),
    (2, 'OrderID003', 'SKU004', 'https://example.com/product4.jpg', '{"color": "Black", "size": "XL"}', 1, 50.00, 50.00),
    (3, 'OrderID003', 'SKU005', 'https://example.com/product5.jpg', '{"color": "White", "size": "Small"}', 2, 40.00, 80.00);

-----------------------仓库相关----------------------
-- Create t_wms_warehouse table
DROP Table IF EXISTS t_wms_warehouse;
CREATE TABLE t_wms_warehouse (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    name VARCHAR(255) COMMENT '仓库名称',
    owner VARCHAR(255) COMMENT '负责人',
    telephone VARCHAR(255) COMMENT '联系电话',
    street_address VARCHAR(255) COMMENT '仓库街道地址',
    city VARCHAR(100) COMMENT '城市',
    state_province VARCHAR(100) COMMENT '州/省',
    postal_code VARCHAR(20) COMMENT '邮政编码',
    country VARCHAR(100) COMMENT '国家',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '仓库创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '仓库最后更新时间',
    deleted_at TIMESTAMP COMMENT '仓库软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

-- Create t_wms_zone Table
DROP Table IF EXISTS t_wms_zone;
CREATE TABLE t_wms_zone (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    name VARCHAR(255) UNIQUE COMMENT '区域名称',
    status int COMMENT '区域状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    deleted_at TIMESTAMP COMMENT '软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

-- Insert values into WMS Zone table
INSERT INTO t_wms_zone (name, status) VALUES 
('拣货区', 1）,
('次品区', 2);

-- Create t_wms_rack table
DROP Table IF EXISTS t_wms_rack;
CREATE TABLE t_wms_rack (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    warehouse_id BIGINT UNSIGNED DEFAULT 0 COMMENT '仓库ID',
    zone_id BIGINT UNSIGNED DEFAULT 1 COMMENT '分区ID',
    rack_code VARCHAR(50) NOT NULL COMMENT '货架编号',
    status ENUM('可分配','停止收货','停用货架') DEFAULT '可分配' COMMENT '货架状态',
    priority INT DEFAULT 1000 COMMENT '拣货权重', -- 拣货权重数值越大权重越小
    remarks TEXT COMMENT '备注说明',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '货架创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '货架最后更新时间',
    deleted_at TIMESTAMP COMMENT '货架软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

-- Create WMS Inventory table
DROP Table IF EXISTS t_wms_inventory;
CREATE TABLE t_wms_inventory (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    goods_sku VARCHAR(50) UNIQUE COMMENT '商品SKU',
    warehouse_id BIGINT UNSIGNED COMMENT '仓库ID',
    zone_id BIGINT UNSIGNED COMMENT '库区ID',
    rack_id  BIGINT UNSIGNED COMMENT '货架ID',
    quantity INT COMMENT '总库存数量',
    -- 安全库存：它是安全交期库存数量，用于仓库剩余库存达到一定数值时，采购提前补货参考
    safety_quantity INT COMMENT '安全库存',
    -- 自发货订单还未发货的商品数，未发=待审核 + 待处理 + 运单号申请 + 待打单
    unshipping_quantity INT COMMENT '预售库存',
    -- 在途数量 = 采购/加工在途 + 库内在途 + 调拨在途
    on_order_quantity INT COMMENT '在途库存',
    -- 预售：它是订单商品SKU预售数量，即订单待打单，已占用库存数
    pre_sale_quantity INT COMMENT '预售库存',
    stock_status ENUM('In Stock', 'Out of Stock', 'Backordered') COMMENT '库存状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '库存创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '库存最后更新时间',
    deleted_at TIMESTAMP COMMENT '库存软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

DROP Table IF EXISTS t_wms_inbound_log;
CREATE TABLE t_wms_inbound_log ( 
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    goods_sku VARCHAR(50) COMMENT '商品SKU',
    warehouse_id BIGINT UNSIGNED COMMENT '仓库ID',
    zone_id BIGINT UNSIGNED COMMENT '库区ID',
    rack_id  BIGINT UNSIGNED COMMENT '货架ID',
    quantity INT COMMENT '入库数量',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '库存创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '库存最后更新时间',
    deleted_at TIMESTAMP COMMENT '库存软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

DROP Table IF EXISTS t_wms_outbound_log;
CREATE TABLE t_wms_outbound_log (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    goods_sku VARCHAR(50) COMMENT '商品SKU',
    warehouse_id BIGINT UNSIGNED COMMENT '仓库ID',
    zone_id BIGINT UNSIGNED COMMENT '库区ID',
    rack_id  BIGINT UNSIGNED COMMENT '货架ID',
    quantity INT COMMENT '出库数量',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '库存创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '库存最后更新时间',
    deleted_at TIMESTAMP COMMENT '库存软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

DROP Table IF EXISTS t_wms_pick_order;
CREATE TABLE t_wms_pick_order (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    order_id BIGINT UNSIGNED DEFAULT 0 COMMENT '订单ID',
    goods_sku VARCHAR(50) COMMENT '商品SKU',
    goods_name VARCHAR(50) COMMENT '商品名称',
    warehouse_id BIGINT UNSIGNED COMMENT '仓库ID',
    zone_id BIGINT UNSIGNED COMMENT '库区ID',
    rack_id  BIGINT UNSIGNED COMMENT '货架ID',
    quantity INT COMMENT '拣货数量',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    deleted_at TIMESTAMP COMMENT '软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

-- Create t_wms_goods table
DROP TABLE IF EXISTS t_wms_goods;
CREATE TABLE t_wms_goods (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    code VARCHAR(20) UNIQUE COMMENT '商品编码', -- 系统生成
    spu VARCHAR(50) UNIQUE COMMENT '商品SPU', -- 可以包含多个子SKU,用户多属性场景
    spu_attributes TEXT COMMENT '属性信息', -- {'material_sku': 'material_quantity'}
    sku VARCHAR(50) UNIQUE COMMENT '商品SKU',
    children_sku TEXT COMMENT '组合商品', -- 子商品SKU列表
    need_additional_process INT COMMENT '需要加工', -- 出货前是否需要加工
    sales_method ENUM('商品', '赠品', '包材') COMMENT '销售方式',
    status ENUM('在售', '热销', '新品', '清仓', '停售') COMMENT '商品状态',
    category_id BIGINT UNSIGNED COMMENT '商品分类',
    chinese_name VARCHAR(1024) COMMENT '中文名称',
    english_name VARCHAR(1024) COMMENT '英文名称',
    identifier_code VARCHAR(50) COMMENT '识别码',
    source_urls TEXT COMMENT '来源URL列表',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '商品创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '商品最后更新时间',
    deleted_at TIMESTAMP COMMENT '商品软删除时间（如果适用’）' -- 请注意这里的注释末尾的特殊字符
) CHARACTER SET utf8mb4;

-- Create SKU Mapping table
DROP Table IF EXISTS t_wms_sku_mapping;
CREATE TABLE t_wms_sku_mapping (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    goods_sku VARCHAR(50) COMMENT '商品SKU',
    product_sku VARCHAR(50) COMMENT '产品SKU',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    deleted_at TIMESTAMP COMMENT '软删除时间（如果适用）',
    UNIQUE KEY u_wms_sku_mapping (goods_sku, product_sku)
) CHARACTER SET utf8mb4;


-- Create t_logistics_packaging table
DROP TABLE IF EXISTS t_wms_logistics_packaging;
CREATE TABLE t_wms_logistics_packaging (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    goods_sku VARCHAR(50) UNIQUE COMMENT '商品SKU',
    declaration_chinese_name VARCHAR(50) COMMENT '报关中文名',
    declaration_english_name VARCHAR(50) COMMENT '报关英文名',
    declaration_price_currency VARCHAR(10) DEFAULT 'USD' COMMENT '报关价格币种',
    declaration_weight_unit VARCHAR(1) DEFAULT 'g' COMMENT '报关重量单位',
    material VARCHAR(255) COMMENT '材质',
    purpose VARCHAR(255) COMMENT '用途',
    customs_code VARCHAR(50) COMMENT '海关编码',
    special_goods ENUM('无', '含电(内电)', '纯电', '液体(特货)', '粉末(特货)', '膏体(特货)', '带磁(特货)', '含非液体化妆品') COMMENT '危险运输品',
    net_weight DECIMAL(10, 2) COMMENT '商品净重 (g)',
    weight_error_tolerance DECIMAL(10, 2) COMMENT '允许称重误差 (g)',
    length_cm DECIMAL(10, 2) COMMENT '商品尺寸 - 长 (cm)',
    width_cm DECIMAL(10, 2) COMMENT '商品尺寸 - 宽 (cm)',
    height_cm DECIMAL(10, 2) COMMENT '商品尺寸 - 高 (cm)',
    packaging TEXT COMMENT '包装信息', -- {'material_sku': 'material_quantity'}
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '包裹创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '包裹最后更新时间',
    deleted_at TIMESTAMP COMMENT '包裹软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

-- Supply and Demand
DROP TABLE IF EXISTS t_sd_supplier;
CREATE TABLE t_sd_supplier ( 
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    name VARCHAR(255) COMMENT '供应商名称',
    company_address VARCHAR(255) COMMENT '公司地址',
    company_website VARCHAR(255) COMMENT '公司网址',
    
    contact_name  VARCHAR(255) COMMENT '联系人',
    contact_telephone  VARCHAR(255) COMMENT '联系电话',
    bank_name  VARCHAR(255) COMMENT '银行名称',
    bank_account_number VARCHAR(255) COMMENT '银行卡号',
    bank_account_name VARCHAR(255) COMMENT '收款人',

    payment_method int COMMENT '付款方式',
    settlement_method int COMMENT '结算方式',

    remarks VARCHAR(1024) COMMENT '备注信息',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '包裹创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '包裹最后更新时间',
    deleted_at TIMESTAMP COMMENT '包裹软删除时间（如果适用）'
) CHARACTER SET utf8mb4;

INSERT INTO t_sd_supplier (
    name, company_address, company_website,
    contact_name, contact_telephone, bank_name,
    bank_account_number, bank_account_name,
    payment_method, settlement_method,
    remarks
) VALUES
    ('Supplier1', 'Address1', 'www.supplier1.com', 'John Doe', '123-456-7890', 'Bank1', '1234567890123456', 'John Doe', 1, 2, 'Sample remarks 1'),
    ('Supplier2', 'Address2', 'www.supplier2.com', 'Jane Smith', '987-654-3210', 'Bank2', '9876543210987654', 'Jane Smith', 2, 1, 'Sample remarks 2'),
    ('Supplier3', 'Address3', 'www.supplier3.com', 'Sam Johnson', '555-123-4567', 'Bank3', '5555123456789012', 'Sam Johnson', 1, 1, 'Sample remarks 3');


-- 规则引擎
DROP TABLE IF EXISTS t_rules;
CREATE TABLE t_rules ( 
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '数据库主键',
    domain VARCHAR(128) COMMENT '领域', --  订单规则， 物流规则，赠品规则等
    name VARCHAR(255) COMMENT '规则名称',
    rule_str TEXT COMMENT '规则条件',
    action_code INT COMMENT '动作代码',
    action_desc VARCHAR(255) COMMENT '动作描述',
    status INT COMMENT '规则状态', -- 启用 停用
    system_default INT COMMENT '系统默认规则', -- 只有系统预设的规则时为 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '包裹创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '包裹最后更新时间',
    deleted_at TIMESTAMP COMMENT '包裹软删除时间（如果适用）'
) CHARACTER SET utf8mb4;