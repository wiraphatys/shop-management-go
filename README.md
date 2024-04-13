# Shop Management System

![ER Diagram](https://i.ibb.co/GTf1K3W/er-diagram-shop-management.png)

## Overview

The Shop Management System is a backend solution designed to streamline operations for shops, supermarkets, and similar businesses. It provides functionalities for managing customers, products, orders, and administrative tasks. This system is built with scalability, efficiency, and ease of use in mind, utilizing the Go programming language, GoFiber framework, GORM for database interaction, PostgreSQL for data storage, Docker for containerization, and Git for version control.

## Features

- **Customer Management**: Track customer information including name, email, address, city, and zip code.
- **Product Management**: Manage product details such as name, description, and unit price.
- **Order Management**: Process customer orders, including order lines specifying products and quantities.
- **Admin Management**: Securely manage administrative accounts with email and password authentication.

## Installation

1. Clone the repository:

   ```bash
   git clone github.com/wiraphatys/shop-management-go
   ```

2. Navigate to the project directory:

   ```bash
   cd shop-management-go
   ```

3. Setup .env
    ```
        # Server settings
        SERVER_HOST=
        SERVER_PORT=

        # Database settings 
        DB_HOST=
        DB_PORT=
        DB_USER=
        DB_PASS=
        DB_NAME=
        DB_SSL_MODE=disable
        DB_TIMEZONE=

        # Logging settings
        LOG_LEVEL=
        LOG_FILE=

        # JWT settings
        JWT_SECRET=
        JWT_EXPIRATION=

        # Database pool settings  
        DB_MAX_OPEN_CONNS=
        DB_MAX_IDLE_CONNS=
    ```

4. Build and run the Docker container:
    ```
        version: '3.8'

        services:
        postgres:
            image: postgres:latest
            container_name: postgres
            environment:
            POSTGRES_DB: <db_name>
            POSTGRES_USER: <username>
            POSTGRES_PASSWORD: <password>
            volumes:
            - postgres_data:/var/lib/postgresql/data
            ports:
            - "5432:5432"
            restart: unless-stopped

        pgadmin:
            image: dpage/pgadmin4:latest
            container_name: pgadmin
            environment:
            PGADMIN_DEFAULT_EMAIL: <email_to_login>
            PGADMIN_DEFAULT_PASSWORD: <password>
            ports:
            - "<external:internal>"
            depends_on:
            - postgres
            restart: unless-stopped

        volumes:
        postgres_data:
    ```

   ```bash
   docker-compose up -d
   ```

5. Access the API at `http://localhost:your_port`.

## Contributors

- [Wiraphat Yodsri](https://www.linkedin.com/in/ywiraphat/)
