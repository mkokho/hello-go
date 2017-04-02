package main

import (
	"time"
	"log"
)

func main() {
	before, err := time.Parse(time.RFC3339, "2016-10-25T15:02:55+13:00")
	if err != nil {
		log.Printf("Unexpected error: %s", err.Error())
	}
	now := time.Now()
	log.Printf("Now is %s, before is %s", now.Format(time.RFC3339), before.Format(time.RFC3339))
	log.Printf("Difference is %v", int(time.Since(before).Seconds()))
}


/*

type CustomTime struct {
	time.Time
}

const ctLayout = "2006/01/02|15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	timestampNanos, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		println(err)
	}
	secs := timestampNanos / 1000000000
	nanos := timestampNanos % 1000000000
	ct.Time = time.Unix(secs, nanos)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Args struct {
	Time CustomTime
}

var data = `
    {"Time": 1472175574000000000}
`

func healthCheckBurst(size int) {
	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		f := func(idx int) int {
			fmt.Printf("Started request %d\n", idx)
			start := time.Now()
			r, _ := http.DefaultClient.Get("http://app5.mc.int.movio.co:8082/vc_ipicus/pub/1/healthcheck")
			//io.Copy(os.Stdout, r.Body)
			elapsed := time.Since(start)
			fmt.Printf("Completed %d request in %dms: %s\n", idx, elapsed.Nanoseconds()/1000000, r.Status)
			wg.Done()
			return 0
		}
		go f(i)
	}

	wg.Wait()
}

func healthCheckInterval() {

}

type Topics struct {
	CISInTopic      string `json:"cis-in"`
	CISSuccessTopic string `json:"cis-out"`
}

type Tenant struct {
	Schema         string `json:"schema"`
	TopicsOverride Topics `json:"topics-override"`
}

type Config struct {
	Tenants       []Tenant `json:"tenants"`
	Topics        Topics   `json:"topics"`
	Brokers       string   `json:"brokers"`
	ConsumerGroup string   `json:"consumerGroup"`
}

func main() {
	inShared := ` {
	  "topics": {
	    "cis-in": "topic1",
	    "cis-out": "topic2",
	    "email-in": "",
	  },
	  "brokers": "va-mq-b1.movio.co:9092,va-mq-b2.movio.co:9092,va-mq-b3.movio.co:9092",
	  "tenants": [
	    {
	      "schema": "vc_regalus",
	      "topics-override": {
	        "cis-in": "topic1-regal"
	      }
	    }
	  ]
	}`

	inLocal := `{
		  "consumerGroup": "cis-requestor",
			"topics": {
	       "cis-in": "topic-override"
	  	}
  }`

	config := &Config{}
	json.Unmarshal([]byte(inShared), &config)
	json.Unmarshal([]byte(inLocal), &config)

	//fmt.Printf("Parse config: %+v\n", config)
	//out, _ := json.MarshalIndent(config, "", "  ")
	//fmt.Printf("Out: \n%s\n", string(out))
	fmt.Printf("Out: %d %d \n", time.Now().Second()  , time.Now().Second() % 20)

	//config.Tenants[0].TopicsOverride.CISInTopic

	//healthCheckBurst(50)
	now := time.Now()
		fmt.Println(now.Format(time.RFC3339Nano))
		fmt.Println(now.UnixNano())

		a := Args{}
		fmt.Println(json.Unmarshal([]byte(data), &a))
		fmt.Println(time.Now().Truncate(time.Minute).UTC().Format("2006-01-02T15:04:05.000"))
		fmt.Println(time.Now().Truncate(time.Minute).UTC().Format("2006-01-02T15:04:05.999"))
		fmt.Println(a.Time.Format(time.RFC3339Nano))
		fmt.Println(time.Now().Format(time.RFC3339Nano))
		fmt.Println(time.Unix(1472175574, 43))
		fmt.Println()


 */
