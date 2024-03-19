-- Insert restaurants in Singapore
INSERT INTO restaurants (name, description, location, cuisine, image_url)
VALUES
  ('Chicken Rice Express', 'Famous for Hainanese Chicken Rice', '123 Orchard Road, Singapore', 'Chicken', './images/dummy.png'),
  ('Noodle Delight', 'Specializing in various noodle dishes', '456 Raffles Place, Singapore', 'Chinese', './images/dummy.png'),
  ('Sushi Haven', 'Authentic Japanese sushi and sashimi', '789 Marina Bay Sands, Singapore', 'Japanese', './images/dummy.png'),
  ('Spicy Noodle House', 'Home of spicy noodle lovers', '101 Tanjong Pagar Road, Singapore', 'Korean', './images/dummy.png'),
  ('Pizza Paradise', 'Delicious pizzas with a variety of toppings', '202 Sentosa Island, Singapore', 'Pizza', './images/dummy.png'),
  ('Curry House', 'Flavorful Indian curries', '303 Little India, Singapore', 'Indian', './images/dummy.png'),
  ('Thai Delights', 'Traditional Thai dishes with a modern twist', '404 Chinatown, Singapore', 'Thai', './images/dummy.png'),
  ('Dessert Dreamland', 'Satisfy your sweet tooth with our desserts', '505 Bugis Street, Singapore', 'Dessert', './images/dummy.png'),
  ('Healthy Bites', 'Nutritious and delicious healthy options', '606 Orchard Boulevard, Singapore', 'Healthy', './images/dummy.png'),
  ('Western Grill', 'Classic Western dishes for all tastes', '707 Clarke Quay, Singapore', 'Western', './images/dummy.png'),
  ('Satay Street', 'Famous for delicious satay skewers', '111 Marina Bay, Singapore', 'Satay', './images/dummy.png'),
  ('Dim Sum Delight', 'Authentic Cantonese dim sum experience', '222 Chinatown, Singapore', 'Chinese', './images/dummy.png'),
  ('Ramen House', 'Specializing in various ramen varieties', '333 Orchard Road, Singapore', 'Japanese', './images/dummy.png');

-- Insert dishes for each restaurant
INSERT INTO dishes (restaurant_id, is_available, name, description, price, cuisine, image_url)
VALUES
  -- Chicken Rice Express
  (1, true, 'Hainanese Chicken Rice', 'Classic Hainanese chicken with fragrant rice', 5.50, 'Chicken', './images/dummy.png'),
  (1, true, 'Roast Chicken Noodles', 'Roasted chicken served with noodles', 6.50, 'Chicken', './images/dummy.png'),
  (1, true, 'Chicken Satay Rice', 'Hainanese-style chicken rice with chicken satay', 8.50, 'Chicken', './images/dummy.png'),
  (1, true, 'Spicy Chicken Noodles', 'Spicy chicken noodles with a kick', 7.00, 'Chicken', './images/dummy.png'),

  -- Noodle Delight
  (2, true, 'Singapore-style Noodles', 'Stir-fried noodles with prawns and vegetables', 8.00, 'Chinese', './images/dummy.png'),
  (2, true, 'Beef Hor Fun', 'Flat rice noodles with beef in a savory sauce', 9.50, 'Chinese', './images/dummy.png'),
  (2, true, 'Char Kway Teow', 'Stir-fried flat rice noodles with prawns and Chinese sausage', 10.00, 'Chinese', './images/dummy.png'),
  (2, true, 'Wonton Noodle Soup', 'Noodle soup with wontons and BBQ pork', 8.50, 'Chinese', './images/dummy.png'),

  -- Sushi Haven
  (3, true, 'Assorted Sushi Platter', 'Chef''s selection of fresh sushi', 12.00, 'Japanese', './images/dummy.png'),
  (3, true, 'Sashimi Deluxe', 'Premium slices of assorted sashimi', 15.00, 'Japanese', './images/dummy.png'),
  (3, true, 'Chirashi Bowl', 'Assorted sashimi over a bed of sushi rice', 18.00, 'Japanese', './images/dummy.png'),
  (3, true, 'Dragon Roll', 'Sushi roll with eel, avocado, and cucumber', 14.00, 'Japanese', './images/dummy.png'),

  -- Spicy Noodle House
  (4, true, 'Spicy Ramen', 'Fiery ramen with spicy broth', 10.00, 'Korean', './images/dummy.png'),
  (4, true, 'Kimchi Udon', 'Udon noodles with kimchi and vegetables', 11.50, 'Korean', './images/dummy.png'),
  (4, true, 'Bibimbap', 'Mixed rice with vegetables, beef, and a spicy sauce', 12.00, 'Korean', './images/dummy.png'),
  (4, true, 'Spicy Kimchi Fried Rice', 'Fried rice with spicy kimchi and pork', 10.50, 'Korean', './images/dummy.png'),
  
  -- Pizza Paradise
  (5, true, 'Margherita Pizza', 'Classic pizza with tomato, mozzarella, and basil', 14.00, 'Pizza', './images/dummy.png'),
  (5, true, 'Pepperoni Delight', 'Pizza topped with pepperoni and cheese', 16.50, 'Pizza', './images/dummy.png'),
  (5, true, 'Seafood Pizza', 'Pizza topped with assorted seafood and herbs', 17.00, 'Pizza', './images/dummy.png'),
  (5, true, 'Vegetarian Supreme', 'Pizza loaded with assorted vegetables and cheese', 15.50, 'Pizza', './images/dummy.png'),

  -- Curry House
  (6, true, 'Chicken Curry', 'Spicy chicken curry with fragrant rice', 9.00, 'Indian', './images/dummy.png'),
  (6, true, 'Vegetable Biryani', 'Flavorful biryani with mixed vegetables', 11.00, 'Indian', './images/dummy.png'),
  (6, true, 'Butter Chicken', 'Creamy and flavorful butter chicken curry', 12.00, 'Indian', './images/dummy.png'),
  (6, true, 'Paneer Tikka Masala', 'Paneer in a spicy tomato-based curry', 13.50, 'Indian', './images/dummy.png'),

  -- Thai Delights
  (7, true, 'Tom Yum Soup', 'Spicy and sour soup with shrimp', 8.50, 'Thai', './images/dummy.png'),
  (7, true, 'Green Curry Chicken', 'Rich green curry with tender chicken', 10.50, 'Thai', './images/dummy.png'),

  -- Dessert Dreamland
  (8, true, 'Mango Sticky Rice', 'Sweet sticky rice topped with fresh mango', 6.00, 'Dessert', './images/dummy.png'),
  (8, true, 'Chocolate Lava Cake', 'Decadent chocolate cake with a gooey center', 7.50, 'Dessert', './images/dummy.png'),
  (8, true, 'Durian Ice Cream', 'Creamy durian-flavored ice cream', 7.00, 'Dessert', './images/dummy.png'),
  (8, true, 'Pandan Cake', 'Light and fluffy pandan-flavored cake', 6.50, 'Dessert', './images/dummy.png'),

  -- Healthy Bites
  (9, true, 'Quinoa Salad', 'Nutrient-packed salad with quinoa and veggies', 11.00, 'Healthy', './images/dummy.png'),
  (9, true, 'Grilled Salmon', 'Grilled salmon fillet with steamed broccoli', 13.50, 'Healthy', './images/dummy.png'),
  (9, true, 'Avocado Salad', 'Salad with avocado, cherry tomatoes, and mixed greens', 12.50, 'Healthy', './images/dummy.png'),
  (9, true, 'Teriyaki Tofu Bowl', 'Bowl with teriyaki tofu, brown rice, and vegetables', 11.00, 'Healthy', './images/dummy.png'),

  -- Western Grill
  (10, true, 'BBQ Ribs', 'Tender BBQ ribs with a savory glaze', 18.00, 'Western', './images/dummy.png'),
  (10, true, 'Classic Cheeseburger', 'Juicy beef patty with melted cheese', 14.50, 'Western', './images/dummy.png'),
  (10, true, 'Surf and Turf', 'Steak and grilled prawns served with a side of mashed potatoes', 20.00, 'Western', './images/dummy.png'),
  (10, true, 'Caesar Salad with Grilled Chicken', 'Classic Caesar salad with grilled chicken', 13.50, 'Western', './images/dummy.png'),
  
  -- Satay Street
  (11, true, 'Chicken Satay', 'Grilled chicken skewers with peanut sauce', 8.00, 'Satay', './images/dummy.png'),
  (11, true, 'Beef Satay', 'Grilled beef skewers with a flavorful marinade', 9.00, 'Satay', './images/dummy.png'),
  (11, true, 'Satay Platter', 'Assortment of grilled chicken and beef satay skewers', 15.00, 'Satay', './images/dummy.png'),
  (11, true, 'Satay Fried Rice', 'Fried rice with chopped chicken and beef satay', 12.50, 'Satay', './images/dummy.png'),

  -- Dim Sum Delight
  (12, true, 'Har Gow', 'Steamed shrimp dumplings with a thin translucent skin', 5.50, 'Chinese', './images/dummy.png'),
  (12, true, 'Siu Mai', 'Steamed pork dumplings with a savory filling', 6.50, 'Chinese', './images/dummy.png'),
  (12, true, 'Xiao Long Bao', 'Soup dumplings with a burst of flavorful broth inside', 7.50, 'Chinese', './images/dummy.png'),
  (12, true, 'Fried Spring Rolls', 'Crispy spring rolls filled with vegetables and shrimp', 6.00, 'Chinese', './images/dummy.png'),

  -- Ramen House
  (13, true, 'Shoyu Ramen', 'Soy-based ramen with tender slices of pork', 10.00, 'Japanese', './images/dummy.png'),
  (13, true, 'Miso Ramen', 'Ramen in a rich miso broth with vegetables', 11.50, 'Japanese', './images/dummy.png'),
  (13, true, 'Spicy Tan Tan Ramen', 'Spicy sesame-based broth with ground pork', 12.00, 'Japanese', './images/dummy.png'),
  (13, true, 'Vegetarian Miso Ramen', 'Ramen in miso broth with tofu and assorted vegetables', 11.00, 'Japanese', './images/dummy.png');


-- Insert sample users
INSERT INTO users (username, email, password_hash, address, uuid)
VALUES
  ('john_doe', 'john@example.com', 'hashed_password', '123 Main Street, Singapore', 'uuid1'),
  ('jane_smith', 'jane@example.com', 'hashed_password', '456 Orchard Road, Singapore', 'uuid2'),
  ('singapore_user', 'user@example.com', 'hashed_password', '789 Marina Bay Sands, Singapore', 'uuid3');

-- Insert sample searches
INSERT INTO searches (user_id, keyword)
VALUES
  (1, 'noodle delight'),
  (1, 'sushi haven'),
  (1, 'pizza paradise');

-- Insert sample playlists
INSERT INTO playlists (name, description, image_url, is_public, delivery_day, category)
VALUES
  ('Best of Singapore', 'A selection of the finest dishes in Singapore', './images/dummy.png', true, 'Sat', 'Local Delights'),
  ('Healthy Choices', 'Nutritious options for a balanced diet', './images/dummy.png', true, 'Sun', 'Healthy Living'),
  ('User 1 Playlist', 'My favorite', './images/dummy.png', false, 'Sun','');

-- Insert sample user playlists | use these seed data hours only: "11:00","11:30","12:00","12:30","13:00","13:30","17:00","17:30","18:00","18:30","19:00","19:30","20:00","20:30"
INSERT INTO user_playlists (user_id, playlist_id, delivery_day, delivery_time, status)
VALUES
  (1,3, 'Sunday', '11:00', 'Pending');

-- Insert sample playlist dishes
INSERT INTO playlist_dishes (order_id, playlist_id, dish_id, dish_quantity)
VALUES
  (1, 1, 1, 2),
  (1, 1, 4, 1),
  (1, 1, 10, 1),
  (2, 2, 7, 1),
  (2, 2, 8, 2),
  (3, 2, 10, 1),
  (3, 2, 12, 1),
  (4, 3, 10, 1),
  (4, 3, 12, 1);