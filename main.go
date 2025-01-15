package main

import (
	"flag"
	"fmt"
)

func main() {
	prefix := flag.String("prefix", "thisguymartin-pit", "Stack name prefix to match")
	region := flag.String("region", "", "AWS region (optional, defaults to AWS_REGION env var)")

	flag.Parse()

	fmt.Println(*prefix)
	fmt.Println(*region)

}
