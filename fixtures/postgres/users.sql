INSERT INTO users (id, created_at, updated_at)
SELECT generate_series, '2023-09-17 12:00:00', '2023-09-17 12:00:00'
FROM generate_series(1, 1000 * 1000)
ON CONFLICT (id) DO UPDATE
    SET updated_at = '2023-09-17 12:00:00';

INSERT INTO user_social_profiles (user_id, social_provider, social_provider_user_id, email, username, name, created_at, updated_at)
VALUES (1, 'github', '63663261', 'yaroslav.podorvanov@readytotouch.com', 'YaroslavPodorvanov', 'Yaroslav Podorvanov', '2023-09-17 12:00:00', '2023-09-17 12:00:00'),
       (1, 'github', '97196828', 'victor.poprozhuk@readytotouch.com', 'VictorPoprozhuk', 'Victor Poprozhuk', '2023-09-17 12:00:00', '2023-09-17 12:00:00'),
       (1, 'github', '123738314', 'anastasiia.sihetii@readytotouch.com', 'AnastasiiaSihetii', 'Anastasiia Sihetii', '2023-09-17 12:00:00', '2023-09-17 12:00:00');
