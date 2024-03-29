// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thrift

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"reflect"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - Dev
//  - Cid
type Devcid struct {
	Dev string `thrift:"Dev,1,required" db:"Dev" json:"Dev"`
	Cid string `thrift:"Cid,2,required" db:"Cid" json:"Cid"`
}

func NewDevcid() *Devcid {
	return &Devcid{}
}

func (p *Devcid) GetDev() string {
	return p.Dev
}

func (p *Devcid) GetCid() string {
	return p.Cid
}
func (p *Devcid) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetDev bool = false
	var issetCid bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
				issetDev = true
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
				issetCid = true
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetDev {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Dev is not set"))
	}
	if !issetCid {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Cid is not set"))
	}
	return nil
}

func (p *Devcid) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Dev = v
	}
	return nil
}

func (p *Devcid) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Cid = v
	}
	return nil
}

func (p *Devcid) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Devcid"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Devcid) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Dev", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Dev: ", p), err)
	}
	if err := oprot.WriteString(string(p.Dev)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Dev (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Dev: ", p), err)
	}
	return err
}

func (p *Devcid) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Cid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Cid: ", p), err)
	}
	if err := oprot.WriteString(string(p.Cid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Cid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Cid: ", p), err)
	}
	return err
}

func (p *Devcid) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Devcid(%+v)", *p)
}

// Attributes:
//  - Title
//  - Text
//  - Appcontent
type Tip struct {
	Title      string `thrift:"Title,1,required" db:"Title" json:"Title"`
	Text       string `thrift:"Text,2,required" db:"Text" json:"Text"`
	Appcontent string `thrift:"Appcontent,3" db:"Appcontent" json:"Appcontent"`
}

func NewTip() *Tip {
	return &Tip{}
}

func (p *Tip) GetTitle() string {
	return p.Title
}

func (p *Tip) GetText() string {
	return p.Text
}

func (p *Tip) GetAppcontent() string {
	return p.Appcontent
}
func (p *Tip) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTitle bool = false
	var issetText bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
				issetTitle = true
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
				issetText = true
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetTitle {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Title is not set"))
	}
	if !issetText {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Text is not set"))
	}
	return nil
}

func (p *Tip) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Title = v
	}
	return nil
}

func (p *Tip) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Text = v
	}
	return nil
}

func (p *Tip) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Appcontent = v
	}
	return nil
}

func (p *Tip) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Tip"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Tip) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Title", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Title: ", p), err)
	}
	if err := oprot.WriteString(string(p.Title)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Title (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Title: ", p), err)
	}
	return err
}

func (p *Tip) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Text", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Text: ", p), err)
	}
	if err := oprot.WriteString(string(p.Text)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Text (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Text: ", p), err)
	}
	return err
}

func (p *Tip) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Appcontent", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:Appcontent: ", p), err)
	}
	if err := oprot.WriteString(string(p.Appcontent)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Appcontent (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:Appcontent: ", p), err)
	}
	return err
}

func (p *Tip) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Tip(%+v)", *p)
}

type PmessService interface {
	// Parameters:
	//  - Oids
	//  - Ptype
	//  - Tip
	Push(ctx context.Context, oids []*Devcid, ptype int32, tip *Tip) (err error)
}

type PmessServiceClient struct {
	c thrift.TClient
}

func NewPmessServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PmessServiceClient {
	return &PmessServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPmessServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PmessServiceClient {
	return &PmessServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPmessServiceClient(c thrift.TClient) *PmessServiceClient {
	return &PmessServiceClient{
		c: c,
	}
}

func (p *PmessServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Oids
//  - Ptype
//  - Tip
func (p *PmessServiceClient) Push(ctx context.Context, oids []*Devcid, ptype int32, tip *Tip) (err error) {
	var _args0 PmessServicePushArgs
	_args0.Oids = oids
	_args0.Ptype = ptype
	_args0.Tip = tip
	var _result1 PmessServicePushResult
	if err = p.Client_().Call(ctx, "Push", &_args0, &_result1); err != nil {
		return
	}
	return nil
}

type PmessServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      PmessService
}

func (p *PmessServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *PmessServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *PmessServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewPmessServiceProcessor(handler PmessService) *PmessServiceProcessor {

	self2 := &PmessServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["Push"] = &pmessServiceProcessorPush{handler: handler}
	return self2
}

func (p *PmessServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x3

}

type pmessServiceProcessorPush struct {
	handler PmessService
}

func (p *pmessServiceProcessorPush) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PmessServicePushArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Push", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := PmessServicePushResult{}
	var err2 error
	if err2 = p.handler.Push(ctx, args.Oids, args.Ptype, args.Tip); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Push: "+err2.Error())
		oprot.WriteMessageBegin("Push", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("Push", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Oids
//  - Ptype
//  - Tip
type PmessServicePushArgs struct {
	Oids  []*Devcid `thrift:"oids,1" db:"oids" json:"oids"`
	Ptype int32     `thrift:"ptype,2" db:"ptype" json:"ptype"`
	Tip   *Tip      `thrift:"tip,3" db:"tip" json:"tip"`
}

func NewPmessServicePushArgs() *PmessServicePushArgs {
	return &PmessServicePushArgs{}
}

func (p *PmessServicePushArgs) GetOids() []*Devcid {
	return p.Oids
}

func (p *PmessServicePushArgs) GetPtype() int32 {
	return p.Ptype
}

var PmessServicePushArgs_Tip_DEFAULT *Tip

func (p *PmessServicePushArgs) GetTip() *Tip {
	if !p.IsSetTip() {
		return PmessServicePushArgs_Tip_DEFAULT
	}
	return p.Tip
}
func (p *PmessServicePushArgs) IsSetTip() bool {
	return p.Tip != nil
}

func (p *PmessServicePushArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PmessServicePushArgs) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Devcid, 0, size)
	p.Oids = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &Devcid{}
		if err := _elem4.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
		}
		p.Oids = append(p.Oids, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *PmessServicePushArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Ptype = v
	}
	return nil
}

func (p *PmessServicePushArgs) ReadField3(iprot thrift.TProtocol) error {
	p.Tip = &Tip{}
	if err := p.Tip.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tip), err)
	}
	return nil
}

func (p *PmessServicePushArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Push_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PmessServicePushArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("oids", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:oids: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Oids)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Oids {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:oids: ", p), err)
	}
	return err
}

func (p *PmessServicePushArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ptype", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ptype: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Ptype)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ptype (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ptype: ", p), err)
	}
	return err
}

func (p *PmessServicePushArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tip", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:tip: ", p), err)
	}
	if err := p.Tip.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tip), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:tip: ", p), err)
	}
	return err
}

func (p *PmessServicePushArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PmessServicePushArgs(%+v)", *p)
}

type PmessServicePushResult struct {
}

func NewPmessServicePushResult() *PmessServicePushResult {
	return &PmessServicePushResult{}
}

func (p *PmessServicePushResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PmessServicePushResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Push_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PmessServicePushResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PmessServicePushResult(%+v)", *p)
}
