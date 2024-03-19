//SELECT
//a.
SELECT name FROM `users` WHERE gender = 1
//b.
SELECT * FROM `products` WHERE id = 3
//c.
SELECT * FROM `users` WHERE created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW()
AND name LIKE '%a%';
//d.
SELECT COUNT(*) FROM users WHERE gender = 0
//e.
SELECT * FROM `users` ORDER BY name ASC
//f.
SELECT * FROM products LIMIT 5;

//UPDATE
//a.
UPDATE `products` SET `name`='[value-5]' WHERE id = 1
//b.
UPDATE `transactions_details` SET `qty`='3' WHERE id = 1

//DELETE
//a.
DELETE FROM `products` WHERE id = 1
//b.
DELETE FROM `products` WHERE product_type_id = 1

//JOIN, UNION, AGREGASI
//1.
SELECT * FROM `transactions` WHERE user_id = 1
UNION
SELECT * FROM `transactions` WHERE user_id = 2
//2.
SELECT
	sum(total_price) as jumlah_harga FROM transactions
WHERE
	user_id =1
//3.
SELECT 
	SUM(total_qty) as total_transaksi 
from 
	transactions 
WHERE 
	user_id = 2
//4.
SELECT 
	t.name,p.operator_id, p.code, p.name, p.status
FROM 
	products p
INNER JOIN
	product_types t
ON
	p.product_type_id = t.id
//5.
SELECT t.*, u.name AS user_name
	FROM 
transactions t
INNER JOIN 
	users u
ON 
	t.user_id = u.id
//6.
DELIMITER $$
CREATE FUNCTION delete_transaction_details(transaction_id INT) RETURNS VOID AS $$
BEGIN
    DELETE FROM transaction_details
    WHERE transaction_id = transaction_id;
END;
$$
DELIMITER ;
//7.
DELIMITER $$
CREATE FUNCTION update_total_qty(transaction_id INT) RETURNS VOID AS $$
DECLARE
    total_qty_update INT;
BEGIN

    SELECT SUM(qty) INTO total_qty_update
    FROM transaction_details
    WHERE transaction_id = transaction_id;

    UPDATE transactions
    SET total_qty = total_qty_update
    WHERE transaction_id = transaction_id;
END;
$$ 
DELIMITER ;
//8.
SELECT *
FROM products 
WHERE id NOT IN (
    SELECT DISTINCT product_id 
    FROM transaction_details
);
