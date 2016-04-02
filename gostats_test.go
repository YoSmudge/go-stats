package gostats

import(
  "testing"
  "os"
  "time"
  "github.com/stretchr/testify/assert"
)

type metricNameTest struct{
  Source    string
  Expected  string
}
func TestSanitizeMetricName(t *testing.T){
  cases := []metricNameTest{
    metricNameTest{"mymetricname", "mymetricname"},
    metricNameTest{"my metric name", "my_metric_name"},
    metricNameTest{"my/metric/name", "my_metric_name"},
    metricNameTest{"my.metric name", "my_metric_name"},
    metricNameTest{"my-metric/name", "my-metric_name"},
    metricNameTest{"my-metric@name", "my-metricname"},
  }

  for _,c := range cases{
    n := sanitizeMetricName(c.Source)
    if n != c.Expected{
      assert.Equal(t, c.Expected, n, "metric name should be sanitized correctly")
    }
  }
}

func TestNew(t *testing.T){
  s := New()
  h, _ := os.Hostname()
  assert.Equal(t, sanitizeMetricName(h), s.Hostname, "hostname should be set")
  assert.Equal(t, "gostats", s.ClientName, "default client name should be set")

  s.Hostname = "localhost"
  assert.Equal(t, "gostats.gostats.localhost.", s.MetricBase(), "metric base should be correct")
}

func TestStart(t *testing.T){
  s, err := Start("localhost:8015", 5, "testclient")
  defer s.Stop()

  assert.Nil(t, err)

  s.Hostname = "localhost"
  assert.Equal(t, "gostats.testclient.localhost.", s.MetricBase(), "metric base should be correct")
  assert.Equal(t, time.Duration(5*time.Second), s.PushInterval, "push interval should be correct")
  assert.Equal(t, "localhost:8015", s.StatsdHost, "statsd host should be correct")
}
