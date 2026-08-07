const express = require("express");
const app = express();

app.get("/", (request, response) => {
  return "Hello World!";
});

console.log("Listening on port 3000");
app.listen(3000);
