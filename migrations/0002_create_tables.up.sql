CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE, 
    reiting INT CHECK (reiting >= 0 AND reiting <= 5) 
);

CREATE TABLE IF NOT EXISTS recipe_ingredients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE 
);

CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    description TEXT,
    author_id INT REFERENCES users(id) 
);

CREATE TABLE IF NOT EXISTS favourites (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ,
    recipe_id INT REFERENCES recipes(id) ,
    UNIQUE(user_id, recipe_id)
);

CREATE TABLE IF NOT EXISTS recipe_ingredient_links (
    recipe_id INT REFERENCES recipes(id),
    ingredient_id INT REFERENCES recipe_ingredients(id),
    PRIMARY KEY (recipe_id, ingredient_id) 
);