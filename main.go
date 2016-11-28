package main

import (
	"fmt"

	"./messages"
	flatbuffers "github.com/google/flatbuffers/go"
)

func makeMessage(builder *flatbuffers.Builder, message []byte, number uint64) []byte {

	builder.Reset()
	messageCursor := builder.CreateByteString(message)
	messages.CommunicationStart(builder)
	messages.CommunicationAddMessage(builder, messageCursor)
	messages.CommunicationAddNumber(builder, number)
	messageCursor = messages.CommunicationEnd(builder)
	builder.Finish(messageCursor)
	return builder.Bytes[builder.Head():]
}

func readMessage(buffer []byte) (message []byte, number uint64) {

	// Initialize User reader from given buffer
	m := messages.GetRootAsCommunication(buffer, 0)
	message = m.Message()
	number = m.Number()
	return
}

func main() {
	builder := flatbuffers.NewBuilder(0) // 0 is size of builder
	buffer := makeMessage(builder, []byte("Dhriti Shikhar"), 1)
	message, number := readMessage(buffer)
	fmt.Printf("Message %d: %s.\nEncoded data is %d bytes long.\n", number, message, len(buffer))
}
