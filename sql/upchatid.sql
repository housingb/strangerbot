ALTER TABLE `users` MODIFY COLUMN `chat_id` BIGINT(20) NOT NULL;
ALTER TABLE `users` MODIFY COLUMN `match_chat_id` BIGINT(20) NOT NULL;
ALTER TABLE `users` MODIFY COLUMN `previous_match` BIGINT(20) NOT NULL;