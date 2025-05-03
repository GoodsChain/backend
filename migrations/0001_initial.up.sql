-- Create tables
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE supplier (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(50),
  address VARCHAR(100),
  phone VARCHAR(20),
  email VARCHAR(100) UNIQUE,
  created_at TIMESTAMPTZ DEFAULT now(),
  created_by VARCHAR(50),
  updated_at TIMESTAMPTZ DEFAULT now(),
  updated_by VARCHAR(50)
);

CREATE TABLE car (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(50),
  supp_id UUID REFERENCES supplier(id),
  price INT,
  created_at TIMESTAMPTZ DEFAULT now(),
  created_by VARCHAR(50),
  updated_at TIMESTAMPTZ DEFAULT now(),
  updated_by VARCHAR(50)
);

CREATE TABLE customer (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100),
  address VARCHAR(100),
  phone VARCHAR(20),
  email VARCHAR(100) UNIQUE,
  created_at TIMESTAMPTZ DEFAULT now(),
  created_by VARCHAR(50),
  updated_at TIMESTAMPTZ DEFAULT now(),
  updated_by VARCHAR(50)
);

CREATE TABLE customer_car (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  car_id UUID REFERENCES car(id),
  cust_id UUID REFERENCES customer(id),
  created_at TIMESTAMPTZ DEFAULT now(),
  created_by VARCHAR(50),
  updated_at TIMESTAMPTZ DEFAULT now(),
  updated_by VARCHAR(5),
  UNIQUE (cust_id, car_id)
);

-- Seed data
INSERT INTO public.supplier (id, name, address, phone, email, created_by, updated_by) VALUES
('8a7d7fae-9c11-4b5c-8974-270a990a8934', 'Toyota Việt Nam', 'Số 15 Phạm Hùng, Hà Nội', '0243123456', 'contact@toyotavn.com.vn', 'admin', 'admin'),
('b2c0f971-5467-4c17-8461-1c0c4f12efb9', 'Honda Việt Nam', '33 Lê Duẩn, Hồ Chí Minh', '0283334444', 'info@hondavn.com.vn', 'admin', 'admin'),
('6d4f8c2e-5f7b-45a9-b9d2-e0e178f1b3c5', 'Mercedes Việt Nam', '2 Ngô Quyền, Hà Nội', '0246668888', 'sales@mercedesvn.com.vn', 'admin', 'admin'),
('9e7b5f3a-1d2c-4e8f-9a7b-3c5d8e9f2a1b', 'Ford Việt Nam', '123 Trần Hưng Đạo, Đà Nẵng', '0236777999', 'support@fordvietnam.vn', 'admin', 'admin'),
('4a8b7c6d-5e4f-3a2b-1c9d-8e7f6a5b4c3d', 'Hyundai Thành Công', '99 Nguyễn Thái Học, Hà Nội', '0245556677', 'info@hyundaitc.vn', 'admin', 'admin'),
('1f2e3d4c-5b6a-7c8d-9e0f-1a2b3c4d5e6f', 'Mazda Việt Nam', '68 Lê Lợi, Hồ Chí Minh', '0282223333', 'contact@mazdavn.com.vn', 'admin', 'admin'),
('1b64b862-e28a-465b-8cc9-54b14cc337ab', 'Kia Việt Nam', '45 Điện Biên Phủ, Hải Phòng', '0225559999', 'sales@kiavn.com.vn', 'admin', 'admin'),
('500b46d9-9fc1-4d2a-9eb4-bd024e740c15', 'Mitsubishi Motors', '55 Hoàng Diệu, Đà Nẵng', '0236444555', 'info@mitsubishivn.vn', 'admin', 'admin'),
('0ea791d0-bb23-42b7-8443-6699b1926756', 'Nissan Việt Nam', '77 Trần Phú, Hồ Chí Minh', '0287778888', 'contact@nissanvn.com.vn', 'admin', 'admin'),
('9fe1abcb-b926-417e-a5c9-c489d87386cd', 'Isuzu Việt Nam', '30 Lê Lai, Hà Nội', '0241112222', 'support@isuzuvn.vn', 'admin', 'admin');

INSERT INTO car (id, name, supp_id, price, created_by, updated_by) VALUES
('b1707beb-bd84-4f43-b501-ff5d0105f19d', 'Toyota Vios', '8a7d7fae-9c11-4b5c-8974-270a990a8934', 450000000, 'admin', 'admin'),
('a43c3181-b85f-4e64-ae3d-6782f5eef9c6', 'Honda City', 'b2c0f971-5467-4c17-8461-1c0c4f12efb9', 520000000, 'admin', 'admin'),
('d7e7d655-2e22-40b3-8f8d-6fa35e338577', 'Mercedes C200', '6d4f8c2e-5f7b-45a9-b9d2-e0e178f1b3c5', 1500000000, 'admin', 'admin'),
('99e69390-5744-44ad-8049-55ac9e1395bc', 'Ford Ranger', '9e7b5f3a-1d2c-4e8f-9a7b-3c5d8e9f2a1b', 850000000, 'admin', 'admin'),
('e06e08f9-e033-43ee-847e-88c53cdb7423', 'Hyundai Accent', '4a8b7c6d-5e4f-3a2b-1c9d-8e7f6a5b4c3d', 480000000, 'admin', 'admin'),
('b99cdf0a-47e0-4592-a851-e187287ffca5', 'Mazda CX-5', '1f2e3d4c-5b6a-7c8d-9e0f-1a2b3c4d5e6f', 900000000, 'admin', 'admin'),
('24f2f286-caef-4514-bc8a-1b0d569f49ad', 'Kia Seltos', '1b64b862-e28a-465b-8cc9-54b14cc337ab', 650000000, 'admin', 'admin'),
('710f8cc5-c0d8-40bb-ba43-476fe7222fec', 'Mitsubishi Xpander', '500b46d9-9fc1-4d2a-9eb4-bd024e740c15', 630000000, 'admin', 'admin'),
('27b9d6e4-6420-444c-ba9e-81d0d854b2d0', 'Nissan Terra', '0ea791d0-bb23-42b7-8443-6699b1926756', 950000000, 'admin', 'admin'),
('6238d49d-59a8-4b59-ba50-8b931bec1905', 'Isuzu D-Max', '9fe1abcb-b926-417e-a5c9-c489d87386cd', 800000000, 'admin', 'admin');

INSERT INTO customer (id, name, address, phone, email, created_by, updated_by) VALUES
('cf85cdb1-0879-4ba9-9fea-1452e344a568', 'Nguyễn Văn An', '56 Lý Thường Kiệt, Hà Nội', '0912345678', 'nguyenvanan@demo.com', 'admin', 'admin'),
('86f8fae9-9e1c-42bc-a59d-60c9ca87337b', 'Trần Thị Bình', '78 Nguyễn Huệ, Hồ Chí Minh', '0923456789', 'tranthiminh@demo.com', 'admin', 'admin'),
('56d7c910-d877-4ca8-9574-9b0df0bbc860', 'Lê Văn Cường', '23 Trần Phú, Đà Nẵng', '0934567890', 'levancuong@demo.com', 'admin', 'admin'),
('d4643aa3-d46f-4c86-b806-60a018092041', 'Phạm Thị Dung', '45 Lê Lợi, Hải Phòng', '0945678901', 'phamthidung@demo.com', 'admin', 'admin'),
('3adb12df-0412-4f35-b232-4419ca6de6e9', 'Hoàng Văn Em', '67 Điện Biên Phủ, Huế', '0956789012', 'hoangvanem@demo.com', 'admin', 'admin'),
('9083463b-92a5-4f43-8b9a-50dfc7d96eb3', 'Võ Thị Phương', '89 Trần Hưng Đạo, Cần Thơ', '0967890123', 'vothiphuong@demo.com', 'admin', 'admin'),
('5d9de7e7-ab3e-4cda-bd66-0f8591580988', 'Đặng Văn Giang', '12 Nguyễn Trãi, Quảng Ninh', '0978901234', 'dangvangiang@demo.com', 'admin', 'admin'),
('8da8acd2-41dd-4b24-90c6-8c2d6816a105', 'Bùi Thị Hương', '34 Bà Triệu, Thanh Hóa', '0989012345', 'buithihuong@demo.com', 'admin', 'admin'),
('6a035b3c-a09d-405b-a264-0e37f4e49ab0', 'Ngô Văn Ích', '56 Quang Trung, Nghệ An', '0990123456', 'ngovanich@demo.com', 'admin', 'admin'),
('b7d03d3a-24bf-4088-b06e-c5e454088a8f', 'Trương Thị Khánh', '78 Lý Thái Tổ, Khánh Hòa', '0901234567', 'truongthikhanh@demo.com', 'admin', 'admin');

INSERT INTO customer_car (id, car_id, cust_id, created_by, updated_by) VALUES
('edd0fd40-3747-44ca-a241-91bdf5ad73bb', 'b1707beb-bd84-4f43-b501-ff5d0105f19d', 'cf85cdb1-0879-4ba9-9fea-1452e344a568', 'admin', 'admin'),
('9fd43993-b03d-44a9-9529-7bb583d27162', 'a43c3181-b85f-4e64-ae3d-6782f5eef9c6', '86f8fae9-9e1c-42bc-a59d-60c9ca87337b', 'admin', 'admin'),
('2d3e874f-5b3e-47e7-aac1-cbcfe5cba5ee', 'd7e7d655-2e22-40b3-8f8d-6fa35e338577', '56d7c910-d877-4ca8-9574-9b0df0bbc860', 'admin', 'admin'),
('0e00782e-e492-4aed-adb0-22a459f8c1fd', '99e69390-5744-44ad-8049-55ac9e1395bc', 'd4643aa3-d46f-4c86-b806-60a018092041', 'admin', 'admin'),
('9c93887e-11ee-4461-8d35-51dc53d17493', 'e06e08f9-e033-43ee-847e-88c53cdb7423', '3adb12df-0412-4f35-b232-4419ca6de6e9', 'admin', 'admin'),
('ae9bc592-4a5e-48c6-b709-bcab8b403906', 'b99cdf0a-47e0-4592-a851-e187287ffca5', '9083463b-92a5-4f43-8b9a-50dfc7d96eb3', 'admin', 'admin'),
('2decd708-9f1a-494c-aaae-b1ca8a2001e2', '24f2f286-caef-4514-bc8a-1b0d569f49ad', '5d9de7e7-ab3e-4cda-bd66-0f8591580988', 'admin', 'admin'),
('fcfbeb46-927f-4123-bce7-1155277a3464', '710f8cc5-c0d8-40bb-ba43-476fe7222fec', '8da8acd2-41dd-4b24-90c6-8c2d6816a105', 'admin', 'admin'),
('84a8a25d-a8f5-45e1-8a37-de8096dcfca1', '27b9d6e4-6420-444c-ba9e-81d0d854b2d0', '6a035b3c-a09d-405b-a264-0e37f4e49ab0', 'admin', 'admin'),
('597b8e47-d540-4be4-a6cb-5ffbeb920a74', '6238d49d-59a8-4b59-ba50-8b931bec1905', 'b7d03d3a-24bf-4088-b06e-c5e454088a8f', 'admin', 'admin');