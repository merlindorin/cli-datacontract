package bigquery

import (
	"errors"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/merlindorin/cli-datacontract/pkg/schema"
)

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

func MapFieldType(t bigquery.FieldType) (*schema.FieldType, error) {
	switch t {
	case bigquery.StringFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeString), nil
	case bigquery.BytesFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeBytes), nil
	case bigquery.IntegerFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeInteger), nil
	case bigquery.FloatFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeFloat), nil
	case bigquery.BooleanFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeBoolean), nil
	case bigquery.TimestampFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeTimestamp), nil
	case bigquery.TimeFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeTimestampNtz), nil
	case bigquery.DateFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeDate), nil
	case bigquery.DateTimeFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeTimestamp), nil
	case bigquery.NumericFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeNumeric), nil
	case bigquery.BigNumericFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeDouble), nil
	case bigquery.GeographyFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeObject), nil
	case bigquery.JSONFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeObject), nil
	case bigquery.RecordFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeObject), nil
	case bigquery.IntervalFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeInt), nil
	case bigquery.RangeFieldType:
		return schema.NewFieldTypeReference(schema.FieldTypeInt), nil
	}

	return nil, errors.New("unknown field type")
}

func MapTableType(t bigquery.TableType) (*schema.Type, error) {
	switch t {
	case bigquery.Snapshot, bigquery.ExternalTable, bigquery.RegularTable:
		return schema.NewTypeReference(schema.TypeTable), nil
	case bigquery.ViewTable, bigquery.MaterializedView:
		return schema.NewTypeReference(schema.TypeView), nil
	default:
		return nil, errors.New("unknown table type")
	}
}
