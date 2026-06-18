-- name: InsertProduct :one
INSERT INTO products (
  name, 
  description, 
  ingredients, 
  nutrition, 
  weight_g, 
  dietary_tags, 
  allergens, 
  price_cents,
  currency,
  image_storage_key,
  image_alt
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
  RETURNING id, name, description, ingredients, nutrition, weight_g, dietary_tags, allergens, price_cents, currency, image_storage_key, image_alt;

-- name: GetProduct :one
SELECT
  id,
  name, 
  description, 
  ingredients, 
  nutrition, 
  dietary_tags, 
  allergens, 
  price_cents,
  currency,
  image_storage_key,
  image_alt
FROM products 
WHERE id = $1;