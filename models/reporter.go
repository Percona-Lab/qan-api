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

package models

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"text/template"
	"time"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/percona/pmm/api/qanpb"
)

// Reporter implements models to select metrics bucket by params.
type Reporter struct {
	db *sqlx.DB
}

// NewReporter initialize Reporter with db instance.
func NewReporter(db *sqlx.DB) Reporter {
	return Reporter{db: db}
}

var funcMap = template.FuncMap{
	"inc":         func(i int) int { return i + 1 },
	"StringsJoin": strings.Join,
}

// M is map for interfaces.
type M map[string]interface{}

const queryReportTmpl = `
	SELECT
	{{ index . "group" }} AS dimension,

	{{ if eq (index . "group") "queryid" }} any(fingerprint) {{ else }} '' {{ end }} AS fingerprint,
	SUM(num_queries) AS num_queries,

	SUM(m_query_time_cnt) AS m_query_time_cnt,
	SUM(m_query_time_sum) AS m_query_time_sum,
	MIN(m_query_time_min) AS m_query_time_min,
	MAX(m_query_time_max) AS m_query_time_max,
	AVG(m_query_time_p99) AS m_query_time_p99,

	{{range $j, $col := index . "common_columns"}}
		SUM(m_{{ $col }}_cnt) AS m_{{ $col }}_cnt,
		SUM(m_{{ $col }}_sum) AS m_{{ $col }}_sum,
		MIN(m_{{ $col }}_min) AS m_{{ $col }}_min,
		MAX(m_{{ $col }}_max) AS m_{{ $col }}_max,
		AVG(m_{{ $col }}_p99) AS m_{{ $col }}_p99,
	{{ end }}
	{{range $j, $col := index . "bool_columns"}}
		SUM(m_{{ $col }}_cnt) AS m_{{ $col }}_cnt,
		SUM(m_{{ $col }}_sum) AS m_{{ $col }}_sum,
	{{ end }}

	rowNumberInAllBlocks() AS total_rows

	FROM metrics
	WHERE period_start > :period_start_from AND period_start < :period_start_to
	{{ if index . "queryids" }} AND queryid IN ( :queryids ) {{ end }}
	{{ if index . "servers" }} AND d_server IN ( :servers ) {{ end }}
	{{ if index . "databases" }} AND d_database IN ( :databases ) {{ end }}
	{{ if index . "schemas" }} AND d_schema IN ( :schemas ) {{ end }}
	{{ if index . "users" }} AND d_username IN ( :users ) {{ end }}
	{{ if index . "hosts" }} AND d_client_host IN ( :hosts ) {{ end }}
	{{ if index . "labels" }}
		AND (
			{{$i := 0}}
			{{range $key, $val := index . "labels"}}
				{{ $i = inc $i}} {{ if gt $i 1}} OR {{ end }}
				has(['{{ StringsJoin $val "','" }}'], labels.value[indexOf(labels.key, '{{ $key }}')])
			{{ end }}
		)
	{{ end }}
	GROUP BY {{ index . "group" }}
		WITH TOTALS
	ORDER BY {{ index . "order" }}
	LIMIT :offset, :limit
`

// Select select metrics for report.
func (r *Reporter) Select(ctx context.Context, periodStartFromSec, periodStartToSec int64,
	dQueryids, dServers, dDatabases, dSchemas, dUsernames, dClientHosts []string,
	dbLabels map[string][]string, group, order string, offset, limit uint32,
	commonColumns, boolColumns []string) ([]M, error) {

	arg := map[string]interface{}{
		"period_start_from": periodStartFromSec,
		"period_start_to":   periodStartToSec,
		"queryids":          dQueryids,
		"servers":           dServers,
		"databases":         dDatabases,
		"schemas":           dSchemas,
		"users":             dUsernames,
		"hosts":             dClientHosts,
		"labels":            dbLabels,
		"group":             group,
		"order":             order,
		"offset":            offset,
		"limit":             limit,
		"common_columns":    commonColumns,
		"bool_columns":      boolColumns,
	}

	var queryBuffer bytes.Buffer
	if tmpl, err := template.New("queryReport").Funcs(funcMap).Parse(queryReportTmpl); err != nil {
		log.Fatalln(err)
	} else if err = tmpl.Execute(&queryBuffer, arg); err != nil {
		log.Fatalln(err)
	}
	var results []M
	query, args, err := sqlx.Named(queryBuffer.String(), arg)
	if err != nil {
		return results, fmt.Errorf("prepare named:%v", err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return results, fmt.Errorf("populate agruments in IN clause:%v", err)
	}
	query = r.db.Rebind(query)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return results, fmt.Errorf("QueryxContext error:%v", err)
	}
	for rows.Next() {
		result := make(M)
		err = rows.MapScan(result)
		if err != nil {
			fmt.Printf("DimensionReport Scan error: %v", err)
		}
		results = append(results, result)
	}
	rows.NextResultSet()
	total := make(M)
	for rows.Next() {
		err = rows.MapScan(total)
		if err != nil {
			fmt.Printf("DimensionReport Scan TOTALS error: %v", err)
		}
		results = append([]M{total}, results...)
	}
	return results, err
}

const queryReportSparklinesTmpl = `
	SELECT
		intDivOrZero(toUnixTimestamp( :period_start_to ) - toUnixTimestamp(period_start), {{ index . "time_frame" }}) AS point,
		toDateTime(toUnixTimestamp( :period_start_to ) - (point * {{ index . "time_frame" }})) AS timestamp,
		{{ index . "time_frame" }} AS time_frame,
		SUM(num_queries) / time_frame AS num_queries_sum_per_sec,
		{{range $j, $col := index . "columns"}}
			if(SUM(m_{{ $col }}_cnt) == 0, NULL, SUM(m_{{ $col }}_sum) / time_frame) AS m_{{ $col }}_sum_per_sec,
		{{ end }}
		if(SUM(m_query_time_cnt) == 0, NULL, SUM(m_query_time_sum) / time_frame) AS m_query_time_sum_per_sec
	FROM metrics
	WHERE period_start >= :period_start_from AND period_start <= :period_start_to
	{{ if index . "dimension_val" }} AND {{ index . "group" }} = '{{ index . "dimension_val" }}' {{ end }}
	{{ if index . "queryids" }} AND queryid IN ( :queryids ) {{ end }}
	{{ if index . "servers" }} AND d_server IN ( :servers ) {{ end }}
	{{ if index . "databases" }} AND d_database IN ( :databases ) {{ end }}
	{{ if index . "schemas" }} AND d_schema IN ( :schemas ) {{ end }}
	{{ if index . "users" }} AND d_username IN ( :users ) {{ end }}
	{{ if index . "hosts" }} AND d_client_host IN ( :hosts ) {{ end }}
	{{ if index . "labels" }}
		AND (
			{{$i := 0}}
			{{range $key, $val := index . "labels"}}
				{{ $i = inc $i}} {{ if gt $i 1}} OR {{ end }}
				has(['{{ StringsJoin $val "','" }}'], labels.value[indexOf(labels.key, '{{ $key }}')])
			{{ end }}
		)
	{{ end }}
	GROUP BY point
	ORDER BY point ASC;
`

//nolint
var tmplQueryReportSparklines = template.Must(template.New("queryReportSparklines").Funcs(funcMap).Parse(queryReportSparklinesTmpl))

// SelectSparklines selects datapoint for sparklines.
func (r *Reporter) SelectSparklines(ctx context.Context, dimensionVal string,
	periodStartFromSec, periodStartToSec int64,
	dQueryids, dServers, dDatabases, dSchemas, dUsernames, dClientHosts []string,
	dbLabels map[string][]string, group string, columns []string) ([]*qanpb.Point, error) {

	// Align to minutes
	periodStartToSec = periodStartToSec / 60 * 60
	periodStartFromSec = periodStartFromSec / 60 * 60

	// If time range is bigger then two hour - amount of sparklines points = 120 to avoid huge data in response.
	// Otherwise amount of sparklines points is equal to minutes in in time range to not mess up calculation.
	amountOfPoints := int64(optimalAmountOfPoint)
	timePeriod := periodStartToSec - periodStartFromSec
	// reduce amount of point if period less then 2h.
	if timePeriod < int64((minFullTimeFrame).Seconds()) {
		// minimum point is 1 minute
		amountOfPoints = timePeriod / 60
	}

	// how many full minutes we can fit into given amount of points.
	minutesInPoint := (periodStartToSec - periodStartFromSec) / 60 / amountOfPoints
	// we need aditional point to show this minutes
	remainder := ((periodStartToSec - periodStartFromSec) / 60) % amountOfPoints
	amountOfPoints += remainder / minutesInPoint
	timeFrame := minutesInPoint * 60

	arg := map[string]interface{}{
		"dimension_val":     dimensionVal,
		"period_start_from": periodStartFromSec,
		"period_start_to":   periodStartToSec,
		"queryids":          dQueryids,
		"servers":           dServers,
		"databases":         dDatabases,
		"schemas":           dSchemas,
		"users":             dUsernames,
		"hosts":             dClientHosts,
		"labels":            dbLabels,
		"group":             group,
		"columns":           columns,
		"time_frame":        timeFrame,
	}

	var results []*qanpb.Point
	var queryBuffer bytes.Buffer

	if err := tmplQueryReportSparklines.Execute(&queryBuffer, arg); err != nil {
		return nil, errors.Wrap(err, "cannot execute tmplQueryReportSparklines")
	}
	query, args, err := sqlx.Named(queryBuffer.String(), arg)
	if err != nil {
		return results, fmt.Errorf("prepare named:%v", err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return results, fmt.Errorf("populate agruments in IN clause:%v", err)
	}
	query = r.db.Rebind(query)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return results, fmt.Errorf("report query:%v", err)
	}
	resultsWithGaps := map[float64]*qanpb.Point{}
	for rows.Next() {
		res := make(map[string]interface{})
		err = rows.MapScan(res)
		if err != nil {
			fmt.Printf("DimensionReport Scan error: %v", err)
		}
		points := qanpb.Point{
			Values: make(map[string]*structpb.Value),
		}
		for k, v := range res {
			switch i := v.(type) {
			case time.Time:
				points.Values[k] = &structpb.Value{
					Kind: &structpb.Value_StringValue{
						StringValue: i.Format(time.RFC3339),
					},
				}
			case nil:
				points.Values[k] = &structpb.Value{
					Kind: &structpb.Value_NullValue{
						NullValue: 0,
					},
				}
			default:
				f, err := getFloat(i)
				if err != nil {
					err = errors.Wrap(err, "cannot get float for sparkline")
					log.Println(err)
				}
				points.Values[k] = &structpb.Value{
					Kind: &structpb.Value_NumberValue{
						NumberValue: f,
					},
				}
			}
		}
		resultsWithGaps[points.Values["point"].GetNumberValue()] = &points
	}

	// fill in gaps in time series.
	for pointN := int64(0); pointN < amountOfPoints; pointN++ {
		point, ok := resultsWithGaps[float64(pointN)]
		if !ok {
			point = &qanpb.Point{
				Values: make(map[string]*structpb.Value),
			}
			point.Values["point"] = &structpb.Value{
				Kind: &structpb.Value_NumberValue{
					NumberValue: float64(pointN),
				},
			}
			point.Values["time_frame"] = &structpb.Value{
				Kind: &structpb.Value_NumberValue{
					NumberValue: float64(timeFrame),
				},
			}
			timeShift := timeFrame * pointN
			ts := periodStartToSec - timeShift
			point.Values["timestamp"] = &structpb.Value{
				Kind: &structpb.Value_StringValue{
					StringValue: time.Unix(ts, 0).UTC().Format(time.RFC3339),
				},
			}
		}
		results = append(results, point)
	}

	return results, err
}

const queryServers = `
	SELECT d_server AS value, count(d_server) AS count
	  FROM metrics
	 WHERE period_start >= ?
	   AND period_start <= ?
  GROUP BY d_server;
`
const queryDatabases = `
	SELECT d_database AS value, count(d_database) AS count
	  FROM metrics
	 WHERE period_start >= ?
	   AND period_start <= ?
  GROUP BY d_database;
`
const querySchemas = `
	SELECT d_schema AS value, count(d_schema) AS count
	  FROM metrics
	 WHERE period_start >= ?
	   AND period_start <= ?
  GROUP BY d_schema;
`
const queryUsernames = `
	SELECT d_username AS value, count(d_username) AS count
	  FROM metrics
	 WHERE period_start >= ?
	   AND period_start <= ?
  GROUP BY d_username;
`
const queryClientHosts = `
	SELECT d_client_host AS value, count(d_client_host) AS count
	  FROM metrics
	 WHERE period_start >= ?
	   AND period_start <= ?
  GROUP BY d_client_host;
`

const queryLabels = `
	SELECT labels.key AS key, labels.value AS value, COUNT(labels.value) AS count
	  FROM metrics
ARRAY JOIN labels
	 WHERE period_start >= ?
	   AND period_start <= ?
  GROUP BY labels.key, labels.value
  ORDER BY labels.key, labels.value;
`

// SelectFilters selects dimension and their values, and also keys and values of labels.
func (r *Reporter) SelectFilters(ctx context.Context, periodStartFrom, periodStartTo time.Time) (*qanpb.FiltersReply, error) {
	result := qanpb.FiltersReply{
		Labels: make(map[string]*qanpb.ListLabels),
	}

	type customLabel struct {
		key   string
		value string
		count int64
	}

	var servers []*qanpb.ValueAndCount
	var databases []*qanpb.ValueAndCount
	var schemas []*qanpb.ValueAndCount
	var users []*qanpb.ValueAndCount
	var hosts []*qanpb.ValueAndCount
	var labels []*customLabel

	err := r.db.SelectContext(ctx, &servers, queryServers, periodStartFrom, periodStartTo)
	if err != nil {
		return nil, fmt.Errorf("cannot select server dimension:%v", err)
	}
	err = r.db.SelectContext(ctx, &databases, queryDatabases, periodStartFrom, periodStartTo)
	if err != nil {
		return nil, fmt.Errorf("cannot select databases dimension:%v", err)
	}
	err = r.db.SelectContext(ctx, &schemas, querySchemas, periodStartFrom, periodStartTo)
	if err != nil {
		return nil, fmt.Errorf("cannot select schemas dimension:%v", err)
	}
	err = r.db.SelectContext(ctx, &users, queryUsernames, periodStartFrom, periodStartTo)
	if err != nil {
		return nil, fmt.Errorf("cannot select usernames dimension:%v", err)
	}
	err = r.db.SelectContext(ctx, &hosts, queryClientHosts, periodStartFrom, periodStartTo)
	if err != nil {
		return nil, fmt.Errorf("cannot select client hosts dimension:%v", err)
	}

	rows, err := r.db.QueryContext(ctx, queryLabels, periodStartFrom, periodStartTo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select labels dimensions")
	}
	defer rows.Close() //nolint:errcheck

	for rows.Next() {
		var label customLabel
		err = rows.Scan(&label.key, &label.value, &label.count)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan labels dimension")
		}
		labels = append(labels, &label)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to select labels dimensions")
	}

	result.Labels["d_server"] = &qanpb.ListLabels{Name: servers}
	result.Labels["d_database"] = &qanpb.ListLabels{Name: databases}
	result.Labels["d_schema"] = &qanpb.ListLabels{Name: schemas}
	result.Labels["d_username"] = &qanpb.ListLabels{Name: users}
	result.Labels["d_client_host"] = &qanpb.ListLabels{Name: hosts}

	for _, label := range labels {
		if _, ok := result.Labels[label.key]; !ok {
			result.Labels[label.key] = &qanpb.ListLabels{
				Name: []*qanpb.ValueAndCount{},
			}
		}
		val := qanpb.ValueAndCount{
			Value: label.value,
			Count: label.count,
		}
		result.Labels[label.key].Name = append(result.Labels[label.key].Name, &val)
	}

	return &result, nil
}
