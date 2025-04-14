import { customProvider } from 'ai';
import { createOllama } from 'ollama-ai-provider';

const ollama = createOllama();

export const myProvider = customProvider({
      languageModels: {
        'deepseek-r1': ollama('deepseek-r1:7b'),
        'qwen2.5:7b': ollama('qwen2.5:7b'),
      }
});
