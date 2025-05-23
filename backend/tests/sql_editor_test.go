package tests

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"

	resourcemysql "github.com/bytebase/bytebase/backend/resources/mysql"
	"github.com/bytebase/bytebase/backend/resources/postgres"
	"github.com/bytebase/bytebase/backend/tests/fake"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
	v1pb "github.com/bytebase/bytebase/proto/generated-go/v1"
)

func TestAdminQueryAffectedRows(t *testing.T) {
	tests := []struct {
		databaseName      string
		dbType            storepb.Engine
		prepareStatements string
		query             string
		want              bool
		affectedRows      []*v1pb.QueryResult
	}{
		{
			databaseName:      "Test1",
			dbType:            storepb.Engine_MYSQL,
			prepareStatements: "CREATE TABLE tbl(id INT PRIMARY KEY);",
			query:             "INSERT INTO tbl VALUES(1);",
			affectedRows: []*v1pb.QueryResult{
				{
					ColumnNames:     []string{"Affected Rows"},
					ColumnTypeNames: []string{"INT"},
					Rows: []*v1pb.QueryRow{
						{
							Values: []*v1pb.RowValue{
								{Kind: &v1pb.RowValue_Int64Value{Int64Value: 1}},
							},
						},
					},
					Statement: "INSERT INTO tbl VALUES(1);",
					RowsCount: 1,
				},
			},
		},
		{
			databaseName:      "Test2",
			dbType:            storepb.Engine_MYSQL,
			prepareStatements: "CREATE TABLE tbl(id INT PRIMARY KEY);",
			query:             "INSERT INTO tbl VALUES(1); DELETE FROM tbl WHERE id = 1;",
			affectedRows: []*v1pb.QueryResult{
				{
					ColumnNames:     []string{"Affected Rows"},
					ColumnTypeNames: []string{"INT"},
					Rows: []*v1pb.QueryRow{
						{
							Values: []*v1pb.RowValue{
								{Kind: &v1pb.RowValue_Int64Value{Int64Value: 1}},
							},
						},
					},
					Statement: "INSERT INTO tbl VALUES(1);",
					RowsCount: 1,
				},
				{
					ColumnNames:     []string{"Affected Rows"},
					ColumnTypeNames: []string{"INT"},
					Rows: []*v1pb.QueryRow{
						{
							Values: []*v1pb.RowValue{
								{Kind: &v1pb.RowValue_Int64Value{Int64Value: 1}},
							},
						},
					},
					Statement: " DELETE FROM tbl WHERE id = 1;",
					RowsCount: 1,
				},
			},
		},
		{
			databaseName:      "Test3",
			dbType:            storepb.Engine_POSTGRES,
			prepareStatements: "CREATE TABLE public.tbl(id INT PRIMARY KEY);",
			query:             "INSERT INTO tbl VALUES(1),(2);",
			affectedRows: []*v1pb.QueryResult{
				{
					ColumnNames:     []string{"Affected Rows"},
					ColumnTypeNames: []string{"INT"},
					Rows: []*v1pb.QueryRow{
						{
							Values: []*v1pb.RowValue{
								{Kind: &v1pb.RowValue_Int64Value{Int64Value: 2}},
							},
						},
					},
					Statement: "INSERT INTO tbl VALUES(1),(2);",
					RowsCount: 1,
				},
			},
		},
		{
			databaseName:      "Test4",
			dbType:            storepb.Engine_POSTGRES,
			prepareStatements: "CREATE TABLE tbl(id INT PRIMARY KEY);",
			query:             "ALTER TABLE tbl ADD COLUMN name VARCHAR(255);",
			affectedRows: []*v1pb.QueryResult{
				{
					ColumnNames:     []string{"Affected Rows"},
					ColumnTypeNames: []string{"INT"},
					Rows: []*v1pb.QueryRow{
						{
							Values: []*v1pb.RowValue{
								{Kind: &v1pb.RowValue_Int64Value{Int64Value: 0}},
							},
						},
					},
					Statement: "ALTER TABLE tbl ADD COLUMN name VARCHAR(255);",
					RowsCount: 1,
				},
			},
		},
	}

	t.Parallel()
	a := require.New(t)
	ctx := context.Background()
	ctl := &controller{}
	dataDir := t.TempDir()
	ctx, err := ctl.StartServerWithExternalPg(ctx, &config{
		dataDir:            dataDir,
		vcsProviderCreator: fake.NewGitLab,
	})
	a.NoError(err)
	defer ctl.Close(ctx)

	mysqlPort := getTestPort()
	mysqlStopInstance := resourcemysql.SetupTestInstance(t, mysqlPort, mysqlBinDir)
	defer mysqlStopInstance()

	// Create a PostgreSQL instance.
	pgPort := getTestPort()
	stopInstance := postgres.SetupTestInstance(pgBinDir, t.TempDir(), pgPort)
	defer stopInstance()

	mysqlInstance, err := ctl.instanceServiceClient.CreateInstance(ctx, &v1pb.CreateInstanceRequest{
		InstanceId: generateRandomString("instance", 10),
		Instance: &v1pb.Instance{
			Title:       "mysqlInstance",
			Engine:      v1pb.Engine_MYSQL,
			Environment: "environments/prod",
			Activation:  true,
			DataSources: []*v1pb.DataSource{{Type: v1pb.DataSourceType_ADMIN, Host: "127.0.0.1", Port: strconv.Itoa(mysqlPort), Username: "root", Password: "", Id: "admin"}},
		},
	})
	a.NoError(err)

	pgInstance, err := ctl.instanceServiceClient.CreateInstance(ctx, &v1pb.CreateInstanceRequest{
		InstanceId: generateRandomString("instance", 10),
		Instance: &v1pb.Instance{
			Title:       "pgInstance",
			Engine:      v1pb.Engine_POSTGRES,
			Environment: "environments/prod",
			Activation:  true,
			DataSources: []*v1pb.DataSource{{Type: v1pb.DataSourceType_ADMIN, Host: "/tmp", Port: strconv.Itoa(pgPort), Username: "root", Id: "admin"}},
		},
	})
	a.NoError(err)

	for _, tt := range tests {
		var instance *v1pb.Instance
		databaseOwner := ""
		switch tt.dbType {
		case storepb.Engine_MYSQL:
			instance = mysqlInstance
		case storepb.Engine_POSTGRES:
			instance = pgInstance
			databaseOwner = "root"
		default:
			a.FailNow("unsupported db type")
		}
		err = ctl.createDatabaseV2(ctx, ctl.project, instance, nil /* environment */, tt.databaseName, databaseOwner, nil)
		a.NoError(err)

		database, err := ctl.databaseServiceClient.GetDatabase(ctx, &v1pb.GetDatabaseRequest{
			Name: fmt.Sprintf("%s/databases/%s", instance.Name, tt.databaseName),
		})
		a.NoError(err)

		sheet, err := ctl.sheetServiceClient.CreateSheet(ctx, &v1pb.CreateSheetRequest{
			Parent: ctl.project.Name,
			Sheet: &v1pb.Sheet{
				Title:   "prepareStatements",
				Content: []byte(tt.prepareStatements),
			},
		})
		a.NoError(err)

		err = ctl.changeDatabase(ctx, ctl.project, database, sheet, v1pb.Plan_ChangeDatabaseConfig_MIGRATE)
		a.NoError(err)

		results, err := ctl.adminQuery(ctx, database, tt.query)
		a.NoError(err)

		a.Equal(len(tt.affectedRows), len(results))
		for idx, result := range results {
			a.Equal("", result.Error)
			result.Latency = nil
			diff := cmp.Diff(tt.affectedRows[idx], result, protocmp.Transform(), protocmp.IgnoreMessages(&durationpb.Duration{}))
			a.Empty(diff)
		}
	}
}
