syntax = "proto3";

package rpc


service Cli {
rpc PublishTx (Transaction) returns (Response) {}
rpc GetTransaction (TransactionKey) returns (Transaction) {}
rpc GetBalance (Key) returns (Value){}
rpc GetState (Key) returns (Value){}
rpc GetBlock (BlockKey) returns (BlockInfo){}
}

message Transaction {
bytes tx = 1;
}

message Response {
int32 code = 1;
}
