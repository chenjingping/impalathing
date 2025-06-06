// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "github.com/chenjingping/thrift/lib/go/thrift"
	"github.com/chenjingping/impalathing/services/status"
	"github.com/chenjingping/impalathing/services/beeswax"
	"github.com/chenjingping/impalathing/services/cli_service"
        "github.com/chenjingping/impalathing/services/impalaservice"
)

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TStatus Cancel(QueryHandle query_id)")
  fmt.Fprintln(os.Stderr, "  TInsertResult CloseInsert(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  void PingImpalaService()")
  fmt.Fprintln(os.Stderr, "  QueryHandle query(Query query)")
  fmt.Fprintln(os.Stderr, "  QueryHandle executeAndWait(Query query, LogContextId clientCtx)")
  fmt.Fprintln(os.Stderr, "  QueryExplanation explain(Query query)")
  fmt.Fprintln(os.Stderr, "  Results fetch(QueryHandle query_id, bool start_over, i32 fetch_size)")
  fmt.Fprintln(os.Stderr, "  QueryState get_state(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  ResultsMetadata get_results_metadata(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  string echo(string s)")
  fmt.Fprintln(os.Stderr, "  string dump_config()")
  fmt.Fprintln(os.Stderr, "  string get_log(LogContextId context)")
  fmt.Fprintln(os.Stderr, "   get_default_configuration(bool include_hadoop)")
  fmt.Fprintln(os.Stderr, "  void close(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  void clean(LogContextId log_context)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := impalaservice.NewImpalaServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "Cancel":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Cancel requires 1 args")
      flag.Usage()
    }
    arg9 := flag.Arg(1)
    mbTrans10 := thrift.NewTMemoryBufferLen(len(arg9))
    defer mbTrans10.Close()
    _, err11 := mbTrans10.WriteString(arg9)
    if err11 != nil {
      Usage()
      return
    }
    factory12 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt13 := factory12.GetProtocol(mbTrans10)
    argvalue0 := beeswax.NewQueryHandle()
    err14 := argvalue0.Read(jsProt13)
    if err14 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Cancel(value0))
    fmt.Print("\n")
    break
  case "CloseInsert":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseInsert requires 1 args")
      flag.Usage()
    }
    arg15 := flag.Arg(1)
    mbTrans16 := thrift.NewTMemoryBufferLen(len(arg15))
    defer mbTrans16.Close()
    _, err17 := mbTrans16.WriteString(arg15)
    if err17 != nil {
      Usage()
      return
    }
    factory18 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt19 := factory18.GetProtocol(mbTrans16)
    argvalue0 := beeswax.NewQueryHandle()
    err20 := argvalue0.Read(jsProt19)
    if err20 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseInsert(value0))
    fmt.Print("\n")
    break
  case "PingImpalaService":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "PingImpalaService requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.PingImpalaService())
    fmt.Print("\n")
    break
  case "query":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Query requires 1 args")
      flag.Usage()
    }
    arg21 := flag.Arg(1)
    mbTrans22 := thrift.NewTMemoryBufferLen(len(arg21))
    defer mbTrans22.Close()
    _, err23 := mbTrans22.WriteString(arg21)
    if err23 != nil {
      Usage()
      return
    }
    factory24 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt25 := factory24.GetProtocol(mbTrans22)
    argvalue0 := beeswax.NewQuery()
    err26 := argvalue0.Read(jsProt25)
    if err26 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Query(value0))
    fmt.Print("\n")
    break
  case "executeAndWait":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ExecuteAndWait requires 2 args")
      flag.Usage()
    }
    arg27 := flag.Arg(1)
    mbTrans28 := thrift.NewTMemoryBufferLen(len(arg27))
    defer mbTrans28.Close()
    _, err29 := mbTrans28.WriteString(arg27)
    if err29 != nil {
      Usage()
      return
    }
    factory30 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt31 := factory30.GetProtocol(mbTrans28)
    argvalue0 := beeswax.NewQuery()
    err32 := argvalue0.Read(jsProt31)
    if err32 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := beeswax.LogContextId(argvalue1)
    fmt.Print(client.ExecuteAndWait(value0, value1))
    fmt.Print("\n")
    break
  case "explain":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Explain requires 1 args")
      flag.Usage()
    }
    arg34 := flag.Arg(1)
    mbTrans35 := thrift.NewTMemoryBufferLen(len(arg34))
    defer mbTrans35.Close()
    _, err36 := mbTrans35.WriteString(arg34)
    if err36 != nil {
      Usage()
      return
    }
    factory37 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt38 := factory37.GetProtocol(mbTrans35)
    argvalue0 := beeswax.NewQuery()
    err39 := argvalue0.Read(jsProt38)
    if err39 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Explain(value0))
    fmt.Print("\n")
    break
  case "fetch":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Fetch requires 3 args")
      flag.Usage()
    }
    arg40 := flag.Arg(1)
    mbTrans41 := thrift.NewTMemoryBufferLen(len(arg40))
    defer mbTrans41.Close()
    _, err42 := mbTrans41.WriteString(arg40)
    if err42 != nil {
      Usage()
      return
    }
    factory43 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt44 := factory43.GetProtocol(mbTrans41)
    argvalue0 := beeswax.NewQueryHandle()
    err45 := argvalue0.Read(jsProt44)
    if err45 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    tmp2, err47 := (strconv.Atoi(flag.Arg(3)))
    if err47 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.Fetch(value0, value1, value2))
    fmt.Print("\n")
    break
  case "get_state":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetState requires 1 args")
      flag.Usage()
    }
    arg48 := flag.Arg(1)
    mbTrans49 := thrift.NewTMemoryBufferLen(len(arg48))
    defer mbTrans49.Close()
    _, err50 := mbTrans49.WriteString(arg48)
    if err50 != nil {
      Usage()
      return
    }
    factory51 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt52 := factory51.GetProtocol(mbTrans49)
    argvalue0 := beeswax.NewQueryHandle()
    err53 := argvalue0.Read(jsProt52)
    if err53 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetState(value0))
    fmt.Print("\n")
    break
  case "get_results_metadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultsMetadata requires 1 args")
      flag.Usage()
    }
    arg54 := flag.Arg(1)
    mbTrans55 := thrift.NewTMemoryBufferLen(len(arg54))
    defer mbTrans55.Close()
    _, err56 := mbTrans55.WriteString(arg54)
    if err56 != nil {
      Usage()
      return
    }
    factory57 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt58 := factory57.GetProtocol(mbTrans55)
    argvalue0 := beeswax.NewQueryHandle()
    err59 := argvalue0.Read(jsProt58)
    if err59 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultsMetadata(value0))
    fmt.Print("\n")
    break
  case "echo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Echo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.Echo(value0))
    fmt.Print("\n")
    break
  case "dump_config":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "DumpConfig requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.DumpConfig())
    fmt.Print("\n")
    break
  case "get_log":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLog requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.GetLog(value0))
    fmt.Print("\n")
    break
  case "get_default_configuration":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDefaultConfiguration requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1) == "true"
    value0 := argvalue0
    fmt.Print(client.GetDefaultConfiguration(value0))
    fmt.Print("\n")
    break
  case "close":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Close requires 1 args")
      flag.Usage()
    }
    arg63 := flag.Arg(1)
    mbTrans64 := thrift.NewTMemoryBufferLen(len(arg63))
    defer mbTrans64.Close()
    _, err65 := mbTrans64.WriteString(arg63)
    if err65 != nil {
      Usage()
      return
    }
    factory66 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt67 := factory66.GetProtocol(mbTrans64)
    argvalue0 := beeswax.NewQueryHandle()
    err68 := argvalue0.Read(jsProt67)
    if err68 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Close(value0))
    fmt.Print("\n")
    break
  case "clean":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Clean requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.Clean(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
