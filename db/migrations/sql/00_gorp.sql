CREATE TABLE gorp_migration (
                                id INT AUTO_INCREMENT PRIMARY KEY,
                                migration_name VARCHAR(255) NOT NULL,
                                applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);