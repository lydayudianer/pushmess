namespace go thrift
namespace java com.xingytech.tczj.pushmess

struct Devcid{
    1: required string Dev,
    2: required string Cid
}

struct Tip{
    1: required string Title,
    2: required string Text,
    3: string Appcontent
}

service PmessService {
	void Push(1: list<Devcid> oids,2: i32 ptype,3: Tip tip)
}