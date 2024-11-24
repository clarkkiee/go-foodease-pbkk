package types

import (
	"context"
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"

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
		SQL: "ST_SetSRID(ST_Point(?, ?), 4326)", // Use SRID 4326 for WGS84
		Vars: []interface{}{
			c.Longitude,
			c.Latitude,
		},
	}
}

func (c *Coordinates) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte: // WKB format
		// Convert WKB to WKT using PostGIS functions in SQL (recommended)
		return fmt.Errorf("WKB parsing not implemented. Use ST_AsText(coordinates) in your SQL query")

	case string: // WKT format: "POINT(longitude latitude)"
		str := strings.TrimSpace(v)
		re := regexp.MustCompile(`POINT\(([-\d.]+) ([-\d.]+)\)`)
		matches := re.FindStringSubmatch(str)
		if len(matches) != 3 {
			return fmt.Errorf("invalid WKT format: %s", str)
		}
		fmt.Sscanf(matches[1], "%f", &c.Longitude)
		fmt.Sscanf(matches[2], "%f", &c.Latitude)
		return nil

	default:
		return fmt.Errorf("unsupported data type: %T", value)
	}
}

func (c Coordinates) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", c.Longitude, c.Latitude), nil
}