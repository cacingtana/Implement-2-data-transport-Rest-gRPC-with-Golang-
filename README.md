# gRPC-sql-test :star:
 gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment.

## Data initialization
 implemented use gRPC gateway, this example will have databases supported. There are MySQL using gorm as library.
 
 gRPC Design
 - Explicitly separate User-Side, Business Logic, and Server-Side
 - Dependencies are going from User-Side and Server-Side to the Business Logic

## How To Consume The API

	//list endpoint
	post: "/voucher/{id}"
	endpoint voucher with param by id (voucher no) to get voucher information, mercant name and valid date time 
	
	
	post: "/mercant/{mercant_id}/{id}"
  	endpoint mercant use to store / transaction voucher to mercant base on mercant_id and voucher_id  

	
## Relation table

	type Voucher struct {
	Id            uint
	MercantId     []Mercant
	VoucherNo     string
	VoucherStatus uint
	VoucherValid  timestamppb.Timestamp
	CreatedAt     time.Time
	UpdatedAt     time.Time
	}
	
	type Mercant struct {
	Id          int
	MercantName string
	MaxVoucher  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	}
	

