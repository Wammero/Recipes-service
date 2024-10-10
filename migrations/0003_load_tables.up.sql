-- Таблица пользователей
INSERT INTO users (name, reiting) VALUES
('Alice', 5),
('Bob', 4),
('Charlie', 3),
('Diana', 5),
('Eve', 2),
('Frank', 4),
('Grace', 3),
('Henry', 1),
('Isabella', 5),
('Jack', 2),
('Katherine', 4),
('Leo', 5),
('Mia', 3),
('Nathan', 1),
('Olivia', 2)
ON CONFLICT DO NOTHING;

-- Таблица ингредиентов
INSERT INTO recipe_ingredients (name) VALUES
('Flour'),
('Sugar'),
('Butter'),
('Milk'),
('Eggs'),
('Vanilla'),
('Salt'),
('Baking powder'),
('Cocoa powder'),
('Olive oil'),
('Garlic'),
('Tomatoes'),
('Basil'),
('Mozzarella'),
('Pasta')
ON CONFLICT DO NOTHING;

-- Таблица рецептов
INSERT INTO recipes (name, description, author_id) VALUES
('Chocolate Cake', 'Rich chocolate cake with layers of buttercream', 1),
('Tomato Pasta', 'Delicious pasta with fresh tomatoes and basil', 2),
('Garlic Bread', 'Crispy bread with garlic butter', 3),
('Vanilla Cupcakes', 'Soft cupcakes with vanilla flavor', 4),
('Cheese Pizza', 'Homemade pizza with fresh mozzarella and basil', 5),
('Pancakes', 'Fluffy pancakes with maple syrup', 6),
('Baked Eggs', 'Eggs baked with a touch of garlic and herbs', 7),
('Tiramisu', 'Classic Italian dessert with layers of coffee and cream', 8),
('Olive Oil Bread', 'Rustic bread with olive oil and garlic', 9),
('Tomato Basil Soup', 'Warm soup with tomatoes and fresh basil', 10),
('Scrambled Eggs', 'Classic scrambled eggs with a hint of butter', 11),
('Chocolate Muffins', 'Moist muffins with cocoa and chocolate chips', 12),
('Lasagna', 'Layers of pasta, cheese, and meat sauce', 13),
('Garlic Butter Pasta', 'Simple pasta dish with garlic butter sauce', 14),
('Mozzarella Sticks', 'Crispy mozzarella sticks with marinara sauce', 15)
ON CONFLICT DO NOTHING;

-- Таблица связей рецептов и ингредиентов
INSERT INTO recipe_ingredient_links (recipe_id, ingredient_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 5), (1, 9),
(2, 13), (2, 12), (2, 14), (2, 15),
(3, 3), (3, 10), (3, 11),
(4, 1), (4, 2), (4, 3), (4, 5), (4, 6),
(5, 12), (5, 13), (5, 14),
(6, 1), (6, 2), (6, 3), (6, 4), (6, 5),
(7, 5), (7, 11), (7, 13),
(8, 1), (8, 2), (8, 3), (8, 5), (8, 7),
(9, 3), (9, 10), (9, 11),
(10, 12), (10, 13), (10, 14),
(11, 5), (11, 3),
(12, 1), (12, 2), (12, 9), (12, 8),
(13, 1), (13, 3), (13, 15),
(14, 3), (14, 10), (14, 11),
(15, 13), (15, 12), (15, 14)
ON CONFLICT DO NOTHING;
