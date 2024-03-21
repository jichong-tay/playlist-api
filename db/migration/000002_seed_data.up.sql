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
  (1, true, 'Hainanese Chicken Rice', 'Classic Hainanese chicken with fragrant rice', 5.50, 'All Chicken Meat', 'https://images.deliveryhero.io/image/fd-sg/Products/5753294.jpg'),
  (1, true, 'Roast Chicken Noodles', 'Roasted chicken served with noodles', 6.50, 'All Chicken Meat', 'https://s3.burpple.com/foods/c7da9ee4c79e2a7641850131_original.?1595942085'),
  (1, true, 'Chicken Satay Rice', 'Hainanese-style chicken rice with chicken satay', 8.50, 'All Chicken Meat', 'https://img.taste.com.au/AnFpsY_0/w643-h428-cfill-q90/taste/2016/11/chicken-satay-sticks-with-jasmine-rice-76144-1.jpeg'),
  (1, true, 'Spicy Chicken Noodles', 'Spicy chicken noodles with a kick', 7.00, 'All Chicken Meat', 'https://tiffycooks.com/wp-content/uploads/2023/04/33EBDA80-8885-40BA-A456-2C80195FD918-768x1024.jpg'),

  -- Noodle Delight
  (2, true, 'Singapore-style Noodles', 'Stir-fried noodles with prawns and vegetables', 8.00, 'All Chinese', 'https://thewoksoflife.com/wp-content/uploads/2020/08/singapore-noodles-20.jpg'),
  (2, true, 'Beef Hor Fun', 'Flat rice noodles with beef in a savory sauce', 9.50, 'All Chinese Meat', 'https://delishar.com/wp-content/uploads/2016/04/Dry-Beef-Hor-Fun-3.jpg'),
  (2, true, 'Char Kway Teow', 'Stir-fried flat rice noodles with prawns and Chinese sausage', 10.00, 'All Chinese', 'https://salu-salo.com/wp-content/uploads/2013/02/Singaporean-Fried-Rice-Noodles-2.jpg'),
  (2, true, 'Wonton Noodle Soup', 'Noodle soup with wontons and BBQ pork', 8.50, 'All Chinese', 'https://cdn.sanity.io/images/2r0kdewr/production/352cdc37e6f98eb9dfe3567bca6886dd2dd9a570-1440x1080.jpg'),

  -- Sushi Haven
  (3, true, 'Assorted Sushi Platter', 'Chef''s selection of fresh sushi', 12.00, 'All Japanese Fish Seafood', 'https://media-cdn.tripadvisor.com/media/photo-s/0d/ae/65/c3/assorted-sushi-platter.jpg'),
  (3, true, 'Sashimi Deluxe', 'Premium slices of assorted sashimi', 15.00, 'All Japanese Fish Seafood', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTS-7OMRf_up9mGAfJeRBSpA6hGtVKnaaKjuP9FowAV6w&s'),
  (3, true, 'Chirashi Bowl', 'Assorted sashimi over a bed of sushi rice', 18.00, 'All Japanese Fish Seafood', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQgfCEVSz0a33WshHR1g4WtZ20wR6xIFbvMA34Gbvwezg&s'),
  (3, true, 'Dragon Roll', 'Sushi roll with eel, avocado, and cucumber', 14.00, 'All Japanese Fish Seafood', 'https://www.afarmgirlsdabbles.com/wp-content/uploads/2022/10/Dragon-Roll38797_1400.jpg'),

  -- Spicy Noodle House
  (4, true, 'Spicy Ramen', 'Fiery ramen with spicy broth', 10.00, 'All Korean Noodle', 'https://foodess.com/wp-content/uploads/2023/11/Spicy-Ramen-819x1024.jpg'),
  (4, true, 'Kimchi Udon', 'Udon noodles with kimchi and vegetables', 11.50, 'All Korean Noodle', 'https://assets.bonappetit.com/photos/57d96bcb5a14a530086ef6aa/1:1/w_2560%2Cc_limit/kimchi-udon-with-scallions.jpg'),
  (4, true, 'Bibimbap', 'Mixed rice with vegetables, beef, and a spicy sauce', 12.00, 'All Korean', 'https://images.squarespace-cdn.com/content/v1/51f7fb1ee4b03d20c9b4c34b/9d5c2e8a-012e-4a19-85a4-26f50b4623da/021_vegetarian-bibimbap-RHK.jpg'),
  (4, true, 'Spicy Kimchi Fried Rice', 'Fried rice with spicy kimchi and pork', 10.50, 'All Korean', 'https://www.recipetineats.com/wp-content/uploads/2021/03/Kimchi-Fried-Rice-Skillet.jpg'),
  
  -- Pizza Paradise
  (5, true, 'Margherita Pizza', 'Classic pizza with tomato, mozzarella, and basil', 14.00, 'All Pizza', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTa8X3YDM_ozN0RVGcb199-_PvkfNxyaVGaXDyEsD06iw&s'),
  (5, true, 'Pepperoni Delight', 'Pizza topped with pepperoni and cheese', 16.50, 'All Pizza', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSa9wV2crmZZJp53q9yuGSQ_hBvjaO6DJjVBp2UYpZZ3Q&s'),
  (5, true, 'Seafood Pizza', 'Pizza topped with assorted seafood and herbs', 17.00, 'All Pizza', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR8IxFAGf5ObWl51dnodEtLNuUICadZU4J4Yve5_kVlSA&s'),
  (5, true, 'Vegetarian Supreme', 'Pizza loaded with assorted vegetables and cheese', 15.50, 'All Pizza', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQbu7ig1ApKey54TTu_GXyk-XSauOu4y5cwR2e0QdyNbA&s'),

  -- Curry House
  (6, true, 'Chicken Curry', 'Spicy chicken curry with fragrant rice', 9.00, 'All Indian', 'https://sinfullyspicy.com/wp-content/uploads/2022/03/Dhaba-Style-Chicken-Curry-Featured-Image-480x270.jpg'),
  (6, true, 'Vegetable Biryani', 'Flavorful biryani with mixed vegetables', 11.00, 'All Indian', 'https://www.cookwithkushi.com/wp-content/uploads/2015/04/best_vegetable_biryani_recipe.jpg'),
  (6, true, 'Butter Chicken', 'Creamy and flavorful butter chicken curry', 12.00, 'All Indian', 'https://cafedelites.com/wp-content/uploads/2019/01/Butter-Chicken-IMAGE-64.jpg'),
  (6, true, 'Paneer Tikka Masala', 'Paneer in a spicy tomato-based curry', 13.50, 'All Indian', 'https://www.indianhealthyrecipes.com/wp-content/uploads/2014/11/paneer-tikka-masala-recipe-1.jpg'),

  -- Thai Delights
  (7, true, 'Tom Yum Soup', 'Spicy and sour soup with shrimp', 8.50, 'All Thai', 'https://nomadette.com/wp-content/uploads/2022/07/Tom-Yum-Soup.jpg'),
  (7, true, 'Green Curry Chicken', 'Rich green curry with tender chicken', 10.50, 'All Thai', 'https://www.kitchensanctuary.com/wp-content/uploads/2019/06/Thai-Green-Curry-square-FS.jpg'),

  -- Dessert Dreamland
  (8, true, 'Mango Sticky Rice', 'Sweet sticky rice topped with fresh mango', 6.00, 'All Dessert', 'https://hot-thai-kitchen.com/wp-content/uploads/2022/09/Mango-sticky-rice-blog.jpg'),
  (8, true, 'Chocolate Lava Cake', 'Decadent chocolate cake with a gooey center', 7.50, 'All Dessert', 'https://preppykitchen.com/wp-content/uploads/2022/03/Chocolate-Lava-Cake-Recipe.jpg'),
  (8, true, 'Durian Ice Cream', 'Creamy durian-flavored ice cream', 7.00, 'All Dessert', 'https://static.wixstatic.com/media/ebc1d6_900b2b8457524e119452303beb94858c~mv2.webp'),
  (8, true, 'Pandan Cake', 'Light and fluffy pandan-flavored cake', 6.50, 'All Dessert', 'https://ucarecdn.com/52208c4f-1597-4168-a233-7d9acfb338c9/-/scale_crop/1280x1280/center/-/quality/normal/-/format/jpeg/pandan-chiffon-cake.jpg'),

  -- Healthy Bites
  (9, true, 'Quinoa Salad', 'Nutrient-packed salad with quinoa and veggies', 11.00, 'All Healthy', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRpV34v0M2fqyLORJnzsT-l5D_MnNZofucWxnjeGV6k-Q&s'),
  (9, true, 'Grilled Salmon', 'Grilled salmon fillet with steamed broccoli', 13.50, 'All Healthy', 'https://www.wholesomeyum.com/wp-content/uploads/2021/06/wholesomeyum-Grilled-Salmon-Recipe-8.jpg'),
  (9, true, 'Avocado Salad', 'Salad with avocado, cherry tomatoes, and mixed greens', 12.50, 'All Healthy', 'https://www.inspiredtaste.net/wp-content/uploads/2019/03/Easy-Avocado-Salad-Recipe-2-1200.jpg'),
  (9, true, 'Teriyaki Tofu Bowl', 'Bowl with teriyaki tofu, brown rice, and vegetables', 11.00, 'All Healthy', 'https://www.justonecookbook.com/wp-content/uploads/2022/03/Teriyaki-Tofu-Bowl-6768-I.jpg'),

  -- Western Grill
  (10, true, 'BBQ Ribs', 'Tender BBQ ribs with a savory glaze', 18.00, 'All Western Meat', 'https://www.allrecipes.com/thmb/IWVelWahUb2gQxixWJC2N-HXp0k=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/22469-Barbecue-Ribs-ddmfs-2x1-210-e799db142f594b00bb317bb357d0971c.jpg'),
  (10, true, 'Classic Cheeseburger', 'Juicy beef patty with melted cheese', 14.50, 'All Western Meat', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS-g6sUPVusMb0BwYriCEWT1FFfGwuPvhyIfF4XYSaiGA&s'),
  (10, true, 'Surf and Turf', 'Steak and grilled prawns served with a side of mashed potatoes', 20.00, 'All Western Meat', 'https://natashaskitchen.com/wp-content/uploads/2022/02/Surf-and-Turf-Steak-and-Shrimp-SQ.jpg'),
  (10, true, 'Caesar Salad with Grilled Chicken', 'Classic Caesar salad with grilled chicken', 13.50, 'All Western Meat', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQMckPIYckZaRQ4WeTwKCk8EX0lttuGnPhC4lpI2xzrIA&s'),
  
  -- Satay Street
  (11, true, 'Chicken Satay', 'Grilled chicken skewers with peanut sauce', 8.00, 'All Satay Meat', 'https://rasamalaysia.com/wp-content/uploads/2019/04/chicken-satay-thumb.jpg'),
  (11, true, 'Beef Satay', 'Grilled beef skewers with a flavorful marinade', 9.00, 'All Satay Meat', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTSlVOhvxd7ijM9DZlIRVoki3fJx-Jj66TL1_b8VFXyVA&s'),
  (11, true, 'Satay Platter', 'Assortment of grilled chicken and beef satay skewers', 15.00, 'All Satay Meat', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTENiDRPJmTeS7l_TA62ggQ5aXk8i8E-1iqATdBFBFxwg&s'),
  (11, true, 'Satay Fried Rice', 'Fried rice with chopped chicken and beef satay', 12.50, 'All Satay Meat', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYPwNdFucSGt0y6GEFV-WYRBhZMrs7gtOYz3bDyDQewA&s'),

  -- Dim Sum Delight
  (12, true, 'Har Gow', 'Steamed shrimp dumplings with a thin translucent skin', 5.50, 'All Chinese', 'https://thewoksoflife.com/wp-content/uploads/2015/10/har-gow-6.jpg'),
  (12, true, 'Siu Mai', 'Steamed pork dumplings with a savory filling', 6.50, 'All Chinese', 'https://recipetineats.com/wp-content/uploads/2020/02/Siu-Mai_Chinese-steamed-dumplings_2.jpg'),
  (12, true, 'Xiao Long Bao', 'Soup dumplings with a burst of flavorful broth inside', 7.50, 'All Chinese', 'https://static.thehoneycombers.com/wp-content/uploads/sites/2/2019/09/xiao-long-bao-singapore.png'),
  (12, true, 'Fried Spring Rolls', 'Crispy spring rolls filled with vegetables and shrimp', 6.00, 'All Chinese', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSTM0nJRZyuAaio0b4JwAqadCOUciUc9qM1QMghuRNL0g&s'),

  -- Ramen House
  (13, true, 'Shoyu Ramen', 'Soy-based ramen with tender slices of pork', 10.00, 'All Japanese', 'https://media-cdn.tripadvisor.com/media/photo-s/1c/a6/36/97/ramen-shoyu-classic.jpg'),
  (13, true, 'Miso Ramen', 'Ramen in a rich miso broth with vegetables', 11.50, 'All Japanese', 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/601FDD58-68E1-47B8-A925-5A7DA7C8E7DE/Derivates/00ab9024-d404-47ef-810b-c48dcf1fdb39.jpg'),
  (13, true, 'Spicy Tan Tan Ramen', 'Spicy sesame-based broth with ground pork', 12.00, 'All Japanese', 'https://www.marionskitchen.com/wp-content/uploads/2020/03/Easy-%E2%80%98Tantanmen%E2%80%99-Ramen1.jpg'),
  (13, true, 'Vegetarian Miso Ramen', 'Ramen in miso broth with tofu and assorted vegetables', 11.00, 'All Japanese', 'https://images.squarespace-cdn.com/content/v1/5e185ce3e56525704fdae715/285175dc-3dda-4f26-8136-397d8d61dad6/DSC02846.jpg');


-- Insert sample users
INSERT INTO users (username, email, password_hash, address, uuid)
VALUES
  ('john_doe', 'john@example.com', 'hashed_password', '123 Main Street, Singapore', 'uuid1'),
  ('jane_smith', 'jane@example.com', 'hashed_password', '456 Orchard Road, Singapore', 'uuid2'),
  ('singapore_user', 'user@example.com', 'hashed_password', '789 Marina Bay Sands, Singapore', 'uuid3');

-- Insert sample searches
INSERT INTO searches (user_id, keyword)
VALUES
  (1, 'all'),
  (1, 'chicken rice'),
  (1, 'sushi'),
  (1, 'chinese'),
  (1, 'pizza');

-- Insert sample playlists
INSERT INTO playlists (name, description, image_url, is_public, delivery_day, category)
VALUES
  ('Best of Singapore', 'A selection of the finest dishes in Singapore', 'https://www.planetware.com/wpimages/2020/03/singapore-in-pictures-beautiful-places-to-photograph-marina-bay-sands.jpg', true,'', ''),
  ('Tasty Delights Delivery Mix', 'Enhance your culinary journey with curated tracks.', 'https://images.unsplash.com/photo-1700937192759-2a86c88128cc?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=400&ixid=MnwxfDB8MXxyYW5kb218MHx8fHx8fHx8MTcxMTAyNDM4OQ&ixlib=rb-4.0.3&q=80&w=400', true,'', ''),
  ('Foodie Favorites Express', 'Enjoy top foodie tunes during your delivery.', 'https://images.unsplash.com/photo-1559598467-8be25b6dc34f?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=400&ixid=MnwxfDB8MXxyYW5kb218MHx8Zm9vZHx8fHx8fDE3MTEwMjUzNjk&ixlib=rb-4.0.3&q=80&utm_campaign=api-credit&utm_medium=referral&utm_source=unsplash_source&w=400', true, '', ''),
  ('Quick Bite Delivery Tunes', 'Upbeat tracks for quick bites on the go.', 'https://images.unsplash.com/photo-1565299585323-38d6b0865b47?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=400&ixid=MnwxfDB8MXxyYW5kb218MHx8Zm9vZHx8fHx8fDE3MTEwMjU1MDI&ixlib=rb-4.0.3&q=80&utm_campaign=api-credit&utm_medium=referral&utm_source=unsplash_source&w=400', true, '', ''),
  ('Flavorful Feast Playlist', 'Dive into flavor sensations with this playlist.', 'https://images.unsplash.com/photo-1556040220-4096d522378d?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=400&ixid=MnwxfDB8MXxyYW5kb218MHx8Zm9vZHx8fHx8fDE3MTEwMjU1MTE&ixlib=rb-4.0.3&q=80&utm_campaign=api-credit&utm_medium=referral&utm_source=unsplash_source&w=400', true, '', ''),
  ('Yum yum Express Hits', 'Nutritious options for a balanced diet', 'https://images.unsplash.com/photo-1543362906-acfc16c67564?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=400&ixid=MnwxfDB8MXxyYW5kb218MHx8Zm9vZHx8fHx8fDE3MTEwMjU1MjQ&ixlib=rb-4.0.3&q=80&utm_campaign=api-credit&utm_medium=referral&utm_source=unsplash_source&w=400', true, '', ''),
  ('Healthy Choices', 'Top hits for waiting on delicious deliveries.', 'https://ch-api.healthhub.sg/api/public/content/2534118f690b4d4d8191831d743cb8c7?v=886c3102', true, 'Sun', ''),
  ('Favourite', 'My favorite', './images/dummy.png', false, 'Sun','');

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
  (2, 2, 11, 2),
  (2, 2, 15, 1),
  (2, 2, 20, 1),
  (3, 3, 12, 1),
  (3, 3, 18, 1),
  (3, 3, 21, 1),
  (4, 4, 5, 1),
  (4, 4, 9, 1),
  (4, 4, 10, 1),
  (5, 5, 11, 1),
  (5, 5, 12, 1),
  (5, 5, 13, 1),
  (6, 6, 40, 1),
  (6, 6, 41, 1),
  (6, 6, 42, 1),
  (7, 7, 10, 1),
  (7, 7, 12, 1),
  (7, 7, 12, 1);
  