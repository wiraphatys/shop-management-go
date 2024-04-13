# Shop Management System

![ER Diagram](https://i.ibb.co/X8fdhk7/shop-management-system-ER-diagram.png)

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

3. Build and run the Docker container:

   ```bash
   docker-compose up -d
   ```

4. Access the API at `http://localhost:your_port`.

## Contributors

- [Wiraphat Yodsri](https://www.linkedin.com/in/ywiraphat/)
