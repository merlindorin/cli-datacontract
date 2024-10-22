//nolint:lll // most of the content here is lengthy for clarity
package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type AdditionalProperties struct {
	// A reference URI to a definition in the specification, internally or externally.
	// Properties will be inherited from the definition.
	Ref *string `json:"$ref,omitempty" yaml:"$ref,omitempty" mapstructure:"$ref,omitempty"`

	// The data class defining the sensitivity level for this field, according to the
	// organization's classification scheme.
	Classification *string `json:"classification,omitempty" yaml:"classification,omitempty" mapstructure:"classification,omitempty"`

	// Additional metadata for field configuration.
	Config *AdditionalPropertiesConfig `json:"config,omitempty" yaml:"config,omitempty" mapstructure:"config,omitempty"`

	// An optional string describing the semantic of the data in this field.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// A value must be equal to one of the elements in this array value. Only
	// evaluated if the value is not null.
	Enum []string `json:"enum,omitempty" yaml:"enum,omitempty" mapstructure:"enum,omitempty"`

	// An example value for this field.
	Example *string `json:"example,omitempty" yaml:"example,omitempty" mapstructure:"example,omitempty"`

	// A value of a number must less than the value of this. Only evaluated if the
	// value is not null. Only applies to numeric values.
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty" mapstructure:"exclusiveMaximum,omitempty"`

	// A value of a number must greater than the value of this. Only evaluated if the
	// value is not null. Only applies to numeric values.
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty" mapstructure:"exclusiveMinimum,omitempty"`

	// The nested fields (e.g. columns) of the object, record, or struct.
	Fields AdditionalPropertiesFields `json:"fields,omitempty" yaml:"fields,omitempty" mapstructure:"fields,omitempty"`

	// A specific format the value must comply with (e.g., 'email', 'uri', 'uuid').
	Format *string `json:"format,omitempty" yaml:"format,omitempty" mapstructure:"format,omitempty"`

	// Items corresponds to the JSON schema field "items".
	Items *AdditionalProperties `json:"items,omitempty" yaml:"items,omitempty" mapstructure:"items,omitempty"`

	// Keys corresponds to the JSON schema field "keys".
	Keys *AdditionalProperties `json:"keys,omitempty" yaml:"keys,omitempty" mapstructure:"keys,omitempty"`

	// Links to external resources.
	Links AdditionalPropertiesLinks `json:"links,omitempty" yaml:"links,omitempty" mapstructure:"links,omitempty"`

	// A value must less than, or equal to, the value of this. Only applies to string
	// types.
	MaxLength *int `json:"maxLength,omitempty" yaml:"maxLength,omitempty" mapstructure:"maxLength,omitempty"`

	// A value of a number must less than, or equal to, the value of this. Only
	// evaluated if the value is not null. Only applies to numeric values.
	Maximum *float64 `json:"maximum,omitempty" yaml:"maximum,omitempty" mapstructure:"maximum,omitempty"`

	// A value must greater than, or equal to, the value of this. Only applies to
	// string types.
	MinLength *int `json:"minLength,omitempty" yaml:"minLength,omitempty" mapstructure:"minLength,omitempty"`

	// A value of a number must greater than, or equal to, the value of this. Only
	// evaluated if the value is not null. Only applies to numeric values.
	Minimum *float64 `json:"minimum,omitempty" yaml:"minimum,omitempty" mapstructure:"minimum,omitempty"`

	// A regular expression the value must match. Only applies to string types.
	Pattern *string `json:"pattern,omitempty" yaml:"pattern,omitempty" mapstructure:"pattern,omitempty"`

	// An indication, if this field contains Personal Identifiable Information (PII).
	Pii *bool `json:"pii,omitempty" yaml:"pii,omitempty" mapstructure:"pii,omitempty"`

	// The maximum number of digits in a number. Only applies to numeric values.
	// Defaults to 38.
	Precision *float64 `json:"precision,omitempty" yaml:"precision,omitempty" mapstructure:"precision,omitempty"`

	// If this field is a primary key.
	Primary bool `json:"primary,omitempty" yaml:"primary,omitempty" mapstructure:"primary,omitempty"`

	// The reference to a field in another model. E.g. use 'orders.order_id' to
	// reference the order_id field of the model orders. Think of defining a foreign
	// key relationship.
	References *string `json:"references,omitempty" yaml:"references,omitempty" mapstructure:"references,omitempty"`

	// An indication, if this field must contain a value and may not be null.
	Required bool `json:"required,omitempty" yaml:"required,omitempty" mapstructure:"required,omitempty"`

	// The maximum number of decimal places in a number. Only applies to numeric
	// values. Defaults to 0.
	Scale *float64 `json:"scale,omitempty" yaml:"scale,omitempty" mapstructure:"scale,omitempty"`

	// Custom metadata to provide additional context.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags,omitempty"`

	// An optional string providing a human readable name for the field. Especially
	// useful if the field name is cryptic or contains abbreviations.
	Title *string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *FieldType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`

	// An indication, if the value must be unique within the model.
	Unique bool `json:"unique,omitempty" yaml:"unique,omitempty" mapstructure:"unique,omitempty"`

	// Values corresponds to the JSON schema field "values".
	Values *AdditionalProperties `json:"values,omitempty" yaml:"values,omitempty" mapstructure:"values,omitempty"`
}

// Additional metadata for field configuration.
type AdditionalPropertiesConfig struct {
	// Specify the logical field type to use when exporting the data model to Apache
	// Avro.
	AvroLogicalType *string `json:"avroLogicalType,omitempty" yaml:"avroLogicalType,omitempty" mapstructure:"avroLogicalType,omitempty"`

	// Specify the field type to use when exporting the data model to Apache Avro.
	AvroType *string `json:"avroType,omitempty" yaml:"avroType,omitempty" mapstructure:"avroType,omitempty"`

	// Specify the physical column type that is used in a Bigquery table, e.g.,
	// `NUMERIC(5, 2)`.
	BigqueryType *string `json:"bigqueryType,omitempty" yaml:"bigqueryType,omitempty" mapstructure:"bigqueryType,omitempty"`

	// Specify the physical column type that is used in a Databricks Unity Catalog
	// table.
	DatabricksType *string `json:"databricksType,omitempty" yaml:"databricksType,omitempty" mapstructure:"databricksType,omitempty"`

	// Specify the physical column type that is used in an AWS Glue Data Catalog
	// table.
	GlueType *string `json:"glueType,omitempty" yaml:"glueType,omitempty" mapstructure:"glueType,omitempty"`

	// Specify the physical column type that is used in a Redshift table, e.g.,
	// `SMALLINT`.
	RedshiftType *string `json:"redshiftType,omitempty" yaml:"redshiftType,omitempty" mapstructure:"redshiftType,omitempty"`

	// Specify the physical column type that is used in a Snowflake table, e.g.,
	// `TIMESTAMP_LTZ`.
	SnowflakeType *string `json:"snowflakeType,omitempty" yaml:"snowflakeType,omitempty" mapstructure:"snowflakeType,omitempty"`

	// Specify the physical column type that is used in a SQL Server table, e.g.,
	// `DATETIME2`.
	SqlserverType *string `json:"sqlserverType,omitempty" yaml:"sqlserverType,omitempty" mapstructure:"sqlserverType,omitempty"`
}

// The nested fields (e.g. columns) of the object, record, or struct.
type AdditionalPropertiesFields map[string]AdditionalProperties

// Links to external resources.
type AdditionalPropertiesLinks map[string]string

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalProperties) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain AdditionalProperties
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["primary"]; !ok || v == nil {
		plain.Primary = false
	}
	if v, ok := raw["required"]; !ok || v == nil {
		plain.Required = false
	}
	if v, ok := raw["unique"]; !ok || v == nil {
		plain.Unique = false
	}
	*j = AdditionalProperties(plain)
	return nil
}

// Additional metadata for model configuration.
type Config struct {
	// The namespace to use when importing and exporting the data model from / to
	// Apache Avro.
	AvroNamespace *string `json:"avroNamespace,omitempty" yaml:"avroNamespace,omitempty" mapstructure:"avroNamespace,omitempty"`
}

type Datacontract struct {
	// Specifies the Data Contract Specification being used.
	Specification DataContractSpecification `json:"dataContractSpecification" yaml:"dataContractSpecification" mapstructure:"dataContractSpecification"`

	// Clear and concise explanations of syntax, semantic, and classification of
	// business objects in a given domain.
	Definitions Definitions `json:"definitions,omitempty" yaml:"definitions,omitempty" mapstructure:"definitions,omitempty"`

	// The Examples Object is an array of Example Objects.
	Examples []ExamplesElem `json:"examples,omitempty" yaml:"examples,omitempty" mapstructure:"examples,omitempty"`

	// Specifies the identifier of the data contract.
	ID string `json:"id" yaml:"id" mapstructure:"id"`

	// Metadata and life cycle information about the data contract.
	Info Info `json:"info" yaml:"info" mapstructure:"info"`

	// Links to external resources.
	Links DataContractLinks `json:"links,omitempty" yaml:"links,omitempty" mapstructure:"links,omitempty"`

	// Specifies the logical data model. Use the models name (e.g., the table name) as
	// the key.
	Models Models `json:"models,omitempty" yaml:"models,omitempty" mapstructure:"models,omitempty"`

	// The quality object contains quality attributes and checks.
	Quality *Quality `json:"quality,omitempty" yaml:"quality,omitempty" mapstructure:"quality,omitempty"`

	// The schema of the data contract describes the syntax and semantics of provided
	// data sets. It supports different schema types.
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty" mapstructure:"schema,omitempty"`

	// Information about the servers.
	Servers *Servers `json:"servers,omitempty" yaml:"servers,omitempty" mapstructure:"servers,omitempty"`

	// Specifies the service level agreements for the provided data, including
	// availability, data retention policies, latency requirements, data freshness,
	// update frequency, support availability, and backup policies.
	Servicelevels *Servicelevels `json:"servicelevels,omitempty" yaml:"servicelevels,omitempty" mapstructure:"servicelevels,omitempty"`

	// Tags to facilitate searching and filtering.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags,omitempty"`

	// The terms and conditions of the data contract.
	Terms *Terms `json:"terms,omitempty" yaml:"terms,omitempty" mapstructure:"terms,omitempty"`
}

type DataContractSpecification string

const DataContractSpecificationA090 DataContractSpecification = "0.9.0"
const DataContractSpecificationA091 DataContractSpecification = "0.9.1"
const DataContractSpecificationA092 DataContractSpecification = "0.9.2"
const DataContractSpecificationA093 DataContractSpecification = "0.9.3"

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataContractSpecification) UnmarshalJSON(b []byte) error {
	var dataContractSpecifications = []string{
		"0.9.3",
		"0.9.2",
		"0.9.1",
		"0.9.0",
	}

	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range dataContractSpecifications {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", dataContractSpecifications, v)
	}
	*j = DataContractSpecification(v)
	return nil
}

// Clear and concise explanations of syntax, semantic, and classification of
// business objects in a given domain.
type Definitions map[string]struct {
	// The data class defining the sensitivity level for this field.
	Classification *string `json:"classification,omitempty" yaml:"classification,omitempty" mapstructure:"classification,omitempty"`

	// Clear and concise explanations related to the domain.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// The domain in which this definition is valid.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty" mapstructure:"domain,omitempty"`

	// An example value.
	Example *string `json:"example,omitempty" yaml:"example,omitempty" mapstructure:"example,omitempty"`

	// A value of a number must less than the value of this. Only evaluated if the
	// value is not null. Only applies to numeric values.
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty" mapstructure:"exclusiveMaximum,omitempty"`

	// A value of a number must greater than the value of this. Only evaluated if the
	// value is not null. Only applies to numeric values.
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty" mapstructure:"exclusiveMinimum,omitempty"`

	// The nested fields (e.g. columns) of the object, record, or struct.
	Fields Fields `json:"fields,omitempty" yaml:"fields,omitempty" mapstructure:"fields,omitempty"`

	// Specific format requirements for the value (e.g., 'email', 'uri', 'uuid').
	Format *string `json:"format,omitempty" yaml:"format,omitempty" mapstructure:"format,omitempty"`

	// Items corresponds to the JSON schema field "items".
	Items *AdditionalProperties `json:"items,omitempty" yaml:"items,omitempty" mapstructure:"items,omitempty"`

	// Keys corresponds to the JSON schema field "keys".
	Keys *AdditionalProperties `json:"keys,omitempty" yaml:"keys,omitempty" mapstructure:"keys,omitempty"`

	// Links to external resources.
	Links Links `json:"links,omitempty" yaml:"links,omitempty" mapstructure:"links,omitempty"`

	// A value must be less than or equal to this value. Applies only to string types.
	MaxLength *int `json:"maxLength,omitempty" yaml:"maxLength,omitempty" mapstructure:"maxLength,omitempty"`

	// A value of a number must less than, or equal to, the value of this. Only
	// evaluated if the value is not null. Only applies to numeric values.
	Maximum *float64 `json:"maximum,omitempty" yaml:"maximum,omitempty" mapstructure:"maximum,omitempty"`

	// A value must be greater than or equal to this value. Applies only to string
	// types.
	MinLength *int `json:"minLength,omitempty" yaml:"minLength,omitempty" mapstructure:"minLength,omitempty"`

	// A value of a number must greater than, or equal to, the value of this. Only
	// evaluated if the value is not null. Only applies to numeric values.
	Minimum *float64 `json:"minimum,omitempty" yaml:"minimum,omitempty" mapstructure:"minimum,omitempty"`

	// The technical name of this definition.
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// A regular expression pattern the value must match. Applies only to string
	// types.
	Pattern *string `json:"pattern,omitempty" yaml:"pattern,omitempty" mapstructure:"pattern,omitempty"`

	// Indicates if the field contains Personal Identifiable Information (PII).
	Pii *bool `json:"pii,omitempty" yaml:"pii,omitempty" mapstructure:"pii,omitempty"`

	// The maximum number of digits in a number. Only applies to numeric values.
	// Defaults to 38.
	Precision *int `json:"precision,omitempty" yaml:"precision,omitempty" mapstructure:"precision,omitempty"`

	// The maximum number of decimal places in a number. Only applies to numeric
	// values. Defaults to 0.
	Scale *int `json:"scale,omitempty" yaml:"scale,omitempty" mapstructure:"scale,omitempty"`

	// Custom metadata to provide additional context.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags,omitempty"`

	// The business name of this definition.
	Title *string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type FieldType `json:"type" yaml:"type" mapstructure:"type"`

	// Values corresponds to the JSON schema field "values".
	Values *AdditionalProperties `json:"values,omitempty" yaml:"values,omitempty" mapstructure:"values,omitempty"`
}

type ExamplesElem struct {
	// Data corresponds to the JSON schema field "data".
	Data interface{} `json:"data" yaml:"data" mapstructure:"data"`

	// An optional string describing the example.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// The reference to the model in the schema, e.g., a table name.
	Model *string `json:"model,omitempty" yaml:"model,omitempty" mapstructure:"model,omitempty"`

	// The type of the example data. Well-known types are csv, json, yaml, custom.
	Type ExamplesElemType `json:"type" yaml:"type" mapstructure:"type"`
}

type ExamplesElemType string

const ExamplesElemTypeCSV ExamplesElemType = "csv"
const ExamplesElemTypeCustom ExamplesElemType = "custom"
const ExamplesElemTypeJSON ExamplesElemType = "json"
const ExamplesElemTypeYAML ExamplesElemType = "yaml"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExamplesElemType) UnmarshalJSON(b []byte) error {
	var types = []string{
		"csv",
		"json",
		"yaml",
		"custom",
	}

	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range types {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", types, v)
	}
	*j = ExamplesElemType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExamplesElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["data"]; raw != nil && !ok {
		return fmt.Errorf("field data in ExamplesElem: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in ExamplesElem: required")
	}
	type Plain ExamplesElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ExamplesElem(plain)
	return nil
}

// Metadata and life cycle information about the data contract.
type Info struct {
	// Contact information for the data contract.
	Contact *InfoContact `json:"contact,omitempty" yaml:"contact,omitempty" mapstructure:"contact,omitempty"`

	// A description of the data contract.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// The owner or team responsible for managing the data contract and providing the
	// data.
	Owner *string `json:"owner,omitempty" yaml:"owner,omitempty" mapstructure:"owner,omitempty"`

	// The status of the data contract. Can be proposed, in development, active,
	// retired.
	Status *string `json:"status,omitempty" yaml:"status,omitempty" mapstructure:"status,omitempty"`

	// The title of the data contract.
	Title string `json:"title" yaml:"title" mapstructure:"title"`

	// The version of the data contract document (which is distinct from the Data
	// Contract Specification version or the Data Product implementation version).
	Version string `json:"version" yaml:"version" mapstructure:"version"`

	AdditionalProperties interface{}
}

// Contact information for the data contract.
type InfoContact struct {
	// The email address of the contact person/organization. This MUST be in the form
	// of an email address.
	Email *string `json:"email,omitempty" yaml:"email,omitempty" mapstructure:"email,omitempty"`

	// The identifying name of the contact person/organization.
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// The URL pointing to the contact information. This MUST be in the form of a URL.
	URL *string `json:"url,omitempty" yaml:"url,omitempty" mapstructure:"url,omitempty"`

	AdditionalProperties interface{}
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Info) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	if _, ok := raw["title"]; raw != nil && !ok {
		return fmt.Errorf("field title in Info: required")
	}

	if _, ok := raw["version"]; raw != nil && !ok {
		return fmt.Errorf("field version in Info: required")
	}

	type Plain Info
	var plain Plain

	//nolint:musttag // false positive
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}

	*j = Info(plain)

	return nil
}

// Links to external resources.
type DataContractLinks map[string]string

// Specifies the logical data model. Use the models name (e.g., the table name) as
// the key.
type Models map[string]Model

type Model struct {
	// Additional metadata for model configuration.
	Config *Config `json:"config,omitempty" yaml:"config,omitempty" mapstructure:"config,omitempty"`

	// Description corresponds to the JSON schema field "description".
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// Specifies a field in the data model. Use the field name (e.g., the column name)
	// as the key.
	Fields Fields `json:"fields,omitempty" yaml:"fields,omitempty" mapstructure:"fields,omitempty"`

	// An optional string providing a human readable name for the model. Especially
	// useful if the model name is cryptic or contains abbreviations.
	Title *string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`

	// The type of the model. Examples: table, view, object. Default: table.
	Type Type `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

// The quality object contains quality attributes and checks.
type Quality struct {
	// Specification corresponds to the JSON schema field "specification".
	Specification interface{} `json:"specification" yaml:"specification" mapstructure:"specification"`

	// The type of the quality check. Typical values are SodaCL, montecarlo,
	// great-expectations, custom.
	Type QualityType `json:"type" yaml:"type" mapstructure:"type"`
}

type QualityType string

const QualityTypeCustom QualityType = "custom"
const QualityTypeGreatExpectations QualityType = "great-expectations"
const QualityTypeMontecarlo QualityType = "montecarlo"
const QualityTypeSodaCL QualityType = "SodaCL"

// UnmarshalJSON implements json.Unmarshaler.
func (j *QualityType) UnmarshalJSON(b []byte) error {
	var types = []string{
		"SodaCL",
		"montecarlo",
		"great-expectations",
		"custom",
	}

	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range types {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", types, v)
	}
	*j = QualityType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Quality) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["specification"]; raw != nil && !ok {
		return fmt.Errorf("field specification in Quality: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in Quality: required")
	}
	type Plain Quality
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Quality(plain)
	return nil
}

// The schema of the data contract describes the syntax and semantics of provided
// data sets. It supports different schema types.
type Schema struct {
	// Specification corresponds to the JSON schema field "specification".
	Specification interface{} `json:"specification" yaml:"specification" mapstructure:"specification"`

	// The type of the schema. Typical values are dbt, bigquery, json-schema, sql-ddl,
	// avro, protobuf, custom.
	Type SchemaType `json:"type" yaml:"type" mapstructure:"type"`
}

//nolint:revive // this is another kind of schema
type SchemaType string

const SchemaTypeAvro SchemaType = "avro"
const SchemaTypeBigquery SchemaType = "bigquery"
const SchemaTypeCustom SchemaType = "custom"
const SchemaTypeDbt SchemaType = "dbt"
const SchemaTypeJSONSchema SchemaType = "json-schema"
const SchemaTypeProtobuf SchemaType = "protobuf"
const SchemaTypeSQLDdl SchemaType = "sql-ddl"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SchemaType) UnmarshalJSON(b []byte) error {
	var types = []string{
		"dbt",
		"bigquery",
		"json-schema",
		"sql-ddl",
		"avro",
		"protobuf",
		"custom",
	}

	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range types {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", types, v)
	}
	*j = SchemaType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Schema) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["specification"]; raw != nil && !ok {
		return fmt.Errorf("field specification in Schema: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in Schema: required")
	}
	type Plain Schema
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Schema(plain)
	return nil
}

// Information about the servers.
type Servers struct {
	// An optional string describing the servers.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// The environment in which the servers are running. Examples: prod, sit, stg.
	Environment *string `json:"environment,omitempty" yaml:"environment,omitempty" mapstructure:"environment,omitempty"`

	AdditionalProperties interface{}
}

// Specifies the service level agreements for the provided data, including
// availability, data retention policies, latency requirements, data freshness,
// update frequency, support availability, and backup policies.
type Servicelevels struct {
	// Availability refers to the promise or guarantee by the service provider about
	// the uptime of the system that provides the data.
	Availability *ServicelevelsAvailability `json:"availability,omitempty" yaml:"availability,omitempty" mapstructure:"availability,omitempty"`

	// Backup specifies details about data backup procedures.
	Backup *ServicelevelsBackup `json:"backup,omitempty" yaml:"backup,omitempty" mapstructure:"backup,omitempty"`

	// Frequency describes how often data is updated.
	Frequency *ServicelevelsFrequency `json:"frequency,omitempty" yaml:"frequency,omitempty" mapstructure:"frequency,omitempty"`

	// The maximum age of the youngest row in a table.
	Freshness *ServicelevelsFreshness `json:"freshness,omitempty" yaml:"freshness,omitempty" mapstructure:"freshness,omitempty"`

	// Latency refers to the maximum amount of time from the source to its
	// destination.
	Latency *ServicelevelsLatency `json:"latency,omitempty" yaml:"latency,omitempty" mapstructure:"latency,omitempty"`

	// Retention covers the period how long data will be available.
	Retention *ServicelevelsRetention `json:"retention,omitempty" yaml:"retention,omitempty" mapstructure:"retention,omitempty"`

	// Support describes the times when support will be available for contact.
	Support *ServicelevelsSupport `json:"support,omitempty" yaml:"support,omitempty" mapstructure:"support,omitempty"`
}

// Availability refers to the promise or guarantee by the service provider about
// the uptime of the system that provides the data.
type ServicelevelsAvailability struct {
	// An optional string describing the availability service level.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// An optional string describing the guaranteed uptime in percent (e.g., `99.9%`)
	Percentage *string `json:"percentage,omitempty" yaml:"percentage,omitempty" mapstructure:"percentage,omitempty"`
}

// Backup specifies details about data backup procedures.
type ServicelevelsBackup struct {
	// An optional cron expression when data will be backed up, e.g., `0 0 * * *`.
	Cron *string `json:"cron,omitempty" yaml:"cron,omitempty" mapstructure:"cron,omitempty"`

	// An optional string describing the backup service level.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// An optional interval that defines how often data will be backed up, e.g.,
	// `daily`.
	Interval *string `json:"interval,omitempty" yaml:"interval,omitempty" mapstructure:"interval,omitempty"`

	// An optional Recovery Point Objective (RPO) defines the maximum acceptable age
	// of files that must be recovered from backup storage for normal operations to
	// resume after a disaster or data loss event. This essentially measures how much
	// data you can afford to lose, measured in time (e.g., 4 hours, 24 hours).
	RecoveryPoint *string `json:"recoveryPoint,omitempty" yaml:"recoveryPoint,omitempty" mapstructure:"recoveryPoint,omitempty"`

	// An optional Recovery Time Objective (RTO) specifies the maximum amount of time
	// allowed to restore data from a backup after a failure or loss event (e.g., 4
	// hours, 24 hours).
	RecoveryTime *string `json:"recoveryTime,omitempty" yaml:"recoveryTime,omitempty" mapstructure:"recoveryTime,omitempty"`
}

// Frequency describes how often data is updated.
type ServicelevelsFrequency struct {
	// Optional. Only for batch: A cron expression when the pipelines is triggered.
	// E.g., `0 0 * * *`.
	Cron *string `json:"cron,omitempty" yaml:"cron,omitempty" mapstructure:"cron,omitempty"`

	// An optional string describing the frequency service level.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// Optional. Only for batch: How often the pipeline is triggered, e.g., `daily`.
	Interval *string `json:"interval,omitempty" yaml:"interval,omitempty" mapstructure:"interval,omitempty"`

	// The method of data processing.
	Type *ServicelevelsFrequencyType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

type ServicelevelsFrequencyType string

const ServicelevelsFrequencyTypeBatch ServicelevelsFrequencyType = "batch"
const ServicelevelsFrequencyTypeManual ServicelevelsFrequencyType = "manual"
const ServicelevelsFrequencyTypeMicroBatching ServicelevelsFrequencyType = "micro-batching"
const ServicelevelsFrequencyTypeStreaming ServicelevelsFrequencyType = "streaming"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ServicelevelsFrequencyType) UnmarshalJSON(b []byte) error {
	var servicelevelsFrequencyType = []string{
		"batch",
		"micro-batching",
		"streaming",
		"manual",
	}

	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range servicelevelsFrequencyType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", servicelevelsFrequencyType, v)
	}
	*j = ServicelevelsFrequencyType(v)
	return nil
}

// The maximum age of the youngest row in a table.
type ServicelevelsFreshness struct {
	// An optional string describing the freshness service level.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// An optional maximum age of the youngest entry. Supported formats: Simple
	// duration (e.g., `24 hours`, `5s`) and ISO 8601 duration (e.g., `PT24H`).
	Threshold *string `json:"threshold,omitempty" yaml:"threshold,omitempty" mapstructure:"threshold,omitempty"`

	// An optional reference to the field that contains the timestamp that the
	// threshold refers to.
	TimestampField *string `json:"timestampField,omitempty" yaml:"timestampField,omitempty" mapstructure:"timestampField,omitempty"`
}

// Latency refers to the maximum amount of time from the source to its destination.
type ServicelevelsLatency struct {
	// An optional string describing the latency service level.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// An optional reference to the field that contains the processing timestamp,
	// which denotes when the data is made available to consumers of this data
	// contract.
	ProcessedTimestampField *string `json:"processedTimestampField,omitempty" yaml:"processedTimestampField,omitempty" mapstructure:"processedTimestampField,omitempty"`

	// An optional reference to the field that contains the timestamp when the data
	// was provided at the source.
	SourceTimestampField *string `json:"sourceTimestampField,omitempty" yaml:"sourceTimestampField,omitempty" mapstructure:"sourceTimestampField,omitempty"`

	// An optional maximum duration between the source timestamp and the processed
	// timestamp. Supported formats: Simple duration (e.g., `24 hours`, `5s`) and ISO
	// 8601 duration (e.g, `PT24H`).
	Threshold *string `json:"threshold,omitempty" yaml:"threshold,omitempty" mapstructure:"threshold,omitempty"`
}

// Retention covers the period how long data will be available.
type ServicelevelsRetention struct {
	// An optional string describing the retention service level.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// An optional period of time, how long data is available. Supported formats:
	// Simple duration (e.g., `1 year`, `30d`) and ISO 8601 duration (e.g, `P1Y`).
	Period *string `json:"period,omitempty" yaml:"period,omitempty" mapstructure:"period,omitempty"`

	// An optional reference to the field that contains the timestamp that the period
	// refers to.
	TimestampField *string `json:"timestampField,omitempty" yaml:"timestampField,omitempty" mapstructure:"timestampField,omitempty"`

	// An optional indicator that data is kept forever.
	Unlimited *bool `json:"unlimited,omitempty" yaml:"unlimited,omitempty" mapstructure:"unlimited,omitempty"`
}

// Support describes the times when support will be available for contact.
type ServicelevelsSupport struct {
	// An optional string describing the support service level.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// An optional string describing the time it takes for the support team to
	// acknowledge a request. This does not mean the issue will be resolved
	// immediately, but it assures users that their request has been received and will
	// be dealt with.
	ResponseTime *string `json:"responseTime,omitempty" yaml:"responseTime,omitempty" mapstructure:"responseTime,omitempty"`

	// An optional string describing the times when support will be available for
	// contact such as `24/7` or `business hours only`.
	Time *string `json:"time,omitempty" yaml:"time,omitempty" mapstructure:"time,omitempty"`
}

// The terms and conditions of the data contract.
type Terms struct {
	// The billing describes the pricing model for using the data, such as whether
	// it's free, having a monthly fee, or metered pay-per-use.
	Billing *string `json:"billing,omitempty" yaml:"billing,omitempty" mapstructure:"billing,omitempty"`

	// The limitations describe the restrictions on how the data can be used, can be
	// technical or restrictions on what the data may not be used for.
	Limitations *string `json:"limitations,omitempty" yaml:"limitations,omitempty" mapstructure:"limitations,omitempty"`

	// The period of time that must be given by either party to terminate or modify a
	// data usage agreement. Uses ISO-8601 period format, e.g., 'P3M' for a period of
	// three months.
	NoticePeriod *string `json:"noticePeriod,omitempty" yaml:"noticePeriod,omitempty" mapstructure:"noticePeriod,omitempty"`

	// The usage describes the way the data is expected to be used. Can contain
	// business and technical information.
	Usage *string `json:"usage,omitempty" yaml:"usage,omitempty" mapstructure:"usage,omitempty"`

	AdditionalProperties interface{}
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Datacontract) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	if _, ok := raw["dataContractSpecification"]; raw != nil && !ok {
		return fmt.Errorf("field dataContractSpecification in Datacontract: required")
	}

	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in Datacontract: required")
	}

	if _, ok := raw["info"]; raw != nil && !ok {
		return fmt.Errorf("field info in Datacontract: required")
	}

	type Plain Datacontract
	var plain Plain

	//nolint:musttag // false positive
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}

	*j = Datacontract(plain)

	return nil
}

func NewFieldTypeReference(f FieldType) *FieldType {
	return &f
}

type FieldType string

const FieldTypeArray FieldType = "array"
const FieldTypeBigint FieldType = "bigint"
const FieldTypeBoolean FieldType = "boolean"
const FieldTypeBytes FieldType = "bytes"
const FieldTypeDate FieldType = "date"
const FieldTypeDecimal FieldType = "decimal"
const FieldTypeDouble FieldType = "double"
const FieldTypeFloat FieldType = "float"
const FieldTypeInt FieldType = "int"
const FieldTypeInteger FieldType = "integer"
const FieldTypeLong FieldType = "long"
const FieldTypeMap FieldType = "map"
const FieldTypeNull FieldType = "null"
const FieldTypeNumber FieldType = "number"
const FieldTypeNumeric FieldType = "numeric"
const FieldTypeObject FieldType = "object"
const FieldTypeRecord FieldType = "record"
const FieldTypeString FieldType = "string"
const FieldTypeStruct FieldType = "struct"
const FieldTypeText FieldType = "text"
const FieldTypeTimestamp FieldType = "timestamp"
const FieldTypeTimestampNtz FieldType = "timestamp_ntz"
const FieldTypeTimestampTz FieldType = "timestamp_tz"
const FieldTypeVarchar FieldType = "varchar"

// UnmarshalJSON implements json.Unmarshaler.
func (j *FieldType) UnmarshalJSON(b []byte) error {
	var types = []string{
		"number",
		"decimal",
		"numeric",
		"int",
		"integer",
		"long",
		"bigint",
		"float",
		"double",
		"string",
		"text",
		"varchar",
		"boolean",
		"timestamp",
		"timestamp_tz",
		"timestamp_ntz",
		"date",
		"array",
		"map",
		"object",
		"record",
		"struct",
		"bytes",
		"null",
	}

	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range types {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", types, v)
	}
	*j = FieldType(v)
	return nil
}

// The nested fields (e.g. columns) of the object, record, or struct.
type Fields map[string]AdditionalProperties

// Links to external resources.
type Links map[string]string

type Type string

func NewTypeReference(t Type) *Type {
	return &t
}

const TypeObject Type = "object"
const TypeTable Type = "table"
const TypeView Type = "view"

// UnmarshalJSON implements json.Unmarshaler.
func (j *Type) UnmarshalJSON(b []byte) error {
	var types = []string{
		"table",
		"view",
		"object",
	}

	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range types {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", types, v)
	}
	*j = Type(v)
	return nil
}
