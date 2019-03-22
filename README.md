# qan-api2

[![Build Status](https://travis-ci.org/percona/qan-api2.svg?branch=master)](https://travis-ci.org/percona/qan-api2)
[![Go Report Card](https://goreportcard.com/badge/github.com/percona/qan-api2)](https://goreportcard.com/report/github.com/percona/qan-api2)
[![pullreminders](https://pullreminders.com/badge.svg)](https://pullreminders.com?ref=badge)

qan-api for PMM 2.x.


# Get Report


Examples:
```bash
curl -s -X POST -d '{"period_start_from": "2019-01-01 00:00:00", "period_start_to": "2019-01-01 01:00:00"}' http://127.0.0.1:9922/v1/qan/GetReport | jq

curl -s -X POST -d '{"period_start_from": "2019-01-01 00:00:00", "period_start_to": "2019-01-01 01:00:00"}' http://127.0.0.1:9922/v1/qan/GetReport | jq

curl -s -X POST -d '{"period_start_from": "2019-01-01 00:00:00", "period_start_to": "2019-01-01 01:00:00", "group_by": "d_client_host"}' http://127.0.0.1:9922/v1/qan/GetReport | jq

curl -X POST -s -d '{"period_start_from": "2019-01-01 00:00:00", "period_start_to": "2019-01-01 23:00:00",  "labels": [{"key": "d_client_host", "value": ["10.11.12.4", "10.11.12.59"]}]}' http://127.0.0.1:9922/v1/qan/GetReport | jq

curl -s -X POST -d '{"period_start_from": "2019-01-01 00:00:00", "period_start_to": "2019-01-01 01:00:00", "group_by": "d_client_host", "offset": 10}' http://127.0.0.1:9922/v1/qan/GetReport | jq

curl -s -X POST -d '{"period_start_from": "2019-01-01 00:00:00", "period_start_to": "2019-01-01 01:00:00", "order_by": "num_queries"}' http://127.0.0.1:9922/v1/qan/GetReport | jq

curl -X POST -s -d '{"period_start_from": "2019-01-01 00:00:00", "period_start_to": "2019-01-01 01:00:00", "filter_by": "7DD5F6760F2D2EBB"}' http://127.0.0.1:9922/v1/qan/GetMetrics | jq

 ```

 ```
 curl -X POST -d '{"from": "2019-01-01T00:00:00Z", "to": "2019-01-01T10:00:00Z"}'  http://127.0.0.1:9922/v1/qan/GetFilters
 ```

# Get list of availible metrics.

`curl -X POST -d '{}' http://127.0.0.1:9922/v1/qan/GetMetricsNames -s | jq`

```json
{
  "data": {
    "bytes_sent": "Bytes Sent",
    "count": "Count",
    "docs_returned": "Docs Returned",
    "docs_scanned": "Docs Scanned",
    "filesort": "Filesort",
    "filesort_on_disk": "Filesort on Disk",
    "full_join": "Full Join",
    "full_scan": "Full Scan",
    "innodb_io_r_bytes": "Innodb IO R Bytes",
    "innodb_io_r_ops": "Innodb IO R Ops",
    "innodb_io_r_wait": "Innodb IO R Wait",
    "innodb_pages_distinct": "Innodb Pages Distinct",
    "innodb_queue_wait": "Innodb Queue Wait",
    "innodb_rec_lock_wait": "Innodb Rec Lock Wait",
    "latancy": "Latancy",
    "load": "Load",
    "lock_time": "Lock Time",
    "merge_passes": "Merge Passes",
    "no_good_index_used": "No Good Index Used",
    "no_index_used": "No Index Used",
    "qc_hit": "Query Cache Hit",
    "query_length": "Query Length",
    "query_time": "Query Time",
    "response_length": "Response Length",
    "rows_affected": "Rows Affected",
    "rows_examined": "Rows Examined",
    "rows_read": "Rows Read",
    "rows_sent": "Rows Sent",
    "select_full_range_join": "Select Full Range Join",
    "select_range": "Select Range",
    "select_range_check": "Select Range Check",
    "sort_range": "Sort Range",
    "sort_rows": "Sort Rows",
    "sort_scan": "Sort Scan",
    "tmp_disk_tables": "Tmp Disk Tables",
    "tmp_table": "Tmp Table",
    "tmp_table_on_disk": "Tmp Table on Disk",
    "tmp_table_sizes": "Tmp Table Sizes",
    "tmp_tables": "Tmp Tables"
  }
}
```
