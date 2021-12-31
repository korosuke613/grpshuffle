import express from "express";
import morgan from "morgan";

/* eslint-disable node/no-unpublished-import */
import { GrpshuffleRequest, callShuffle } from "../grpshuffle-client";
/* eslint-enable node/no-unpublished-import */

const port = 8080;
const app = express();

// Setup the logger https://expressjs.com/en/resources/middleware/morgan.html
app.use(morgan("combined"));

// eslint-disable-next-line no-empty-pattern
app.get("/", ({}, res) => {
  res.json({ health: "ok" });
});

app.get(
  "/shuffle",
  // eslint-disable-next-line @typescript-eslint/ban-types
  async (request: express.Request<{}, {}, {}, GrpshuffleRequest>, response) => {
    const { partition, sequential } = request.query;
    const targets = request.query.targets;

    try {
      const result = await callShuffle({ partition, sequential, targets });
      response.json({ result });
    } catch (error) {
      response.status(500).json({ error });
    }
  }
);

app.listen(port, () => console.log(`Express server listening on port ${port}`));
