package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	bq "cloud.google.com/go/bigquery"
	"github.com/merlindorin/cli-datacontract/bigquery"
	"github.com/merlindorin/cli-datacontract/schema"
	"github.com/merlindorin/go-shared/pkg/cmd"
	"google.golang.org/api/iterator"
	"gopkg.in/yaml.v3"
)

type BigqueryCMD struct {
	All    BigqueryAllCMD    `cmd:"remote" help:"import datacontracts from all datasets and dataname in project"`
	Single BigquerySingleCMD `cmd:"single" help:"import datacontract from a single remote tablename"`
	File   BigqueryFileCMD   `cmd:"file" help:"import datacontract from a file"`

	Format string `name:"format" default:"yaml" short:"f" enum:"json,yaml" help:"Specify the output format"`

	DatacontractID               string `name:"datacontract-id" help:"Unique identifier for the datacontract"`
	DataContractSpecification    string `name:"datacontract-specification" default:"0.9.3" enum:"0.9.0,0.9.1,0.9.2,0.9.3" help:"Version of the datacontract specification"`
	DatacontractInfoTitle        string `name:"datacontract-info-title" help:"Title of the BigQuery table for datacontract information"`
	DatacontractInfoVersion      string `name:"datacontract-info-version" help:"Version of the BigQuery table for datacontract information"`
	DatacontractModelName        string `name:"datacontract-model-name" help:"Name of the BigQuery table for the datacontract model"`
	DatacontractModelDescription string `name:"datacontract-model-description" help:"Description for the BigQuery table of the datacontract model"`
}

type BigqueryAllCMD struct {
	Directory string `arg:"directory" optional:"" short:"d" help:"Output directory; defaults to the stdout"`
	ProjectID string `name:"bigquery-projectid" required:"" help:"Unique identifier for the BigQuery project"`
}

func (i *BigqueryAllCMD) Run(parent *BigqueryCMD, co *cmd.Commons) error {
	ctx := context.Background()
	p := co.Printer()

	c, err := bq.NewClient(ctx, i.ProjectID)
	if err != nil {
		return fmt.Errorf("cannot create bigquery client: %w", err)
	}

	datasetsIterator := c.Datasets(ctx)
	for {
		ds, derr := datasetsIterator.Next()
		if errors.Is(derr, iterator.Done) {
			break
		}

		if derr != nil {
			return fmt.Errorf("error iterating datasets: %w", derr)
		}

		tablesIterator := ds.Tables(ctx)
		for {
			table, terr := tablesIterator.Next()
			if errors.Is(terr, iterator.Done) {
				break
			}

			if terr != nil {
				return fmt.Errorf("error iterating tables: %w", terr)
			}

			out := os.Stdout

			if i.Directory != "" {
				name := fmt.Sprintf("%s-%s.%s", table.DatasetID, table.TableID, "json")
				filename := filepath.Join(filepath.Clean(i.Directory), name)

				out, terr = os.Create(filename)
				if terr != nil {
					return fmt.Errorf("cannot create output file: %w", terr)
				}
				defer func() {
					terr = errors.Join(terr, out.Close())
				}()
			}

			metadata, terr := table.Metadata(ctx)
			if terr != nil {
				return fmt.Errorf("cannot get table metadata: %w", terr)
			}

			tableName := metadata.FullID
			tableDescription := metadata.Description
			tableType := metadata.Type
			tableSchema := metadata.Schema

			terr = process(parent, out, tableName, tableDescription, tableType, tableSchema)
			if terr != nil {
				return fmt.Errorf("cannot import datacontract from remote: %w", terr)
			}

			p.Printf(
				"Table collected: projectID=%s, datasetID=%s, tableID=%s, file=%s\n",
				table.ProjectID,
				table.DatasetID,
				table.TableID,
				out.Name(),
			)
		}
	}

	return nil
}

type BigquerySingleCMD struct {
	Out string `name:"out" short:"o" help:"Output filename; defaults to stdout if not specified"`

	ProjectID string `name:"bigquery-projectid" required:"" help:"Unique identifier for the BigQuery project"`
	DatasetID string `name:"bigquery-datasetid" required:"" help:"Unique identifier for the BigQuery dataset"`
	TableName string `name:"bigquery-tablename" required:"" help:"Unique identifier for the BigQuery table"`
}

func (i *BigquerySingleCMD) Run(parent *BigqueryCMD, co *cmd.Commons) error {
	ctx := context.Background()
	p := co.Printer()

	c, err := bq.NewClient(ctx, i.ProjectID)
	if err != nil {
		return fmt.Errorf("cannot create bigquery client: %w", err)
	}

	metadata, err := c.Dataset(i.DatasetID).Table(i.TableName).Metadata(ctx)
	if err != nil {
		return fmt.Errorf("cannot get table metadata: %w", err)
	}

	tableName := metadata.Name
	tableDescription := metadata.Description
	tableType := metadata.Type
	tableSchema := metadata.Schema

	out := os.Stdout

	if i.Out != "" {
		out, err = os.Create(filepath.Clean(i.Out))
		if err != nil {
			return fmt.Errorf("cannot create output file: %w", err)
		}
		defer func() {
			err = errors.Join(err, out.Close())
		}()
	}

	err = process(parent, out, tableName, tableDescription, tableType, tableSchema)
	if err != nil {
		return fmt.Errorf("cannot import datacontract from remote: %w", err)
	}

	p.Printf(
		"Table collected: projectID=%s, datasetID=%s, tableID=%s, file=%s\n",
		i.ProjectID,
		i.DatasetID,
		tableName,
		out.Name(),
	)

	return nil
}

type BigqueryFileCMD struct {
	Out        string `name:"out" short:"o" help:"Output filename; defaults to stdout if not specified"`
	SchemaOnly bool   `name:"bigquery-schema-only" help:"Bigquery schema only"`

	Filename string `name:"bigquery-filename" arg:"" required:"" help:"Bigquery filename"`
}

func (i *BigqueryFileCMD) Run(parent *BigqueryCMD, co *cmd.Commons) error {
	f := BigqueryFile{}
	p := co.Printer()

	file, err := os.Open(i.Filename)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}

	if i.SchemaOnly {
		err = json.NewDecoder(file).Decode(&f.Schema.Fields)
		f.Type = bq.ViewTable
	} else {
		err = json.NewDecoder(file).Decode(&f)
	}

	if err != nil {
		return fmt.Errorf("cannot parse file: %w", err)
	}

	tableName := f.TableReference.TableID
	tableType := f.Type
	tableSchema := f.Schema.Fields

	out := os.Stdout

	if i.Out != "" {
		out, err = os.Create(filepath.Clean(i.Out))
		if err != nil {
			return fmt.Errorf("cannot create output file: %w", err)
		}
		defer func() {
			err = errors.Join(err, out.Close())
		}()
	}

	err = process(parent, out, tableName, "", tableType, tableSchema)
	if err != nil {
		return fmt.Errorf("cannot process bigquery file: %w", err)
	}

	p.Printf(
		"Table collected: tableID=%s, file=%s\n",
		tableName,
		out.Name(),
	)

	return nil
}

func process(
	parent *BigqueryCMD,
	out io.Writer,
	tableName string,
	tableDescription string,
	tableType bq.TableType,
	tableSchema bq.Schema,
) error {
	var datacontractSpecification schema.DataContractSpecification
	err := datacontractSpecification.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", parent.DataContractSpecification)))
	if err != nil {
		return fmt.Errorf("error unmarshalling BigQuery Specification version: %w", err)
	}

	d := schema.Datacontract{
		ID:            parent.DatacontractID,
		Specification: datacontractSpecification,
		Info: schema.Info{
			Title:   parent.DatacontractInfoTitle,
			Version: parent.DatacontractInfoVersion,
		},
	}

	if parent.DatacontractModelName != "" {
		tableName = parent.DatacontractModelName
	}

	if parent.DatacontractModelDescription != "" {
		tableDescription = parent.DatacontractModelDescription
	}

	err = bigquery.ImportSchema(tableName, tableDescription, tableType, tableSchema, &d)
	if err != nil {
		return fmt.Errorf("cannot import: %w", err)
	}

	var indent []byte

	if parent.Format == "json" {
		indent, err = json.MarshalIndent(d, "", "    ") //nolint:musttag // no idea why
		if err != nil {
			return fmt.Errorf("cannot marshal json: %w", err)
		}
	}

	if parent.Format == "yaml" {
		indent, err = yaml.Marshal(d) //nolint:musttag // no idea why
		if err != nil {
			return fmt.Errorf("cannot marshal json: %w", err)
		}
	}

	_, err = io.Copy(out, bytes.NewReader(indent))
	if err != nil {
		return fmt.Errorf("cannot write to output: %w", err)
	}

	return nil
}

type BigqueryFile struct {
	Kind           string `json:"kind"`
	Etag           string `json:"etag"`
	ID             string `json:"id"`
	SelfLink       string `json:"selfLink"`
	TableReference struct {
		ProjectID string `json:"projectId"`
		DatasetID string `json:"datasetId"`
		TableID   string `json:"tableId"`
	} `json:"tableReference"`
	Schema struct {
		Fields []*bq.FieldSchema `json:"fields"`
	} `json:"schema"`
	NumBytes                   string       `json:"numBytes"`
	NumLongTermBytes           string       `json:"numLongTermBytes"`
	NumRows                    string       `json:"numRows"`
	CreationTime               string       `json:"creationTime"`
	LastModifiedTime           string       `json:"lastModifiedTime"`
	Type                       bq.TableType `json:"type"`
	Location                   string       `json:"location"`
	NumTimeTravelPhysicalBytes string       `json:"numTimeTravelPhysicalBytes"`
	NumTotalLogicalBytes       string       `json:"numTotalLogicalBytes"`
	NumActiveLogicalBytes      string       `json:"numActiveLogicalBytes"`
	NumLongTermLogicalBytes    string       `json:"numLongTermLogicalBytes"`
	NumTotalPhysicalBytes      string       `json:"numTotalPhysicalBytes"`
	NumActivePhysicalBytes     string       `json:"numActivePhysicalBytes"`
	NumLongTermPhysicalBytes   string       `json:"numLongTermPhysicalBytes"`
	NumCurrentPhysicalBytes    string       `json:"numCurrentPhysicalBytes"`
}
