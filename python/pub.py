import time
import zmq

def send_data(socket, data, itr: int):
    start = time.time()
    for _ in range(itr):
        socket.send(data)
    elapsed = time.time() - start
    print(f"Time took: {elapsed}")

base = 10_000
num_tests = 4
data_size = 1_0

data = b"1"*data_size

context = zmq.Context()
socket = context.socket(zmq.PUB)
socket.bind("tcp://*:5555")

for i in range(num_tests):
    iterations = base * (10 ** i)
    send_data(socket, data, iterations)
