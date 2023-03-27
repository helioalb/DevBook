insert into users (name, nick, email, password)
values
('User1', 'user1', 'user1@mail.com', '$2a$10$XsPRjVelxzHv9/7w9ewXK.eC8KaI9ETXzf9pgAiSzsQ3XPo599IN2A'),
('User2', 'user2', 'user2@mail.com', '$2a$10$XsPRjVelxzHv9/7w9ewXK.eC8KaI9ETXzf9pgAiSzsQ3XPo599IN2A'),
('User3', 'user3', 'user3@mail.com', '$2a$10$XsPRjVelxzHv9/7w9ewXK.eC8KaI9ETXzf9pgAiSzsQ3XPo599IN2A');

insert into followers (user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);
