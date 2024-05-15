FROM php:8.3-fpm

# Обновление пакетов и установка зависимостей
RUN apt-get update && apt-get install -y \
    libpng-dev \
    libjpeg-dev \
    libzip-dev \
    libpq-dev \
    && docker-php-ext-configure gd --with-jpeg \
    && docker-php-ext-install gd \
    && docker-php-ext-install pdo_mysql \
    && docker-php-ext-install zip \
    && docker-php-ext-install pdo_pgsql

# Установка Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# Копирование файлов проекта
COPY . /var/www/html

# Установка зависимостей Composer
RUN composer install --no-scripts --no-autoloader

# Генерация ключа приложения
RUN php artisan key:generate

# Установка прав доступа
RUN chown -R www-data:www-data /var/www/html/storage /var/www/html/bootstrap/cache

# Запуск Apache с PHP-FPM
CMD ["php-fpm"]
