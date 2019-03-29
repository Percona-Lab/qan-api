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
	"github.com/kylelemons/godebug/pretty"
	"github.com/percona/pmm/api/qanpb"

	"github.com/percona/qan-api2/models"
)

const expectedDataFile = "../../test_data/profile.json"

func TestService_GetReport(t *testing.T) {
	dsn, ok := os.LookupEnv("QANAPI_DSN_TEST")
	if !ok {
		dsn = "clickhouse://127.0.0.1:19000?database=pmm_test"
	}
	db, err := sqlx.Connect("clickhouse", dsn)
	if err != nil {
		log.Fatal("Connection: ", err)
	}

	rm := models.NewReporter(db)
	mm := models.NewMetrics(db)
	t1, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2019-01-01T10:00:00Z")
	var want qanpb.ReportReply
	expectedData, err := ioutil.ReadFile(expectedDataFile)

	fmt.Println("expectedData", string(expectedData))
	if err != nil {
		log.Fatal("read file with expected filtering data: ", err)
	}
	err = json.Unmarshal(expectedData, &want)
	if err != nil {
		log.Fatal("cannot unmarshal expected result: ", err)
	}
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
				},
			},
			&want,
			false,
		},
		{
			"wrong time range",
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
			"empty fail",
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
			gotJSON, _ := json.MarshalIndent(got, "", "\t")
			fmt.Println("expectedData", string(gotJSON))

			if diff := pretty.Compare(got, *tt.want); diff != "" {
				t.Errorf("%s: Service.GetReport() = diff: (-got +want)\n%s", tt.name, diff)
			}
		})
	}
}
