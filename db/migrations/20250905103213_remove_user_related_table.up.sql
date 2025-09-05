-- ===========================
-- UP: Remove users, roles, user_roles
-- ===========================

-- ===========================
-- Remove foreign keys referencing users
-- ===========================
ALTER TABLE products DROP CONSTRAINT products_seller_id_fkey;
ALTER TABLE orders DROP CONSTRAINT orders_buyer_id_fkey;
ALTER TABLE cart DROP CONSTRAINT cart_buyer_id_fkey;
ALTER TABLE reviews DROP CONSTRAINT reviews_buyer_id_fkey;

-- ===========================
-- Drop user_roles, roles, users
-- ===========================
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
