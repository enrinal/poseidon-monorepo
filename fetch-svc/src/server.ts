import http from "http";

import app from "./app";

require("dotenv").config();

const startServer = () => {
  const PORT = process.env.PORT || 7000;
  const server = http.createServer(app);
  return server
    .listen(PORT)
    .on("listening", () => console.log("Server is listening on port", PORT))
    .on("error", (err) => {
      console.log("Failed start sever. Err: ", err);
    });
};

startServer();
