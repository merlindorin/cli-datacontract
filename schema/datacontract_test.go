package schema_test

import (
	_ "embed"
	"testing"

	"github.com/merlindorin/cli-datacontract/schema"
	"github.com/stretchr/testify/assert"
)

//go:embed sample_datacontract.json
var sample []byte

func TestSchema_UnmarshalJSON(t *testing.T) {
	got := schema.Datacontract{}
	want := schema.Datacontract{
		Specification: schema.DataContractSpecificationA093,
		ID:            "my-data-contract-id",
		Info: schema.Info{
			Title:   "My Data Contract",
			Version: "0.0.1",
		},
		Models: schema.Models{
			"incident": schema.Model{
				Type: schema.TypeTable,
				Fields: map[string]schema.AdditionalProperties{
					"id": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"_fivetran_deleted": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeBoolean),
					},
					"_fivetran_synced": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeTimestamp),
					},
					"call_url": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"created_at": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeTimestamp),
					},
					"creator_user_id": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"external_issue_reference_issue_name": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"external_issue_reference_issue_permalink": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"external_issue_reference_provider": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"incident_status_id": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"incident_type_id": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"mode": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"name": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"permalink": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"postmortem_document_url": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"reference": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"severity_id": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"slack_channel_id": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"slack_channel_name": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"slack_team_id": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"summary": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"updated_at": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeTimestamp),
					},
					"visibility": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeString),
					},
					"workload_minutes_late": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeFloat),
					},
					"workload_minutes_sleeping": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeFloat),
					},
					"workload_minutes_total": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeFloat),
					},
					"workload_minutes_working": {
						Type: schema.NewFieldTypeReference(schema.FieldTypeFloat),
					},
				},
			},
		},
	}

	err := got.UnmarshalJSON(sample)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
