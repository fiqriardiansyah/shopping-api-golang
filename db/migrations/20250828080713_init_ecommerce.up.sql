-- Enable extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ===========================
-- Users & Roles
-- ===========================
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(150) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE roles (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name VARCHAR(50) UNIQUE NOT NULL,
                       alt_name VARCHAR(50) NOT NULL,
                       description TEXT
);

CREATE TABLE user_roles (
                            user_id UUID NOT NULL,
                            role_id UUID NOT NULL,
                            PRIMARY KEY (user_id, role_id),
                            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                            FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

-- ===========================
-- Products & Categories
-- ===========================
CREATE TABLE categories (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            name VARCHAR(100) NOT NULL,
                            description TEXT
);

CREATE TABLE products (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          seller_id UUID NOT NULL,
                          name VARCHAR(200) NOT NULL,
                          description TEXT,
                          price DECIMAL(10,2) NOT NULL,
                          stock INT DEFAULT 0,
                          category_id UUID,
                          created_at TIMESTAMP DEFAULT NOW(),
                          updated_at TIMESTAMP DEFAULT NOW(),
                          FOREIGN KEY (seller_id) REFERENCES users(id) ON DELETE CASCADE,
                          FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);

-- ===========================
-- Orders
-- ===========================
CREATE TABLE orders (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        buyer_id UUID NOT NULL,
                        status VARCHAR(20) CHECK (status IN ('PENDING','PAID','SHIPPED','COMPLETED','CANCELLED')) DEFAULT 'PENDING',
                        total_amount DECIMAL(10,2) NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW(),
                        FOREIGN KEY (buyer_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE order_items (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             order_id UUID NOT NULL,
                             product_id UUID NOT NULL,
                             quantity INT NOT NULL CHECK (quantity > 0),
                             price DECIMAL(10,2) NOT NULL,
                             FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
                             FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- ===========================
-- Cart
-- ===========================
CREATE TABLE cart (
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      buyer_id UUID NOT NULL UNIQUE,
                      created_at TIMESTAMP DEFAULT NOW(),
                      FOREIGN KEY (buyer_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE cart_items (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            cart_id UUID NOT NULL,
                            product_id UUID NOT NULL,
                            quantity INT NOT NULL CHECK (quantity > 0),
                            FOREIGN KEY (cart_id) REFERENCES cart(id) ON DELETE CASCADE,
                            FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- ===========================
-- Reviews
-- ===========================
CREATE TABLE reviews (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         buyer_id UUID NOT NULL,
                         product_id UUID NOT NULL,
                         rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
                         comment TEXT,
                         created_at TIMESTAMP DEFAULT NOW(),
                         FOREIGN KEY (buyer_id) REFERENCES users(id) ON DELETE CASCADE,
                         FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
