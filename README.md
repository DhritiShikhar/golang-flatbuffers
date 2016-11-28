Author: Dhriti Shikhar

Serialization
--------------

The process of converting an object into stream of bytes.

Purpose of serialization
-------------------------

- To recreate the object whenever needed
- retain state of object
- provide storage of object
- to provide support for data exchange

FlatBuffers
------------

- serialization format from Google
- its fast

This is the implementation of FlatBuffers in golang

How to run
-----------

1. Install flatbuffer from source. Reference --> https://rwinslow.com/posts/how-to-install-flatbuffers/

2. write a schema file. Here, myschema.fbs is the schema file.
Run it using the command:

./flatbuffers/flatc -g myschema.fbs

where flatc is the binary

3. write your program in main.go

4. run it using the command:

GOPATH=$(pwd) go run main.go
