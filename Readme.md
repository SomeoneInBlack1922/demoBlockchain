# Blockchein study
This folder contains a knock-off demo implementation of a blockchain for study.

This module contains thre executable code folders each containing several go files where one definces a main() function.
List of those folders with files in them containing main() function:
- example with main() in exampleMain.go
- rpc_client with main() in rpc_client.go
- web_api with main() in web_api.go

Each can be executed or compiled by passing the file containing main() function to go's cli tool with terminal command like this
`go run example/exampleMain.go` (assuming having repository's root directory as woking directory)

Folder bch contains a blockhain library where i implement the blockain itself with simple types, it is used by all executable folders and they all expect it to be in the same folder as they are to compile.

example is a non interractive cli program with no inputs that creates a lockchain of two blocks and prints them to stdout. It is just to show how the algorythm works.

rpc_client is a cli tool. It seems it was supposed to work as a cli blockchain client, but it is not finished, and i don't remember what i wanted it to do at this point. I also don't know why is it called RPC.

web_api is a http client that maintains a blockhain manipulates it on a request. This one actually implements something like RPC.