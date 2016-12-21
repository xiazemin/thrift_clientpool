// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package entity

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - BookId
//  - BookName
//  - Author
//  - Price
//  - Date
//  - Cover
type Book struct {
	BookId   string  `thrift:"BookId,1,required" json:"BookId"`
	BookName string  `thrift:"BookName,2,required" json:"BookName"`
	Author   string  `thrift:"Author,3,required" json:"Author"`
	Price    *string `thrift:"Price,4" json:"Price,omitempty"`
	Date     *string `thrift:"Date,5" json:"Date,omitempty"`
	Cover    []byte  `thrift:"Cover,6" json:"Cover,omitempty"`
}

func NewBook() *Book {
	return &Book{}
}

func (p *Book) GetBookId() string {
	return p.BookId
}

func (p *Book) GetBookName() string {
	return p.BookName
}

func (p *Book) GetAuthor() string {
	return p.Author
}

var Book_Price_DEFAULT string

func (p *Book) GetPrice() string {
	if !p.IsSetPrice() {
		return Book_Price_DEFAULT
	}
	return *p.Price
}

var Book_Date_DEFAULT string

func (p *Book) GetDate() string {
	if !p.IsSetDate() {
		return Book_Date_DEFAULT
	}
	return *p.Date
}

var Book_Cover_DEFAULT []byte

func (p *Book) GetCover() []byte {
	return p.Cover
}
func (p *Book) IsSetPrice() bool {
	return p.Price != nil
}

func (p *Book) IsSetDate() bool {
	return p.Date != nil
}

func (p *Book) IsSetCover() bool {
	return p.Cover != nil
}

func (p *Book) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetBookId bool = false
	var issetBookName bool = false
	var issetAuthor bool = false

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
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetBookId = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetBookName = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetAuthor = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
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
	if !issetBookId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field BookId is not set"))
	}
	if !issetBookName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field BookName is not set"))
	}
	if !issetAuthor {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Author is not set"))
	}
	return nil
}

func (p *Book) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.BookId = v
	}
	return nil
}

func (p *Book) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.BookName = v
	}
	return nil
}

func (p *Book) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Author = v
	}
	return nil
}

func (p *Book) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Price = &v
	}
	return nil
}

func (p *Book) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Date = &v
	}
	return nil
}

func (p *Book) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Cover = v
	}
	return nil
}

func (p *Book) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Book"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Book) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("BookId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:BookId: ", p), err)
	}
	if err := oprot.WriteString(string(p.BookId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.BookId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:BookId: ", p), err)
	}
	return err
}

func (p *Book) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("BookName", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:BookName: ", p), err)
	}
	if err := oprot.WriteString(string(p.BookName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.BookName (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:BookName: ", p), err)
	}
	return err
}

func (p *Book) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Author", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:Author: ", p), err)
	}
	if err := oprot.WriteString(string(p.Author)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Author (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:Author: ", p), err)
	}
	return err
}

func (p *Book) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetPrice() {
		if err := oprot.WriteFieldBegin("Price", thrift.STRING, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:Price: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Price)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Price (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:Price: ", p), err)
		}
	}
	return err
}

func (p *Book) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetDate() {
		if err := oprot.WriteFieldBegin("Date", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:Date: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Date)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Date (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:Date: ", p), err)
		}
	}
	return err
}

func (p *Book) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetCover() {
		if err := oprot.WriteFieldBegin("Cover", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:Cover: ", p), err)
		}
		if err := oprot.WriteBinary(p.Cover); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Cover (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:Cover: ", p), err)
		}
	}
	return err
}

func (p *Book) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Book(%+v)", *p)
}

// Attributes:
//  - UserId
//  - UserName
//  - BookNum
//  - BookList
//  - Avatar
type User struct {
	UserId   string          `thrift:"UserId,1,required" json:"UserId"`
	UserName string          `thrift:"UserName,2,required" json:"UserName"`
	BookNum  *int64          `thrift:"BookNum,3" json:"BookNum,omitempty"`
	BookList map[string]bool `thrift:"BookList,4" json:"BookList,omitempty"`
	Avatar   []byte          `thrift:"Avatar,5" json:"Avatar,omitempty"`
}

func NewUser() *User {
	return &User{}
}

func (p *User) GetUserId() string {
	return p.UserId
}

func (p *User) GetUserName() string {
	return p.UserName
}

var User_BookNum_DEFAULT int64

func (p *User) GetBookNum() int64 {
	if !p.IsSetBookNum() {
		return User_BookNum_DEFAULT
	}
	return *p.BookNum
}

var User_BookList_DEFAULT map[string]bool

func (p *User) GetBookList() map[string]bool {
	return p.BookList
}

var User_Avatar_DEFAULT []byte

func (p *User) GetAvatar() []byte {
	return p.Avatar
}
func (p *User) IsSetBookNum() bool {
	return p.BookNum != nil
}

func (p *User) IsSetBookList() bool {
	return p.BookList != nil
}

func (p *User) IsSetAvatar() bool {
	return p.Avatar != nil
}

func (p *User) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetUserId bool = false
	var issetUserName bool = false

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
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetUserId = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetUserName = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
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
	if !issetUserId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field UserId is not set"))
	}
	if !issetUserName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field UserName is not set"))
	}
	return nil
}

func (p *User) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.UserId = v
	}
	return nil
}

func (p *User) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.UserName = v
	}
	return nil
}

func (p *User) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.BookNum = &v
	}
	return nil
}

func (p *User) readField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadSetBegin()
	if err != nil {
		return thrift.PrependError("error reading set begin: ", err)
	}
	tSet := make(map[string]bool, size)
	p.BookList = tSet
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.BookList[_elem0] = true
	}
	if err := iprot.ReadSetEnd(); err != nil {
		return thrift.PrependError("error reading set end: ", err)
	}
	return nil
}

func (p *User) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Avatar = v
	}
	return nil
}

func (p *User) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("User"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *User) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("UserId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:UserId: ", p), err)
	}
	if err := oprot.WriteString(string(p.UserId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.UserId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:UserId: ", p), err)
	}
	return err
}

func (p *User) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("UserName", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:UserName: ", p), err)
	}
	if err := oprot.WriteString(string(p.UserName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.UserName (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:UserName: ", p), err)
	}
	return err
}

func (p *User) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetBookNum() {
		if err := oprot.WriteFieldBegin("BookNum", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:BookNum: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.BookNum)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.BookNum (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:BookNum: ", p), err)
		}
	}
	return err
}

func (p *User) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetBookList() {
		if err := oprot.WriteFieldBegin("BookList", thrift.SET, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:BookList: ", p), err)
		}
		if err := oprot.WriteSetBegin(thrift.STRING, len(p.BookList)); err != nil {
			return thrift.PrependError("error writing set begin: ", err)
		}
		for v, _ := range p.BookList {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteSetEnd(); err != nil {
			return thrift.PrependError("error writing set end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:BookList: ", p), err)
		}
	}
	return err
}

func (p *User) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetAvatar() {
		if err := oprot.WriteFieldBegin("Avatar", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:Avatar: ", p), err)
		}
		if err := oprot.WriteBinary(p.Avatar); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Avatar (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:Avatar: ", p), err)
		}
	}
	return err
}

func (p *User) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("User(%+v)", *p)
}
