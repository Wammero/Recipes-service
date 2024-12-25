-- Создание таблицы пользователей, если она не существует
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Создание таблицы рецептов, если она не существует
CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    cooking_time INT,
    servings INT,
    cost INT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы категорий рецептов, если она не существует
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

-- Создание таблицы связи рецептов и категорий, если она не существует
CREATE TABLE IF NOT EXISTS recipe_categories (
    recipe_id INT NOT NULL REFERENCES recipes(id),
    category_id INT NOT NULL REFERENCES categories(id),
    UNIQUE (recipe_id, category_id)
);

-- Создание таблицы избранных рецептов, если она не существует
CREATE TABLE IF NOT EXISTS favorites (
    user_id INT NOT NULL REFERENCES users(id),
    recipe_id INT NOT NULL REFERENCES recipes(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, recipe_id)
);

-- Создание таблицы отзывов и оценок, если она не существует
CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    recipe_id INT NOT NULL REFERENCES recipes(id),
    rating INT CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, recipe_id)
);
