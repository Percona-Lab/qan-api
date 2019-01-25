// qan-api
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
	"fmt"

	metricspb "github.com/Percona-Lab/qan-api/api/metrics"
)

// GetMetricsByDigest implements rpc to exchange data between API and agent.
func (s *Service) GetMetricsByDigest(ctx context.Context, in *metricspb.MetricsRequest) (*metricspb.MetricsReply, error) {
	fmt.Println("Call GetMetricsByDigest")
	labels := in.GetLabels()
	dbServers := []string{}
	dbSchemas := []string{}
	dbUsernames := []string{}
	clientHosts := []string{}
	dbLabels := map[string][]string{}
	for _, label := range labels {
		fmt.Printf("label: %v, : %v \n", label.Key, label.Value)
		switch label.Key {
		case "db_server":
			dbServers = label.Value
		case "db_schema":
			dbSchemas = label.Value
		case "db_username":
			dbUsernames = label.Value
		case "client_host":
			clientHosts = label.Value
		default:
			dbLabels[label.Key] = label.Value
		}
	}
	metrics, err := s.mm.Get(in.From, in.To, in.Digest, dbServers, dbSchemas, dbUsernames, clientHosts, dbLabels)
	return metrics, err
}
