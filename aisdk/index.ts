import { createOllama } from "ollama-ai-provider";

const ollama = createOllama();

import { streamText } from "ai";
import { createServer } from "http";
import { log } from "console";
log("hello");
createServer(async (req, res) => {
  const result = streamText({
    model: ollama("qwen2.5:7b"),
    onError({ error }) {
      console.error(error); // your error logging logic here
    },
    prompt: "所以爱是什么",
  });

  result.pipeDataStreamToResponse(res);
  // example: use textStream as an async iterable
  for await (const textPart of result.textStream) {
    log(textPart);
  }
}).listen(3030);
