const express = require("express");
const albumArt = require("album-art");
const cors = require("cors");

const app = express();
const PORT = 4000;

// Advanced CORS configuration (optional)
const corsOptions = {
  origin: "http://localhost:3000",
  methods: "GET,HEAD,PUT,PATCH,POST,DELETE",
  allowedHeaders: ["Content-Type", "Authorization"],
  credentials: true,
  optionsSuccessStatus: 204,
};
app.use(cors(corsOptions));

app.get("/:artist/:album", async (req, res) => {
  const { artist, album } = req.params;
  const image = await albumArt(artist, { album, size: "medium" });
  res.send(image);
});

app.get("/:artist/:album/:size?", async (req, res) => {
  const { artist, album, size } = req.params;
  const image = await albumArt(artist, { album, size });
  res.send({ image });
});

app.listen(PORT, () => {
  console.log(`Album cover API listening on port ${PORT}`);
});
