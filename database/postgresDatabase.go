package database

import (
	"fmt"

	"github.com/wiraphatys/shop-management-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Name,
		cfg.Db.Port,
		cfg.Db.SSLMode,
		cfg.Db.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{
		Db: db,
	}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}

func CreateCustomerIDTrigger(db *gorm.DB) error {
	// Trigger function SQL
	triggerSQL := `
		CREATE OR REPLACE FUNCTION generate_customer_id()
		RETURNS TRIGGER AS $$
		DECLARE
			max_id BIGINT;
			new_id TEXT;
		BEGIN
			SELECT COALESCE(MAX(SUBSTRING(id, 2)::BIGINT), 0) INTO max_id FROM customers;
			new_id := 'C' || LPAD(CAST((max_id + 1) AS TEXT), 6, '0');
			NEW.id := new_id;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER customer_id_trigger
		BEFORE INSERT ON customers
		FOR EACH ROW
		EXECUTE PROCEDURE generate_customer_id();
	`

	// Run the trigger SQL
	return db.Exec(triggerSQL).Error
}

// CreateProductIDTrigger creates a trigger to automatically generate Product IDs before inserting into the Product table.
func CreateProductIDTrigger(db *gorm.DB) error {
	// Trigger function SQL
	triggerSQL := `
		CREATE OR REPLACE FUNCTION generate_product_id()
		RETURNS TRIGGER AS $$
		DECLARE
			max_id BIGINT;
			new_id TEXT;
		BEGIN
			SELECT COALESCE(MAX(SUBSTRING(id, 2)::BIGINT), 0) INTO max_id FROM products;
			new_id := 'P' || LPAD(CAST((max_id + 1) AS TEXT), 6, '0');
			NEW.id := new_id;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER product_id_trigger
		BEFORE INSERT ON products
		FOR EACH ROW
		EXECUTE PROCEDURE generate_product_id();
	`

	// Run the trigger SQL
	return db.Exec(triggerSQL).Error
}

// CreateOrderIDTrigger creates a trigger to automatically generate Order IDs before inserting into the Order table.
func CreateOrderIDTrigger(db *gorm.DB) error {
	// Trigger function SQL
	triggerSQL := `
		CREATE OR REPLACE FUNCTION generate_order_id()
		RETURNS TRIGGER AS $$
		DECLARE
			max_id BIGINT;
			new_id TEXT;
		BEGIN
			SELECT COALESCE(MAX(SUBSTRING(id, 2)::BIGINT), 0) INTO max_id FROM orders;
			new_id := 'O' || LPAD(CAST((max_id + 1) AS TEXT), 6, '0');
			NEW.id := new_id;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER order_id_trigger
		BEFORE INSERT ON orders
		FOR EACH ROW
		EXECUTE PROCEDURE generate_order_id();
	`

	// Run the trigger SQL
	return db.Exec(triggerSQL).Error
}

// CreateOrderLineIDTrigger creates a trigger to automatically generate OrderLine IDs before inserting into the OrderLine table.
func CreateOrderLineIDTrigger(db *gorm.DB) error {
	// Trigger function SQL
	triggerSQL := `
		CREATE OR REPLACE FUNCTION generate_order_line_id()
		RETURNS TRIGGER AS $$
		DECLARE
			max_id BIGINT;
			new_id TEXT;
		BEGIN
			SELECT COALESCE(MAX(SUBSTRING(id, 2)::BIGINT), 0) INTO max_id FROM order_lines;
			new_id := 'OL' || LPAD(CAST((max_id + 1) AS TEXT), 6, '0');
			NEW.id := new_id;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER order_line_id_trigger
		BEFORE INSERT ON order_lines
		FOR EACH ROW
		EXECUTE PROCEDURE generate_order_line_id();
	`

	// Run the trigger SQL
	return db.Exec(triggerSQL).Error
}

// CreateAdminIDTrigger creates a trigger to automatically generate Admin IDs before inserting into the Admin table.
func CreateAdminIDTrigger(db *gorm.DB) error {
	// Trigger function SQL
	triggerSQL := `
		CREATE OR REPLACE FUNCTION generate_admin_id()
		RETURNS TRIGGER AS $$
		DECLARE
			max_id BIGINT;
			new_id TEXT;
		BEGIN
			SELECT COALESCE(MAX(SUBSTRING(id, 2)::BIGINT), 0) INTO max_id FROM admins;
			new_id := 'A' || LPAD(CAST((max_id + 1) AS TEXT), 6, '0');
			NEW.id := new_id;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER admin_id_trigger
		BEFORE INSERT ON admins
		FOR EACH ROW
		EXECUTE PROCEDURE generate_admin_id();
	`

	// Run the trigger SQL
	return db.Exec(triggerSQL).Error
}
