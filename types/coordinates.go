package types

import (
	"context"
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

func (c Coordinates) GormDataType() string {
	return "geometry"
}

func (c Coordinates) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL: "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", c.Longitude, c.Latitude)},
	}
}

func (c *Coordinates) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan coordinates: invalid data type %T", value)
	}

	// Parse nilai dari WKT (Well-Known Text)
	_, err := fmt.Sscanf(str, "POINT(%f %f)", &c.Longitude, &c.Latitude)
	if err != nil {
		return fmt.Errorf("failed to parse POINT: %v", err)
	}
	return nil
}

func (c Coordinates) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%f %f)", c.Longitude, c.Latitude), nil
}