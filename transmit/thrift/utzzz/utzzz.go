// Code generated by Thrift Compiler (0.18.1). DO NOT EDIT.

package utzzz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

// Attributes:
//  - Head
//  - Src
//  - Targets
//  - Content
type PMAMsg struct {
  Head map[string]string `thrift:"head,1,required" db:"head" json:"head"`
  Src string `thrift:"src,2,required" db:"src" json:"src"`
  Targets []string `thrift:"targets,3,required" db:"targets" json:"targets"`
  Content string `thrift:"content,4,required" db:"content" json:"content"`
}

func NewPMAMsg() *PMAMsg {
  return &PMAMsg{}
}


func (p *PMAMsg) GetHead() map[string]string {
  return p.Head
}

func (p *PMAMsg) GetSrc() string {
  return p.Src
}

func (p *PMAMsg) GetTargets() []string {
  return p.Targets
}

func (p *PMAMsg) GetContent() string {
  return p.Content
}
func (p *PMAMsg) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetHead bool = false;
  var issetSrc bool = false;
  var issetTargets bool = false;
  var issetContent bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.MAP {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetHead = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetSrc = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetTargets = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
        issetContent = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetHead{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Head is not set"));
  }
  if !issetSrc{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Src is not set"));
  }
  if !issetTargets{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Targets is not set"));
  }
  if !issetContent{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Content is not set"));
  }
  return nil
}

func (p *PMAMsg)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[string]string, size)
  p.Head =  tMap
  for i := 0; i < size; i ++ {
var _key0 string
    if v, err := iprot.ReadString(ctx); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
var _val1 string
    if v, err := iprot.ReadString(ctx); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val1 = v
}
    p.Head[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(ctx); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *PMAMsg)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Src = v
}
  return nil
}

func (p *PMAMsg)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.Targets =  tSlice
  for i := 0; i < size; i ++ {
var _elem2 string
    if v, err := iprot.ReadString(ctx); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem2 = v
}
    p.Targets = append(p.Targets, _elem2)
  }
  if err := iprot.ReadListEnd(ctx); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *PMAMsg)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Content = v
}
  return nil
}

func (p *PMAMsg) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "PMAMsg"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PMAMsg) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "head", thrift.MAP, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:head: ", p), err) }
  if err := oprot.WriteMapBegin(ctx, thrift.STRING, thrift.STRING, len(p.Head)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.Head {
    if err := oprot.WriteString(ctx, string(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := oprot.WriteString(ctx, string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteMapEnd(ctx); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:head: ", p), err) }
  return err
}

func (p *PMAMsg) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "src", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:src: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Src)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.src (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:src: ", p), err) }
  return err
}

func (p *PMAMsg) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "targets", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:targets: ", p), err) }
  if err := oprot.WriteListBegin(ctx, thrift.STRING, len(p.Targets)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Targets {
    if err := oprot.WriteString(ctx, string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(ctx); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:targets: ", p), err) }
  return err
}

func (p *PMAMsg) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "content", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:content: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Content)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.content (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:content: ", p), err) }
  return err
}

func (p *PMAMsg) Equals(other *PMAMsg) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if len(p.Head) != len(other.Head) { return false }
  for k, _tgt := range p.Head {
    _src3 := other.Head[k]
    if _tgt != _src3 { return false }
  }
  if p.Src != other.Src { return false }
  if len(p.Targets) != len(other.Targets) { return false }
  for i, _tgt := range p.Targets {
    _src4 := other.Targets[i]
    if _tgt != _src4 { return false }
  }
  if p.Content != other.Content { return false }
  return true
}

func (p *PMAMsg) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PMAMsg(%+v)", *p)
}

func (p *PMAMsg) Validate() error {
  return nil
}
type PMAService interface {
  // Parameters:
  //  - PmaMsg
  RequestFunc(ctx context.Context, pmaMsg *PMAMsg) (_err error)
}

type PMAServiceClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewPMAServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PMAServiceClient {
  return &PMAServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewPMAServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PMAServiceClient {
  return &PMAServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewPMAServiceClient(c thrift.TClient) *PMAServiceClient {
  return &PMAServiceClient{
    c: c,
  }
}

func (p *PMAServiceClient) Client_() thrift.TClient {
  return p.c
}

func (p *PMAServiceClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *PMAServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - PmaMsg
func (p *PMAServiceClient) RequestFunc(ctx context.Context, pmaMsg *PMAMsg) (_err error) {
  var _args5 PMAServiceRequestFuncArgs
  _args5.PmaMsg = pmaMsg
  p.SetLastResponseMeta_(thrift.ResponseMeta{})
  if _, err := p.Client_().Call(ctx, "requestFunc", &_args5, nil); err != nil {
    return err
  }
  return nil
}

type PMAServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler PMAService
}

func (p *PMAServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *PMAServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *PMAServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewPMAServiceProcessor(handler PMAService) *PMAServiceProcessor {

  self6 := &PMAServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["requestFunc"] = &pMAServiceProcessorRequestFunc{handler:handler}
return self6
}

func (p *PMAServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x7.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x7

}

type pMAServiceProcessorRequestFunc struct {
  handler PMAService
}

func (p *pMAServiceProcessorRequestFunc) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := PMAServiceRequestFuncArgs{}
  if err2 := args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  _ = tickerCancel

  if err2 := p.handler.RequestFunc(ctx, args.PmaMsg); err2 != nil {
    tickerCancel()
    err = thrift.WrapTException(err2)
  }
  tickerCancel()
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - PmaMsg
type PMAServiceRequestFuncArgs struct {
  PmaMsg *PMAMsg `thrift:"pmaMsg,1" db:"pmaMsg" json:"pmaMsg"`
}

func NewPMAServiceRequestFuncArgs() *PMAServiceRequestFuncArgs {
  return &PMAServiceRequestFuncArgs{}
}

var PMAServiceRequestFuncArgs_PmaMsg_DEFAULT *PMAMsg
func (p *PMAServiceRequestFuncArgs) GetPmaMsg() *PMAMsg {
  if !p.IsSetPmaMsg() {
    return PMAServiceRequestFuncArgs_PmaMsg_DEFAULT
  }
return p.PmaMsg
}
func (p *PMAServiceRequestFuncArgs) IsSetPmaMsg() bool {
  return p.PmaMsg != nil
}

func (p *PMAServiceRequestFuncArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *PMAServiceRequestFuncArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.PmaMsg = &PMAMsg{}
  if err := p.PmaMsg.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.PmaMsg), err)
  }
  return nil
}

func (p *PMAServiceRequestFuncArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "requestFunc_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PMAServiceRequestFuncArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "pmaMsg", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pmaMsg: ", p), err) }
  if err := p.PmaMsg.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.PmaMsg), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pmaMsg: ", p), err) }
  return err
}

func (p *PMAServiceRequestFuncArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PMAServiceRequestFuncArgs(%+v)", *p)
}


