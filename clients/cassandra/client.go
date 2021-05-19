package cassandra_client

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/scylladb/go-reflectx"
	"github.com/scylladb/gocqlx/v2"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (gocqlx.Session, error) {
	wrappedSession, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return wrappedSession, err
	}

	wrappedSession.Mapper = reflectx.NewMapperFunc("cassandra", strings.ToLower)

	return wrappedSession, nil
}
