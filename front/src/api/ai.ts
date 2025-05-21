import axios, { type AxiosInstance } from "axios";
import { streamText } from "ai";
import { myProvider } from "../lib/providers";
import { systemPrompt } from "../lib/prompts";
import { ollama } from "ollama-ai-provider";

// 创建 axios 实例
const OllamaInstance: AxiosInstance = axios.create({
  baseURL: `http://${import.meta.env.VITE_BACK_URL}/api`,
  timeout: 10000,
  headers: {
    Accept: "application/vnd.github.v3+json",
  },
});

// 定义聊天接口函数
export const qwenAPI = async (messages: any[]) => {
  try {
    // 使用OllamaInstance发送请求
    await OllamaInstance.post('/chat', {
      messages,
      stream: true
    });

    const result = streamText({
      model: myProvider.languageModel("qwen2.5:7b"),
      system: systemPrompt("qwen2.5:7b"),
      messages,
    });
    // 返回响应数据
    return result.toDataStreamResponse();
  } catch (error) {
    console.error("Chat API Error:", error);
    throw error;
  }
};

//分析
export const ana = (file:any) => {
  return OllamaInstance.post("/ana",{
    files:file
  });
};