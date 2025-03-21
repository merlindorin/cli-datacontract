package bigquery

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/merlindorin/cli-datacontract/schema"
)

// ImportSchema imports the schema for a BigQuery table into the provided
// datacontract. It maps the BigQuery table and field types to the corresponding
// schema types and stores them in the datacontract's models.
func ImportSchema(
	tableName string,
	tableDescription string,
	tableType bigquery.TableType,
	tableSchema []*bigquery.FieldSchema,
	datacontract *schema.Datacontract,
) error {
	t, err := MapTableType(tableType)
	if err != nil {
		return fmt.Errorf("cannot map table type: %w", err)
	}

	fields := map[string]schema.AdditionalProperties{}

	for _, bqField := range tableSchema {
		name := bqField.Name
		dtype, er := MapFieldType(bqField.Type)
		if er != nil {
			return fmt.Errorf("cannot map field type: %w", er)
		}
		fields[name] = schema.AdditionalProperties{
			Title: &name,
			Type:  dtype,
		}
	}

	if datacontract.Models == nil {
		datacontract.Models = map[string]schema.Model{}
	}

	datacontract.Models[tableName] = schema.Model{
		Title:       &tableName,
		Description: &tableDescription,
		Type:        *t,
		Fields:      fields,
	}

	return nil
}

// MapFieldType converts a BigQuery field type to the corresponding
// schema.FieldType. It returns an error if the field type is unknown.
func MapFieldType(t bigquery.FieldType) (*schema.FieldType, error) {
	switch t {
	case bigquery.StringFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeString), nil
	case bigquery.BytesFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeBytes), nil
	case bigquery.IntegerFieldType, "INT64":
		return schema.NewFieldTypeReference(schema.FieldTypeInteger), nil
	case bigquery.FloatFieldType, "FLOAT64":
		return schema.NewFieldTypeReference(schema.FieldTypeFloat), nil
	case bigquery.BooleanFieldType, "BOOL":
		return schema.NewFieldTypeReference(schema.FieldTypeBoolean), nil
	case bigquery.TimestampFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeTimestamp), nil
	case bigquery.TimeFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeTimestampNtz), nil
	case bigquery.DateFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeDate), nil
	case bigquery.DateTimeFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeTimestamp), nil
	case bigquery.NumericFieldType, "DECIMAL":
		return schema.NewFieldTypeReference(schema.FieldTypeNumeric), nil
	case bigquery.BigNumericFieldType, "BIGDECIMAL":
		return schema.NewFieldTypeReference(schema.FieldTypeDouble), nil
	case bigquery.GeographyFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeObject), nil
	case bigquery.JSONFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeObject), nil
	case bigquery.RecordFieldType, "STRUCT":
		return schema.NewFieldTypeReference(schema.FieldTypeObject), nil
	case bigquery.IntervalFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeInt), nil
	case bigquery.RangeFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeInt), nil
	}

	return nil, fmt.Errorf("unknown field type: %s", t)
}

// MapTableType converts a BigQuery table type to the corresponding
// schema.Type. It returns an error if the table type is unknown.
func MapTableType(t bigquery.TableType) (*schema.Type, error) {
	switch t {
	case bigquery.Snapshot, bigquery.ExternalTable, bigquery.RegularTable:
		return schema.NewTypeReference(schema.TypeTable), nil
	case bigquery.ViewTable, bigquery.MaterializedView:
		return schema.NewTypeReference(schema.TypeView), nil
	default:
		return nil, fmt.Errorf("unknown table type: `%s`", t)
	}
}
