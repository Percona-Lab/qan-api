// qan-api2
// Copyright (C) 2019 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package analitycs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jmoiron/sqlx"
	"github.com/percona/pmm/api/qanpb"

	"github.com/percona/qan-api2/models"
)

func setup() *sqlx.DB {
	dsn, ok := os.LookupEnv("QANAPI_DSN_TEST")
	if !ok {
		dsn = "clickhouse://127.0.0.1:19000?database=pmm_test&debug=true"
	}
	db, err := sqlx.Connect("clickhouse", dsn)
	if err != nil {
		log.Fatal("Connection: ", err)
	}

	return db
}

func expectedData(t *testing.T, got, want interface{}, filename string) {
	if os.Getenv("REFRESH_TEST_DATA") != "" {
		json, err := json.MarshalIndent(got, "", "\t")
		if err != nil {
			t.Errorf("cannot marshal:%v", err)
		}
		err = ioutil.WriteFile(filename, json, 0644)
		if err != nil {
			t.Errorf("cannot write:%v", err)
		}
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("cannot read data from file:%v", err)
	}

	err = json.Unmarshal(data, want)
	if err != nil {
		t.Errorf("cannot read data from file:%v", err)
	}
}

func getExpectedJSON(t *testing.T, got interface{}, filename string) []byte {
	if os.Getenv("REFRESH_TEST_DATA") != "" {
		json, err := json.MarshalIndent(got, "", "\t")
		if err != nil {
			t.Errorf("cannot marshal:%v", err)
		}
		err = ioutil.WriteFile(filename, json, 0644)
		if err != nil {
			t.Errorf("cannot write:%v", err)
		}
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("cannot read data from file:%v", err)
	}
	return data
}

func TestService_GetReport(t *testing.T) {
	db := setup()
	rm := models.NewReporter(db)
	mm := models.NewMetrics(db)
	t1, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2019-01-01T10:00:00Z")
	var want qanpb.ReportReply
	type fields struct {
		rm models.Reporter
		mm models.Metrics
	}
	type args struct {
		ctx context.Context
		in  *qanpb.ReportRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *qanpb.ReportReply
		wantErr bool
	}{
		{
			"success",
			fields{rm: rm, mm: mm},
			args{
				context.TODO(),
				&qanpb.ReportRequest{
					PeriodStartFrom: &timestamp.Timestamp{Seconds: t1.Unix()},
					PeriodStartTo:   &timestamp.Timestamp{Seconds: t2.Unix()},
					GroupBy:         "queryid",
					Columns:         []string{"lock_time", "sort_scan"},
					Offset:          0,
					Limit:           10,
				},
			},
			&want,
			false,
		},
		{
			"wrong_time_range",
			fields{rm: rm, mm: mm},
			args{
				context.TODO(),
				&qanpb.ReportRequest{
					PeriodStartFrom: &timestamp.Timestamp{Seconds: t2.Unix()},
					PeriodStartTo:   &timestamp.Timestamp{Seconds: t1.Unix()},
				},
			},
			&qanpb.ReportReply{},
			true,
		},
		{
			"empty_fail",
			fields{rm: rm, mm: mm},
			args{
				context.TODO(),
				&qanpb.ReportRequest{},
			},
			&qanpb.ReportReply{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				rm: tt.fields.rm,
				mm: tt.fields.mm,
			}
			got, err := s.GetReport(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.want = nil
			expectedData(t, got, &tt.want, "../../test_data/TestService_GetReport_"+tt.name+".json")
			// TODO: why travis-ci return other values then expected?
			if got.TotalRows != tt.want.TotalRows {
				t.Errorf("got.TotalRows (%v) != *tt.want.TotalRows (%v)", got.TotalRows, tt.want.TotalRows)
			}

			for i, v := range got.Rows {
				if v.NumQueries != tt.want.Rows[i].NumQueries {
					t.Errorf("got.Rows[0].NumQueries (%v) != *tt.want.Rows[0].NumQueries (%v)", v.NumQueries, tt.want.Rows[i].NumQueries)
				}
			}
		})
	}
}

func TestService_GetReport_Mix(t *testing.T) {
	db := setup()
	rm := models.NewReporter(db)
	mm := models.NewMetrics(db)
	t1, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2019-01-01T10:00:00Z")
	var want qanpb.ReportReply
	type fields struct {
		rm models.Reporter
		mm models.Metrics
	}
	type args struct {
		ctx context.Context
		in  *qanpb.ReportRequest
	}
	test := struct {
		name    string
		fields  fields
		args    args
		want    *qanpb.ReportReply
		wantErr bool
	}{
		"reverce_order",
		fields{rm: rm, mm: mm},
		args{
			context.TODO(),
			&qanpb.ReportRequest{
				PeriodStartFrom: &timestamp.Timestamp{Seconds: t1.Unix()},
				PeriodStartTo:   &timestamp.Timestamp{Seconds: t2.Unix()},
				GroupBy:         "queryid",
				Columns:         []string{"lock_time", "sort_scan"},
				OrderBy:         "-load",
				Offset:          10,
				Limit:           10,
				Labels: []*qanpb.ReportMapFieldEntry{
					{
						Key:   "label1",
						Value: []string{"value1", "value2"},
					},
					{
						Key:   "d_server",
						Value: []string{"db1", "db2", "db3", "db4", "db5", "db6", "db7"},
					},
				},
			},
		},
		&want,
		false,
	}
	t.Run(test.name, func(t *testing.T) {
		s := &Service{
			rm: test.fields.rm,
			mm: test.fields.mm,
		}
		got, err := s.GetReport(test.args.ctx, test.args.in)
		if (err != nil) != test.wantErr {
			t.Errorf("Service.GetReport() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		test.want = nil
		expectedData(t, got, &test.want, "../../test_data/TestService_GetReport_Mix_"+test.name+".json")

		for i, v := range got.Rows {
			if v.NumQueries != test.want.Rows[i].NumQueries {
				t.Errorf("got.Rows[%d].NumQueries (%v) != *tt.want.Rows[%d].NumQueries (%v)", i, v.NumQueries, i, test.want.Rows[i].NumQueries)
			}
		}
	})

	test.name = "correct_load"
	t.Run(test.name, func(t *testing.T) {
		s := &Service{
			rm: test.fields.rm,
			mm: test.fields.mm,
		}
		got, err := s.GetReport(test.args.ctx, test.args.in)
		if (err != nil) != test.wantErr {
			t.Errorf("Service.GetReport() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		test.want = nil
		expectedData(t, got, &test.want, "../../test_data/TestService_GetReport_Mix_"+test.name+".json")

		for i, v := range got.Rows {
			if v.Load != test.want.Rows[i].Load {
				t.Errorf("got.Rows[%d].Load (%v) != *tt.want.Rows[%d].Load (%v)", i, v.NumQueries, i, test.want.Rows[i].NumQueries)
			}
		}
	})

	test.name = "correct_latency"
	t.Run(test.name, func(t *testing.T) {
		s := &Service{
			rm: test.fields.rm,
			mm: test.fields.mm,
		}
		got, err := s.GetReport(test.args.ctx, test.args.in)
		if (err != nil) != test.wantErr {
			t.Errorf("Service.GetReport() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		test.want = nil
		expectedData(t, got, &test.want, "../../test_data/TestService_GetReport_Mix_"+test.name+".json")

		for i, v := range got.Rows {
			if v.Metrics["latency"].Stats.Sum != test.want.Rows[i].Metrics["latency"].Stats.Sum {
				t.Errorf(
					"got.Rows[%d].Metrics[latency].Stats.Sum (%v) != *tt.want.Rows[%d].Metrics[latency].Stats.Sum (%v)",
					i,
					v.Metrics["latency"].Stats.Sum,
					i,
					test.want.Rows[i].Metrics["latency"].Stats.Sum,
				)
			}
		}
	})

	t.Run("no error on limit is 0", func(t *testing.T) {
		s := &Service{
			rm: test.fields.rm,
			mm: test.fields.mm,
		}

		test.args.in.Limit = 0
		_, err := s.GetReport(test.args.ctx, test.args.in)
		if err != nil {
			t.Errorf("Service.GetReport() error = %v, wantErr %v", err, test.wantErr)
			return
		}
	})

	t.Run("Limit is 0", func(t *testing.T) {
		s := &Service{
			rm: test.fields.rm,
			mm: test.fields.mm,
		}

		test.args.in.GroupBy = "unknown dimension"
		expectedErr := fmt.Errorf("unknown group dimension: %s", "unknown dimension")
		_, err := s.GetReport(test.args.ctx, test.args.in)
		if err.Error() != expectedErr.Error() {
			t.Errorf("Service.GetReport() unexpected error = %v, wantErr %v", err, expectedErr)
			return
		}
	})
}

func TestService_GetReport_AllLabels(t *testing.T) {
	db := setup()
	rm := models.NewReporter(db)
	mm := models.NewMetrics(db)
	t1, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2019-01-01T10:00:00Z")
	type fields struct {
		rm models.Reporter
		mm models.Metrics
	}
	type args struct {
		ctx context.Context
		in  *qanpb.ReportRequest
	}

	genDimensionvalues := func(dimKey string, amount int) []string {
		arr := []string{}
		for i := 0; i < amount; i++ {
			arr = append(arr, fmt.Sprintf("%s%d", dimKey, i))
		}
		return arr
	}
	test := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		"",
		fields{rm: rm, mm: mm},
		args{
			context.TODO(),
			&qanpb.ReportRequest{
				PeriodStartFrom: &timestamp.Timestamp{Seconds: t1.Unix()},
				PeriodStartTo:   &timestamp.Timestamp{Seconds: t2.Unix()},
				GroupBy:         "queryid",
				Columns:         []string{"lock_time", "sort_scan"},
				OrderBy:         "-load",
				Offset:          10,
				Limit:           10,
				Labels: []*qanpb.ReportMapFieldEntry{
					{
						Key:   "label1",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label2",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label3",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label4",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label5",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label6",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label7",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label8",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "label9",
						Value: genDimensionvalues("value", 100),
					},
					{
						Key:   "d_server",
						Value: genDimensionvalues("db", 10),
					},
					{
						Key:   "d_database",
						Value: genDimensionvalues("schema", 100),
					},
					{
						Key:   "d_schema",
						Value: []string{},
					},
					{
						Key:   "d_username",
						Value: genDimensionvalues("user", 100),
					},
					{
						Key:   "d_client_host",
						Value: genDimensionvalues("10.11.12.", 100),
					},
				},
			},
		},
		false,
	}
	t.Run("Use all label keys", func(t *testing.T) {
		s := &Service{
			rm: test.fields.rm,
			mm: test.fields.mm,
		}
		got, err := s.GetReport(test.args.ctx, test.args.in)
		if (err != nil) != test.wantErr {
			t.Errorf("Service.GetReport() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		expectedRows := 1
		gotRows := len(got.Rows)
		if gotRows != expectedRows {
			t.Errorf("Got rows count: %d - expected, %d", gotRows, expectedRows)
		}
	})

}
