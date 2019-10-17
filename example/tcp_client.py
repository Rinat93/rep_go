import socket
sock = socket.socket()
sock.connect(('localhost', 3333))
sock.send(b'hello, world!')
sock.close()