#!/usr/bin/env node

import express from "express";
import morgan from "morgan";

/* eslint-disable node/no-unpublished-import */
import { GrpshuffleRequest, callShuffle, callHealth } from "./GrpshuffleClient";
/* eslint-enable node/no-unpublished-import */

const port = 8080;
const app = express();

// Setup the logger https://expressjs.com/en/resources/middleware/morgan.html
app.use(morgan("combined"));

// eslint-disable-next-line no-empty-pattern
app.get("/", async ({}, res) => {
  try {
    const result = await callHealth();
    res.json(result);
  } catch (error) {
    res.status(500).json({ error });
  }
});

app.get(
  "/shuffle",
  // eslint-disable-next-line @typescript-eslint/ban-types
  async (request: express.Request<{}, {}, {}, GrpshuffleRequest>, res) => {
    const { divide, sequential } = request.query;
    const targets = request.query.targets;

    try {
      const result = await callShuffle({ divide, sequential, targets });
      res.json(result);
    } catch (error) {
      res.status(500).json({ error });
    }
  }
);

app.listen(port, () => console.log(`Express server listening on port ${port}`));
