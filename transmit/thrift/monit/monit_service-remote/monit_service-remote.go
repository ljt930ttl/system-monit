// Code generated by Thrift Compiler (0.18.1). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"monit"
)

var _ = monit.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void Ping()")
  fmt.Fprintln(os.Stderr, "  Response Regist(RegInfo request)")
  fmt.Fprintln(os.Stderr, "  Response GetBaseInfo(Info request)")
  fmt.Fprintln(os.Stderr, "  Response LoadCurrentInfo(Info request)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
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
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  var cfg *thrift.TConfiguration = nil
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransportConf(trans, cfg)
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
    protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := monit.NewMonitServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "Ping":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Ping requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Ping(context.Background()))
    fmt.Print("\n")
    break
  case "Regist":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Regist requires 1 args")
      flag.Usage()
    }
    arg25 := flag.Arg(1)
    mbTrans26 := thrift.NewTMemoryBufferLen(len(arg25))
    defer mbTrans26.Close()
    _, err27 := mbTrans26.WriteString(arg25)
    if err27 != nil {
      Usage()
      return
    }
    factory28 := thrift.NewTJSONProtocolFactory()
    jsProt29 := factory28.GetProtocol(mbTrans26)
    argvalue0 := monit.NewRegInfo()
    err30 := argvalue0.Read(context.Background(), jsProt29)
    if err30 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Regist(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetBaseInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetBaseInfo requires 1 args")
      flag.Usage()
    }
    arg31 := flag.Arg(1)
    mbTrans32 := thrift.NewTMemoryBufferLen(len(arg31))
    defer mbTrans32.Close()
    _, err33 := mbTrans32.WriteString(arg31)
    if err33 != nil {
      Usage()
      return
    }
    factory34 := thrift.NewTJSONProtocolFactory()
    jsProt35 := factory34.GetProtocol(mbTrans32)
    argvalue0 := monit.NewInfo()
    err36 := argvalue0.Read(context.Background(), jsProt35)
    if err36 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetBaseInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "LoadCurrentInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "LoadCurrentInfo requires 1 args")
      flag.Usage()
    }
    arg37 := flag.Arg(1)
    mbTrans38 := thrift.NewTMemoryBufferLen(len(arg37))
    defer mbTrans38.Close()
    _, err39 := mbTrans38.WriteString(arg37)
    if err39 != nil {
      Usage()
      return
    }
    factory40 := thrift.NewTJSONProtocolFactory()
    jsProt41 := factory40.GetProtocol(mbTrans38)
    argvalue0 := monit.NewInfo()
    err42 := argvalue0.Read(context.Background(), jsProt41)
    if err42 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.LoadCurrentInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
