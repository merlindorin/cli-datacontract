package bigquery_test

import (
	_ "embed"
	"testing"

	bq "cloud.google.com/go/bigquery"
	"github.com/merlindorin/cli-datacontract/pkg/bigquery"
	"github.com/merlindorin/cli-datacontract/pkg/schema"
	"github.com/stretchr/testify/assert"
)

//go:embed sample-bigquery-table-fields.json
var testData []byte

func TestMapType(t *testing.T) {
	type fields struct {
		schema []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]schema.FieldType
		wantErr bool
	}{
		{
			name: "sample",
			fields: fields{
				schema: testData,
			},
			want: map[string]schema.FieldType{
				"id":                                  schema.FieldTypeString,
				"_fivetran_deleted":                   schema.FieldTypeBoolean,
				"_fivetran_synced":                    schema.FieldTypeTimestamp,
				"call_url":                            schema.FieldTypeString,
				"created_at":                          schema.FieldTypeTimestamp,
				"creator_user_id":                     schema.FieldTypeString,
				"external_issue_reference_issue_name": schema.FieldTypeString,
				"external_issue_reference_issue_permalink": schema.FieldTypeString,
				"external_issue_reference_provider":        schema.FieldTypeString,
				"incident_status_id":                       schema.FieldTypeString,
				"incident_type_id":                         schema.FieldTypeString,
				"mode":                                     schema.FieldTypeString,
				"name":                                     schema.FieldTypeString,
				"permalink":                                schema.FieldTypeString,
				"postmortem_document_url":                  schema.FieldTypeString,
				"reference":                                schema.FieldTypeString,
				"severity_id":                              schema.FieldTypeString,
				"slack_channel_id":                         schema.FieldTypeString,
				"slack_channel_name":                       schema.FieldTypeString,
				"slack_team_id":                            schema.FieldTypeString,
				"summary":                                  schema.FieldTypeString,
				"updated_at":                               schema.FieldTypeTimestamp,
				"visibility":                               schema.FieldTypeString,
				"workload_minutes_late":                    schema.FieldTypeFloat,
				"workload_minutes_sleeping":                schema.FieldTypeFloat,
				"workload_minutes_total":                   schema.FieldTypeFloat,
				"workload_minutes_working":                 schema.FieldTypeFloat,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc, err := bq.SchemaFromJSON(testData)
			assert.Nil(t, err)

			for _, s := range sc {
				bt, er := bigquery.MapFieldType(s.Type)
				assert.Nil(t, er)
				assert.Equal(t, tt.want[s.Name], *bt)
			}
		})
	}
}

func TestMapTableType(t *testing.T) {
	type args struct {
		t bq.TableType
	}
	tests := []struct {
		name    string
		args    args
		want    schema.Type
		wantErr bool
	}{
		{
			name: "regularTable",
			args: args{
				t: bq.RegularTable,
			},
			want: schema.TypeTable,
		},
		{
			name: "external",
			args: args{
				t: bq.ExternalTable,
			},
			want: schema.TypeTable,
		},
		{
			name: "snapshot",
			args: args{
				t: bq.Snapshot,
			},
			want: schema.TypeTable,
		},
		{
			name: "view",
			args: args{
				t: bq.ViewTable,
			},
			want: schema.TypeView,
		},
		{
			name: "materialized view",
			args: args{
				t: bq.MaterializedView,
			},
			want: schema.TypeView,
		},
		{
			name: "unknown",
			args: args{
				t: bq.TableType("lorem"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bigquery.MapTableType(tt.args.t)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, *got)
			}
		})
	}
}
