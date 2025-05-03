# Tables

```sql
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
```

# Seed

```sql
INSERT INTO supplier (id, name, address, phone, email, created_by, updated_by) VALUES
('8a7d7fae-9c11-4b5c-8974-270a990a8934', 'Toyota Việt Nam', 'Số 15 Phạm Hùng, Hà Nội', '0243123456', 'contact@toyotavn.com.vn', 'admin', 'admin'),
('b2c0f971-5467-4c17-8461-1c0c4f12efb9', 'Honda Việt Nam', '33 Lê Duẩn, Hồ Chí Minh', '0283334444', 'info@hondavn.com.vn', 'admin', 'admin'),
('6d4f8c2e-5f7b-45a9-b9d2-e0e178f1b3c5', 'Mercedes Việt Nam', '2 Ngô Quyền, Hà Nội', '0246668888', 'sales@mercedesvn.com.vn', 'admin', 'admin'),
('9e7b5f3a-1d2c-4e8f-9a7b-3c5d8e9f2a1b', 'Ford Việt Nam', '123 Trần Hưng Đạo, Đà Nẵng', '0236777999', 'support@fordvietnam.vn', 'admin', 'admin'),
('4a8b7c6d-5e4f-3a2b-1c9d-8e7f6a5b4c3d', 'Hyundai Thành Công', '99 Nguyễn Thái Học, Hà Nội', '0245556677', 'info@hyundaitc.vn', 'admin', 'admin'),
('1f2e3d4c-5b6a-7c8d-9e0f-1a2b3c4d5e6f', 'Mazda Việt Nam', '68 Lê Lợi, Hồ Chí Minh', '0282223333', 'contact@mazdavn.com.vn', 'admin', 'admin'),
('7g6f5e4d-3c2b-1a9f-8e7d-6c5b4a3f2e1d', 'Kia Việt Nam', '45 Điện Biên Phủ, Hải Phòng', '0225559999', 'sales@kiavn.com.vn', 'admin', 'admin'),
('2h3i4j5k-6l7m-8n9o-0p1q-2r3s4t5u6v7w', 'Mitsubishi Motors', '55 Hoàng Diệu, Đà Nẵng', '0236444555', 'info@mitsubishivn.vn', 'admin', 'admin'),
('8x7y6z5a-4b3c-2d1e-9f8g-7h6i5j4k3l2m', 'Nissan Việt Nam', '77 Trần Phú, Hồ Chí Minh', '0287778888', 'contact@nissanvn.com.vn', 'admin', 'admin'),
('3n4o5p6q-7r8s-9t0u-1v2w-3x4y5z6a7b8c', 'Isuzu Việt Nam', '30 Lê Lai, Hà Nội', '0241112222', 'support@isuzuvn.vn', 'admin', 'admin');

INSERT INTO car (id, name, supp_id, price, created_by, updated_by) VALUES
('5e6f7a8b-9c0d-1e2f-3a4b-5c6d7e8f9a0b', 'Toyota Vios', '8a7d7fae-9c11-4b5c-8974-270a990a8934', 450000000, 'admin', 'admin'),
('1c2d3e4f-5g6h-7i8j-9k0l-1m2n3o4p5q6r', 'Honda City', 'b2c0f971-5467-4c17-8461-1c0c4f12efb9', 520000000, 'admin', 'admin'),
('7s8t9u0v-1w2x-3y4z-5a6b-7c8d9e0f1g2h', 'Mercedes C200', '6d4f8c2e-5f7b-45a9-b9d2-e0e178f1b3c5', 1500000000, 'admin', 'admin'),
('3i4j5k6l-7m8n-9o0p-1q2r-3s4t5u6v7w8x', 'Ford Ranger', '9e7b5f3a-1d2c-4e8f-9a7b-3c5d8e9f2a1b', 850000000, 'admin', 'admin'),
('9y0z1a2b-3c4d-5e6f-7g8h-9i0j1k2l3m4n', 'Hyundai Accent', '4a8b7c6d-5e4f-3a2b-1c9d-8e7f6a5b4c3d', 480000000, 'admin', 'admin'),
('5o6p7q8r-9s0t-1u2v-3w4x-5y6z7a8b9c0d', 'Mazda CX-5', '1f2e3d4c-5b6a-7c8d-9e0f-1a2b3c4d5e6f', 900000000, 'admin', 'admin'),
('1e2f3g4h-5i6j-7k8l-9m0n-1o2p3q4r5s6t', 'Kia Seltos', '7g6f5e4d-3c2b-1a9f-8e7d-6c5b4a3f2e1d', 650000000, 'admin', 'admin'),
('7u8v9w0x-1y2z-3a4b-5c6d-7e8f9g0h1i2j', 'Mitsubishi Xpander', '2h3i4j5k-6l7m-8n9o-0p1q-2r3s4t5u6v7w', 630000000, 'admin', 'admin'),
('3k4l5m6n-7o8p-9q0r-1s2t-3u4v5w6x7y8z', 'Nissan Terra', '8x7y6z5a-4b3c-2d1e-9f8g-7h6i5j4k3l2m', 950000000, 'admin', 'admin'),
('9a0b1c2d-3e4f-5g6h-7i8j-9k0l1m2n3o4p', 'Isuzu D-Max', '3n4o5p6q-7r8s-9t0u-1v2w-3x4y5z6a7b8c', 800000000, 'admin', 'admin');

INSERT INTO customer (id, name, address, phone, email, created_by, updated_by) VALUES
('a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6', 'Nguyễn Văn An', '56 Lý Thường Kiệt, Hà Nội', '0912345678', 'nguyenvanan@demo.com', 'admin', 'admin'),
('q7r8s9t0-u1v2-w3x4-y5z6-a7b8c9d0e1f2', 'Trần Thị Bình', '78 Nguyễn Huệ, Hồ Chí Minh', '0923456789', 'tranthiminh@demo.com', 'admin', 'admin'),
('g3h4i5j6-k7l8-m9n0-o1p2-q3r4s5t6u7v8', 'Lê Văn Cường', '23 Trần Phú, Đà Nẵng', '0934567890', 'levancuong@demo.com', 'admin', 'admin'),
('w9x0y1z2-a3b4-c5d6-e7f8-g9h0i1j2k3l4', 'Phạm Thị Dung', '45 Lê Lợi, Hải Phòng', '0945678901', 'phamthidung@demo.com', 'admin', 'admin'),
('m5n6o7p8-q9r0-s1t2-u3v4-w5x6y7z8a9b0', 'Hoàng Văn Em', '67 Điện Biên Phủ, Huế', '0956789012', 'hoangvanem@demo.com', 'admin', 'admin'),
('c1d2e3f4-g5h6-i7j8-k9l0-m1n2o3p4q5r6', 'Võ Thị Phương', '89 Trần Hưng Đạo, Cần Thơ', '0967890123', 'vothiphuong@demo.com', 'admin', 'admin'),
('s7t8u9v0-w1x2-y3z4-a5b6-c7d8e9f0g1h2', 'Đặng Văn Giang', '12 Nguyễn Trãi, Quảng Ninh', '0978901234', 'dangvangiang@demo.com', 'admin', 'admin'),
('i3j4k5l6-m7n8-o9p0-q1r2-s3t4u5v6w7x8', 'Bùi Thị Hương', '34 Bà Triệu, Thanh Hóa', '0989012345', 'buithihuong@demo.com', 'admin', 'admin'),
('y9z0a1b2-c3d4-e5f6-g7h8-i9j0k1l2m3n4', 'Ngô Văn Ích', '56 Quang Trung, Nghệ An', '0990123456', 'ngovanich@demo.com', 'admin', 'admin'),
('o5p6q7r8-s9t0-u1v2-w3x4-y5z6a7b8c9d0', 'Trương Thị Khánh', '78 Lý Thái Tổ, Khánh Hòa', '0901234567', 'truongthikhanh@demo.com', 'admin', 'admin');

INSERT INTO customer_car (id, car_id, cust_id, created_by, updated_by) VALUES
('e1f2g3h4-i5j6-k7l8-m9n0-o1p2q3r4s5t6', '5e6f7a8b-9c0d-1e2f-3a4b-5c6d7e8f9a0b', 'a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6', 'admin', 'admin'),
('u7v8w9x0-y1z2-a3b4-c5d6-e7f8g9h0i1j2', '1c2d3e4f-5g6h-7i8j-9k0l-1m2n3o4p5q6r', 'q7r8s9t0-u1v2-w3x4-y5z6-a7b8c9d0e1f2', 'admin', 'admin'),
('k3l4m5n6-o7p8-q9r0-s1t2-u3v4w5x6y7z8', '7s8t9u0v-1w2x-3y4z-5a6b-7c8d9e0f1g2h', 'g3h4i5j6-k7l8-m9n0-o1p2-q3r4s5t6u7v8', 'admin', 'admin'),
('a9b0c1d2-e3f4-g5h6-i7j8-k9l0m1n2o3p4', '3i4j5k6l-7m8n-9o0p-1q2r-3s4t5u6v7w8x', 'w9x0y1z2-a3b4-c5d6-e7f8-g9h0i1j2k3l4', 'admin', 'admin'),
('q5r6s7t8-u9v0-w1x2-y3z4-a5b6c7d8e9f0', '9y0z1a2b-3c4d-5e6f-7g8h-9i0j1k2l3m4n', 'm5n6o7p8-q9r0-s1t2-u3v4-w5x6y7z8a9b0', 'admin', 'admin'),
('g1h2i3j4-k5l6-m7n8-o9p0-q1r2s3t4u5v6', '5o6p7q8r-9s0t-1u2v-3w4x-5y6z7a8b9c0d', 'c1d2e3f4-g5h6-i7j8-k9l0-m1n2o3p4q5r6', 'admin', 'admin'),
('w7x8y9z0-a1b2-c3d4-e5f6-g7h8i9j0k1l2', '1e2f3g4h-5i6j-7k8l-9m0n-1o2p3q4r5s6t', 's7t8u9v0-w1x2-y3z4-a5b6-c7d8e9f0g1h2', 'admin', 'admin'),
('m3n4o5p6-q7r8-s9t0-u1v2-w3x4y5z6a7b8', '7u8v9w0x-1y2z-3a4b-5c6d-7e8f9g0h1i2j', 'i3j4k5l6-m7n8-o9p0-q1r2-s3t4u5v6w7x8', 'admin', 'admin'),
('c9d0e1f2-g3h4-i5j6-k7l8-m9n0o1p2q3r4', '3k4l5m6n-7o8p-9q0r-1s2t-3u4v5w6x7y8z', 'y9z0a1b2-c3d4-e5f6-g7h8-i9j0k1l2m3n4', 'admin', 'admin'),
('s5t6u7v8-w9x0-y1z2-a3b4-c5d6e7f8g9h0', '9a0b1c2d-3e4f-5g6h-7i8j-9k0l1m2n3o4p', 'o5p6q7r8-s9t0-u1v2-w3x4-y5z6a7b8c9d0', 'admin', 'admin');
```