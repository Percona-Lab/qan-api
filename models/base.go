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
	"time"

	"github.com/percona/pmm/api/inventorypb"
	"github.com/percona/pmm/api/qanpb"
)

const queryTimeout = 30 * time.Second

//nolint
var sparklinePointAllFields = []string{
	"point",
	"timestamp",
	"time_frame",
	"num_queries_per_sec",
	"m_query_time_sum_per_sec",
	"m_lock_time_sum_per_sec",
	"m_rows_sent_sum_per_sec",
	"m_rows_examined_sum_per_sec",
	"m_rows_affected_sum_per_sec",
	"m_rows_read_sum_per_sec",
	"m_merge_passes_sum_per_sec",
	"m_innodb_io_r_ops_sum_per_sec",
	"m_innodb_io_r_bytes_sum_per_sec",
	"m_innodb_io_r_wait_sum_per_sec",
	"m_innodb_rec_lock_wait_sum_per_sec",
	"m_innodb_queue_wait_sum_per_sec",
	"m_innodb_pages_distinct_sum_per_sec",
	"m_query_length_sum_per_sec",
	"m_bytes_sent_sum_per_sec",
	"m_tmp_tables_sum_per_sec",
	"m_tmp_disk_tables_sum_per_sec",
	"m_tmp_table_sizes_sum_per_sec",
	"m_qc_hit_sum_per_sec",
	"m_full_scan_sum_per_sec",
	"m_full_join_sum_per_sec",
	"m_tmp_table_sum_per_sec",
	"m_tmp_table_on_disk_sum_per_sec",
	"m_filesort_sum_per_sec",
	"m_filesort_on_disk_sum_per_sec",
	"m_select_full_range_join_sum_per_sec",
	"m_select_range_sum_per_sec",
	"m_select_range_check_sum_per_sec",
	"m_sort_range_sum_per_sec",
	"m_sort_rows_sum_per_sec",
	"m_sort_scan_sum_per_sec",
	"m_no_index_used_sum_per_sec",
	"m_no_good_index_used_sum_per_sec",
	"m_docs_returned_sum_per_sec",
	"m_response_length_sum_per_sec",
	"m_docs_scanned_sum_per_sec",
}

func getPointFieldsList(point *qanpb.Point, fields []string) []interface{} {
	sparklinePointValuesMap := map[string]interface{}{
		"point":                                &point.Point,
		"timestamp":                            &point.Timestamp,
		"time_frame":                           &point.TimeFrame,
		"load":                                 &point.Load,
		"num_queries_per_sec":                  &point.NumQueriesPerSec,
		"m_query_time_sum_per_sec":             &point.MQueryTimeSumPerSec,
		"m_lock_time_sum_per_sec":              &point.MLockTimeSumPerSec,
		"m_rows_sent_sum_per_sec":              &point.MRowsSentSumPerSec,
		"m_rows_examined_sum_per_sec":          &point.MRowsExaminedSumPerSec,
		"m_rows_affected_sum_per_sec":          &point.MRowsAffectedSumPerSec,
		"m_rows_read_sum_per_sec":              &point.MRowsReadSumPerSec,
		"m_merge_passes_sum_per_sec":           &point.MMergePassesSumPerSec,
		"m_innodb_io_r_ops_sum_per_sec":        &point.MInnodbIoROpsSumPerSec,
		"m_innodb_io_r_bytes_sum_per_sec":      &point.MInnodbIoRBytesSumPerSec,
		"m_innodb_io_r_wait_sum_per_sec":       &point.MInnodbIoRWaitSumPerSec,
		"m_innodb_rec_lock_wait_sum_per_sec":   &point.MInnodbRecLockWaitSumPerSec,
		"m_innodb_queue_wait_sum_per_sec":      &point.MInnodbQueueWaitSumPerSec,
		"m_innodb_pages_distinct_sum_per_sec":  &point.MInnodbPagesDistinctSumPerSec,
		"m_query_length_sum_per_sec":           &point.MQueryLengthSumPerSec,
		"m_bytes_sent_sum_per_sec":             &point.MBytesSentSumPerSec,
		"m_tmp_tables_sum_per_sec":             &point.MTmpTablesSumPerSec,
		"m_tmp_disk_tables_sum_per_sec":        &point.MTmpDiskTablesSumPerSec,
		"m_tmp_table_sizes_sum_per_sec":        &point.MTmpTableSizesSumPerSec,
		"m_qc_hit_sum_per_sec":                 &point.MQcHitSumPerSec,
		"m_full_scan_sum_per_sec":              &point.MFullScanSumPerSec,
		"m_full_join_sum_per_sec":              &point.MFullJoinSumPerSec,
		"m_tmp_table_sum_per_sec":              &point.MTmpTableSumPerSec,
		"m_tmp_table_on_disk_sum_per_sec":      &point.MTmpTableOnDiskSumPerSec,
		"m_filesort_sum_per_sec":               &point.MFilesortSumPerSec,
		"m_filesort_on_disk_sum_per_sec":       &point.MFilesortOnDiskSumPerSec,
		"m_select_full_range_join_sum_per_sec": &point.MSelectFullRangeJoinSumPerSec,
		"m_select_range_sum_per_sec":           &point.MSelectRangeSumPerSec,
		"m_select_range_check_sum_per_sec":     &point.MSelectRangeCheckSumPerSec,
		"m_sort_range_sum_per_sec":             &point.MSortRangeSumPerSec,
		"m_sort_rows_sum_per_sec":              &point.MSortRowsSumPerSec,
		"m_sort_scan_sum_per_sec":              &point.MSortScanSumPerSec,
		"m_no_index_used_sum_per_sec":          &point.MNoIndexUsedSumPerSec,
		"m_no_good_index_used_sum_per_sec":     &point.MNoGoodIndexUsedSumPerSec,
		"m_docs_returned_sum_per_sec":          &point.MDocsReturnedSumPerSec,
		"m_response_length_sum_per_sec":        &point.MResponseLengthSumPerSec,
		"m_docs_scanned_sum_per_sec":           &point.MDocsScannedSumPerSec,
	}

	sparklinePointValuesList := []interface{}{}
	for _, v := range fields {
		sparklinePointValuesList = append(sparklinePointValuesList, sparklinePointValuesMap[v])
	}

	return sparklinePointValuesList
}

func isValidMetricColumn(name string) bool {
	fields := map[string]struct{}{
		"num_queries":                  {},
		"m_query_time_cnt":             {},
		"m_query_time_sum":             {},
		"m_query_time_min":             {},
		"m_query_time_max":             {},
		"m_query_time_p99":             {},
		"m_lock_time_cnt":              {},
		"m_lock_time_sum":              {},
		"m_lock_time_min":              {},
		"m_lock_time_max":              {},
		"m_lock_time_p99":              {},
		"m_rows_sent_cnt":              {},
		"m_rows_sent_sum":              {},
		"m_rows_sent_min":              {},
		"m_rows_sent_max":              {},
		"m_rows_sent_p99":              {},
		"m_rows_examined_cnt":          {},
		"m_rows_examined_sum":          {},
		"m_rows_examined_min":          {},
		"m_rows_examined_max":          {},
		"m_rows_examined_p99":          {},
		"m_rows_affected_cnt":          {},
		"m_rows_affected_sum":          {},
		"m_rows_affected_min":          {},
		"m_rows_affected_max":          {},
		"m_rows_affected_p99":          {},
		"m_rows_read_cnt":              {},
		"m_rows_read_sum":              {},
		"m_rows_read_min":              {},
		"m_rows_read_max":              {},
		"m_rows_read_p99":              {},
		"m_merge_passes_cnt":           {},
		"m_merge_passes_sum":           {},
		"m_merge_passes_min":           {},
		"m_merge_passes_max":           {},
		"m_merge_passes_p99":           {},
		"m_innodb_io_r_ops_cnt":        {},
		"m_innodb_io_r_ops_sum":        {},
		"m_innodb_io_r_ops_min":        {},
		"m_innodb_io_r_ops_max":        {},
		"m_innodb_io_r_ops_p99":        {},
		"m_innodb_io_r_bytes_cnt":      {},
		"m_innodb_io_r_bytes_sum":      {},
		"m_innodb_io_r_bytes_min":      {},
		"m_innodb_io_r_bytes_max":      {},
		"m_innodb_io_r_bytes_p99":      {},
		"m_innodb_io_r_wait_cnt":       {},
		"m_innodb_io_r_wait_sum":       {},
		"m_innodb_io_r_wait_min":       {},
		"m_innodb_io_r_wait_max":       {},
		"m_innodb_io_r_wait_p99":       {},
		"m_innodb_rec_lock_wait_cnt":   {},
		"m_innodb_rec_lock_wait_sum":   {},
		"m_innodb_rec_lock_wait_min":   {},
		"m_innodb_rec_lock_wait_max":   {},
		"m_innodb_rec_lock_wait_p99":   {},
		"m_innodb_queue_wait_cnt":      {},
		"m_innodb_queue_wait_sum":      {},
		"m_innodb_queue_wait_min":      {},
		"m_innodb_queue_wait_max":      {},
		"m_innodb_queue_wait_p99":      {},
		"m_innodb_pages_distinct_cnt":  {},
		"m_innodb_pages_distinct_sum":  {},
		"m_innodb_pages_distinct_min":  {},
		"m_innodb_pages_distinct_max":  {},
		"m_innodb_pages_distinct_p99":  {},
		"m_query_length_cnt":           {},
		"m_query_length_sum":           {},
		"m_query_length_min":           {},
		"m_query_length_max":           {},
		"m_query_length_p99":           {},
		"m_bytes_sent_cnt":             {},
		"m_bytes_sent_sum":             {},
		"m_bytes_sent_min":             {},
		"m_bytes_sent_max":             {},
		"m_bytes_sent_p99":             {},
		"m_tmp_tables_cnt":             {},
		"m_tmp_tables_sum":             {},
		"m_tmp_tables_min":             {},
		"m_tmp_tables_max":             {},
		"m_tmp_tables_p99":             {},
		"m_tmp_disk_tables_cnt":        {},
		"m_tmp_disk_tables_sum":        {},
		"m_tmp_disk_tables_min":        {},
		"m_tmp_disk_tables_max":        {},
		"m_tmp_disk_tables_p99":        {},
		"m_tmp_table_sizes_cnt":        {},
		"m_tmp_table_sizes_sum":        {},
		"m_tmp_table_sizes_min":        {},
		"m_tmp_table_sizes_max":        {},
		"m_tmp_table_sizes_p99":        {},
		"m_qc_hit_cnt":                 {},
		"m_qc_hit_sum":                 {},
		"m_full_scan_cnt":              {},
		"m_full_scan_sum":              {},
		"m_full_join_cnt":              {},
		"m_full_join_sum":              {},
		"m_tmp_table_cnt":              {},
		"m_tmp_table_sum":              {},
		"m_tmp_table_on_disk_cnt":      {},
		"m_tmp_table_on_disk_sum":      {},
		"m_filesort_cnt":               {},
		"m_filesort_sum":               {},
		"m_filesort_on_disk_cnt":       {},
		"m_filesort_on_disk_sum":       {},
		"m_select_full_range_join_cnt": {},
		"m_select_full_range_join_sum": {},
		"m_select_range_cnt":           {},
		"m_select_range_sum":           {},
		"m_select_range_check_cnt":     {},
		"m_select_range_check_sum":     {},
		"m_sort_range_cnt":             {},
		"m_sort_range_sum":             {},
		"m_sort_rows_cnt":              {},
		"m_sort_rows_sum":              {},
		"m_sort_scan_cnt":              {},
		"m_sort_scan_sum":              {},
		"m_no_index_used_cnt":          {},
		"m_no_index_used_sum":          {},
		"m_no_good_index_used_cnt":     {},
		"m_no_good_index_used_sum":     {},
		"m_docs_returned_cnt":          {},
		"m_docs_returned_sum":          {},
		"m_docs_returned_min":          {},
		"m_docs_returned_max":          {},
		"m_docs_returned_p99":          {},
		"m_response_length_cnt":        {},
		"m_response_length_sum":        {},
		"m_response_length_min":        {},
		"m_response_length_max":        {},
		"m_response_length_p99":        {},
		"m_docs_scanned_cnt":           {},
		"m_docs_scanned_sum":           {},
		"m_docs_scanned_min":           {},
		"m_docs_scanned_max":           {},
		"m_docs_scanned_p99":           {},
	}
	_, isValid := fields[name]
	return isValid
}

func agentTypeToClickHouseEnum(agentType inventorypb.AgentType) string {
	// agentTypes represents Agent type as stored in database.
	// Should be same as in pmm/inventorypb/agents.proto
	agentTypes := map[inventorypb.AgentType]string{
		inventorypb.AgentType_AGENT_TYPE_INVALID:         "agent_type_invalid",
		inventorypb.AgentType_QAN_MYSQL_PERFSCHEMA_AGENT: "mysql-perfschema",
		inventorypb.AgentType_QAN_MYSQL_SLOWLOG_AGENT:    "mysql-slowlog",
		inventorypb.AgentType_QAN_MONGODB_PROFILER_AGENT: "mongodb-profiler",
	}

	if val, ok := agentTypes[agentType]; ok {
		return val
	}
	return agentTypes[inventorypb.AgentType_AGENT_TYPE_INVALID]
}
