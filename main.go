package main

import (
	"fmt"

	"./messages"
	flatbuffers "github.com/google/flatbuffers/go"
)

func makeMessage(builder *flatbuffers.Builder, message []byte, id uint64) []byte {

	builder.Reset() // reuse buffer
	messageCursor := builder.CreateByteString(message)
	messages.CommunicationStart(builder)
	messages.CommunicationAddMessage(builder, messageCursor)
	messages.CommunicationAddId(builder, id)
	messageCursor = messages.CommunicationEnd(builder)
	builder.Finish(messageCursor)
	return builder.Bytes[builder.Head():]
}

func readMessage(buffer []byte) (message []byte, id uint64) {

	// Initialize User reader from given buffer
	message = messages.GetRootAsCommunication(buffer, 0)
	message = messages.Communication()
	id = message.Id()
	return
}

func main() {
	builder := flatbuffers.NewBuilder(0) // 0 is size of builder
	buffer := makeMessage(builder, []byte("Dhriti Shikhar"), 1)
	message, number := readMessage(buffer)
	fmt.Printf("Message %d: %s.\nEncoded data is %d bytes long.\n", id, message, len(buffer))
}
