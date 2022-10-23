import chat_service_pb2_grpc, chat_service_pb2
import time
import grpc
from concurrent import futures

class MockServer(chat_service_pb2_grpc.ChatServicer):
    def getChats(self, request, context):
       pass

    def getMessages(self, request, context):
       pass

    def readMessage(self, request, context):
       pass

    def editChat(self, request, context):
       pass

    def chatChanged(self, request, context):
       pass

    def sendMessage(self, request, context):
       pass

    def recieveMessage(self, request, context):
        messages = [chat_service_pb2.ExchangedMessage(chat_id="1"), chat_service_pb2.ExchangedMessage(chat_id="2"), chat_service_pb2.ExchangedMessage(chat_id="3")]
        for message in messages:
            yield message

        time.sleep(5)

        messages = [chat_service_pb2.ExchangedMessage(chat_id="4"), chat_service_pb2.ExchangedMessage(chat_id="5"), chat_service_pb2.ExchangedMessage(chat_id="6")]
        for message in messages:
            yield message


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_service_pb2_grpc.add_ChatServicer_to_server(MockServer(), server)
    server.add_insecure_port('[::]:8080')
    server.start()
    server.wait_for_termination()


serve()