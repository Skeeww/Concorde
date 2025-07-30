const ws = require("ws");

const server = new ws.Server({
  port: 5500,
});

server.on("connection", (socket, _) => {
  setInterval(() => {
    socket.send("hey");
  }, 1000);
  socket.on("message", (data, _) => {
    console.log(data);
  });
});
