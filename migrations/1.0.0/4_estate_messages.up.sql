CREATE TABLE IF NOT EXISTS `estate_messages` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sender_id` bigint(20) unsigned DEFAULT NULL,
  `recipient_id` bigint(20) unsigned DEFAULT NULL,
  `estate_id` bigint(20) unsigned DEFAULT NULL,
  `rent_id` bigint(20) unsigned DEFAULT NULL,
  `message` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `estate_messages_sender_id_foreign` (`sender_id`),
  KEY `estate_messages_recipient_id_foreign` (`recipient_id`),
  KEY `estate_messages_estate_id_foreign` (`estate_id`),
  KEY `estate_messages_rent_id_foreign` (`rent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;