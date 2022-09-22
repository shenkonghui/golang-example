package main

import (
	"bytes"
	"flag"
	"fmt"
	"k8s.io/klog/v2"
)

func main() {
	//flag.CommandLine.Set("-log_dir", "logs")
	//flag.Parse()
	////klog.Fatalln("Fatalln")
	//klog.Errorln("Errorln")
	//klog.Warningln("Warningln")
	//klog.Infoln("Infoln")
	//klog.Flush()

	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("alsologtostderr", "false")
	flag.Set("log_file", "./log")
	flag.Parse()

	buf := new(bytes.Buffer)
	klog.SetOutput(buf)
	klog.Info("nice to meet you")
	klog.Flush()

	fmt.Printf("LOGGED: %s", buf.String())
}
